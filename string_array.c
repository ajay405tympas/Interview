// Online C compiler to run C program online
#include <stdio.h>

void func(char *str){
  printf("%lu\n", sizeof(str));
  printf("%c", *str++);
}


int main() {
    char a[] = "hello";
    char *q = "Ajay";
    char str[] = "world";

    char *p = a;

    printf("%c ", *(a+3));
    printf("%c \n", *a);
    
    printf("%c ", q[0]);
    printf("%c ", q[1]);

    printf("%lu", sizeof(str));
    
    func(str);
    
    
}