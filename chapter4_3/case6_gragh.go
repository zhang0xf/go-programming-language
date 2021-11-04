package chapter4_3

// Map的value类型也可以是一个聚合类型，比如是一个map或slice。

// 有向图!

var graph = make(map[string]map[string]bool)

// ddEdge函数惰性初始化map是一个惯用方式，也就是说在每个值首次作为key时才初始化。
func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
