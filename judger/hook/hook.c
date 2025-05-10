
#define _BSD_SOURCE // readlink
#include <dlfcn.h>
#include <link.h>  // ElfW
#include <seccomp.h>
#include <stdio.h>
#include <stdlib.h> // exit
#include <string.h> // strstr, memset
#include <unistd.h> // readlink
int syscalls_blacklist[] = {
    // SCMP_SYS(read),
    // SCMP_SYS(write),
    SCMP_SYS(execve),
    SCMP_SYS(fork),
    SCMP_SYS(vfork),
    SCMP_SYS(kill),
    SCMP_SYS(creat),
    SCMP_SYS(unlink),
    SCMP_SYS(chmod),
    SCMP_SYS(rename),
    SCMP_SYS(ptrace),
    SCMP_SYS(prctl),
    SCMP_SYS(socket),
    SCMP_SYS(connect),
    SCMP_SYS(bind),
    SCMP_SYS(accept),
    SCMP_SYS(clone),
    };
typedef int (*main_t)(int, char **, char **);

#ifndef __unbounded
#define __unbounded
#endif

int __libc_start_main(main_t main, int argc, char *__unbounded *__unbounded ubp_av, ElfW(auxv_t) * __unbounded auxvec,
                      __typeof(main) init, void (*fini)(void), void (*rtld_fini)(void), void *__unbounded stack_end)
{

    int i;
    ssize_t len;
    void *libc;
    int blacklist_length = sizeof(syscalls_blacklist) / sizeof(int);
    scmp_filter_ctx ctx = NULL;
    int (*libc_start_main)(main_t main, int, char *__unbounded *__unbounded, ElfW(auxv_t) *, __typeof(main),
                           void (*fini)(void), void (*rtld_fini)(void), void *__unbounded stack_end);

    // Get __libc_start_main entry point
    libc = dlopen("libc.so.6", RTLD_LOCAL | RTLD_LAZY);
    if (!libc)
    {
        exit(1);
    }

    libc_start_main = dlsym(libc, "__libc_start_main");
    if (!libc_start_main)
    {
        exit(2);
    }

    ctx = seccomp_init(SCMP_ACT_ALLOW);
    if (!ctx)
    {
        exit(3);
    }
    for (i = 0; i < blacklist_length; i++)
    {
        if (seccomp_rule_add(ctx, SCMP_ACT_KILL, syscalls_blacklist[i], 0))
        {
            exit(4);
        }
    }


    if (seccomp_load(ctx))
    {
        exit(5);
    }
    seccomp_release(ctx);
    return ((*libc_start_main)(main, argc, ubp_av, auxvec, init, fini, rtld_fini, stack_end));
}
