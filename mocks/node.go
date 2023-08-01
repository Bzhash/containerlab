// Code generated by MockGen. DO NOT EDIT.
// Source: nodes/node.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	ns "github.com/containernetworking/plugins/pkg/ns"
	gomock "github.com/golang/mock/gomock"
	exec "github.com/srl-labs/containerlab/clab/exec"
	nodes "github.com/srl-labs/containerlab/nodes"
	runtime "github.com/srl-labs/containerlab/runtime"
	types "github.com/srl-labs/containerlab/types"
	netlink "github.com/vishvananda/netlink"
)

// MockNode is a mock of Node interface.
type MockNode struct {
	ctrl     *gomock.Controller
	recorder *MockNodeMockRecorder
}

// MockNodeMockRecorder is the mock recorder for MockNode.
type MockNodeMockRecorder struct {
	mock *MockNode
}

// NewMockNode creates a new mock instance.
func NewMockNode(ctrl *gomock.Controller) *MockNode {
	mock := &MockNode{ctrl: ctrl}
	mock.recorder = &MockNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNode) EXPECT() *MockNodeMockRecorder {
	return m.recorder
}

// AddEndpoint mocks base method.
func (m *MockNode) AddEndpoint(e types.Endpt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEndpoint", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEndpoint indicates an expected call of AddEndpoint.
func (mr *MockNodeMockRecorder) AddEndpoint(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEndpoint", reflect.TypeOf((*MockNode)(nil).AddEndpoint), e)
}

// AddNetlinkLinkToContainer mocks base method.
func (m *MockNode) AddNetlinkLinkToContainer(ctx context.Context, link netlink.Link, f func(ns.NetNS) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNetlinkLinkToContainer", ctx, link, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNetlinkLinkToContainer indicates an expected call of AddNetlinkLinkToContainer.
func (mr *MockNodeMockRecorder) AddNetlinkLinkToContainer(ctx, link, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNetlinkLinkToContainer", reflect.TypeOf((*MockNode)(nil).AddNetlinkLinkToContainer), ctx, link, f)
}

// CheckDeploymentConditions mocks base method.
func (m *MockNode) CheckDeploymentConditions(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckDeploymentConditions", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckDeploymentConditions indicates an expected call of CheckDeploymentConditions.
func (mr *MockNodeMockRecorder) CheckDeploymentConditions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDeploymentConditions", reflect.TypeOf((*MockNode)(nil).CheckDeploymentConditions), arg0)
}

// CheckInterfaceName mocks base method.
func (m *MockNode) CheckInterfaceName() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckInterfaceName")
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckInterfaceName indicates an expected call of CheckInterfaceName.
func (mr *MockNodeMockRecorder) CheckInterfaceName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckInterfaceName", reflect.TypeOf((*MockNode)(nil).CheckInterfaceName))
}

// Config mocks base method.
func (m *MockNode) Config() *types.NodeConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*types.NodeConfig)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockNodeMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockNode)(nil).Config))
}

// Delete mocks base method.
func (m *MockNode) Delete(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockNodeMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNode)(nil).Delete), arg0)
}

// DeleteNetnsSymlink mocks base method.
func (m *MockNode) DeleteNetnsSymlink() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNetnsSymlink")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNetnsSymlink indicates an expected call of DeleteNetnsSymlink.
func (mr *MockNodeMockRecorder) DeleteNetnsSymlink() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNetnsSymlink", reflect.TypeOf((*MockNode)(nil).DeleteNetnsSymlink))
}

// Deploy mocks base method.
func (m *MockNode) Deploy(arg0 context.Context, arg1 *nodes.DeployParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy.
func (mr *MockNodeMockRecorder) Deploy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockNode)(nil).Deploy), arg0, arg1)
}

// GenerateConfig mocks base method.
func (m *MockNode) GenerateConfig(dst, templ string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateConfig", dst, templ)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateConfig indicates an expected call of GenerateConfig.
func (mr *MockNodeMockRecorder) GenerateConfig(dst, templ interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateConfig", reflect.TypeOf((*MockNode)(nil).GenerateConfig), dst, templ)
}

// GetContainers mocks base method.
func (m *MockNode) GetContainers(ctx context.Context) ([]runtime.GenericContainer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainers", ctx)
	ret0, _ := ret[0].([]runtime.GenericContainer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainers indicates an expected call of GetContainers.
func (mr *MockNodeMockRecorder) GetContainers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainers", reflect.TypeOf((*MockNode)(nil).GetContainers), ctx)
}

// GetEndpoints mocks base method.
func (m *MockNode) GetEndpoints() []types.Endpt {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEndpoints")
	ret0, _ := ret[0].([]types.Endpt)
	return ret0
}

// GetEndpoints indicates an expected call of GetEndpoints.
func (mr *MockNodeMockRecorder) GetEndpoints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEndpoints", reflect.TypeOf((*MockNode)(nil).GetEndpoints))
}

// GetImages mocks base method.
func (m *MockNode) GetImages(arg0 context.Context) map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImages", arg0)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetImages indicates an expected call of GetImages.
func (mr *MockNodeMockRecorder) GetImages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImages", reflect.TypeOf((*MockNode)(nil).GetImages), arg0)
}

