package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    result, sum := new(ListNode), 0
    for tmp:= result ; l1 != nil || l2 != nil || sum !=0 ;{
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        tmp.Next = &ListNode{Val: sum % 10}
        sum /= 10
        tmp = tmp.Next
    }

    return result.Next
}

//Another version
/*func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    result := &ListNode{}
    tmp := result
    for l1 != nil || l2 != nil {
        if l1 != nil {
            tmp.Val += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            tmp.Val += l2.Val
            l2 = l2.Next
        }
        if tmp.Val > 9 {
            tmp.Val -= 10
            tmp.Next = &ListNode{Val: 1}
        } else if l1 != nil || l2 != nil {
            tmp.Next = &ListNode{}
        }
        tmp = tmp.Next
    }
    return result
}*/