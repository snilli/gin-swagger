package testutil

import (
	"context"

	"github.com/example/ormprovider"
	"github.com/example/ormprovider/ent/enttest"
	_ "github.com/mattn/go-sqlite3"
)

// TestingT is a minimal interface for testing that matches both testing.TB and Ginkgo's interface
type TestingT interface {
	Fatalf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	FailNow()
	Helper()
}

// NewTestDBClient creates a new ORM client for testing using SQLite in-memory database
func NewTestDBClient(t TestingT) *ormprovider.Client {
	opts := []enttest.Option{
		enttest.WithOptions(),
	}

	// Create a wrapper that implements testing.TB
	wrapper := &testWrapper{t: t}

	client := enttest.Open(wrapper, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)

	// Run migrations
	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}

	return &ormprovider.Client{Client: client}
}

// testWrapper wraps TestingT to implement testing.TB
type testWrapper struct {
	t TestingT
}

func (w *testWrapper) Cleanup(func())            {}
func (w *testWrapper) Error(args ...interface{}) {}
func (w *testWrapper) Errorf(format string, args ...interface{}) {
	w.t.Errorf(format, args...)
}
func (w *testWrapper) Fail()                     {}
func (w *testWrapper) FailNow()                  { w.t.FailNow() }
func (w *testWrapper) Failed() bool              { return false }
func (w *testWrapper) Fatal(args ...interface{}) {}
func (w *testWrapper) Fatalf(format string, args ...interface{}) {
	w.t.Fatalf(format, args...)
}
func (w *testWrapper) Helper()                                  { w.t.Helper() }
func (w *testWrapper) Log(args ...interface{})                  {}
func (w *testWrapper) Logf(format string, args ...interface{})  {}
func (w *testWrapper) Name() string                             { return "test" }
func (w *testWrapper) Setenv(key, value string)                 {}
func (w *testWrapper) Skip(args ...interface{})                 {}
func (w *testWrapper) SkipNow()                                 {}
func (w *testWrapper) Skipf(format string, args ...interface{}) {}
func (w *testWrapper) Skipped() bool                            { return false }
func (w *testWrapper) TempDir() string                          { return "" }
