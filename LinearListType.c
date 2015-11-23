// linear list type example
# include <stdio.h>
# include <stdlib.h>

typedef int ElemType;
# define MaxSize 100

typedef struct{
    ElemType data[MaxSize];
    int length;
}SqList;

void List_Init(SqList *L)
{
    L->length = 0;
}

int GetLength(SqList *L)
{
    return L->length;
}

int GetElem(SqList *L, int n)
{
    if(n<0 || n>L->length){
        fprintf(stderr, "Error: out of index");
        return 0;
    }else{
        return L->data[n-1];
    }
}

int Locate(SqList *L, ElemType x)
{
    int i=0;
    while(i<L->length){
        if(L->data[i] == x) break;
        i++;
    }
    if(i>=L->length)
        return 0;
    else
        return i+1;
}

int Insert(SqList *L, ElemType x, int i)
{
    ElemType tmp;
    if(i<1 || i>L->length+1){
        return 0;
    }
    if(L->length >= MaxSize){
        return 0;
    }
    L->length++;
    for(int j=i-1;j<L->length;j++){
        tmp = L->data[j];
        L->data[j] = x;
        i = tmp;
    }
    return 1;
}

int Delete(SqList *L, int n)
{
    ElemType tmp;
    int i=n;
    if(n<1 || n>L->length)
        return 0;
    while(i<L->length){
        L->data[i-1] = L->data[i];
        i++;
    }
    L->length--;
    return 1;
}

void DispList(SqList *L)
{
    int i=0;
    while(i<L->length){
        printf("%d ", L->data[i]);
        i++;
    }
    printf("\n");
}

void GetCommen(SqList *A, SqList *B, SqList *C)
{
    int i=0, j=0, k=1;
    while(i<A->length){
        while(j<B->length && A->data[i]!=B->data[j])
            j++;
        if(j<B->length){
            Insert(C, B->data[j], k);
            k++;
        }
        i++;
        j=0;
    }
}

int main()
{
    SqList S, K, T;
    List_Init(&S);
    List_Init(&K);
    List_Init(&T);
    Insert(&S, 1, 1);
    Insert(&S, 2, 2);
    Insert(&S, 3, 3);
    Insert(&S, 4, 4);
    Insert(&S, 5, 5);
    Insert(&K, 1, 1);
    Insert(&K, 7, 2);
    Insert(&K, 3, 3);
    Insert(&K, 8, 4);
    Insert(&K, 5, 5);
    printf("S:");
    DispList(&S);
    printf("K:");
    DispList(&K);
    GetCommen(&S, &K, &T);
    printf("T:");
    DispList(&T);
    return 0;
}