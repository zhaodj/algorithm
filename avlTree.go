package algorithm

//平衡二叉树
type AVLTree struct {
	root *AVLNode
}

type AVLNode struct {
	left, right *AVLNode
	value       Comparable
	height      int
}

//添加节点
func (this *AVLTree) Add(v Comparable) {
	this.root = add(this.root, v)
}

func add(t *AVLNode, v Comparable) *AVLNode {
	if t == nil {
		t = &AVLNode{nil, nil, v, 0}
	} else if v.CompareTo(t.value) < 0 {
		t.left = add(t.left, v)
		if height(t.left)-height(t.right) == 2 {
			if v.CompareTo(t.left.value) < 0 {
				t = rotateLL(t)
			} else {
				t = rotateLR(t)
			}
		}
	} else {
		t.right = add(t.right, v)
		if height(t.right)-height(t.left) == 2 {
			if v.CompareTo(t.right.value) > 0 {
				t = rotateRR(t)
			} else {
				t = rotateRL(t)
			}
		}
	}
	t.height = max(height(t.left), height(t.right)) + 1
	return t
}

func height(node *AVLNode) int {
	if node == nil {
		return -1
	}
	return node.height
}

func rotateLL(node *AVLNode) *AVLNode {
	top := node.left
	node.left = top.right
	top.right = node
	node.height = max(height(node.left), height(node.right)) + 1
	top.height = max(height(top.left), height(top.right)) + 1
	return top
}

func rotateRR(node *AVLNode) *AVLNode {
	top := node.right
	node.right = top.left
	top.left = node
	node.height = max(height(node.left), height(node.right)) + 1
	top.height = max(height(top.left), height(top.right)) + 1
	return top
}

func rotateLR(node *AVLNode) *AVLNode {
	node.left = rotateRR(node.left)
	return rotateLL(node)
}

func rotateRL(node *AVLNode) *AVLNode {
	node.right = rotateLL(node.right)
	return rotateRR(node)
}

func max(first int, second int) int {
	if first > second {
		return first
	}
	return second
}

func (node *AVLNode) Left() Node {
	return node.left
}

func (node *AVLNode) Right() Node {
	return node.right
}

func (node *AVLNode) Value() Comparable {
	return node.value
}

func (tree *AVLTree) Print() {
	var nodes []Node = []Node{tree.root}
	PrintNode(nodes, 1, maxLevel(tree.root))
}

func (this *AVLTree) SearchRange(min Comparable, max Comparable) []Comparable {
	var res []Comparable
	res = searchAVLRange(min, max, this.root, res)
	return res
}

func searchAVLRange(min Comparable, max Comparable, node *AVLNode, res []Comparable) []Comparable {
	if node == nil {
		return res
	}
	if min.CompareTo(node.value) < 0 {
		res = searchAVLRange(min, max, node.left, res)
	}
	if min.CompareTo(node.value) <= 0 && max.CompareTo(node.value) >= 0 {
		res = append(res, node.value)
	}
	if min.CompareTo(node.value) > 0 || max.CompareTo(node.value) > 0 {
		res = searchAVLRange(min, max, node.right, res)
	}
	return res
}

//是否包含值
func (this *AVLTree) Contain(v Comparable) bool {
	return this.contain(this.root, v)
}

func (this *AVLTree) contain(t *AVLNode, v Comparable) bool {
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
