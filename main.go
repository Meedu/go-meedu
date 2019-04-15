package main

import (
	"flag"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"sync"
)

var mu sync.Mutex

func main() {
	var address = flag.String("address", "0.0.0.0:8089", "please input listen address.")
	var appKey = flag.String("key", "", "please input app key for sec.")
	flag.Parse()
	http.HandleFunc("/install", func(w http.ResponseWriter, r *http.Request) {
		// 解析参数
		form := r.URL.Query()
		php := form.Get("php")
		composer := form.Get("composer")
		action := form.Get("action")
		pkg := form.Get("pkg")
		dir := form.Get("dir")
		addons := form.Get("addons")
		notify := form.Get("notify")
		key := form.Get("key")
		if php == "" || composer == "" || action == "" || pkg == "" {
			fmt.Println("params error")
			return
		}
		pkgs := strings.Split(pkg, "|")
		if *appKey != key {
			fmt.Println("key error.")
			return
		}
		// 执行命令
		go func() {
			// 防止并发
			mu.Lock()

			// 打印log
			fmt.Printf("prepare run [%s %s %s %s] command\n", php, composer, action, pkg)

			// 组装执行的参数
			arg := []string{composer, action, "--no-interaction", "--no-update", "--no-suggest"}
			arg = append(arg, pkgs...)
			cmd := exec.Command(php, arg...)

			// 设置执行目录
			cmd.Dir = dir

			// 执行
			output, err := cmd.CombinedOutput()
			fmt.Printf("%s %v", output)
			mu.Unlock()

			// 回调
			status := "success"
			if err != nil {
				status = "fail"
			}
			_, _ = http.Get(notify + "?addons=" + addons + "&status=" + status + "&key=" + key)
		}()
	})
	fmt.Printf("listen:%s\n", *address)
	err := http.ListenAndServe(*address, nil)
	if err != nil {
		panic(err)
	}
}
