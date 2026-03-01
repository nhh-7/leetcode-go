package graph

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)
	for i, e := range equations {
		u, v := e[0], e[1]
		if graph[u] == nil {
			graph[u] = make(map[string]float64)
		}
		if graph[v] == nil {
			graph[v] = make(map[string]float64)
		}
		graph[u][v] = values[i]
		graph[v][u] = 1.0 / values[i]
	}

	var dfs func(curr, target string, visited map[string]bool) float64
	dfs = func(curr, target string, visited map[string]bool) float64 {
		if _, ok := graph[curr]; !ok {
			return -1.0
		}
		if curr == target {
			return 1.0
		}
		visited[curr] = true
		for nxt, w := range graph[curr] {
			if !visited[nxt] {
				res := dfs(nxt, target, visited)
				if res != -1.0 {
					return res * w
				}
			}
		}
		return -1.0
	}

	ans := make([]float64, len(queries))
	for i, q := range queries {
		if _, ok := graph[q[0]]; !ok {
			ans[i] = -1.0
		} else if _, ok := graph[q[1]]; !ok {
			ans[i] = -1.0
		} else {
			ans[i] = dfs(q[0], q[1], make(map[string]bool))
		}
	}
	return ans
}
