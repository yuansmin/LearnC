# include <stdio.h>
# include <stdlib.h>

#define MaxSize 50

// recursive a string
// char *s = "sdfsdfsdfs";
// while (*s != '\0') {
//     printf("%C", *s);
//     *s++;
// }

typedef struct node
{
    int data;
    struct node *lchild, *rchild;
} BiTree;

void create_tree(BiTree **bt, char *str)                        //create the tree by a string like "(2(3,4))"~
{
    BiTree *ST[MaxSize], *p;
    int i=0, top=-1, k;
    *bt = NULL;
    char ch=str[i];
    while(ch!='\0'){
        switch(ch){
            case '(': k=1; top++; ST[top]=p; break;             //save the parent
            case ',': k=2; break;
            case ')': top--; break;
            default: p=(BiTree *)malloc(sizeof(BiTree));
                    p->data=ch; p->rchild=p->lchild=NULL;
                    if(*bt==NULL){                             //check if it is root
                        *bt = p;
                    }else{
                        switch(k){
                            case 1: ST[top]->lchild=p; break;
                            case 2: ST[top]->rchild=p; break;
                        }
                    }
        }
        i++;
        ch = str[i];
    }
}

void visit_tree(BiTree *bt)
{
    printf("%d ", bt->data);
    if(bt->lchild!=NULL)
        visit_tree(bt->lchild);
    if(bt->rchild!=NULL)
        visit_tree(bt->rchild);
}

int bt_height(BiTree *bt)
{
    if(bt == NULL)
        return 0;
    else
        return bt_height(bt->lchild)>bt_height(bt->rchild)?
                    bt_height(bt->lchild)+1:bt_height(bt->rchild)+1;     // return the max by < ? : >
}

int node_count(BiTree *bt)
{
    int left_node, right_node;
    if(bt == NULL)
        return 0;
    else{
        left_node = node_count(bt->lchild);
        right_node = node_count(bt->rchild);
        return left_node + right_node + 1;
    }
}

int node_2_count(BiTree *bt)
{
    int l_two_node=0, r_two_node=0, count=0;
    if(bt==NULL)
        return 0;
    if(bt->lchild!=NULL && bt->rchild!=NULL){
        count++;
    }
    if(bt->lchild!=NULL)
        l_two_node = node_2_count(bt->lchild);
    if(bt->rchild!=NULL)
        r_two_node = node_2_count(bt->rchild);
    return l_two_node + r_two_node + count;
}

int leaf_count(BiTree *bt)
{
    return node_2_count(bt) + 1;
}

void disp_tree(BiTree *bt)
{
    if(bt!=NULL){
        printf("%d", bt->data);
        if(bt->lchild!=NULL || bt->rchild!=NULL){
            printf("(");
            disp_tree(bt->lchild);
            if(bt->rchild!=NULL){
                printf(",");
            }
            disp_tree(bt->rchild);
            printf(")");
        }
    }
}

int main()
{
    BiTree *bt;
    char str[20] = "(2(3(5,6),4(3)))";
    create_tree(&bt, str);
    visit_tree(bt);
    printf("height %d\n", bt_height(bt));
    printf("node count:%d\n", node_count(bt));
    printf("leaf node:%d\n", leaf_count(bt));
    disp_tree(bt);
    printf("\n");
    return 0;
}