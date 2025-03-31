package sandbox

/*
#include <sys/resource.h>
#include <sys/time.h>
#include <unistd.h>

int set_limit(int resource, rlim_t limit) {
    struct rlimit rl;
    rl.rlim_cur = limit;
    rl.rlim_max = limit;
    return setrlimit(resource, &rl);
}
*/
import "C"
import "fmt"

// SetLimits 设置资源限制
func SetLimits(cpuTime, memoryLimit, fileSize, processLimit int) error {
    // 限制 CPU 时间 (秒)
    if C.set_limit(C.RLIMIT_CPU, C.rlim_t(cpuTime)) != 0 {
        return fmt.Errorf("无法设置 CPU 时间限制")
    }

    // 限制内存使用 (字节)
    if C.set_limit(C.RLIMIT_AS, C.rlim_t(memoryLimit)) != 0 {
        return fmt.Errorf("无法设置内存限制")
    }

    // 限制文件大小 (字节)
    if C.set_limit(C.RLIMIT_FSIZE, C.rlim_t(fileSize)) != 0 {
        return fmt.Errorf("无法设置文件大小限制")
    }

    // 限制最大进程数
    if C.set_limit(C.RLIMIT_NPROC, C.rlim_t(processLimit)) != 0 {
        return fmt.Errorf("无法设置进程数量限制")
    }

    return nil
}
