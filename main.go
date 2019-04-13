package main

import (
	"flag"
	"fmt"
	"net/http"
	"os/exec"
	"sync"
)

var mu sync.Mutex

func main() {
	var address = flag.String("address", "0.0.0.0:8089", "please input listen address.")
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
		if php == "" || composer == "" || action == "" || pkg == "" {
			fmt.Println("params error")
			return
		}
		// 执行命令
		go func() {
			mu.Lock()
			fmt.Printf("prepare run [%s %s %s %s] command\n", php, composer, action, pkg)
			cmd := exec.Command(php, composer, action, pkg)
			cmd.Dir = dir
			output, err := cmd.CombinedOutput()
			fmt.Printf("%s %v", output)
			mu.Unlock()

			// 回调
			status := "success"
			if err != nil {
				status = "fail"
			}
			_, _ = http.Get(notify + "?addons=" + addons + "&status=" + status)
		}()
	})
	fmt.Printf("listen:%s\n", *address)
	err := http.ListenAndServe(*address, nil)
	if err != nil {
		panic(err)
	}
}
