package helpers

import "fmt"

type Index2D struct {
	X int
	Y int
}

func PrintMaze(mazeArr [][]int, width, height int) {
	UNEXPLORED := 0
	WALL := 2
	PATH := 3
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

func BreadthFirstSolve(maze *[][]int, targetX, targetY int, nodeQueue []Index2D) bool {
  mazeArr := *maze
  if len(nodeQueue) < 1 {
    return false
  }
  x := nodeQueue[0].X
  y := nodeQueue[0].Y
  nodeQueue = nodeQueue[1:]
  EMPTY := 1
  PATH := 3
  N := 10
  S := 11
  E := 12
  W := 13

  if x == targetX && y == targetY {
    mazeArr[x][y] = PATH
    return true
  }

  if mazeArr[x-1][y] == EMPTY {
    mazeArr[x-1][y] = E
    nodeQueue = append(nodeQueue, Index2D{x-1, y})
  }
  if mazeArr[x+1][y] == EMPTY {
    mazeArr[x+1][y] = W
    nodeQueue = append(nodeQueue, Index2D{x+1, y})
  }
  if mazeArr[x][y-1] == EMPTY {
    mazeArr[x][y-1] = N
    nodeQueue = append(nodeQueue, Index2D{x, y-1})
  }
  if mazeArr[x][y+1] == EMPTY {
    mazeArr[x][y+1] = S
    nodeQueue = append(nodeQueue, Index2D{x, y+1})
  }
  out := BreadthFirstSolve(maze, targetX, targetY, nodeQueue)
  return out
}

// Start at the end and work backwards
func PrintSolutionPath(maze *[][]int, x, y int) {
  mazeArr := *maze
  PATH := 3
  N := 10
  S := 11
  E := 12
  W := 13
  switch mazeArr[x][y] {
  case N:
    PrintSolutionPath(maze, x, y+1)
  case S:
    PrintSolutionPath(maze, x, y-1)
  case E:
    PrintSolutionPath(maze, x+1, y)
  case W:
    PrintSolutionPath(maze, x-1, y)
  }
  mazeArr[x][y] = PATH
}
