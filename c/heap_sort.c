# include <stdio.h>
# include <malloc.h>

# define leftChild(i) (2*i+1)
# define maxSize 30

typedef struct {
    int a[maxSize];
    int n;
} Heap;

// 堆排序

void printList(int a[], int n)
{
    for(int i=0; i<n; i++) {
        printf("%d ", a[i]);
    }
    printf("\n");
}

// 小数上升
void MinHeapFixup(int a[], int i)
{
    int tmp, p;
    tmp = a[i];
    p = (i - 1) / 2;
    while (p > 0 && i != 0) {
        if (a[p] > a[i]) {
            a[i] = a[p];
            i = p;
            p = (i - 1) / 2;
        } else {
            break;
        }
    }
    a[i] = tmp;
}

void insert(int a[], int n, int num)
{
    a[n] = num;
    MinHeapFixup(a, n);
}

// 大数下沉
void MinHeapFixdown(int a[], int i, int n)
{
    int j, tmp;
    tmp = a[i];
    j = 2*i + 1;
    while(j < n) {
        if (j+1 < n && a[j] > a[j+1])
            j++;
        if (a[j] >= tmp)
            break;
        a[i] = a[j];
        i = j;
        j = 2*i +1;
    }
    a[i] = tmp;
}

void makeMinHeap(int a[], int n)
{
    int i, tmp;
    for (i=n/2 -1; i >= 0; i--) {
        MinHeapFixdown(a, i, n);
    }
}

void reverse(int a[], int n)
{
    int i=0, j, tmp;
    j = n -i -1;
    while (i < j) {
        tmp = a[i]; a[i] = a[j]; a[j] = tmp;
        i++;
        j = n -i -1;
    }
}

void heapSort(int a[], int n)
{
    int tmp, j=n-1;
    printList(a, n);
    makeMinHeap(a, n);
    while (j > 0) {
        tmp = a[0];
        a[0] = a[j];
        a[j] = tmp;
        MinHeapFixdown(a, 0, --j);
    }
    printList(a, n);
    reverse(a, n);
    printList(a, n);
}


// void heapSort(int a[], int n)
// {
//     int i;
//     for (i=n/2; i >= 0; i--) {
//         percDown(a, i, n)
//     }
//     for (i=n-1; i>0; i--) {

//     }
// }

int main()
{
    int a[] = {9, 12, 17, 30, -50, 20, 60, 65, 4, 49};
    // int a[] = {1, 2, -3, 4, 5, 7};
    heapSort(a, 10);
    return 0;
}
