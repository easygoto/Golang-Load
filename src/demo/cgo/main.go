package main

/*
#include <stdio.h>
#include <stdlib.h>

typedef struct _node {
    int value;
    struct _node* next;
} Node;

Node* create()
{
    Node* node = (Node*)malloc(sizeof(Node));
    node->value = 123;
    node->next = NULL;
    return node;
}

void test()
{
    Node *head, *list, *temp;
    head = (Node*)calloc(1, sizeof(Node));
    list = head;

    int arr[] = { 1, 2, 3, 4, 5, 6, 7, 8, 9 };
    for(int i = 0, len = sizeof(arr) / sizeof(int); i < len; i++) {
        temp = (Node*)calloc(1, sizeof(Node));
        temp->value = arr[i];
        list->next = temp;
        list = list->next;
    }

    list = head->next;
    while(list != NULL) {
        temp = list;
        printf("%d ", list->value);
        list = list->next;
        free(temp);
        temp = NULL;
    }
	printf("\n");
    free(head);
    head = NULL;
}

void forceFree(Node *node)
{
	free(node);
	node = NULL;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	C.test()

	node := C.create()
	_, _ = fmt.Printf("addr=%v, node=%v, node.value=%v\n", &node, node, node.value)
	C.free(unsafe.Pointer(node))
	_, _ = fmt.Printf("addr=%v, node=%v, node.value=%v\n", &node, node, node.value)
	C.forceFree(node)
	_, _ = fmt.Printf("addr=%v, node=%v, node.value=%v\n", &node, node, node.value)
}
