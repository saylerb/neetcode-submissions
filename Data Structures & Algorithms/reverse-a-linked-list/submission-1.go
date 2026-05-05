/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    curr := head

    if curr == nil {
        return nil
    }
    if curr.Next == nil {
        return curr
    } 
    tail := reverseList(curr.Next)

    curr.Next.Next = curr
    curr.Next = nil
    return tail
}
