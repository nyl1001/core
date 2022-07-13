// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	enginetypes "github.com/projecteru2/core/engine/types"
	mock "github.com/stretchr/testify/mock"

	resources "github.com/projecteru2/core/resources"

	types "github.com/projecteru2/core/types"
)

// Plugin is an autogenerated mock type for the Plugin type
type Plugin struct {
	mock.Mock
}

// AddNode provides a mock function with given fields: ctx, nodename, resourceOpts, nodeInfo
func (_m *Plugin) AddNode(ctx context.Context, nodename string, resourceOpts types.NodeResourceOpts, nodeInfo *enginetypes.Info) (*resources.AddNodeResponse, error) {
	ret := _m.Called(ctx, nodename, resourceOpts, nodeInfo)

	var r0 *resources.AddNodeResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, types.NodeResourceOpts, *enginetypes.Info) *resources.AddNodeResponse); ok {
		r0 = rf(ctx, nodename, resourceOpts, nodeInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resources.AddNodeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.NodeResourceOpts, *enginetypes.Info) error); ok {
		r1 = rf(ctx, nodename, resourceOpts, nodeInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConvertNodeResourceInfoToMetrics provides a mock function with given fields: ctx, podname, nodename, nodeResourceInfo
func (_m *Plugin) ConvertNodeResourceInfoToMetrics(ctx context.Context, podname string, nodename string, nodeResourceInfo *resources.NodeResourceInfo) (*resources.ConvertNodeResourceInfoToMetricsResponse, error) {
	ret := _m.Called(ctx, podname, nodename, nodeResourceInfo)

	var r0 *resources.ConvertNodeResourceInfoToMetricsResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, string, *resources.NodeResourceInfo) *resources.ConvertNodeResourceInfoToMetricsResponse); ok {
		r0 = rf(ctx, podname, nodename, nodeResourceInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resources.ConvertNodeResourceInfoToMetricsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, *resources.NodeResourceInfo) error); ok {
		r1 = rf(ctx, podname, nodename, nodeResourceInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeployArgs provides a mock function with given fields: ctx, nodename, deployCount, resourceOpts
func (_m *Plugin) GetDeployArgs(ctx context.Context, nodename string, deployCount int, resourceOpts types.WorkloadResourceOpts) (*resources.GetDeployArgsResponse, error) {
	ret := _m.Called(ctx, nodename, deployCount, resourceOpts)

	var r0 *resources.GetDeployArgsResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, int, types.WorkloadResourceOpts) *resources.GetDeployArgsResponse); ok {
		r0 = rf(ctx, nodename, deployCount, resourceOpts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resources.GetDeployArgsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int, types.WorkloadResourceOpts) error); ok {
		r1 = rf(ctx, nodename, deployCount, resourceOpts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMetricsDescription provides a mock function with given fields: ctx
func (_m *Plugin) GetMetricsDescription(ctx context.Context) (*resources.GetMetricsDescriptionResponse, error) {
	ret := _m.Called(ctx)

	var r0 *resources.GetMetricsDescriptionResponse
	if rf, ok := ret.Get(0).(func(context.Context) *resources.GetMetricsDescriptionResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resources.GetMetricsDescriptionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMostIdleNode provides a mock function with given fields: ctx, nodeNames
func (_m *Plugin) GetMostIdleNode(ctx context.Context, nodeNames []string) (*resources.GetMostIdleNodeResponse, error) {
	ret := _m.Called(ctx, nodeNames)

	var r0 *resources.GetMostIdleNodeResponse
	if rf, ok := ret.Get(0).(func(context.Context, []string) *resources.GetMostIdleNodeResponse); ok {
		r0 = rf(ctx, nodeNames)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resources.GetMostIdleNodeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, nodeNames)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeResourceInfo provides a mock function with given fields: ctx, nodename, workloads, fix
func (_m *Plugin) GetNodeResourceInfo(ctx context.Context, nodename string, workloads []*types.Workload, fix bool) (*resources.GetNodeResourceInfoResponse, error) {
	ret := _m.Called(ctx, nodename, workloads, fix)

	var r0 *resources.GetNodeResourceInfoResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, []*types.Workload, bool) *resources.GetNodeResourceInfoResponse); ok {
		r0 = rf(ctx, nodename, workloads, fix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resources.GetNodeResourceInfoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []*types.Workload, bool) error); ok {
		r1 = rf(ctx, nodename, workloads, fix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodesDeployCapacity provides a mock function with given fields: ctx, nodeNames, resourceOpts
func (_m *Plugin) GetNodesDeployCapacity(ctx context.Context, nodeNames []string, resourceOpts types.WorkloadResourceOpts) (*resources.GetNodesDeployCapacityResponse, error) {
	ret := _m.Called(ctx, nodeNames, resourceOpts)

	var r0 *resources.GetNodesDeployCapacityResponse
	if rf, ok := ret.Get(0).(func(context.Context, []string, types.WorkloadResourceOpts) *resources.GetNodesDeployCapacityResponse); ok {
		r0 = rf(ctx, nodeNames, resourceOpts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resources.GetNodesDeployCapacityResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string, types.WorkloadResourceOpts) error); ok {
		r1 = rf(ctx, nodeNames, resourceOpts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReallocArgs provides a mock function with given fields: ctx, nodename, originResourceArgs, resourceOpts
func (_m *Plugin) GetReallocArgs(ctx context.Context, nodename string, originResourceArgs types.WorkloadResourceArgs, resourceOpts types.WorkloadResourceOpts) (*resources.GetReallocArgsResponse, error) {
	ret := _m.Called(ctx, nodename, originResourceArgs, resourceOpts)

	var r0 *resources.GetReallocArgsResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, types.WorkloadResourceArgs, types.WorkloadResourceOpts) *resources.GetReallocArgsResponse); ok {
		r0 = rf(ctx, nodename, originResourceArgs, resourceOpts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resources.GetReallocArgsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.WorkloadResourceArgs, types.WorkloadResourceOpts) error); ok {
		r1 = rf(ctx, nodename, originResourceArgs, resourceOpts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRemapArgs provides a mock function with given fields: ctx, nodename, workloadMap
func (_m *Plugin) GetRemapArgs(ctx context.Context, nodename string, workloadMap map[string]*types.Workload) (*resources.GetRemapArgsResponse, error) {
	ret := _m.Called(ctx, nodename, workloadMap)

	var r0 *resources.GetRemapArgsResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]*types.Workload) *resources.GetRemapArgsResponse); ok {
		r0 = rf(ctx, nodename, workloadMap)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resources.GetRemapArgsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]*types.Workload) error); ok {
		r1 = rf(ctx, nodename, workloadMap)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Name provides a mock function with given fields:
func (_m *Plugin) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RemoveNode provides a mock function with given fields: ctx, nodename
func (_m *Plugin) RemoveNode(ctx context.Context, nodename string) (*resources.RemoveNodeResponse, error) {
	ret := _m.Called(ctx, nodename)

	var r0 *resources.RemoveNodeResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) *resources.RemoveNodeResponse); ok {
		r0 = rf(ctx, nodename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resources.RemoveNodeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nodename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetNodeResourceCapacity provides a mock function with given fields: ctx, nodename, nodeResourceOpts, nodeResourceArgs, delta, incr
func (_m *Plugin) SetNodeResourceCapacity(ctx context.Context, nodename string, nodeResourceOpts types.NodeResourceOpts, nodeResourceArgs types.NodeResourceArgs, delta bool, incr bool) (*resources.SetNodeResourceCapacityResponse, error) {
	ret := _m.Called(ctx, nodename, nodeResourceOpts, nodeResourceArgs, delta, incr)

	var r0 *resources.SetNodeResourceCapacityResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, types.NodeResourceOpts, types.NodeResourceArgs, bool, bool) *resources.SetNodeResourceCapacityResponse); ok {
		r0 = rf(ctx, nodename, nodeResourceOpts, nodeResourceArgs, delta, incr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resources.SetNodeResourceCapacityResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.NodeResourceOpts, types.NodeResourceArgs, bool, bool) error); ok {
		r1 = rf(ctx, nodename, nodeResourceOpts, nodeResourceArgs, delta, incr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetNodeResourceInfo provides a mock function with given fields: ctx, nodename, resourceCapacity, resourceUsage
func (_m *Plugin) SetNodeResourceInfo(ctx context.Context, nodename string, resourceCapacity types.NodeResourceArgs, resourceUsage types.NodeResourceArgs) (*resources.SetNodeResourceInfoResponse, error) {
	ret := _m.Called(ctx, nodename, resourceCapacity, resourceUsage)

	var r0 *resources.SetNodeResourceInfoResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, types.NodeResourceArgs, types.NodeResourceArgs) *resources.SetNodeResourceInfoResponse); ok {
		r0 = rf(ctx, nodename, resourceCapacity, resourceUsage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resources.SetNodeResourceInfoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.NodeResourceArgs, types.NodeResourceArgs) error); ok {
		r1 = rf(ctx, nodename, resourceCapacity, resourceUsage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetNodeResourceUsage provides a mock function with given fields: ctx, nodename, nodeResourceOpts, nodeResourceArgs, workloadResourceArgs, delta, incr
func (_m *Plugin) SetNodeResourceUsage(ctx context.Context, nodename string, nodeResourceOpts types.NodeResourceOpts, nodeResourceArgs types.NodeResourceArgs, workloadResourceArgs []types.WorkloadResourceArgs, delta bool, incr bool) (*resources.SetNodeResourceUsageResponse, error) {
	ret := _m.Called(ctx, nodename, nodeResourceOpts, nodeResourceArgs, workloadResourceArgs, delta, incr)

	var r0 *resources.SetNodeResourceUsageResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, types.NodeResourceOpts, types.NodeResourceArgs, []types.WorkloadResourceArgs, bool, bool) *resources.SetNodeResourceUsageResponse); ok {
		r0 = rf(ctx, nodename, nodeResourceOpts, nodeResourceArgs, workloadResourceArgs, delta, incr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resources.SetNodeResourceUsageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.NodeResourceOpts, types.NodeResourceArgs, []types.WorkloadResourceArgs, bool, bool) error); ok {
		r1 = rf(ctx, nodename, nodeResourceOpts, nodeResourceArgs, workloadResourceArgs, delta, incr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPlugin interface {
	mock.TestingT
	Cleanup(func())
}

// NewPlugin creates a new instance of Plugin. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPlugin(t mockConstructorTestingTNewPlugin) *Plugin {
	mock := &Plugin{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
