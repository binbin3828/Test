/*
 * @Autor: 郭彬
 * @Description: 单链表
 * @Date: 2022-06-15 15:19:23
 * @LastEditTime: 2022-06-20 09:45:48
 * @FilePath: \Test\util\List.go
 */
package util

import "fmt"

type Object interface{}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	headNode *Node
}

// 判断链表是否为空
func (t *List) IsEmpty() bool {
	if t.headNode == nil {
		return true
	} else {
		return false
	}
}

// 获取链表长度
func (t *List) Length() int {
	cur := t.headNode
	count := 0
	for cur != nil {
		cur = cur.Next
		count++
	}
	return count
}

//从链表头部添加元素
func (t *List) Add(data Object) *Node {
	node := &Node{Data: data}
	node.Next = t.headNode
	t.headNode = node
	return node
}

// 从链表尾部添加元素
func (t *List) Append(data Object) *Node {
	node := &Node{Data: data}
	if t.IsEmpty() {
		t.headNode = node
	} else {
		cur := t.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
		cur.Next.Next = nil
	}
	return node
}

// 在指定位置添加元素，此处index只下标，t.headNode的index
func (t *List) Insert(index int, data Object) {
	if index < 0 {
		t.Add(data)
	} else if index > t.Length() {
		t.Append(data)
	} else {
		pre := t.headNode
		count := 0
		for count < (index - 1) { //控制移动的数量
			pre = pre.Next
			count++
		}
		//当循环退出后，pre指向index -1的位置
		node := &Node{Data: data}
		node.Next = pre.Next
		pre.Next = node
	}
}

//删除指定元素
func (t *List) Remove(data Object) {
	pre := t.headNode
	if pre.Data == data {
		t.headNode = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Next.Data == data {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
	}
}

//删除指定位置元素
func (t *List) RemoveAtIndex(index int) {
	pre := t.headNode
	if index <= 1 {
		t.headNode = pre.Next
	} else if index > t.Length() {
		fmt.Println("超出链表的长度")
		return
	} else {
		count := 1
		for count != index-1 && pre.Next != nil {
			count++
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
}

//查看是否包含某一个元素
func (t *List) Contain(data Object) bool {
	cur := t.headNode

	for cur != nil {
		if cur.Data == data {
			return true
		}
		cur = cur.Next
	}
	return false
}

//遍历所有节点
func (t *List) ShowList() {
	if !t.IsEmpty() {
		cur := t.headNode
		for cur != nil {
			fmt.Println("===", cur.Data)
			cur = cur.Next
		}
	}
}

// 前后两个元素依次交叉重排
func (t *List) ReSortList() {
	cur := t.headNode
	var ls []*Node
	for ; cur != nil; cur = cur.Next {
		ls = append(ls, cur)
	}
	i := 0
	j := len(ls) - 1
	for i < j {
		ls[i].Next = ls[j]
		i++
		if i >= j {
			break
		}
		ls[j].Next = ls[i]
		j--
	}
	ls[i].Next = nil
}

// 反转
func (t *List) Reverse() {
	var pre *Node
	cur := t.headNode
	for cur != nil {
		tmp := cur.Next
		//修改指针指向
		cur.Next = pre
		//pre， cur 后移
		pre = cur
		cur = tmp
	}
	t.headNode = pre
}

// 合并两个有序链表
func (t *List) MergeTwoLists(list1 *Node, list2 *Node) *Node {
	head := &Node{}
	tail := head
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Data.(int) < list2.Data.(int) {
				tail.Next = list1
				list1 = list1.Next
			} else {
				tail.Next = list2
				list2 = list2.Next
			}
			tail = tail.Next
		} else if list1 != nil && list2 == nil {
			tail.Next = list1
			break
		} else if list1 == nil && list2 != nil {
			tail.Next = list2
			break
		}
	}
	return head.Next
}
