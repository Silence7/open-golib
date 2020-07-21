package public

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

type ShellResult struct {
	OutMsg string
	OutErr string
	Err    error
}

func BaseCmd(ctx context.Context, shellCmd string, ch chan *ShellResult) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var result ShellResult

	cmd := exec.CommandContext(ctx, "bash", "-c", shellCmd)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	result.Err = err
	result.OutMsg = strings.TrimRight(stdout.String(), "\n")
	result.OutErr = strings.TrimRight(stderr.String(), "\n")

	ch <- &result
}
func ShellCmd(cmd string, timeout int64) string {
	var ch = make(chan *ShellResult)
	var out string

	ctx, cancel := context.WithCancel(context.Background())
	timeTick := time.NewTicker(time.Duration(timeout) * time.Second)
	if nil == timeTick {
		log.Println("NewTicker error")
		return ""
	}

	defer timeTick.Stop()

	go BaseCmd(ctx, cmd, ch)
	select {
	case <-timeTick.C:
		// 超时时间到，杀死执行脚本
		out = fmt.Sprintf("shell exec timeout")
		cancel()
	case result := <-ch:
		if nil == result {
			out = fmt.Sprintf("shell exec error: result is nil")
			return out
		}

		if nil != result.Err {
			out = fmt.Sprintf("shell exec error:%v msg:%s", result.Err, result.OutErr)
			return out
		}

		out = result.OutMsg
	}

	return out
}