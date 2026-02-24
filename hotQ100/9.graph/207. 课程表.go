package graph

func canFinish1(numCourses int, prerequisites [][]int) bool {
	inDegree := make([]int, numCourses)
	g := make([][]int, numCourses)
	for _, row := range prerequisites {
		g[row[1]] = append(g[row[1]], row[0])
		inDegree[row[0]]++
	}

	// 将入度为0的加入队列
	q := []int{}
	for i, d := range inDegree {
		if d == 0 {
			q = append(q, i)
		}
	}

	cnt := 0
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		cnt++
		for _, j := range g[i] {
			inDegree[j]--
			if inDegree[j] == 0 {
				q = append(q, j)
			}
		}
	}
	return cnt == numCourses
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree := make([]int, numCourses)
	g := make([][]int, numCourses) // 记录课程之间的指向关系
	for _, row := range prerequisites {
		inDegree[row[0]]++
		g[row[1]] = append(g[row[1]], row[0])
	}

	q := []int{}
	for i, d := range inDegree {
		if d == 0 {
			q = append(q, i)
		}
	}

	cnt := 0
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		cnt++
		for _, c := range g[i] {
			inDegree[c]--
			if inDegree[c] == 0 {
				q = append(q, c)
			}
		}
	}
	return cnt == numCourses
}
