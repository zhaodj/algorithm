package algorithm

import (
	"fmt"
	"math"
)

type Comparable interface {
	CompareTo(Comparable) int
}

type Tree struct {
	root *TreeNode
}

type TreeNode struct {
	left, right *TreeNode
	value       Comparable
}

//添加节点
func (this *Tree) Add(v Comparable) {
	this.root = this.add(this.root, v)
}

func (this *Tree) add(t *TreeNode, v Comparable) *TreeNode {
	if t == nil {
		t = &TreeNode{nil, nil, v}
		return t
	}
	if v.CompareTo(t.value) < 0 {
		t.left = this.add(t.left, v)
		return t
	}
	t.right = this.add(t.right, v)
	return t
}

//是否包含值
func (this *Tree) Contain(v Comparable) bool {
	return this.contain(this.root, v)
}

func (this *Tree) contain(t *TreeNode, v Comparable) bool {
	if t == nil {
		return false
	}
	if v.CompareTo(t.value) < 0 {
		return this.contain(t.left, v)
	}
	if v.CompareTo(t.value) > 0 {
		return this.contain(t.right, v)
	}
	return true
}

func (this *Tree) SearchRange(min Comparable, max Comparable) []Comparable {
	var res []Comparable
	res = searchRange(min, max, this.root, res)
	return res
}

func searchRange(min Comparable, max Comparable, node *TreeNode, res []Comparable) []Comparable {
	if node == nil {
		return res
	}
	if min.CompareTo(node.value) < 0 {
		res = searchRange(min, max, node.left, res)
	}
	if min.CompareTo(node.value) <= 0 && max.CompareTo(node.value) >= 0 {
		res = append(res, node.value)
	}
	if min.CompareTo(node.value) > 0 || max.CompareTo(node.value) > 0 {
		res = searchRange(min, max, node.right, res)
	}
	return res
}

//打印树
func (this *Tree) Print() {
	maxLevel := maxLevel(this.root)
	var nodes []*TreeNode = []*TreeNode{this.root}
	printNode(nodes, 1, maxLevel)
}

func printNode(nodes []*TreeNode, level int, maxLevel int) {
	if len(nodes) == 0 || isAllElementsNull(nodes) {
		return
	}
	var floor int = maxLevel - level
	var endgeLines = int(math.Pow(2.0, float64(floor-1)))
	var firstSpaces = int(math.Pow(2.0, float64(floor)) - 1)
	var betweenSpaces = int(math.Pow(2.0, float64(floor+1)) - 1)
	printSpace(firstSpaces)
	var newNodes []*TreeNode
	for _, node := range nodes {
		if node == nil {
			newNodes = append(newNodes, nil)
			newNodes = append(newNodes, nil)
			fmt.Print(" ")
		} else {
			newNodes = append(newNodes, node.left)
			newNodes = append(newNodes, node.right)
			fmt.Print(node.value)
		}
		printSpace(betweenSpaces)
	}
	fmt.Println("")

	for i := 1; i <= endgeLines; i++ {
		for j := 0; j < len(nodes); j++ {
			printSpace(firstSpaces - i)
			if nodes[j] == nil {
				printSpace(endgeLines + endgeLines + i + 1)
				continue
			}
			if nodes[j].left != nil {
				fmt.Print("/")
			} else {
				printSpace(1)
			}
			printSpace(i + i - 1)
			if nodes[j].right != nil {
				fmt.Print("\\")
			} else {
				printSpace(1)
			}
			printSpace(endgeLines + endgeLines - i)
		}
		fmt.Println("")
	}
	printNode(newNodes, level+1, maxLevel)
}

func (this *Tree) Level() int {
	return maxLevel(this.root)
}

func maxLevel(tn *TreeNode) int {
	if tn == nil {
		return 0
	}
	leftLevel := maxLevel(tn.left)
	rightLevel := maxLevel(tn.right)
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

func isAllElementsNull(nodes []*TreeNode) bool {
	for _, v := range nodes {
		if v != nil {
			return false
		}
	}
	return true
}
