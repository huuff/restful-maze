package maze

import (
	"github.com/samber/lo"
	"xyz.haff/maze/pkg/grid"
  "fmt"
)

type Direction byte

const (
  North Direction = iota
  East
  West
  South
)

func (d Direction) inverse() Direction {
  switch d {
    case North:
      return South
    case East:
      return West
    case South:
      return North
    case West:
      return East
  }

  panic(fmt.Sprintf("No inverse found for direction %d", d))
}

func directionBetween(p1 grid.Point, p2 grid.Point) Direction {
  switch {
  case p2.X == (p1.X - 1):
    return West
  case p2.Y == (p1.Y - 1):
    return North
  case p2.X == (p1.X + 1):
    return East
  case p2.Y == (p1.Y + 1):
    return South
  }

  panic(fmt.Sprintf("No cardinal direction between %v and %v!", p1, p2))

}

type Room struct {
  Location grid.Point
  Connections map[Direction]*Room
}

func newRoom(location grid.Point) *Room {
  return &Room {
    Location: location,
    Connections: make(map[Direction]*Room),
  }
}

func (thisRoom *Room) addConnection(otherRoom *Room) {
  direction := directionBetween(thisRoom.Location, otherRoom.Location)
  thisRoom.Connections[direction] = otherRoom
  otherRoom.Connections[direction.inverse()] = thisRoom
} 

type Maze struct {
  rooms map[grid.Point]*Room
}

func NewMaze(g grid.Grid, edges [][2]grid.Point) *Maze {
  rooms := make(map[grid.Point]*Room) 

  for x := range lo.Range(g.Width) {
    for y := range lo.Range(g.Height) {
      location := grid.Point { x, y }
      rooms[location] = newRoom(location)
    }
  }

  for _, edge := range edges {
    r1 := rooms[edge[0]]
    r2 := rooms[edge[1]]
    r1.addConnection(r2)
  }

  return &Maze { rooms: rooms }
}