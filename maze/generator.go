package maze

import "fmt"
import "time"
import "math/rand"

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

// Return the stack's length
func (s *Stack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

func filter(s []GameObject, fn func(GameObject) bool) []GameObject {
	var p []GameObject // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

type point struct {
	x int
	y int
}

const width int = 5
const height int = 6

func checkVisited(visited [width][height]bool, x, y int) bool {
	if x >= width || y >= height || x < 0 || y < 0 {
		return true
	}
	return visited[x][y]
}

func (game *Game) DeleteWall(x, y int, direction Direction) {
	var vertical bool
	switch direction {
	case Up:
		vertical = true
	case Left:
		vertical = false
	case Right:
		vertical = false
		x++
	case Down:
		vertical = true
		y++
	}

	var objs []GameObject // == nil
	for _, obj := range game.objects {
		wall, ok := obj.(*Wall)
		if !(ok && wall.x == x && wall.y == y && wall.vertical != vertical) {
			objs = append(objs, obj)
		}
	}
	game.objects = objs
}

func (game *Game) GenerateMaze() {
	rand.Seed(time.Now().UnixNano())
	// delete all existing walls
	game.objects = filter(game.objects, func(obj GameObject) bool {
		_, ok := obj.(*Wall)
		return !ok
	})
	currentx := 0
	currenty := 0
	visited := [width][height]bool{}
	stack := new(Stack)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			visited[x][y] = false
			game.objects = append(game.objects, MakeWall(x, y, false))
			game.objects = append(game.objects, MakeWall(x, y, true))
		}
	}
	stack.Push(point{x: currentx, y: currenty})
	for {
		visited[currentx][currenty] = true
		options := []Direction{}
		if !checkVisited(visited, currentx+1, currenty) {
			options = append(options, Right)
		}
		if !checkVisited(visited, currentx, currenty+1) {
			options = append(options, Down)
		}
		if !checkVisited(visited, currentx-1, currenty) {
			options = append(options, Left)
		}
		if !checkVisited(visited, currentx, currenty-1) {
			options = append(options, Up)
		}
		if len(options) > 0 {
			chosenDirection := options[rand.Intn(len(options))]
			game.DeleteWall(currentx, currenty, chosenDirection)
			switch chosenDirection {
			case Up:
				currenty -= 1
			case Down:
				currenty += 1
			case Left:
				currentx -= 1
			case Right:
				currentx += 1
			}
			stack.Push(point{x: currentx, y: currenty})
		} else if stack.Len() > 0 {
			val := stack.Pop()
			p := val.(point)
			currentx = p.x
			currenty = p.y
		} else {
			break
		}
	}
}
