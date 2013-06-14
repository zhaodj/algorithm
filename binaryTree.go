package algorithm

type Comparable interface {
	CompareTo(Comparable) int
}

type Node interface {
	Left() Node
	Right() Node
	Value() Comparable
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

func (node *TreeNode) Left() Node {
	return node.left
}

func (node *TreeNode) Right() Node {
	return node.right
}

func (node *TreeNode) Value() Comparable {
	return node.value
}

//打印树
func (this *Tree) Print() {
	maxLevel := maxLevel(this.root)
	var nodes []Node = []Node{this.root}
	PrintNode(nodes, 1, maxLevel)
}

func (this *Tree) Level() int {
	return maxLevel(this.root)
}
