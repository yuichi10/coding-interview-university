package uf

type unionFind struct {
	per []int
	siz []int
}

func New(n int) *unionFind {
	per := make([]int, n)
	siz := make([]int, n)
	for i := range per {
		per[i] = -1
		siz[i] = 1
	}
	return &unionFind{
		per: per,
		siz: siz,
	}
}

func (u *unionFind) Root(x int) int {
	if u.per[x] == -1 {
		return x
	}
	return u.Root(u.per[x])
}

func (u *unionFind) IsSame(x, y int) bool {
	return u.Root(x) == u.Root(y)
}

func (u *unionFind) Unite(x, y int) bool {
	x = u.Root(x)
	y = u.Root(y)

	// rootが同じならすでに統合されている。
	if x == y {
		return false
	}

	if u.siz[x] < u.siz[y] {
		a := x
		x = y
		y = a
	}

	u.per[y] = x
	u.siz[x] += u.siz[y]
	return true
}

func (u *unionFind) Size(x int) int {
	return u.siz[u.Root(x)]
}
