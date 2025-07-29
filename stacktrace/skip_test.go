package stacktrace_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"stacktrace"
)

func helper(err error) error {
	return stacktrace.PropagateSkip(err, 1, "wrapped")
}

func TestPropagateSkip(t *testing.T) {
	root := errors.New("root")
	err := helper(root)
	errStr := err.Error()

	assert.Contains(t, errStr, "TestPropagateSkip", "Error string should contain the caller site (TestPropagateSkip)")
	assert.NotContains(t, errStr, "helper", "Error string should NOT contain the helper wrapper function")
}

func builderHelper(err error) error {
	b := stacktrace.NewErrorBuilder().WithCode(EcodeInvalidVillain).WithSkip(1)
	return b.Propagate(err, "builder wrapped")
}

func TestBuilderSkip(t *testing.T) {
	root := errors.New("root")
	err := builderHelper(root)
	errStr := err.Error()

	assert.Contains(t, errStr, "TestBuilderSkip", "Error string should contain the caller site (TestBuilderSkip)")
	assert.NotContains(t, errStr, "builderHelper", "Error string should NOT contain the builderHelper wrapper function")
	assert.Equal(t, EcodeInvalidVillain, stacktrace.GetCode(err), "Error code should be EcodeInvalidVillain")
}