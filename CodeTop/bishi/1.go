package main

import (
	"bufio"
	"fmt"
	"os"
)

type E struct {
	u, v int
}

type O struct {
	t, x int
}

type H struct {
	v, id int
	l, r  *H
}

func cmp(a, b *H) bool {
	if a.v != b.v {
		return a.v > b.v
	}
	return a.id > b.id
}

func mg(a, b *H) *H {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if !cmp(a, b) {
		a, b = b, a
	}
	a.r = mg(a.r, b)
	a.l, a.r = a.r, a.l
	return a
}

func pop(a *H) *H {
	return mg(a.l, a.r)
}

var (
	in      = bufio.NewReaderSize(os.Stdin, 1<<20)
	out     = bufio.NewWriterSize(os.Stdout, 1<<20)
	n, m, q int
	e       []E
	op      []O

	fa, sz []int
	heap   []*H
	val    []int
)

func find(x int) int {
	if fa[x] != x {
		fa[x] = find(fa[x])
	}
	return fa[x]
}

func union(x, y int) {
	x, y = find(x), find(y)
	if x == y {
		return
	}
	if sz[x] < sz[y] {
		x, y = y, x
	}
	fa[y] = x
	sz[x] += sz[y]
	heap[x] = mg(heap[x], heap[y])
}

func clean(x int) {
	x = find(x)
	for heap[x] != nil && heap[x].v != val[heap[x].id] {
		heap[x] = pop(heap[x])
	}
}

func add(x int) {
	val[x]++
	r := find(x)
	heap[r] = mg(heap[r], &H{v: val[x], id: x})
}

func main() {
	defer out.Flush()

	fmt.Fscan(in, &n, &m, &q)
	e = make([]E, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &e[i].u, &e[i].v)
	}

	op = make([]O, q+1)
	del := make([]bool, m+1)
	for i := 1; i <= q; i++ {
		fmt.Fscan(in, &op[i].t, &op[i].x)
		if op[i].t == 1 {
			del[op[i].x] = true
		}
	}

	deg := make([]int, n+1)
	for i := 1; i <= m; i++ {
		if !del[i] {
			u, v := e[i].u, e[i].v
			deg[u]++
			deg[v]++
		}
	}

	fa = make([]int, n+1)
	sz = make([]int, n+1)
	heap = make([]*H, n+1)
	val = make([]int, n+1)

	for i := 1; i <= n; i++ {
		fa[i] = i
		sz[i] = 1
		val[i] = deg[i] + i
		heap[i] = &H{v: val[i], id: i}
	}

	for i := 1; i <= m; i++ {
		if !del[i] {
			union(e[i].u, e[i].v)
		}
	}

	ans := make([]int, 0, q)
	for i := q; i >= 1; i-- {
		t, x := op[i].t, op[i].x
		if t == 2 {
			r := find(x)
			clean(r)
			ans = append(ans, heap[r].v)
		} else {
			u, v := e[x].u, e[x].v
			add(u)
			add(v)
			union(u, v)
		}
	}

	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Fprintln(out, ans[i])
	}
}
