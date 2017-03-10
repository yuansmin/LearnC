# include "stdio.h"

// 将数字单个输出，可输出任意实数

void PrintOut(int x);
void PrintDigit(int x);
void PreDispose(double x);

int main()
{
    PreDispose(4321.5125);
    // printDigit(4321);
    // printf("%d\n", haha(-10));
    return 0;
}

void printDigit(int x)
{
    // if (x<10){
    //     printf("%d", x);
    // } else {
    //     printDigit(x/10);
    //     printf("%d", x%10);
    // }

    if (x>10){
        printDigit(x/10);
    }
    printf("%d", x%10);
}

void PreDispose(double x)
{
    if (x < 0){
        putchar('-');
        x = -x;
    }
    int int_part=(int)x;
    double y = x - int_part;
    PrintOut(int_part);
    int n=10;
    putchar('.');
    while (n) {
        y *= 10;
        PrintDigit((int)y % 10);
        y = y - (int)y;
        n--;
    }
    printf("\n");
}

void PrintOut(int x)
{
    if (x > 10) {
        PrintDigit(x/10);
    }
    PrintDigit(x%10);
}

void PrintDigit(int x)
{
    printf("%d", x);
}
