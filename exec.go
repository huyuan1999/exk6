// @Time    : 2023/11/30 7:38 PM
// @Author  : HuYuan
// @File    : exec.go
// @Email   : hy2803660215@163.com

package exk6

import (
	"log"
	"os/exec"
	"strings"

	"go.k6.io/k6/js/modules"
)

// EXEC 模块实现
type EXEC struct {
	workDir string
}

func init() {
	modules.Register("k6/x/exec", new(EXEC))
}

// SetWorkDir 设置命令行工作目录
func (e *EXEC) SetWorkDir(workDir string) {
	e.workDir = workDir
}

// Command 执行命令
func (e *EXEC) Command(name string, args ...string) string {
	cmd := exec.Command(name, args...)
	cmd.Dir = e.workDir
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err.Error() + " on command: " + name + " " + strings.Join(args, " "))
	}
	return string(out)
}
