package ext

/////////////////////////////////////////////////////////////////////////////
// Error handling with the `ok` and `err` struct for `Result` type.
/////////////////////////////////////////////////////////////////////////////

type ok[T any] struct {
	t T
}

type err[E any] struct {
	e E
}

/////////////////////////////////////////////////////////////////////////////
// Error handling with the `Result` type.
/////////////////////////////////////////////////////////////////////////////

type Result[T, E any] interface {
	implResultForOkAndErr()
}

func (_ ok[T]) implResultForOkAndErr()  {}
func (_ err[E]) implResultForOkAndErr() {}

/////////////////////////////////////////////////////////////////////////
// Wrap the values
/////////////////////////////////////////////////////////////////////////

func Ok[T, E any](t T) Result[T, E] {
	return ok[T]{t}
}

func Err[T, E any](e E) Result[T, E] {
	return err[E]{e}
}

/////////////////////////////////////////////////////////////////////////
// Querying the contained values
/////////////////////////////////////////////////////////////////////////

func IsOk[T, E any](r Result[T, E]) bool {
	var res bool
	switch r.(type) {
	case ok[T]:
		res = true
	case err[E]:
		res = false
	}
	return res
}

func IsErr[T, E any](r Result[T, E]) bool {
	return !IsOk[T](r)
}

func IsOkAnd[T, E any](r Result[T, E], f func(T) bool) bool {
	var res bool
	switch x := r.(type) {
	case ok[T]:
		res = f(x.t)
	case err[E]:
		res = false
	}
	return res
}

func IsErrAnd[T, E any](r Result[T, E], f func(E) bool) bool {
	var res bool
	switch x := r.(type) {
	case ok[T]:
		res = false
	case err[E]:
		res = f(x.e)
	}
	return res
}

/////////////////////////////////////////////////////////////////////////
// Adapter for each variant
/////////////////////////////////////////////////////////////////////////

func OkToOpt[T, E any](r Result[T, E]) Option[T] {
	var res Option[T]
	switch x := r.(type) {
	case ok[T]:
		res = Some(x.t)
	case err[E]:
		res = None[T]()
	}
	return res
}

func ErrToOpt[T, E any](r Result[T, E]) Option[E] {
	var res Option[E]
	switch x := r.(type) {
	case ok[T]:
		res = None[E]()
	case err[E]:
		res = Some(x.e)
	}
	return res
}

/////////////////////////////////////////////////////////////////////////
// Transforming contained values
/////////////////////////////////////////////////////////////////////////

func MapOk[T, U, E any](r Result[T, E], f func(T) U) Result[U, E] {
	var res Result[U, E]
	switch x := r.(type) {
	case ok[T]:
		res = Ok[U, E](f(x.t))
	case err[E]:
		res = x
	}
	return res
}

func MapErr[T, E, F any](r Result[T, E], f func(E) F) Result[T, F] {
	var res Result[T, F]
	switch x := r.(type) {
	case ok[T]:
		res = x
	case err[E]:
		res = Err[T, F](f(x.e))
	}
	return res
}

func MapOkOrElse[T, E, U, F any](r Result[T, E], fail func(E) F, succ func(T) U) Result[U, F] {
	var res Result[U, F]
	switch x := r.(type) {
	case ok[T]:
		res = Ok[U, F](succ(x.t))
	case err[E]:
		res = Err[U, F](fail(x.e))
	}
	return res
}

func UnwrapOk[T, E any](r Result[T, E]) T {
	var res T
	switch x := r.(type) {
	case ok[T]:
		res = x.t
	}
	return res
}

func UnwrapErr[T, E any](r Result[T, E]) E {
	var res E
	switch x := r.(type) {
	case err[E]:
		res = x.e
	}
	return res
}

func UnwrapOkOr[T, E any](r Result[T, E], v T) T {
	var res T
	switch x := r.(type) {
	case ok[T]:
		res = x.t
	case err[E]:
		res = v
	}
	return res
}

func UnwrapErrOr[T, E any](r Result[T, E], v E) E {
	var res E
	switch x := r.(type) {
	case ok[T]:
		res = v
	case err[E]:
		res = x.e
	}
	return res
}

func UnwrapOkOrElse[T, E any](r Result[T, E], f func() T) T {
	var res T
	switch x := r.(type) {
	case ok[T]:
		res = x.t
	case err[E]:
		res = f()
	}
	return res
}

func UnwrapErrOrElse[T, E any](r Result[T, E], f func() E) E {
	var res E
	switch x := r.(type) {
	case ok[T]:
		res = f()
	case err[E]:
		res = x.e
	}
	return res
}

func UnwrapOkOrDefault[T, E any](r Result[T, E]) T {
	var res T
	switch x := r.(type) {
	case ok[T]:
		res = x.t
	case err[E]:
		res = *new(T)
	}
	return res
}

func UnwrapErrOrDefault[T, E any](r Result[T, E]) E {
	var res E
	switch x := r.(type) {
	case ok[T]:
		res = *new(E)
	case err[E]:
		res = x.e
	}
	return res
}
