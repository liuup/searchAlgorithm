package bfs

import (
	"fmt"
)

type pair struct {
	x, y int
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

// bfs实现
func BFS(board [][]string, startx, starty, endx, endy int) {
	bfs(board, startx, starty, endx, endy)

	showPath(pre, startx, starty, endx, endy)
}

func bfs(board [][]string, startx, starty, endx, endy int) {
	// 初始化
	visited = make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	pre = make([][]pair, len(board))
	for i := 0; i < len(pre); i++ {
		pre[i] = make([]pair, len(board[0]))
	}

	q := []pair{{startx, starty}}
	visited[startx][starty] = true

	for len(q) != 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]

			if cur.x == endx && cur.y == endy {
				return
			}

			for _, d := range direction {
				dx, dy := cur.x+d.x, cur.y+d.y
				if !isok(board, dx, dy) { // 越界
					continue
				}
				if board[dx][dy] == "1" { // 撞墙
					continue
				}
				if visited[dx][dy] { // 已访问
					continue
				}

				visited[dx][dy] = true
				pre[dx][dy] = pair{cur.x, cur.y}
				q = append(q, pair{dx, dy})
			}
		}
	}
}

func isok(board [][]string, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func showPath(pre [][]pair, startx, starty, endx, endy int) {
	if endx == startx && endy == starty {
		fmt.Println(endx, endy)
		return
	}
	showPath(pre, startx, starty, pre[endx][endy].x, pre[endx][endy].y)
	fmt.Println(endx, endy)
}
