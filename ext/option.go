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

type Option[T any] interface {
	implOptionForSomeAndNone()
}

func (_ some[T]) implOptionForSomeAndNone() {}
func (_ none) implOptionForSomeAndNone()    {}

/////////////////////////////////////////////////////////////////////////
// Wrap the values
/////////////////////////////////////////////////////////////////////////

func Some[T any](t T) Option[T] {
	return some[T]{t}
}

func None[T any]() Option[T] {
	return none{}
}

/////////////////////////////////////////////////////////////////////////
// Querying the contained values
/////////////////////////////////////////////////////////////////////////

func IsSome[T any](o Option[T]) bool {
	var res bool
	switch o.(type) {
	case some[T]:
		res = true
	case none:
		res = false
	}
	return res
}

func IsNone[T any](o Option[T]) bool {
	return !IsSome[T](o)
}

func IsSomeAnd[T any](o Option[T], f func(T) bool) bool {
	var res bool
	switch x := o.(type) {
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

func UnwrapOpt[T any](o Option[T]) T {
	var res T
	switch x := o.(type) {
	case some[T]:
		res = x.t
	}
	return res
}

func UnwrapOptOr[T any](o Option[T], v T) T {
	var res T
	switch x := o.(type) {
	case some[T]:
		res = x.t
	case none:
		res = v
	}
	return res
}

func UnwrapOptOrElse[T any](o Option[T], f func() T) T {
	var res T
	switch x := o.(type) {
	case some[T]:
		res = x.t
	case none:
		res = f()
	}
	return res
}

func UnwrapOptOrDefault[T any](o Option[T]) T {
	var res T
	switch x := o.(type) {
	case some[T]:
		res = x.t
	case none:
		res = *new(T)
	}
	return res
}

func MapSome[T, U any](o Option[T], fn func(T) U) Option[U] {
	var res Option[U]
	switch x := o.(type) {
	case some[T]:
		res = Some(fn(x.t))
	case none:
		res = None[U]()
	}
	return res
}

func OkOr[T, E any](o Option[T], fail E) Result[T, E] {
	var res Result[T, E]
	switch x := o.(type) {
	case some[T]:
		res = Ok[T, E](x.t)
	case none:
		res = Err[T, E](fail)
	}
	return res
}

func OkOrElse[T, E any](o Option[T], fail func() E) Result[T, E] {
	var res Result[T, E]
	switch x := o.(type) {
	case some[T]:
		res = Ok[T, E](x.t)
	case none:
		res = Err[T, E](fail())
	}
	return res
}
