#include <stdio.h>
#include <malloc.h>

typedef struct nodes{
    int val;
    struct nodes *next;
} node;

void createStack(node **p)
{
    *p = (node *)malloc(sizeof(node));
    (*p)->next = NULL;
}

void push(node **p, int val)
{
    node *a;
    printf("%d ", val);
    a = (node *)malloc(sizeof(node));
    a->val = val;
    a->next = *p;
    *p = a;
}

int pop(node **p)
{
    node *a;
    int val;
    a = *p;
    val = a->val;
    if (a->next != NULL) {
        *p = a->next;
        free(a);
    }
    return val;
}

int main()
{
    int a[] = {1, 2, 3, 4, 5, 6};
    node *p;
    createStack(&p);
    int i=0;
    for (i=0; i<6; i++) {
        push(&p, a[i]);
    }
    printf("\n");
    for (i=0; i<6; i++) {
        printf("%d ", pop(&p));
    }
    printf("\n");
}
