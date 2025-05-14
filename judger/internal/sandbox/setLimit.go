package sandbox

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/containerd/cgroups"
)

type cgroup struct {
	v2         bool
	cgroupName string
}

const cgroupRoot = "/sys/fs/cgroup"

func NewCgroup(cgroupName string) (*cgroup, error) {

	res := new(cgroup)
	res.cgroupName = cgroupName

	if cgroups.Mode() == cgroups.Unified {
		if err := os.Mkdir(path.Join(cgroupRoot, cgroupName), 0755); err != nil {
			return nil, fmt.Errorf("can't create new sub cgroup: %v", err)
		}
		res.v2 = true
	} else {
		if err := os.Mkdir(path.Join(cgroupRoot, "memory", cgroupName), 0755); err != nil {
			return nil, fmt.Errorf("can't create new sub cgroup: %v", err)
		}
		res.v2 = false
	}
	return res, nil
}

func (c *cgroup) AddProc(pid int) error {
	if c.v2 {
		os.WriteFile(
			path.Join(cgroupRoot,c.cgroupName,"cgroup.procs"),
			[]byte(strconv.Itoa(pid)),
			0655,
		)
	}else {

		os.WriteFile(
			path.Join(cgroupRoot,"memory",c.cgroupName,"tasks"),
			[]byte(strconv.Itoa(pid)),
			0655,
		)

	}

	return nil
}

func (c *cgroup) SetMemoryLimit(memoryLimit int) error {

	if c.v2 {
		if err := os.WriteFile(
			path.Join(cgroupRoot, c.cgroupName, "memory.max"),
			[]byte(strconv.Itoa(memoryLimit*1024*1024)),
			0755,
		); err != nil {
			return fmt.Errorf("set memory.max failed: %v", err)
		}
		if err := os.WriteFile(
			path.Join(cgroupRoot, c.cgroupName, "memory.swap.max"),
			[]byte("0"),
			0755,
		); err != nil {
			return fmt.Errorf("set memory.swap.max failed: %v", err)
		}
	} else {
		if err := os.WriteFile(
			path.Join(cgroupRoot, "memory", c.cgroupName, "memory.limit_in_bytes"),
			[]byte(strconv.Itoa(memoryLimit*1024*1024)),
			0755,
		); err != nil {
			return fmt.Errorf("set memory.limit_in_bytes failed: %v", err)
		}

		if err := os.WriteFile(
			path.Join(cgroupRoot, "memory", c.cgroupName, "memory.swappiness"),
			[]byte(strconv.Itoa(memoryLimit*1024*1024)),
			0755,
		); err != nil {
			return fmt.Errorf("set memory.limit_in_bytes failed: %v", err)
		}
	}
	return nil
}
func (c *cgroup) DelCgroup() error {
	if c.v2 {
		if err := os.RemoveAll(path.Join(cgroupRoot, c.cgroupName)); err != nil {
			return fmt.Errorf("delete cgroup failed: %v", err)
		}
	}
	return nil
}

type MemoryEvents struct {
	Low      uint64 `json:"low"`
	High     uint64 `json:"high"`
	Max      uint64 `json:"max"`
	OOM      uint64 `json:"oom"`
	OOMKill  uint64 `json:"oom_kill"`
}

func parseMemoryEvents(r io.Reader) (*MemoryEvents, error) {
	events := &MemoryEvents{}
	scanner := bufio.NewScanner(r)
	foundFields := make(map[string]bool)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("malformed line: %q", line)
		}

		key := strings.TrimSpace(parts[0])
		valueStr := strings.TrimSpace(parts[1])

		value, err := strconv.ParseUint(valueStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for %s: %v", key, err)
		}

		switch key {
		case "low":
			events.Low = value
			foundFields["low"] = true
		case "high":
			events.High = value
			foundFields["high"] = true
		case "max":
			events.Max = value
			foundFields["max"] = true
		case "oom":
			events.OOM = value
			foundFields["oom"] = true
		case "oom_kill":
			events.OOMKill = value
			foundFields["oom_kill"] = true
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// 检查必要字段是否存在
	requiredFields := []string{"oom", "oom_kill"}
	for _, field := range requiredFields {
		if !foundFields[field] {
			return nil, fmt.Errorf("missing required field: %s", field)
		}
	}

	return events, nil
}

func (c *cgroup) CheckOom() (bool, error) {

	var res bool

	if c.v2 {
		file, err := os.Open(path.Join(cgroupRoot,c.cgroupName,"memory.events"))
		if err != nil {
			return false,fmt.Errorf("read oom event failed: %v",err)
		}

		event,err := parseMemoryEvents(file)
		if err != nil {
			return false,fmt.Errorf("read oom event failed: %v",err)
		}

		if event.OOM>0 || event.OOMKill>0 {
			res = true
		}else {
			res = false
		}

	} else {
		cntByte, err := os.ReadFile(path.Join(cgroupRoot, "memory", c.cgroupName, "memory.failcnt"))	
		if err != nil {
			return false, fmt.Errorf("read oom event failed: %v", err)
		}

		// 逆天末尾换行符
		fuckendl := strings.TrimSpace(string(cntByte))
		cnt, err := strconv.Atoi(string(fuckendl))
		if err != nil {
			return false, fmt.Errorf("read oom event failed: %v", err)
		}
		if cnt > 0 {
			res = true
		}else {
			res = false
		}
	}

	return res,nil
}

// 获取内存使用量（KB）
func (c *cgroup) GetMemoryUsage() (int, error) {
	var memFile string
	
	if c.v2 {
		memFile = path.Join(cgroupRoot, c.cgroupName, "memory.current")
	} else {
		memFile = path.Join(cgroupRoot, "memory", c.cgroupName, "memory.max_usage_in_bytes")
	}
	
	memBytes, err := os.ReadFile(memFile)
	if err != nil {
		return 0, fmt.Errorf("读取内存使用失败: %v", err)
	}
	
	memStr := strings.TrimSpace(string(memBytes))
	memUsage, err := strconv.Atoi(memStr)
	if err != nil {
		return 0, fmt.Errorf("解析内存使用失败: %v", err)
	}
	
	// 转换为 KB
	return memUsage / 1024, nil
}
