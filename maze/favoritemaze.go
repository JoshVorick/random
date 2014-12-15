package main

import "fmt"
import "math/rand"
import "golang.org/x/crypto/ssh/terminal"
import "time"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var height, width int
	var err error
	height, width, err = terminal.GetSize(0)
	fmt.Printf("%d %d\n", height, width)
	if width % 2 == 0 {
		width++
	}
	height = (height + 3) / 2
	if height % 2 == 0 {
		height++
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d %d\n", height, width)

	UNEXPLORED := 0 //Still need to figure out if its a wall or not
	EMPTY := 1 // white space
	WALL := 2  // wall
	PATH := 3 // correct path

	mazeArr := make([][]int, width, width)
	for i:=0; i<width; i++ {
		mazeArr[i] = make([]int, height, height)
	}

	// Make walls around the edges
	for i:=0; i<width; i++ {
		mazeArr[i][0] = WALL
		mazeArr[i][1] = WALL
		mazeArr[i][height-1] = WALL
		mazeArr[i][height-2] = WALL
	}
	for i:=0; i<height; i++ {
		mazeArr[0][i] = WALL
		mazeArr[1][i] = WALL
		mazeArr[width-1][i] = WALL
		mazeArr[width-2][i] = WALL
	}

	// Generate maze starting at (2,2)
	mazeArr[0][2] = EMPTY
	mazeArr[1][2] = EMPTY
	mazeArr[2][2] = EMPTY
	breadthFirstGenerate(&mazeArr, 2, 2)

	// Make an exit
	for i:=width-1; i>1; i-- {
		mazeArr[i][height-3] = EMPTY
		if mazeArr[i-1][height-3] == EMPTY || mazeArr[i][height-4] == EMPTY{
			break
		}
	}

	mazeArr[0][2] = PATH
	mazeArr[width-1][height-3] = PATH
	if !depthFirstSolve(&mazeArr, 1, 2, width-2, height-3) {
		fmt.Println("CANNOT SOLVE. THIS MAZE IS SHITTY!!!!")
	}

	// Print maze
	for i:=1; i<width-1; i++ {
		for j:=1; j<height-1; j++ {
			if mazeArr[i][j] == UNEXPLORED || mazeArr[i][j] == WALL {
				fmt.Print("OO")
			} else if mazeArr[i][j] == PATH {
				fmt.Print("..")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

type Index2D struct {
	x int
	y int
}

func (i Index2D) Strin() string {
	return fmt.Sprint(i.x, i.y)
}

func breadthFirstGenerate(maze *[][]int, x, y int) {
	mazeArr := *maze

	UNEXPLORED := 0
	EMPTY := 1
	//WALL := 2

	var nodeQueue []Index2D
	nodeQueue = make([]Index2D, 0, 30)
	nodeQueue = append(nodeQueue, Index2D{x, y})

	for len(nodeQueue) > 0 {
		x = nodeQueue[0].x
		y = nodeQueue[0].y

		nodeQueue = nodeQueue[1:]

		// Go up
		if mazeArr[x][y+2] == UNEXPLORED {
			mazeArr[x][y+2] = EMPTY
			mazeArr[x][y+1] = EMPTY
			//mazeArr[x-1][y+1] = WALL
			//mazeArr[x+1][y+1] = WALL
			nodeQueue = append(nodeQueue, Index2D{x, y+2})
		}
		// Go down
		if mazeArr[x][y-2] == UNEXPLORED {
			mazeArr[x][y-2] = EMPTY
			mazeArr[x][y-1] = EMPTY
			//mazeArr[x-1][y-1] = WALL
			//mazeArr[x+1][y-1] = WALL
			nodeQueue = append(nodeQueue, Index2D{x, y-2})
		}
		// Go left
		if mazeArr[x-2][y] == UNEXPLORED {
			mazeArr[x-2][y] = EMPTY
			mazeArr[x-1][y] = EMPTY
			//mazeArr[x-1][y+1] = WALL
			//mazeArr[x-1][y-1] = WALL
			nodeQueue = append(nodeQueue, Index2D{x-2, y})
		}
		// Go right
		if mazeArr[x+2][y] == UNEXPLORED {
			mazeArr[x+2][y] = EMPTY
			mazeArr[x+1][y] = EMPTY
			//mazeArr[x+1][y-1] = WALL
			//mazeArr[x+1][y+1] = WALL
			nodeQueue = append(nodeQueue, Index2D{x+2, y})
		}

		// Shuffle the array of nodes
		temp := make([]Index2D, len(nodeQueue))
		_ = copy(temp, nodeQueue)
		perm := rand.Perm(len(nodeQueue))
		fmt.Println(len(nodeQueue), perm)
		for i, v := range perm {
			nodeQueue[v] = temp[i]
		}
	}
}

func depthFirstSolve(maze *[][]int, x, y, targetX, targetY int) bool {
	mazeArr := *maze
	EMPTY := 1
	PATH := 3

	if x == targetX && y == targetY {
		mazeArr[x][y] = PATH
		return true
	}

	mazeArr[x][y] = PATH

	out := false
	if mazeArr[x-1][y] == EMPTY {
		out = depthFirstSolve(maze, x-1, y, targetX, targetY)
		if out {
			return true
		}
	}
	if mazeArr[x+1][y] == EMPTY {
		out = depthFirstSolve(maze, x+1, y, targetX, targetY)
		if out {
			return true
		}
	}
	if mazeArr[x][y-1] == EMPTY {
		out = depthFirstSolve(maze, x, y-1, targetX, targetY)
		if out {
			return true
		}
	}
	if mazeArr[x][y+1] == EMPTY {
		out = depthFirstSolve(maze, x, y+1, targetX, targetY)
		if out {
			return true
		}
	}
	mazeArr[x][y] = 9
	return false
}
