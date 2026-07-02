package runtime_test

import (
	"context"
	"testing"

	"github.com/xukin10/AetherOS/core/runtime"
)

type mockEngine struct {
	name string
	log  []string
}

func (m *mockEngine) Name() string { return m.name }

func (m *mockEngine) Start(ctx context.Context) error {
	m.log = append(m.log, "start")
	return nil
}

func (m *mockEngine) Stop(ctx context.Context) error {
	m.log = append(m.log, "stop")
	return nil
}

func TestRuntimeStartStop(t *testing.T) {
	r := runtime.Bootstrap()

	e1 := &mockEngine{name: "e1"}
	e2 := &mockEngine{name: "e2"}

	r.Register(e1)
	r.Register(e2)

	if err := r.Start(context.Background()); err != nil {
		t.Fatal(err)
	}

	if r.State() != runtime.Running {
		t.Fatal("expected running state")
	}

	if err := r.Stop(context.Background()); err != nil {
		t.Fatal(err)
	}

	if r.State() != runtime.Stopped {
		t.Fatal("expected stopped state")
	}
}