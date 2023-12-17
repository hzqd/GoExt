package ext

/////////////////////////////////////////////////////////////////////////////
// The `some` and `none` struct for `Option` type.
/////////////////////////////////////////////////////////////////////////////

type some[T any] struct {
	t T
}

type none struct{}

/////////////////////////////////////////////////////////////////////////////
// The `Option` type.
/////////////////////////////////////////////////////////////////////////////

type Option interface {
	implOptionForSomeAndNone()
}

func (_ some[T]) implOptionForSomeAndNone() {}
func (_ none) implOptionForSomeAndNone()    {}

/////////////////////////////////////////////////////////////////////////
// Wrap the values
/////////////////////////////////////////////////////////////////////////

func Some[T any](t T) Option {
	return some[T]{t}
}

func None() Option {
	return none{}
}

/////////////////////////////////////////////////////////////////////////
// Querying the contained values
/////////////////////////////////////////////////////////////////////////

func IsSome[T any](o Option) bool {
	var res bool
	switch any(o).(type) {
	case some[T]:
		res = true
	case none:
		res = false
	}
	return res
}

func IsNone[T any](o Option) bool {
	return !IsSome[T](o)
}

func IsSomeAnd[T any](o Option, f func(T) bool) bool {
	var res bool
	switch x := any(o).(type) {
	case some[T]:
		res = f(x.t)
	case none:
		res = false
	}
	return res
}

/////////////////////////////////////////////////////////////////////////
// Getting to contained values
/////////////////////////////////////////////////////////////////////////

func UnwrapOption[T any](o Option) T {
	var res T
	switch x := any(o).(type) {
	case some[T]:
		res = x.t
	}
	return res
}

func MapSome[T, U any](o Option, fn func(T) U) Option {
	var res Option
	switch x := any(o).(type) {
	case some[T]:
		res = Some(fn(x.t))
	case none:
		res = None()
	}
	return res
}

func OkOr[T any](o Option, fail T) Result {
	var res Result
	switch x := any(o).(type) {
	case some[T]:
		res = Ok(x.t)
	case none:
		res = Err(fail)
	}
	return res
}

func OkOrElse[T any](o Option, fail func() T) Result {
	var res Result
	switch x := any(o).(type) {
	case some[T]:
		res = Ok(x.t)
	case none:
		res = Err(fail())
	}
	return res
}
