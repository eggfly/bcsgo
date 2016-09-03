#include<stdio.h>
#include<string.h>

// 字符串翻转函数
void reverse_str(char* strInput, int nStart, int nEnd)  
{  
    if (nStart >= nEnd || nStart < 0 || nEnd >= strlen(strInput)) {  
        return;  
    }  
    while(nStart < nEnd) {  
        char cTemp = strInput[nStart];  
        strInput[nStart] = strInput[nEnd];  
        strInput[nEnd] = cTemp;  
        nStart++;  
        nEnd--;  
    }  
}  
  
void reverse_domain(char* strInput)  
{  
    reverse_str(strInput, 0, strlen(strInput) - 1);  
  
    // 域名中的每个单词反转  
    char* strStart = strInput;  
    int nStart = 0;  
    int nEnd = 0;  
    while( *strInput != '\0')  
    {  
        if ( *strInput == '.')  
        {  
            reverse_str(strStart, nStart, nEnd - 1);  
            nStart = nEnd + 1;  
            strStart = strInput;  
        }  
        nEnd ++;  
        strInput ++;  
    }  
}

int main()
{
    char a[] = "www.mi.com";
    reverse_domain(a);
    printf("%s", a);   
}
