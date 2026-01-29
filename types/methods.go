package types

type rect struct {
	width, height int
}

func area(r *rect) int {
	return r.width * r.height
}
