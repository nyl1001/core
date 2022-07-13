// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	enginetypes "github.com/projecteru2/core/engine/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/projecteru2/core/types"
)

// Cluster is an autogenerated mock type for the Cluster type
type Cluster struct {
	mock.Mock
}

// AddNode provides a mock function with given fields: _a0, _a1
func (_m *Cluster) AddNode(_a0 context.Context, _a1 *types.AddNodeOptions) (*types.Node, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.Node
	if rf, ok := ret.Get(0).(func(context.Context, *types.AddNodeOptions) *types.Node); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.AddNodeOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddPod provides a mock function with given fields: ctx, podname, desc
func (_m *Cluster) AddPod(ctx context.Context, podname string, desc string) (*types.Pod, error) {
	ret := _m.Called(ctx, podname, desc)

	var r0 *types.Pod
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *types.Pod); ok {
		r0 = rf(ctx, podname, desc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, podname, desc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildImage provides a mock function with given fields: ctx, opts
func (_m *Cluster) BuildImage(ctx context.Context, opts *types.BuildOptions) (chan *types.BuildImageMessage, error) {
	ret := _m.Called(ctx, opts)

	var r0 chan *types.BuildImageMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.BuildOptions) chan *types.BuildImageMessage); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.BuildImageMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.BuildOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CacheImage provides a mock function with given fields: ctx, opts
func (_m *Cluster) CacheImage(ctx context.Context, opts *types.ImageOptions) (chan *types.CacheImageMessage, error) {
	ret := _m.Called(ctx, opts)

	var r0 chan *types.CacheImageMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.ImageOptions) chan *types.CacheImageMessage); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.CacheImageMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.ImageOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CalculateCapacity provides a mock function with given fields: _a0, _a1
func (_m *Cluster) CalculateCapacity(_a0 context.Context, _a1 *types.DeployOptions) (*types.CapacityMessage, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.CapacityMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.DeployOptions) *types.CapacityMessage); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.CapacityMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.DeployOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConnectNetwork provides a mock function with given fields: ctx, network, target, ipv4, ipv6
func (_m *Cluster) ConnectNetwork(ctx context.Context, network string, target string, ipv4 string, ipv6 string) ([]string, error) {
	ret := _m.Called(ctx, network, target, ipv4, ipv6)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) []string); ok {
		r0 = rf(ctx, network, target, ipv4, ipv6)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, network, target, ipv4, ipv6)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControlWorkload provides a mock function with given fields: ctx, ids, t, force
func (_m *Cluster) ControlWorkload(ctx context.Context, ids []string, t string, force bool) (chan *types.ControlWorkloadMessage, error) {
	ret := _m.Called(ctx, ids, t, force)

	var r0 chan *types.ControlWorkloadMessage
	if rf, ok := ret.Get(0).(func(context.Context, []string, string, bool) chan *types.ControlWorkloadMessage); ok {
		r0 = rf(ctx, ids, t, force)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.ControlWorkloadMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string, string, bool) error); ok {
		r1 = rf(ctx, ids, t, force)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Copy provides a mock function with given fields: ctx, opts
func (_m *Cluster) Copy(ctx context.Context, opts *types.CopyOptions) (chan *types.CopyMessage, error) {
	ret := _m.Called(ctx, opts)

	var r0 chan *types.CopyMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.CopyOptions) chan *types.CopyMessage); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.CopyMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.CopyOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateWorkload provides a mock function with given fields: ctx, opts
func (_m *Cluster) CreateWorkload(ctx context.Context, opts *types.DeployOptions) (chan *types.CreateWorkloadMessage, error) {
	ret := _m.Called(ctx, opts)

	var r0 chan *types.CreateWorkloadMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.DeployOptions) chan *types.CreateWorkloadMessage); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.CreateWorkloadMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.DeployOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisconnectNetwork provides a mock function with given fields: ctx, network, target, force
