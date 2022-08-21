/*
Detect a cycle in a linked list. Note that the head pointer may be 'NULL' if the list is empty.

A Node is defined as: 
    struct Node {
        int data;
        struct Node* next;
    }
*/

bool has_cycle(Node* head) {
    Node* k1 = head;
    Node* k2 = head;
    for (;;) {
        if (k1 == NULL || k2 == NULL || k2->next == NULL) {
            return false;
        }
        k1 = k1->next;
        k2 = k2->next->next;
        if (k1 == k2) {
            return true;
        }
    }
    return false;
}
