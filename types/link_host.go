package types

import (
	"context"
	"fmt"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"
)

// LinkHostRaw is the raw (string) representation of a host link as defined in the topology file.
type LinkHostRaw struct {
	LinkCommonParams `yaml:",inline"`
	HostInterface    string       `yaml:"host-interface"`
	Endpoint         *EndpointRaw `yaml:"endpoint"`
}

// ToLinkConfig converts the raw link into a LinkConfig.
func (r *LinkHostRaw) ToLinkConfig() *LinkConfig {
	lc := &LinkConfig{
		Vars:      r.Vars,
		Labels:    r.Labels,
		MTU:       r.Mtu,
		Endpoints: make([]string, 2),
	}

	lc.Endpoints[0] = fmt.Sprintf("%s:%s", r.Endpoint.Node, r.Endpoint.Iface)
	lc.Endpoints[1] = fmt.Sprintf("%s:%s", "host", r.HostInterface)

	return lc
}

func hostFromLinkConfig(lc LinkConfig, specialEPIndex int) (*LinkHostRaw, error) {
	_, hostIf, node, nodeIf := extractHostNodeInterfaceData(lc, specialEPIndex)

	result := &LinkHostRaw{
		LinkCommonParams: LinkCommonParams{
			Mtu:    lc.MTU,
			Labels: lc.Labels,
			Vars:   lc.Vars,
		},
		HostInterface: hostIf,
		Endpoint:      NewEndpointRaw(node, nodeIf, ""),
	}
	return result, nil
}

func (r *LinkHostRaw) GetType() LinkType {
	return LinkTypeHost
}

func (r *LinkHostRaw) Resolve(params *ResolveParams) (LinkInterf, error) {
	link := &LinkHost{
		LinkCommonParams: r.LinkCommonParams,
		HostInterface:    r.HostInterface,
	}
	// resolve and populate the endpoint
	ep, err := r.Endpoint.Resolve(params.Nodes, link)
	if err != nil {
		return nil, err
	}
	// set the end point in the link
	link.Endpoint = ep
	return link, nil
}

type LinkHost struct {
	LinkCommonParams `yaml:",inline"`
	HostInterface    string        `yaml:"host-interface"`
	Endpoint         *EndptGeneric `yaml:"endpoint"`
}

func (l *LinkHost) Deploy(ctx context.Context) error {
	// build the netlink.Veth struct for the link provisioning
	link := &netlink.Veth{
		LinkAttrs: netlink.LinkAttrs{
			Name: l.Endpoint.GetRandIfaceName(),
			MTU:  l.Mtu,
			// Mac address is set later on
		},
		PeerName: l.HostInterface,
		// PeerMac address is set later on
	}

	// add the link
	err := netlink.LinkAdd(link)
	if err != nil {
		return err
	}

	// add link to node, rename, set mac and Up
	err = l.Endpoint.GetNode().AddNetlinkLinkToContainer(ctx, link, SetNameMACAndUpInterface(link, l.Endpoint))
	if err != nil {
		return err
	}

	// get the link on the host side
	hostLink, err := netlink.LinkByName(l.HostInterface)
	if err != nil {
		return err
	}

	// set the host side link to up
	err = netlink.LinkSetUp(hostLink)
	if err != nil {
		return err
	}

	return nil
}

func (l *LinkHost) GetType() LinkType {
	return LinkTypeHost
}

func (l *LinkHost) Remove(ctx context.Context) error {
	// TODO
	return nil
}

func (l *LinkHost) GetEndpoints() []Endpt {
	return []Endpt{
		l.Endpoint,
		&EndptGeneric{
			state: EndptDeployStateDeployed,
			Node:  GetFakeHostLinkNode(),
			Iface: l.HostInterface,
			Link:  l,
		}}
}

type GenericLinkNode struct {
	shortname string
	endpoints []Endpt
	nspath    string
}

func (g *GenericLinkNode) AddNetlinkLinkToContainer(ctx context.Context, link netlink.Link, f func(ns.NetNS) error) error {
	if g.nspath != "" {
		return nil
	}

	// retrieve the namespace handle
	netns, err := ns.GetNS(g.nspath)
	if err != nil {
		return err
	}
	// move veth endpoint to namespace
	if err = netlink.LinkSetNsFd(link, int(netns.Fd())); err != nil {
		return err
	}
	// execute the given function
	return netns.Do(f)
}
func (g *GenericLinkNode) AddEndpoint(e Endpt) error {
	g.endpoints = append(g.endpoints, e)
	return nil
}
func (g *GenericLinkNode) GetLinkEndpointType() LinkEndpointType {
	return LinkEndpointTypeRegular
}
func (g *GenericLinkNode) GetShortName() string {
	return g.shortname
}