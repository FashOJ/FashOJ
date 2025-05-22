// #include <stdio.h>
// #include <iostream>
// #include <unistd.h>

// int main() {

//     std::cout << "hello world" << std::endl;

//     // 尝试执行被禁止的系统调用

//     char *args[] = {"/bin/ls", NULL};
//     execve(args[0], args, NULL);
    
//     // 如果沙箱工作正常，上面的调用应该失败
//     printf("Execve was blocked as expected\n");
//     return 0;
// }

#include <bits/stdc++.h>

using namespace std;
void solve()
{
    int n;
    cin >> n;
    int ans = n;
    for (int i = 0; i < n; i++)
    {
        int x;
        cin >> x;
        if (x == 1)
            ans--;
    }
    cout << ans << "\n";
}
int main()
{
    ios::sync_with_stdio(false), cin.tie(0), cout.tie(0);
    int _ = 1;
    // std::cin >> _;
    while (_--)
    {
        solve();
    }
    return 0;
}