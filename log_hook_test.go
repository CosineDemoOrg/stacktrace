package stacktrace_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/palantir/stacktrace"
)

func TestLogHook(t *testing.T) {
	var capturedErr error
	var capturedCause error

	stacktrace.LogHook = func(e, c error) {
		capturedErr = e
		capturedCause = c
	}
	defer func() { stacktrace.LogHook = nil }()

	rootErr := errors.New("root")
	wrapped := stacktrace.PropagateSkip(rootErr, 0, "wrap")

	assert.Equal(t, wrapped, capturedErr, "LogHook should capture the newly created error")
	assert.Equal(t, rootErr, capturedCause, "LogHook should capture the cause error")
}