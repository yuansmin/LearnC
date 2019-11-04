# include "stdio.h"

// int fib2(int x, int y[]);
// void printFibonacci(int n);

// int fib(int x)
// {
//     int y[] = {0, 1};
//     return fib2(x, y);
//     // if (x<=1){
//     //     return 1;
//     // }
//     // else
//     //     return fib(x-1) + fib(x-2);
// }

// int fib2(int x, int y[])
// {
//     if (x==0){
//         return y[0];
//     }
//     int tmp1 = y[0];
//     int tmp=y[1];
//     y[1] = tmp + y[0];
//     y[0] = tmp;
//     return fib2(x-1, y) + tmp1;
// }

void simple_fib(int n)
{
    int a=0, b=1, tmp;
    while(n > 0){
        printf("%d ", b);
        tmp = b;
        b = a+ b;
        a = tmp;
        n--;
    }
}


void printFibonacci(int n){

    static long int first=0,second=1,sum;

    if(n>0){
         sum = first + second;
         first = second;
         second = sum;
         printf("%ld ",sum);
         printFibonacci(n-1);
    }

}


int main(){
    printf("%d\n", fib(2));
    // printFibonacci(5);
    return 0;
}
