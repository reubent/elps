package delveserver

import (
	"github.com/go-delve/delve/service/api"
	"github.com/go-delve/delve/service/rpc2"
	"github.com/google/go-dap"
	"time"
)

type ServerDebugger interface {
	GetBreakpoint(id int) (*api.Breakpoint, error)
	GetBreakpointByName(name string) (*api.Breakpoint, error)
	GetStacktrace(st *rpc2.StacktraceOut)
	GetDapStacktrace() []dap.StackFrame
	GetAllBreakpoints() map[int]*api.Breakpoint
	CreateBreakpoint(breakpoint *api.Breakpoint) *api.Breakpoint
	RemoveBreakpoint(id int) error
	AmendBreakpoint(bp *api.Breakpoint) error
	GetThread() *api.Thread
	GetVariables() []api.Variable
	GetFunctionArgs() []api.Variable
	SetVariableInScope(scope api.EvalScope, symbol string, value string) error
	Sources(filter string) ([]string, error)
	Functions(filter string) ([]string, error)
	FindLocation(scope api.EvalScope, loc string, lines bool) ([]api.Location, error)
	ListPackagesBuildInfo(files bool) []api.PackageBuildInfo
	State(blocking bool) (*api.DebuggerState, error)
	Command(a *api.DebuggerCommand) (*api.DebuggerState, error)
	LastModified() time.Time
	Complete() error
	GetGoRoutine() *api.Goroutine
	Step()
	Continue()
	Halt()
	IsStopped() bool
	GetModules() []dap.Module
	Eval(text string) string
	GetArguments() []dap.Variable
}