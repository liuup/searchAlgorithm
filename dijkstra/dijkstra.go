package dijkstra

import (
	"container/heap"
	"fmt"
)

type pair struct {
	x, y int
}

type pairWithValue struct {
	x, y, value int
}

// 优先队列的实现
type Pqueue []pairWithValue

func (h Pqueue) Len() int           { return len(h) }
func (h Pqueue) Less(i, j int) bool { return h[i].value < h[j].value }
func (h Pqueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Pqueue) Push(x interface{}) {
	*h = append(*h, x.(pairWithValue))
}
func (h *Pqueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var (
	direction = []pair{ // 4个搜索方向
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}

	visited = [][]bool{} // 记录是否访问过

	pre = [][]pair{} // 记录前驱节点
)

func Dijkstra(board [][]string, startx, starty, endx, endy int) {
	dijkstra(board, startx, starty, endx, endy)

	showPath(pre, startx, starty, endx, endy)
}

// dijkstra实现
func dijkstra(board [][]string, startx, starty, endx, endy int) {
	// 初始化
	visited = make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	pre = make([][]pair, len(board))
	for i := 0; i < len(pre); i++ {
		pre[i] = make([]pair, len(board[0]))
	}

	// 初始化优先队列
	pq := &Pqueue{pairWithValue{startx, starty, 0}}
	heap.Init(pq)
	visited[startx][starty] = true

	for len(*pq) != 0 {
		size := len(*pq)

		// fmt.Println(pq)

		for i := 0; i < size; i++ {
			cur := (*pq)[0]
			heap.Pop(pq)

			// fmt.Println(cur)

			if cur.x == endx && cur.y == endy {
				return
			}

			// 再遍历邻居
			for _, d := range direction {
				dx, dy := cur.x+d.x, cur.y+d.y
				if !isok(board, dx, dy) { // 越界
					continue
				}
				if board[dx][dy] == "1" { // 撞墙
					continue
				}
				if visited[dx][dy] { // 已经访问
					continue
				}

				visited[dx][dy] = true
				v := heuristic(dx, dy, startx, starty) // 与起点的距离
				pre[dx][dy] = pair{cur.x, cur.y}
				heap.Push(pq, pairWithValue{dx, dy, v})
			}
		}
	}
}

func isok(board [][]string, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

// 曼哈顿距离
func heuristic(startx, starty, endx, endy int) int {
	return abs(endx-startx) + abs(endy-starty)
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func showPath(pre [][]pair, startx, starty, endx, endy int) {
	if endx == startx && endy == starty {
		fmt.Println(endx, endy)
		return
	}
	showPath(pre, startx, starty, pre[endx][endy].x, pre[endx][endy].y)
	fmt.Println(endx, endy)
}
