package main

import (
	"fmt"
	"os"
)

func main() {

	// 注意读取的maze.in文件是LF换行，否则读取时会将\r\n转换为0 \n
	maze := readMaze("maze/maze.in")

	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

	fmt.Println("--------------")

	start := point{0, 0}
	end := point{len(maze) - 1, len(maze[0]) - 1}
	steps := walk(maze, start, end)

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}


	fmt.Println("--------------")
	// 输出多少步可以走出，以及路线
	routes := route(steps, start, end)
	for i:= len(routes) - 1;i>=0;i--{
		fmt.Println(routes[i])
	}

}
func route(steps [][]int, start point, end point) []point {
	cur := end
	routes := []point{cur}

	for {
		nextValue := cur.valueAt(steps) - 1
		// 从周边的点遍历
		for _, dir := range dirs {
			next := cur.add(dir)
			val, ok := next.at(steps)
			if !ok{
				continue
			}
			if val == nextValue {
				routes = append(routes, next)
				cur = next
				break
			}
		}
		if cur.equals(start) {
			break
		}
	}

	return routes
}

func readMaze(filename string) [][]int {

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var row, col int
	// 读取行和列
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row) // 行
	// 列初始化
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

// 读取当前节点的4个方向位置
var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}, // 上左下右
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

// 表示p是否在grid中
// 第二个参数表示有没有值
func (p point) at(grid [][]int) (int, bool) {

	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func (p point) valueAt(grid [][]int) int {
	val, ok := p.at(grid)
	if !ok {
		panic("无效的点" + string(p.i) + " " + string(p.j))
	}
	return val
}

func (p point) equals(d point) bool {
	return p.i == d.i && p.j == d.j
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze)) // 用于填写值,走过的路径
	// 对step进行初始化，全部都是0
	// 除了起点的0，其他的0表示无效
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	// 将起点放入队列中
	Q := []point{start}

	var cur point

	for len(Q) > 0 { // 如果Q不为空，表示还有点需要探索
		cur = Q[0]
		Q = Q[1:]

		if cur.equals(end) { // 发现终点
			break
		}

		for _, dir := range dirs {
			//next := cur + dir // 没有操作符重载，需要自定义方法
			next := cur.add(dir)
			// 探索下一个节点 maze中下一个节点是0 and steps at next is 0 and next != start
			// maze 中节点为0表示可以走，steps中节点为0表示没有走过，可以走
			// next节点也不能等于start
			val, ok := next.at(maze)
			if !ok || val == 1 {
				// 说明该点不能探索
				continue
			}
			val, ok = next.at(steps)
			if !ok || val != 0 {
				// 说明走过了
				continue
			}
			if next.equals(start) {
				// 到原点了
				continue
			}
			// 获取当前在steps中的值+1
			steps[next.i][next.j] = cur.valueAt(steps) + 1
			// 将next添加到队列中
			Q = append(Q, next)
		}
	}

	if cur.equals(end) {
		fmt.Println("需要走", cur.valueAt(steps), "步到达终点")
	} else {
		fmt.Println("没有走到终点")
	}

	return steps
}
