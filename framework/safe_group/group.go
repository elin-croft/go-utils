package safe_group

import (
	"context"
	"sync"
)

type SafeGroup interface {
	Go(f func() error)
	Wait() error
}

type Group struct {
	name    string
	cancel  func()
	wg      sync.WaitGroup
	err     error
	errOnce sync.Once
}

func NewGroup(name string) *Group {
	return &Group{name: name}
}

func NewGroupWithContext(ctx context.Context, name string) (*Group, context.Context) {
	gCtx, cancel := context.WithCancel(ctx)
	g := &Group{name: name, cancel: cancel}
	return g, gCtx
}

func (g *Group) Go(fn func() error) {
	if len(g.name) == 0 {
		panic("safe group: empty name")
	}
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		if err := fn(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				g.cancel()
			})
		}
	}()
}

func (g *Group) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}
