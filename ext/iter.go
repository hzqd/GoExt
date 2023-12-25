package ext

type Iterator[T any] interface {
	Next() Option[T]
}

func Fold[A, B any](iter Iterator[A], acc B, f func(B, A) B) B {
	for it := iter.Next(); IsSome(it); {
		acc = f(acc, UnwrapOpt(it))
	}
	return acc
}

func Count[T any](iter Iterator[T]) uint {
	return Fold(iter, 0, func(acc uint, _ T) uint {
		return acc + 1
	})
}

func Last[T any](iter Iterator[T]) Option[T] {
	return Fold(iter, None[T](), func(_ Option[T], x T) Option[T] {
		return Some(x)
	})
}

func AdvanceBy[T any](iter Iterator[T], n uint) Result[Unit, uint] {
	for i := uint(0); i < n; i++ {
		if IsNone(iter.Next()) {
			return Err[Unit](i)
		}
	}
	return Ok[Unit, uint](Unit{})
}

func Nth[T any](iter Iterator[T], n uint) Option[T] {
	if IsNone(OkToOpt(AdvanceBy(iter, n))) {
		return None[T]()
	}
	return iter.Next()
}
