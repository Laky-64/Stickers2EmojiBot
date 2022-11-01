package utils

func Reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue

	for _, v := range s {
		acc = f(acc, v)
	}

	return acc
}
