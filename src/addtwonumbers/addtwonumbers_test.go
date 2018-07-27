package addtwonumbers

import "testing"

func Test_addtwonumbers(t *testing.T)  {
	node1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	node2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	nodeResult := addTwoNumbers(node1, node2)
	if nodeResult.Val == 7 && nodeResult.Next.Val == 0 && nodeResult.Next.Next.Val == 8 {
		t.Log("OK")
	} else {
		t.Log("Failed")
	}
}