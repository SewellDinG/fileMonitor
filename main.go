package main

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
	"time"
)

var path = flag.String("path", "./", "Input Your fileMonitorPath")

func watchDir(watcher *fsnotify.Watcher, dir string) {
	// 遍历所有子目录
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			path, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			err = watcher.Add(path)
			if err != nil {
				return err
			}
			fmt.Println("*", time.Now().Format("2006-01-02 15:04:05"), "Monitor Dir:", path)
		}
		return nil
	})
	go func() {
		for {
			select {
			case ev := <-watcher.Events:
				{
					if ev.Op&fsnotify.Create == fsnotify.Create {
						fmt.Println("  |--", time.Now().Format("2006-01-02 15:04:05"), "Create:", ev.Name)
						// 如果是目录，则加入到监控中
						fi, err := os.Stat(ev.Name)
						if err == nil && fi.IsDir() {
							_ = watcher.Add(ev.Name)
							fmt.Println("*", time.Now().Format("2006-01-02 15:04:05"), "Add Dir:", ev.Name)
						}
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						fmt.Println("  |--", time.Now().Format("2006-01-02 15:04:05"), "Write:", ev.Name)
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						fmt.Println("  |--", time.Now().Format("2006-01-02 15:04:05"), "Remove:", ev.Name)
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						fmt.Println("  |--", time.Now().Format("2006-01-02 15:04:05"), "Rename:", ev.Name)
					}
					// 权限修改，输出冗余，视情况开启
					//if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
					//	fmt.Println("  |--", time.Now().Format("2006-01-02 15:04:05"), "Chmod:", ev.Name)
					//}
				}
			case err := <-watcher.Errors:
				{
					log.Fatal("Watcher chan err:", err)
					return
				}
			}
		}
	}()
}

func main() {
	flag.Parse()
	fileMonitorPath:=*path
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("fsnotify.NewWatcher err:", err)
	}
	watchDir(watcher, fileMonitorPath)
	// 阻止主goroutine退出
	select {}
}
