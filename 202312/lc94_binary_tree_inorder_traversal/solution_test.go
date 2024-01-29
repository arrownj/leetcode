package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var node1 = &Node{
	Value: 1,
	Left:  nil,
	Right: nil,
}

var node3 = &Node{
	Value: 3,
	Left:  nil,
	Right: nil,
}

var node2 = &Node{
	Value: 2,
	Left:  node1,
	Right: node3,
}

var node5 = &Node{
	Value: 5,
	Left:  nil,
	Right: nil,
}

var node7 = &Node{
	Value: 7,
	Left:  nil,
	Right: nil,
}

var node6 = &Node{
	Value: 6,
	Left:  node5,
	Right: node7,
}

var node4 = &Node{
	Value: 4,
	Left:  node2,
	Right: node6,
}

// Inorder
func TestInorderRecursive1(t *testing.T) {
	assert.Equal(t, []int{}, InorderRecursive(nil))
}

func TestInorderRecursive2(t *testing.T) {
	assert.Equal(t, []int{1}, InorderRecursive(node1))
}

func TestInorderRecursive3(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, InorderRecursive(node2))
}

func TestInorderRecursive4(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, InorderRecursive(node4))
}

func TestInorderIterative1(t *testing.T) {
	assert.Equal(t, []int{}, InorderRecursive(nil))
}

func TestInorderIterative2(t *testing.T) {
	assert.Equal(t, []int{1}, InorderRecursive(node1))
}

func TestInorderIterative3(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, InorderRecursive(node2))
}

func TestInorderIterative4(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, InorderRecursive(node4))
}

// Preorder
func TestPreorderRecursive1(t *testing.T) {
	assert.Equal(t, []int{}, PreorderRecursive(nil))
}

func TestPreorderRecursive2(t *testing.T) {
	assert.Equal(t, []int{1}, PreorderRecursive(node1))
}

func TestPreorderRecursive3(t *testing.T) {
	assert.Equal(t, []int{2, 1, 3}, PreorderRecursive(node2))
}

func TestPreorderRecursive4(t *testing.T) {
	assert.Equal(t, []int{4, 2, 1, 3, 6, 5, 7}, PreorderRecursive(node4))
}

func TestPreorderIterative1(t *testing.T) {
	assert.Equal(t, []int{}, PreorderRecursive(nil))
}

func TestPreorderIterative2(t *testing.T) {
	assert.Equal(t, []int{1}, PreorderRecursive(node1))
}

func TestPreorderIterative3(t *testing.T) {
	assert.Equal(t, []int{2, 1, 3}, PreorderRecursive(node2))
}

func TestPreorderIterative4(t *testing.T) {
	assert.Equal(t, []int{4, 2, 1, 3, 6, 5, 7}, PreorderRecursive(node4))
}

// Postorder
func TestPostorderRecursive1(t *testing.T) {
	assert.Equal(t, []int{}, PostorderRecursive(nil))
}

func TestPostorderRecursive2(t *testing.T) {
	assert.Equal(t, []int{1}, PostorderRecursive(node1))
}

func TestPostorderRecursive3(t *testing.T) {
	assert.Equal(t, []int{1, 3, 2}, PostorderRecursive(node2))
}

func TestPostorderRecursive4(t *testing.T) {
	assert.Equal(t, []int{1, 3, 2, 5, 7, 6, 4}, PostorderRecursive(node4))
}

func TestPostorderIterative1(t *testing.T) {
	assert.Equal(t, []int{}, PostorderRecursive(nil))
}

func TestPostorderIterative2(t *testing.T) {
	assert.Equal(t, []int{1}, PostorderRecursive(node1))
}

func TestPostorderIterative3(t *testing.T) {
	assert.Equal(t, []int{1, 3, 2}, PostorderRecursive(node2))
}

func TestPostorderIterative4(t *testing.T) {
	assert.Equal(t, []int{1, 3, 2, 5, 7, 6, 4}, PostorderRecursive(node4))
}

//LevelOrder
func TestLevelOrder1(t *testing.T) {
	assert.Equal(t, []int{}, LevelOrder(nil))
}

func TestLevelOrder2(t *testing.T) {
	assert.Equal(t, []int{1}, LevelOrder(node1))
}

func TestLevelOrder3(t *testing.T) {
	assert.Equal(t, []int{2, 1, 3}, LevelOrder(node2))
}

func TestLevelOrder4(t *testing.T) {
	assert.Equal(t, []int{4, 2, 6, 1, 3, 5, 7}, LevelOrder(node4))
}
