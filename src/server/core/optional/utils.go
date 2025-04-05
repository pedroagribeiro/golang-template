package optional

func OptionalEnumeratedToOptionalInt[T optionable](src Optional[T]) (dst Optional[Int]) {
	if src.Present() {
		if value, ok := any(src.OrZero()).(Int); ok {
			dst = New(Int(value))
		} else {
			dst = NewUndefined[Int]()
		}
		return
	}
	dst = NewUndefined[Int]()
	return
}
