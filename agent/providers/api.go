package providers

import (
	"errors"

	"github.com/synadia-io/nex/agent/providers/lib"
	agentapi "github.com/synadia-io/nex/internal/agent-api"
)

// NexExecutionProviderELF Executable Linkable Format execution provider
const NexExecutionProviderELF = "elf"

// NexExecutionProviderV8 V8 execution provider
const NexExecutionProviderV8 = "v8"

// NexExecutionProviderOCI OCI execution provider
const NexExecutionProviderOCI = "oci"

// NexExecutionProviderWasm Wasm execution provider
const NexExecutionProviderWasm = "wasm"

// ExecutionProvider implementations provide support for a specific
// execution environment pattern -- e.g., statically-linked ELF
// binaries, serverless JavaScript functions, OCI images, Wasm, etc.
type ExecutionProvider interface {
	// Deploy a service (e.g., "elf" and "oci" types) or executable function (e.g., "v8" and "wasm" types)
	Deploy() error

	// Undeploy a workload, giving it a chance to gracefully clean up after itself (if applicable)
	UnDeploy() error

	// Execute a deployed function, if supported by the execution provider implementation (e.g., "v8" and "wasm" types)
	Execute(subject string, payload []byte) ([]byte, error)

	// Validate the executable artifact, e.g., specific characteristics of a
	// statically-linked binary or raw source code, depending on provider implementation
	Validate() error
}

// NewExecutionProvider initializes and returns an execution provider for a given work request
func NewExecutionProvider(params *agentapi.ExecutionProviderParams) (ExecutionProvider, error) {
	if params.WorkloadType == nil {
		return nil, errors.New("execution provider factory requires a workload type parameter")
	}

	switch *params.WorkloadType {
	case NexExecutionProviderELF:
		return lib.InitNexExecutionProviderELF(params)
	case NexExecutionProviderV8:
		return lib.InitNexExecutionProviderV8(params)
	case NexExecutionProviderOCI:
		// TODO-- return lib.InitNexExecutionProviderOCI(params), nil
		return nil, errors.New("oci execution provider not yet implemented")
	case NexExecutionProviderWasm:
		return lib.InitNexExecutionProviderWasm(params)
	default:
		break
	}

	return nil, errors.New("invalid execution provider specified")
}