// GetLinkEndpointType mocks base method.
func (m *MockNode) GetLinkEndpointType() types.LinkEndpointType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLinkEndpointType")
	ret0, _ := ret[0].(types.LinkEndpointType)
	return ret0
}

// GetLinkEndpointType indicates an expected call of GetLinkEndpointType.
func (mr *MockNodeMockRecorder) GetLinkEndpointType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLinkEndpointType", reflect.TypeOf((*MockNode)(nil).GetLinkEndpointType))
}

// GetRuntime mocks base method.
func (m *MockNode) GetRuntime() runtime.ContainerRuntime {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRuntime")
	ret0, _ := ret[0].(runtime.ContainerRuntime)
	return ret0
}

// GetRuntime indicates an expected call of GetRuntime.
func (mr *MockNodeMockRecorder) GetRuntime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRuntime", reflect.TypeOf((*MockNode)(nil).GetRuntime))
}

// GetShortName mocks base method.
func (m *MockNode) GetShortName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShortName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetShortName indicates an expected call of GetShortName.
func (mr *MockNodeMockRecorder) GetShortName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShortName", reflect.TypeOf((*MockNode)(nil).GetShortName))
}

// Init mocks base method.
func (m *MockNode) Init(arg0 *types.NodeConfig, arg1 ...nodes.NodeOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Init", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockNodeMockRecorder) Init(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockNode)(nil).Init), varargs...)
}

// PostDeploy mocks base method.
func (m *MockNode) PostDeploy(ctx context.Context, params *nodes.PostDeployParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostDeploy", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostDeploy indicates an expected call of PostDeploy.
func (mr *MockNodeMockRecorder) PostDeploy(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostDeploy", reflect.TypeOf((*MockNode)(nil).PostDeploy), ctx, params)
}

// PreDeploy mocks base method.
func (m *MockNode) PreDeploy(ctx context.Context, params *nodes.PreDeployParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreDeploy", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// PreDeploy indicates an expected call of PreDeploy.
func (mr *MockNodeMockRecorder) PreDeploy(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreDeploy", reflect.TypeOf((*MockNode)(nil).PreDeploy), ctx, params)
}

// RunExec mocks base method.
func (m *MockNode) RunExec(ctx context.Context, execCmd *exec.ExecCmd) (*exec.ExecResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunExec", ctx, execCmd)
	ret0, _ := ret[0].(*exec.ExecResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunExec indicates an expected call of RunExec.
func (mr *MockNodeMockRecorder) RunExec(ctx, execCmd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunExec", reflect.TypeOf((*MockNode)(nil).RunExec), ctx, execCmd)
}

// SaveConfig mocks base method.
func (m *MockNode) SaveConfig(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveConfig indicates an expected call of SaveConfig.
func (mr *MockNodeMockRecorder) SaveConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveConfig", reflect.TypeOf((*MockNode)(nil).SaveConfig), arg0)
}

// SetupNetworking mocks base method.
func (m *MockNode) SetupNetworking(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupNetworking", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupNetworking indicates an expected call of SetupNetworking.
func (mr *MockNodeMockRecorder) SetupNetworking(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupNetworking", reflect.TypeOf((*MockNode)(nil).SetupNetworking), ctx)
}

// UpdateConfigWithRuntimeInfo mocks base method.
func (m *MockNode) UpdateConfigWithRuntimeInfo(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConfigWithRuntimeInfo", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConfigWithRuntimeInfo indicates an expected call of UpdateConfigWithRuntimeInfo.
func (mr *MockNodeMockRecorder) UpdateConfigWithRuntimeInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfigWithRuntimeInfo", reflect.TypeOf((*MockNode)(nil).UpdateConfigWithRuntimeInfo), arg0)
}

// VerifyStartupConfig mocks base method.
func (m *MockNode) VerifyStartupConfig(topoDir string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyStartupConfig", topoDir)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyStartupConfig indicates an expected call of VerifyStartupConfig.
func (mr *MockNodeMockRecorder) VerifyStartupConfig(topoDir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyStartupConfig", reflect.TypeOf((*MockNode)(nil).VerifyStartupConfig), topoDir)
}

// WithMgmtNet mocks base method.
func (m *MockNode) WithMgmtNet(arg0 *types.MgmtNet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WithMgmtNet", arg0)
}

// WithMgmtNet indicates an expected call of WithMgmtNet.
func (mr *MockNodeMockRecorder) WithMgmtNet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithMgmtNet", reflect.TypeOf((*MockNode)(nil).WithMgmtNet), arg0)
}

// WithRuntime mocks base method.
func (m *MockNode) WithRuntime(arg0 runtime.ContainerRuntime) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WithRuntime", arg0)
}

// WithRuntime indicates an expected call of WithRuntime.
func (mr *MockNodeMockRecorder) WithRuntime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithRuntime", reflect.TypeOf((*MockNode)(nil).WithRuntime), arg0)
}
