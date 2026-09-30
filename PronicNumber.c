// Online C compiler to run C program online
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

bool searchArray(int key, int* arr, int counter){
    bool value = false;
    for(int i=0; i<counter; i++){
        if(key == arr[i]){
            value = true;
            break;
        }
        else{
            continue;
            value = false;
        }
            
    }
    return value;
}

int main() {
    // Write C code here
    int n=45;
    int arr[45];
    int counter = 0;
    bool value = false;
    for(int i=2;i<=n/2; i++){
        if(n%i == 0){
            arr[counter] = i;
            counter = counter+1;
        }
    }
    for (int j=0; j<counter; j++){
        value = searchArray((arr[j]+1), arr,counter);
    }
    
    if(value == true){
            printf("number is pronic");
        }
        else{
            printf("number is not pronic");
        }

    return 0;
}