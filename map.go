package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type connection struct {
	distance int
	node     *node
}

type node struct {
	x           int
	y           int
	connections [4]*connection
}

type level struct {
	w     int
	h     int
	nodes *map[[2]int]*node
}

func (l *level) getNode(x, y int) *node {
	return (*l.nodes)[[2]int{x, y}]
}

func distance(x1, y1, x2, y2 int) int {
	dx := x2 - x1
	if dx < 0 {
		dx = -dx
	}
	dy := y2 - y1
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

func connnectNodesLR(n1, n2 *node) {
	if n1 != nil {
		dist := distance(n1.x, n1.y, n2.x, n2.y)
		n1.connections[1] = &connection{dist, n2}
		n2.connections[3] = &connection{dist, n1}
	}
}

func connectNodesTB(n1, n2 *node) {
	if n1 != nil {
		dist := distance(n1.x, n1.y, n2.x, n2.y)
		n1.connections[2] = &connection{dist, n2}
		n2.connections[0] = &connection{dist, n1}
	}
}

func newLevel(path string) *level {
	// Read map file
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("could not read %s: %v", path, err))
	}

	// Get text, width and height
	text := string(bytes)
	w := strings.Index(text, "\n")
	h := strings.Count(text, "\n")
	text = strings.Replace(text, "\n", "", -1)

	// Initialize containers
	nodes := make(map[[2]int]*node)
	topNodes := make([]*node, w)

	for y := 1; y < h-1; y++ {
		prv := false
		cur := text[y*w] == 48
		nxt := text[y*w+1] == 48

		var leftNode *node

		for x := 1; x < w-1; x++ {
			prv = cur
			cur = nxt
			nxt = text[y*w+x+1] == 48

			if !cur {
				continue
			}

			var n *node

			if prv && nxt && x == 1 {
				n = &node{0, y, [4]*connection{}}
				leftNode = n
			}

			if prv && nxt && x == w-2 {
				n = &node{w - 1, y, [4]*connection{}}
				connnectNodesLR(leftNode, n)
				leftNode = nil
			}

			if prv {
				if nxt { // 000
					if text[(y-1)*w+x] == 48 || text[(y+1)*w+x] == 48 {
						n = &node{x, y, [4]*connection{}}
						connnectNodesLR(leftNode, n)
						leftNode = n
					}
				} else { // 001
					n = &node{x, y, [4]*connection{}}
					connnectNodesLR(leftNode, n)
					leftNode = nil
				}
			} else {
				if nxt { // 100
					n = &node{x, y, [4]*connection{}}
					leftNode = n
				} else { // 101
					if text[(y-1)*w+x] != 48 || text[(y+1)*w+x] != 48 {
						n = &node{x, y, [4]*connection{}}
					}
				}
			}

			if n != nil {
				nodes[[2]int{n.x, n.y}] = n

				if text[(y-1)*w+x] == 48 {
					connectNodesTB(topNodes[x], n)
				}

				if text[(y+1)*w+x] == 48 {
					topNodes[x] = n
				} else {
					topNodes[x] = nil
				}
			}
		}
	}

	for y := 1; y < h-1; y++ {
		n1 := nodes[[2]int{0, y}]
		n2 := nodes[[2]int{w - 1, y}]
		if n1 != nil && n2 != nil {
			n1.connections[3] = &connection{0, n2}
			n2.connections[1] = &connection{0, n1}
		}
	}

	return &level{w, h, &nodes}
}
