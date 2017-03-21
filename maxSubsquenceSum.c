# include <stdio.h>

// ### 要求：通过函数找出最大子序列

// **分析**：最大子序列可能出现在三处
// 1. 整个出现在数组的左半部
// 2. 整个出现在数组的右半部
// 3. 出现在数组的中间，横跨数据左边和右边

// 解决办法： 求出分别位于数组左边、右边和中间的最大子序列。使用“**分治**”策略
// 1. 1, 2 可使用递归来解决
// 2. 第三中情况可通过求出前半部分(包含最后一个元素)最大子序列和后半部分(包含第一个元素)最大子序列，然后求和即可。
// 3. 最后算出最大值

int roughMaxSubsequenceSum(int a[], int n)
{
    int sum,thisSum,i,j,k;
    for(i=0;i<n;i++){
        for(j=i; j<n; j++){
            thisSum = 0;
            for(k=i;k<=j;k++){
                thisSum += a[k];
            }
            if(thisSum>sum){
                sum = thisSum;
            }
        }
    }
    printf("sum = %d\n", sum);
    return sum
}

int max3(int a, int b, int c)
{
    int max=0;
    if (a > b){
        max = a;
    }else{
        max = b;
    }
    if (c > max){
        return c;
    }else{
        return max;
    }
}

int recursiveMaxSubsequenceSum(int a[], int left, int right)
{
    int maxLeft=0, maxRight=0, center=0, max=0;
    int maxLeftBorder=0, maxRightBorder=0, sum=0;
    if (left == right){
        if (a[left] > 0){
            return a[left];
        }else{
            return 0;
        }
    }

    center = (left + right) / 2;
    maxLeft = recursiveMaxSubsequenceSum(a, left, center);
    maxRight = recursiveMaxSubsequenceSum(a, center+1, right);

    for(int i=center; i>=0; i--){
        sum += a[i];
        if (sum > maxLeftBorder){
            maxLeftBorder = sum;
        }
    }

    sum = 0;
    for (int i=center+1; i<=right; i++){
        sum += a[i];
        if (sum > maxRightBorder){
            maxRightBorder = sum;
        }
    }

    max = max3(maxLeftBorder+maxRightBorder, maxRight, maxLeft);
    printf("%d\n", max);
    return max;
}

int main()
{
    int a[] = {1, 2, -3, 4, 5};
    // getMaxSubsequenceSum(a, 5);
    recursiveMaxSubsequenceSum(a, 0, 4);
    return 0;
}