package ext

func Pipe[T, R any](x *T, f func(*T) R) R {
	return f(x)
}

func Tap[T any](x *T, f func(*T)) *T {
	f(x)
	return x
}
