## fileMonitor

基于Golang的文件监控小工具，使用[fsnotify](https://github.com/fsnotify/fsnotify)包，在**代码审计**或**CTF-AWD**中，无需依赖直接运行，使用十分方便。

```
[Sewell]: ~/Documents/fileMonitor
➜  ./fileMonitor_mac -h
Usage of ./fileMonitor_mac:
  -path string
    	Input Your fileMonitorPath (default "./")

[Sewell]: ~/Documents/fileMonitor
➜  ./fileMonitor_mac ./
* 2020-01-17 21:36:08 Monitor Dir: /Users/sewellding/Documents/fileMonitor
  |-- 2020-01-17 21:36:54 Create: /Users/sewellding/Documents/fileMonitor/test
* 2020-01-17 21:36:54 Add Dir: /Users/sewellding/Documents/fileMonitor/test
  |-- 2020-01-17 21:37:02 Create: /Users/sewellding/Documents/fileMonitor/test/test.go
  |-- 2020-01-17 21:37:02 Create: /Users/sewellding/Documents/fileMonitor/test/test.go~
  |-- 2020-01-17 21:37:02 Write: /Users/sewellding/Documents/fileMonitor/test/test.go
  |-- 2020-01-17 21:37:02 Remove: /Users/sewellding/Documents/fileMonitor/test/test.go~
  |-- 2020-01-17 21:38:14 Remove: /Users/sewellding/Documents/fileMonitor/test/test.go
  |-- 2020-01-17 21:38:14 Remove: /Users/sewellding/Documents/fileMonitor/test
```

