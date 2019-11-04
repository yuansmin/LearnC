# include "stdio.h"

# define GET_ARRAY_LEN(array, len) {len = sizeof(array)/sizeof(array[0]);}

// C语言插入排序-升序

void my_print(int a[], int n)
{
    int i=0;
    while(i<n){
        printf("%d ", a[i]);
        i++;
    }   
    printf("\n");
}


int main()
{
    int a[]={5, 2, 4, 6, 1, 3}, key, len;
    GET_ARRAY_LEN(a, len);
    my_print(a, len);
    for(int i=1; i<len; i++){
        key = a[i];
        int j = i;
        while((j>0) && (a[j-1]>key)){
            a[j] = a[j-1];
            j--;
        }
        a[j] = key;
    }
    my_print(a, len);
//    printf("%s\n", a);
    return 0;
}

