// Online C compiler to run C program online
#include <stdio.h>

int main() {
    
    int p[] = {0,0,1,1,1,2,2,3,3,4};
    int q[50] = {0};
    
    int length;
    
    length = sizeof(p)/sizeof(p[0]);
    
    for(int i=0; i<length ; i++){
        for (int j=0;j<i;j++){
            if(q[j]!=p[i]){
                if(q[j+1]!= 0)
                    continue;
                else{
                    q[j+1] = p[i];
                }
            }
            else{
                break;
            }
            
        }
                
    }
    
    for (int k=0; k<50; k++){
        printf("Ajay-- %d\n", q[k]);
    }
    // Write C code here
    printf("Try programiz.pro");

    return 0;
}