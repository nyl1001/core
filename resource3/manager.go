package resource3

import (
	"context"

	enginetypes "github.com/projecteru2/core/engine/types"
	plugintypes "github.com/projecteru2/core/resource3/plugins/types"
	"github.com/projecteru2/core/types"
)

// Manager indicate manages
// coretypes --> manager to rawparams --> plugins types
type Manager interface {
	AddNode(context.Context, string, types.Resources, *enginetypes.Info) (types.Resources, error)
	RemoveNode(context.Context, string) error
	GetMostIdleNode(context.Context, []string) (string, error)
	GetNodeResourceInfo(context.Context, string, []*types.Workload, bool) (types.Resources, types.Resources, []string, error)
	SetNodeResourceUsage(context.Context, string, types.Resources, types.Resources, []types.Resources, bool, bool) (types.Resources, types.Resources, error)
	GetNodesDeployCapacity(context.Context, []string, types.Resources) (map[string]*plugintypes.NodeDeployCapacity, int, error)
	SetNodeResourceCapacity(context.Context, string, types.Resources, types.Resources, bool, bool) (types.Resources, types.Resources, error)

	GetMetricsDescription(context.Context) ([]*plugintypes.MetricsDescription, error)
	GetNodeMetrics(context.Context, *types.Node) ([]*plugintypes.Metrics, error)

	Remap(context.Context, string, []*types.Workload) (types.Resources, error)

	Alloc(context.Context, string, int, types.Resources) ([]types.Resources, []types.Resources, error)
	RollbackAlloc(context.Context, string, []types.Resources) error

	Realloc(context.Context, string, types.Resources, types.Resources) (*plugintypes.EngineParams, types.Resources, types.Resources, error)
	RollbackRealloc(context.Context, string, types.Resources) error
}