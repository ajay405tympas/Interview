// Online C compiler to run C program online
#include <stdio.h>

int main() {
    
    char *p = "this is a very beautiful world";
    char q[50] = {0};
    
    for(int i=0; p[i]!= '\0'; i++){
        
        if (q[0] == '\0'){
            q[0] = p[i];
            continue;
        }
        for (int j=0;j<q[j]!='\0';j++){
            if(q[j]!=p[i]){
                if(q[j+1] != '\0'){
                    continue;
                }
                else{
                    q[j+1]=p[i];
                    break;
                }
            }
            else{
                break;
            }
        }
    }
    
    for (int k=0; q[k] != '\0'; k++){
        printf("Ajay-- %c\n", q[k]);
    }
    // Write C code here
    printf("Try programiz.pro");

    return 0;
}