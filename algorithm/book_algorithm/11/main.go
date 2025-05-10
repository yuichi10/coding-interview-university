package main

import "fmt"

type UnionFind struct {
	par []int
	siz []int
}

func NewUnionFind(n int) UnionFind {
	par := make([]int, n)
	siz := make([]int, n)
	for i := range par {
		par[i] = -1
		siz[i] = 1
	}
	return UnionFind{
		par,
		siz,
	}
}

func (u *UnionFind) Root(x int) int {
	if u.par[x] == -1 {
		return x
	}
	u.par[x] = u.Root(u.par[x])
	return u.par[x]
}

func (u *UnionFind) IsSame(x, y int) bool {
	return u.Root(x) == u.Root(y)
}

func (u *UnionFind) Unite(x, y int) bool {
	x = u.Root(x)
	y = u.Root(y)

	if x == y {
		return false
	}

	if u.siz[x] < u.siz[y] {
		a := x
		x = y
		y = a
	}
	u.par[y] = x
	u.siz[x] += u.siz[y]

	return true
}

func (u *UnionFind) Size(x int) int {
	return u.siz[u.Root(x)]
}

func main() {

	uf := NewUnionFind(7)
	fmt.Printf("%+v\n", uf)

	uf.Unite(2, 3)
	fmt.Printf("%+v\n", uf)

	uf.Unite(5, 6)
	fmt.Printf("%+v\n", uf)

	uf.Unite(1, 2)
	fmt.Printf("%+v\n", uf)

	fmt.Println(uf.Size(1))
	fmt.Println(uf.Size(2))
	fmt.Println(uf.Size(3))

}
