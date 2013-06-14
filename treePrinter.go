package algorithm

import (
	"fmt"
	"math"
	"reflect"
)

func PrintNode(nodes []Node, level int, maxLevel int) {
	if len(nodes) == 0 || isAllElementsNull(nodes) {
		return
	}
	var floor int = maxLevel - level
	var endgeLines = int(math.Pow(2.0, float64(floor-1)))
	var firstSpaces = int(math.Pow(2.0, float64(floor)) - 1)
	var betweenSpaces = int(math.Pow(2.0, float64(floor+1)) - 1)
	printSpace(firstSpaces)
	var newNodes []Node
	for _, node := range nodes {
		if node == nil || reflect.ValueOf(node).IsNil() {
			newNodes = append(newNodes, nil)
			newNodes = append(newNodes, nil)
			fmt.Print(" ")
		} else {
			newNodes = append(newNodes, node.Left())
			newNodes = append(newNodes, node.Right())
			fmt.Print(node.Value())
		}
		printSpace(betweenSpaces)
	}
	fmt.Println("")

	for i := 1; i <= endgeLines; i++ {
		for j := 0; j < len(nodes); j++ {
			printSpace(firstSpaces - i)
			if nodes[j] == nil || reflect.ValueOf(nodes[j]).IsNil() {
				printSpace(endgeLines + endgeLines + i + 1)
				continue
			}
			if nodes[j].Left() != nil && !reflect.ValueOf(nodes[j].Left()).IsNil() {
				fmt.Print("/")
			} else {
				printSpace(1)
			}
			printSpace(i + i - 1)
			if nodes[j].Right() != nil && !reflect.ValueOf(nodes[j].Right()).IsNil() {
				fmt.Print("\\")
			} else {
				printSpace(1)
			}
			printSpace(endgeLines + endgeLines - i)
		}
		fmt.Println("")
	}
	PrintNode(newNodes, level+1, maxLevel)
}

func maxLevel(tn Node) int {
	if tn == nil || reflect.ValueOf(tn).IsNil() {
		return 0
	}
	leftLevel := maxLevel(tn.Left())
	rightLevel := maxLevel(tn.Right())
	if leftLevel > rightLevel {
		return leftLevel + 1
	}
	return rightLevel + 1
}

func printSpace(count int) {
	for i := 0; i < count; i++ {
		fmt.Print(" ")
	}
}

func isAllElementsNull(nodes []Node) bool {
	for _, v := range nodes {
		if v != nil && !reflect.ValueOf(v).IsNil() {
			return false
		}
	}
	return true
}
