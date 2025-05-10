#include <stdio.h>
#include <iostream>
#include <unistd.h>

int main() {

    std::cout << "hello world" << std::endl;

    // 尝试执行被禁止的系统调用

    char *args[] = {"/bin/ls", NULL};
    execve(args[0], args, NULL);
    
    // 如果沙箱工作正常，上面的调用应该失败
    printf("Execve was blocked as expected\n");
    return 0;
}