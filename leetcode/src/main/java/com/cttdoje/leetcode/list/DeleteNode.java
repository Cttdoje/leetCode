package com.cttdoje.leetcode.list;

/**
 * @author cttdoje
 * @date 2020/3/21 下午4:49
 */
public class DeleteNode {
    /**
     * 解题思路：
     * 1. 首先在这个链表中无法访问上个节点
     * 2. 把后一个节点的值复制过来
     * 3. 把后一个节点删除
     */
    public void deleteNode(ListNode node) {
        ListNode next = node.next;
        node.val = next.val;
        node.next = next.next;
    }
}

class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
    }
}
