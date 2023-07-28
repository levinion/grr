package grr

import (
	"runtime"
)

type Result[T any] struct {
	value T
	err   error
}

// 0 - OK; 1 - Err
func (r *Result[T]) t() uint {
	if r.err != nil {
		return 1
	}
	return 0
}

type Handler[T any] struct {
	r *Result[T]
	ch chan struct{}
}

func (h *Handler[T]) Err(err error) {
	if err != nil {
		h.r.err = err
		h.ch<-struct{}{}
		runtime.Goexit()
	}
}

func (h *Handler[T]) OK(value T) {
	h.r.value = value
}

func Try[T any](f func(h *Handler[T])) *Result[T] {
	r := &Result[T]{}
	ch := make(chan struct{})
	handler := &Handler[T]{r,ch}
	go func() {
		f(handler)
		ch <- struct{}{}
	}()
	<-ch
	return r
}

func (r *Result[T]) Unwarp() T {
	if r.t() == 0 {
		return r.value
	}
	panic(r.err)
}

func (r *Result[T]) Expect(f func(v T)) *ErrHandler {
	if r.t() == 0 {
		f(r.value)
	}
	return &ErrHandler{r.err}
}

type ErrHandler struct {
	err error
}

func (e *ErrHandler) Else(f func(err error)) {
	if e.err != nil {
		f(e.err)
	}
}
