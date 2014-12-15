package main

import "fmt"
import "math/rand"
import "golang.org/x/crypto/ssh/terminal"
import "time"

import "./helpers"

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

	startArr := make([]helpers.Index2D, 1, 1)
	startArr[0].X = 1
	startArr[0].Y = 2
	mazeArr[0][2] = PATH
	mazeArr[width-1][height-3] = PATH
	if !helpers.BreadthFirstSolve(&mazeArr, width-2, height-3, startArr) {
		fmt.Println("CANNOT SOLVE THIS MAZE.")
	} else {
		mazeArr[1][2] = PATH
		helpers.PrintSolutionPath(&mazeArr, width-3, height-3)
	}

	helpers.PrintMaze(mazeArr, width, height)
}

func primGenerate(maze *[][]int, x, y int) {
	mazeArr := *maze

	UNEXPLORED := 0
	EMPTY := 1

	var nodeQueue []helpers.Index2D
	nodeQueue = make([]helpers.Index2D, 0, 30)
	nodeQueue = append(nodeQueue, helpers.Index2D{x, y})

	for len(nodeQueue) > 0 {
		x = nodeQueue[0].X
		y = nodeQueue[0].Y

		nodeQueue = nodeQueue[1:]

		// Go up
		if mazeArr[x][y+2] == UNEXPLORED {
			mazeArr[x][y+2] = EMPTY
			mazeArr[x][y+1] = EMPTY
			nodeQueue = append(nodeQueue, helpers.Index2D{x, y+2})
		}
		// Go down
		if mazeArr[x][y-2] == UNEXPLORED {
			mazeArr[x][y-2] = EMPTY
			mazeArr[x][y-1] = EMPTY
			nodeQueue = append(nodeQueue, helpers.Index2D{x, y-2})
		}
		// Go left
		if mazeArr[x-2][y] == UNEXPLORED {
			mazeArr[x-2][y] = EMPTY
			mazeArr[x-1][y] = EMPTY
			nodeQueue = append(nodeQueue, helpers.Index2D{x-2, y})
		}
		// Go right
		if mazeArr[x+2][y] == UNEXPLORED {
			mazeArr[x+2][y] = EMPTY
			mazeArr[x+1][y] = EMPTY
			nodeQueue = append(nodeQueue, helpers.Index2D{x+2, y})
		}

		// Shuffle the array of nodes
		temp := make([]helpers.Index2D, len(nodeQueue))
		_ = copy(temp, nodeQueue)
		perm := rand.Perm(len(nodeQueue))
		for i, v := range perm {
			nodeQueue[v] = temp[i]
		}
	}
}
