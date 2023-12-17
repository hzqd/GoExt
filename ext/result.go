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

type Result interface {
	implResultForOkAndErr()
}

func (_ ok[T]) implResultForOkAndErr()  {}
func (_ err[E]) implResultForOkAndErr() {}

/////////////////////////////////////////////////////////////////////////
// Wrap the values
/////////////////////////////////////////////////////////////////////////

func Ok[T any](t T) Result {
	return ok[T]{t}
}

func Err[E any](e E) Result {
	return err[E]{e}
}

/////////////////////////////////////////////////////////////////////////
// Querying the contained values
/////////////////////////////////////////////////////////////////////////

func IsOk[A any](r Result) bool {
	var res bool
	switch any(r).(type) {
	case ok[A]:
		res = true
	case err[A]:
		res = false
	}
	return res
}

func IsErr[A any](r Result) bool {
	return !IsOk[A](r)
}

func IsOkAnd[A any](r Result, f func(A) bool) bool {
	var res bool
	switch x := any(r).(type) {
	case ok[A]:
		res = f(x.t)
	case err[A]:
		res = false
	}
	return res
}

func IsErrAnd[A any](r Result, f func(A) bool) bool {
	var res bool
	switch x := any(r).(type) {
	case err[A]:
		res = f(x.e)
	case ok[A]:
		res = false
	}
	return res
}

/////////////////////////////////////////////////////////////////////////
// Adapter for each variant
/////////////////////////////////////////////////////////////////////////

func OkToOption[A any](r Result) Option {
	var res Option
	switch x := any(r).(type) {
	case ok[A]:
		res = Some(x.t)
	case err[A]:
		res = None()
	}
	return res
}

func ErrToOption[A any](r Result) Option {
	var res Option
	switch x := any(r).(type) {
	case ok[A]:
		res = None()
	case err[A]:
		res = Some(x.e)
	}
	return res
}

/////////////////////////////////////////////////////////////////////////
// Transforming contained values
/////////////////////////////////////////////////////////////////////////

func MapOk[A, U any](r Result, f func(A) U) Result {
	var res Result
	switch x := any(r).(type) {
	case ok[A]:
		res = Ok(f(x.t))
	case err[A]:
		res = x
	}
	return res
}

func MapErr[A, U any](r Result, f func(A) U) Result {
	var res Result
	switch x := any(r).(type) {
	case ok[A]:
		res = x
	case err[A]:
		res = Err(f(x.e))
	}
	return res
}

func MapOkOrElse[T, U any](r Result, fail func(T) U, succ func(T) U) Result {
	var res Result
	switch x := any(r).(type) {
	case ok[T]:
		res = Ok(succ(x.t))
	case err[T]:
		res = Err(fail(x.e))
	}
	return res
}

func UnwrapOk[T any](r Result) T {
	var res T
	switch x := any(r).(type) {
	case ok[T]:
		res = x.t
	}
	return res
}

func UnwrapErr[E any](r Result) E {
	var res E
	switch x := any(r).(type) {
	case err[E]:
		res = x.e
	}
	return res
}
