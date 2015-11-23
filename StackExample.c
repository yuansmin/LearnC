//stack example
# include <stdio.h>

# define STACK_MAX 100

typedef struct
{
   int data[STACK_MAX];
   int size;
}Stack;

void Stack_Init(Stack *S)
{
   S->size=0;
}

int Stack_Top(Stack *S)
{
   if(S->size == 0){
      fprintf(stderr, "Error: stack empty\n");
      return -1;
   }
   return S->data[S->size - 1];
}

void Stack_Push(Stack *S, int a)
{
   if(S->size >= STACK_MAX-1){
      fprintf(stderr, "Error: stack full\n");
   }else{
        S->data[S->size] = a;
        S->size++;
   }
}

void Stack_Pop(Stack *S)
{
    if(S->size > 0){
        S->size--;
    }else{
        fprintf(stderr, "Error: stack empty\n");
    }
}

int main()
{
    Stack S;
    Stack_Init(&S);
    Stack_Push(&S, 5);
    Stack_Push(&S, 6);
    Stack_Push(&S, 7);
    Stack_Pop(&S);
    printf("top is %d\n", Stack_Top(&S));
    printf("%d", S.data[S.size-2]);
    return 0;
}