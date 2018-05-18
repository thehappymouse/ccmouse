package main

// 迷宫 广度优先算法
import (
	"os"
	"fmt"
)

// 从文件读取迷宫数据
func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)

	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	i, j int
}

// 探索一个方向点
func (p point) Add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

// 取一点位，在 grid ( 内的值
func (p point) At(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[0]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

// 代表四个方向
var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}
// 走迷宫，从start开始， 向四个方向探索，走到结束位置
// 广度优先，退出条件：1，没有路了。 2，回到原点了。3，到达终点。（如果有多条路到终点呢）
func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for row := range steps {
		steps[row] = make([]int, len(maze[0]))
	}

	Q := []point{start}
	for  len(Q) > 0 {

		curr := Q[0]
		Q = Q[1:]
		if curr == end {
			break
		}

		for _, dir := range dirs {
			next := curr.Add(dir)
			// 迷宫中，1表示墙，不能继续
			val, ok := next.At(maze)
			if !ok || val == 1 {
				continue
			}
			// 已走过的列表中，如果不是1，表示走过了，跳过
			val, ok = next.At(steps)
			if !ok || val != 0 {
				continue
			}
			if next == start {
				continue
			}

			setVal, _ := curr.At(steps)
			steps[next.i][next.j] = setVal + 1
			Q = append(Q, next)
		}

	}
	return steps
}
func main() {
	maze := readMaze("maze/maze.in")
	steps := walk(maze, point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
