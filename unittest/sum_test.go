package unittest_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/liuli404/VBlog/unittest"
)

func TestSum(t *testing.T) {
	// 读取环境变量作为入参，VScode配置环境变量文件：Test Env File: ${workspaceFolder}/etc/unit_test.env
	a1, _ := strconv.Atoi(os.Getenv("ARG1"))
	a2, _ := strconv.Atoi(os.Getenv("ARG2"))
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	t.Log(unittest.Sum(a1, a2))
}