func (_m *Cluster) DisconnectNetwork(ctx context.Context, network string, target string, force bool) error {
	ret := _m.Called(ctx, network, target, force)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) error); ok {
		r0 = rf(ctx, network, target, force)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DissociateWorkload provides a mock function with given fields: ctx, ids
func (_m *Cluster) DissociateWorkload(ctx context.Context, ids []string) (chan *types.DissociateWorkloadMessage, error) {
	ret := _m.Called(ctx, ids)

	var r0 chan *types.DissociateWorkloadMessage
	if rf, ok := ret.Get(0).(func(context.Context, []string) chan *types.DissociateWorkloadMessage); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.DissociateWorkloadMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteWorkload provides a mock function with given fields: ctx, opts, inCh
func (_m *Cluster) ExecuteWorkload(ctx context.Context, opts *types.ExecuteWorkloadOptions, inCh <-chan []byte) chan *types.AttachWorkloadMessage {
	ret := _m.Called(ctx, opts, inCh)

	var r0 chan *types.AttachWorkloadMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.ExecuteWorkloadOptions, <-chan []byte) chan *types.AttachWorkloadMessage); ok {
		r0 = rf(ctx, opts, inCh)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.AttachWorkloadMessage)
		}
	}

	return r0
}

// Finalizer provides a mock function with given fields:
func (_m *Cluster) Finalizer() {
	_m.Called()
}

// GetIdentifier provides a mock function with given fields:
func (_m *Cluster) GetIdentifier() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNode provides a mock function with given fields: ctx, nodename
func (_m *Cluster) GetNode(ctx context.Context, nodename string) (*types.Node, error) {
	ret := _m.Called(ctx, nodename)

	var r0 *types.Node
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.Node); ok {
		r0 = rf(ctx, nodename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Node)
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

// GetNodeEngineInfo provides a mock function with given fields: ctx, nodename
func (_m *Cluster) GetNodeEngineInfo(ctx context.Context, nodename string) (*enginetypes.Info, error) {
	ret := _m.Called(ctx, nodename)

	var r0 *enginetypes.Info
	if rf, ok := ret.Get(0).(func(context.Context, string) *enginetypes.Info); ok {
		r0 = rf(ctx, nodename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*enginetypes.Info)
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

// GetNodeStatus provides a mock function with given fields: ctx, nodename
func (_m *Cluster) GetNodeStatus(ctx context.Context, nodename string) (*types.NodeStatus, error) {
	ret := _m.Called(ctx, nodename)

	var r0 *types.NodeStatus
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.NodeStatus); ok {
		r0 = rf(ctx, nodename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.NodeStatus)
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

// GetPod provides a mock function with given fields: ctx, podname
func (_m *Cluster) GetPod(ctx context.Context, podname string) (*types.Pod, error) {
	ret := _m.Called(ctx, podname)

	var r0 *types.Pod
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.Pod); ok {
		r0 = rf(ctx, podname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, podname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkload provides a mock function with given fields: ctx, id
func (_m *Cluster) GetWorkload(ctx context.Context, id string) (*types.Workload, error) {
	ret := _m.Called(ctx, id)

	var r0 *types.Workload
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.Workload); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Workload)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkloads provides a mock function with given fields: ctx, ids
func (_m *Cluster) GetWorkloads(ctx context.Context, ids []string) ([]*types.Workload, error) {
	ret := _m.Called(ctx, ids)

	var r0 []*types.Workload
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*types.Workload); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Workload)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkloadsStatus provides a mock function with given fields: ctx, ids
func (_m *Cluster) GetWorkloadsStatus(ctx context.Context, ids []string) ([]*types.StatusMeta, error) {
	ret := _m.Called(ctx, ids)

	var r0 []*types.StatusMeta
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*types.StatusMeta); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.StatusMeta)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListImage provides a mock function with given fields: ctx, opts
func (_m *Cluster) ListImage(ctx context.Context, opts *types.ImageOptions) (chan *types.ListImageMessage, error) {
	ret := _m.Called(ctx, opts)

	var r0 chan *types.ListImageMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.ImageOptions) chan *types.ListImageMessage); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.ListImageMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.ImageOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListNetworks provides a mock function with given fields: ctx, podname, driver
func (_m *Cluster) ListNetworks(ctx context.Context, podname string, driver string) ([]*enginetypes.Network, error) {
	ret := _m.Called(ctx, podname, driver)

	var r0 []*enginetypes.Network
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []*enginetypes.Network); ok {
		r0 = rf(ctx, podname, driver)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*enginetypes.Network)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, podname, driver)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListNodeWorkloads provides a mock function with given fields: ctx, nodename, labels
func (_m *Cluster) ListNodeWorkloads(ctx context.Context, nodename string, labels map[string]string) ([]*types.Workload, error) {
	ret := _m.Called(ctx, nodename, labels)

	var r0 []*types.Workload
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) []*types.Workload); ok {
		r0 = rf(ctx, nodename, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Workload)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]string) error); ok {
		r1 = rf(ctx, nodename, labels)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPodNodes provides a mock function with given fields: _a0, _a1
func (_m *Cluster) ListPodNodes(_a0 context.Context, _a1 *types.ListNodesOptions) (<-chan *types.Node, error) {
	ret := _m.Called(_a0, _a1)

	var r0 <-chan *types.Node
	if rf, ok := ret.Get(0).(func(context.Context, *types.ListNodesOptions) <-chan *types.Node); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *types.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.ListNodesOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPods provides a mock function with given fields: ctx
func (_m *Cluster) ListPods(ctx context.Context) ([]*types.Pod, error) {
	ret := _m.Called(ctx)

	var r0 []*types.Pod
	if rf, ok := ret.Get(0).(func(context.Context) []*types.Pod); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Pod)
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

// ListWorkloads provides a mock function with given fields: ctx, opts
func (_m *Cluster) ListWorkloads(ctx context.Context, opts *types.ListWorkloadsOptions) ([]*types.Workload, error) {
	ret := _m.Called(ctx, opts)

	var r0 []*types.Workload
	if rf, ok := ret.Get(0).(func(context.Context, *types.ListWorkloadsOptions) []*types.Workload); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Workload)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.ListWorkloadsOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LogStream provides a mock function with given fields: ctx, opts
func (_m *Cluster) LogStream(ctx context.Context, opts *types.LogStreamOptions) (chan *types.LogStreamMessage, error) {
	ret := _m.Called(ctx, opts)

	var r0 chan *types.LogStreamMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.LogStreamOptions) chan *types.LogStreamMessage); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.LogStreamMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.LogStreamOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeResource provides a mock function with given fields: ctx, nodename, fix
func (_m *Cluster) NodeResource(ctx context.Context, nodename string, fix bool) (*types.NodeResource, error) {
	ret := _m.Called(ctx, nodename, fix)

	var r0 *types.NodeResource
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) *types.NodeResource); ok {
		r0 = rf(ctx, nodename, fix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.NodeResource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, nodename, fix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeStatusStream provides a mock function with given fields: ctx
func (_m *Cluster) NodeStatusStream(ctx context.Context) chan *types.NodeStatus {
	ret := _m.Called(ctx)

	var r0 chan *types.NodeStatus
	if rf, ok := ret.Get(0).(func(context.Context) chan *types.NodeStatus); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.NodeStatus)
		}
	}

	return r0
}

