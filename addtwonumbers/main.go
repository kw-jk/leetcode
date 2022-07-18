package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func newNode(val int) *ListNode {
	node := &ListNode{
		Val: val,
	}

	node.Next = nil

	return node
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addNumbers(l1, l2, false)
}

func addNumbers(l1 *ListNode, l2 *ListNode, carry bool) *ListNode {
	var (
		node  *ListNode
		next1 *ListNode
		next2 *ListNode
	)
	if l1 != nil || l2 != nil {
		val1 := 0
		val2 := 0
		if l1 != nil {
			val1 = l1.Val
		}
		if l2 != nil {
			val2 = l2.Val
		}
		sum := 0
		if carry {
			sum = val1 + val2 + 1
		} else {
			sum = val1 + val2
		}
		node = newNode(sum % 10)

		if l1 != nil {
			next1 = l1.Next
		} else {
			next1 = nil
		}

		if l2 != nil {
			next2 = l2.Next
		} else {
			next2 = nil
		}
		if sum >= 10 {
			carry = true
		} else {
			carry = false
		}
		node.Next = addNumbers(next1, next2, carry)
	} else if carry {
		node = newNode(1)
	}

	return node
}

func main() {
	l1 := newNode(9)
	l1.Next = newNode(9)
	l1.Next.Next = newNode(9)
	l1.Next.Next.Next = newNode(9)
	l1.Next.Next.Next.Next = newNode(9)
	l1.Next.Next.Next.Next.Next = newNode(9)
	l1.Next.Next.Next.Next.Next.Next = newNode(9)

	l2 := newNode(9)
	l2.Next = newNode(9)
	l2.Next.Next = newNode(9)
	l2.Next.Next.Next = newNode(9)

	addTwoNumbers(l1, l2)
}
