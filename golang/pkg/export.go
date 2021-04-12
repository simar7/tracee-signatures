package main

import (
	"github.com/aquasecurity/tracee/tracee-rules/types"
	traceesigs "github.com/simar7/tracee-signatures/golang"
)

// ExportedSignatures fulfills the goplugins contract required by the rule-engine
// this is a list of signatures that this plugin exports
var ExportedSignatures = []types.Signature{
	&traceesigs.StdioOverSocket{},
}
