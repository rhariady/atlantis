// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/runatlantis/atlantis/server/events (interfaces: Executor)

package mocks

import (
	"reflect"

	pegomock "github.com/petergtz/pegomock"
	events "github.com/runatlantis/atlantis/server/events"
)

type MockExecutor struct {
	fail func(message string, callerSkip ...int)
}

func NewMockExecutor() *MockExecutor {
	return &MockExecutor{fail: pegomock.GlobalFailHandler}
}

func (mock *MockExecutor) Execute(ctx *events.CommandContext) events.CommandResponse {
	params := []pegomock.Param{ctx}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Execute", params, []reflect.Type{reflect.TypeOf((*events.CommandResponse)(nil)).Elem()})
	var ret0 events.CommandResponse
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(events.CommandResponse)
		}
	}
	return ret0
}

func (mock *MockExecutor) VerifyWasCalledOnce() *VerifierExecutor {
	return &VerifierExecutor{mock, pegomock.Times(1), nil}
}

func (mock *MockExecutor) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierExecutor {
	return &VerifierExecutor{mock, invocationCountMatcher, nil}
}

func (mock *MockExecutor) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierExecutor {
	return &VerifierExecutor{mock, invocationCountMatcher, inOrderContext}
}

type VerifierExecutor struct {
	mock                   *MockExecutor
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierExecutor) Execute(ctx *events.CommandContext) *Executor_Execute_OngoingVerification {
	params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Execute", params)
	return &Executor_Execute_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Executor_Execute_OngoingVerification struct {
	mock              *MockExecutor
	methodInvocations []pegomock.MethodInvocation
}

func (c *Executor_Execute_OngoingVerification) GetCapturedArguments() *events.CommandContext {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *Executor_Execute_OngoingVerification) GetAllCapturedArguments() (_param0 []*events.CommandContext) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*events.CommandContext, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*events.CommandContext)
		}
	}
	return
}
