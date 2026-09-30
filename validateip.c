// Online C compiler to run C program online
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <ctype.h>

int main() {
    // Write C code here
    
    char ip[] = ".192.168.0.1";
    char ch = '.';
    int count = 0;
    int i = 0;
    char *result;
    
    //printf("Ajay");
    for(i = 0; ip[i] != '\0'; i++) {
        if (ip[i] == ch) {
            count++;
        }
    }
    // printf("Ajayx");
    if ((count >= 4) || (ip[0] == '.')){
        printf("Ip address is invalid");
       return 0;
    }
     //printf("Ajayy");
     for(i = 0; i<strlen(ip)-1; i++) {
        if (ip[i] == ch && ip[i+1] == ch) {
            printf("Ip address is invalid");
            return 0; 
        }
    }
    //printf("Ajayz");
    result = strtok(ip, ".");
    while (result != NULL){
        printf("result%s\n", result);
        
        //printf("AjayzA\n");
       for(i=0;i<strlen(result);i++){
        if (!isdigit(result[i])){
            printf("Ip address is invalid contains alpha numeric");
            return 0; 
        }
       }
        
        int temp = atoi(result);
        printf("temp%d",temp);
        if(temp<0 || temp>255){
            printf("Ip address is invalid out of bounds");
            return 0; 
        }
        temp = 0;
        result = strtok(NULL, ".");
    }
    
    //printf("Try programiz.pro %d" , ip);

    return 0;
}