// PodResource provides a mock function with given fields: ctx, podname
func (_m *Cluster) PodResource(ctx context.Context, podname string) (chan *types.NodeResource, error) {
	ret := _m.Called(ctx, podname)

	var r0 chan *types.NodeResource
	if rf, ok := ret.Get(0).(func(context.Context, string) chan *types.NodeResource); ok {
		r0 = rf(ctx, podname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.NodeResource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, podname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReallocResource provides a mock function with given fields: ctx, opts
func (_m *Cluster) ReallocResource(ctx context.Context, opts *types.ReallocOptions) error {
	ret := _m.Called(ctx, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.ReallocOptions) error); ok {
		r0 = rf(ctx, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveImage provides a mock function with given fields: ctx, opts
func (_m *Cluster) RemoveImage(ctx context.Context, opts *types.ImageOptions) (chan *types.RemoveImageMessage, error) {
	ret := _m.Called(ctx, opts)

	var r0 chan *types.RemoveImageMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.ImageOptions) chan *types.RemoveImageMessage); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.RemoveImageMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.ImageOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveNode provides a mock function with given fields: ctx, nodename
func (_m *Cluster) RemoveNode(ctx context.Context, nodename string) error {
	ret := _m.Called(ctx, nodename)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, nodename)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemovePod provides a mock function with given fields: ctx, podname
func (_m *Cluster) RemovePod(ctx context.Context, podname string) error {
	ret := _m.Called(ctx, podname)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, podname)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveWorkload provides a mock function with given fields: ctx, ids, force
func (_m *Cluster) RemoveWorkload(ctx context.Context, ids []string, force bool) (chan *types.RemoveWorkloadMessage, error) {
	ret := _m.Called(ctx, ids, force)

	var r0 chan *types.RemoveWorkloadMessage
	if rf, ok := ret.Get(0).(func(context.Context, []string, bool) chan *types.RemoveWorkloadMessage); ok {
		r0 = rf(ctx, ids, force)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.RemoveWorkloadMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string, bool) error); ok {
		r1 = rf(ctx, ids, force)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveWorkloadSync provides a mock function with given fields: ctx, ids
func (_m *Cluster) RemoveWorkloadSync(ctx context.Context, ids []string) error {
	ret := _m.Called(ctx, ids)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) error); ok {
		r0 = rf(ctx, ids)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReplaceWorkload provides a mock function with given fields: ctx, opts
func (_m *Cluster) ReplaceWorkload(ctx context.Context, opts *types.ReplaceOptions) (chan *types.ReplaceWorkloadMessage, error) {
	ret := _m.Called(ctx, opts)

	var r0 chan *types.ReplaceWorkloadMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.ReplaceOptions) chan *types.ReplaceWorkloadMessage); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.ReplaceWorkloadMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.ReplaceOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunAndWait provides a mock function with given fields: ctx, opts, inCh
func (_m *Cluster) RunAndWait(ctx context.Context, opts *types.DeployOptions, inCh <-chan []byte) ([]string, <-chan *types.AttachWorkloadMessage, error) {
	ret := _m.Called(ctx, opts, inCh)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, *types.DeployOptions, <-chan []byte) []string); ok {
		r0 = rf(ctx, opts, inCh)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 <-chan *types.AttachWorkloadMessage
	if rf, ok := ret.Get(1).(func(context.Context, *types.DeployOptions, <-chan []byte) <-chan *types.AttachWorkloadMessage); ok {
		r1 = rf(ctx, opts, inCh)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(<-chan *types.AttachWorkloadMessage)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *types.DeployOptions, <-chan []byte) error); ok {
		r2 = rf(ctx, opts, inCh)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Send provides a mock function with given fields: ctx, opts
func (_m *Cluster) Send(ctx context.Context, opts *types.SendOptions) (chan *types.SendMessage, error) {
	ret := _m.Called(ctx, opts)

	var r0 chan *types.SendMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.SendOptions) chan *types.SendMessage); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.SendMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.SendOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetNode provides a mock function with given fields: ctx, opts
func (_m *Cluster) SetNode(ctx context.Context, opts *types.SetNodeOptions) (*types.Node, error) {
	ret := _m.Called(ctx, opts)

	var r0 *types.Node
	if rf, ok := ret.Get(0).(func(context.Context, *types.SetNodeOptions) *types.Node); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.SetNodeOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetNodeStatus provides a mock function with given fields: ctx, nodename, ttl
func (_m *Cluster) SetNodeStatus(ctx context.Context, nodename string, ttl int64) error {
	ret := _m.Called(ctx, nodename, ttl)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) error); ok {
		r0 = rf(ctx, nodename, ttl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetWorkloadsStatus provides a mock function with given fields: ctx, status, ttls
func (_m *Cluster) SetWorkloadsStatus(ctx context.Context, status []*types.StatusMeta, ttls map[string]int64) ([]*types.StatusMeta, error) {
	ret := _m.Called(ctx, status, ttls)

	var r0 []*types.StatusMeta
	if rf, ok := ret.Get(0).(func(context.Context, []*types.StatusMeta, map[string]int64) []*types.StatusMeta); ok {
		r0 = rf(ctx, status, ttls)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.StatusMeta)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []*types.StatusMeta, map[string]int64) error); ok {
		r1 = rf(ctx, status, ttls)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchServiceStatus provides a mock function with given fields: _a0
func (_m *Cluster) WatchServiceStatus(_a0 context.Context) (<-chan types.ServiceStatus, error) {
	ret := _m.Called(_a0)

	var r0 <-chan types.ServiceStatus
	if rf, ok := ret.Get(0).(func(context.Context) <-chan types.ServiceStatus); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan types.ServiceStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WorkloadStatusStream provides a mock function with given fields: ctx, appname, entrypoint, nodename, labels
func (_m *Cluster) WorkloadStatusStream(ctx context.Context, appname string, entrypoint string, nodename string, labels map[string]string) chan *types.WorkloadStatus {
	ret := _m.Called(ctx, appname, entrypoint, nodename, labels)

	var r0 chan *types.WorkloadStatus
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, map[string]string) chan *types.WorkloadStatus); ok {
		r0 = rf(ctx, appname, entrypoint, nodename, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.WorkloadStatus)
		}
	}

	return r0
}

type mockConstructorTestingTNewCluster interface {
	mock.TestingT
	Cleanup(func())
}

// NewCluster creates a new instance of Cluster. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCluster(t mockConstructorTestingTNewCluster) *Cluster {
	mock := &Cluster{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
