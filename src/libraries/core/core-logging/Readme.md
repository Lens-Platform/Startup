## Blackspace Core Logging Library
---
This document outlines how to effectively make use of this library.

```bash
go get github.com/Lens-Platform/Platform/src/libraries/core/core-logging
```

Requires:

* Go >= 1.12

This document outlines how to effectively make use of this library.

```go
import (
	log "github.com/Lens-Platform/Platform/src/libraries/core/core-logging"
)

// Log output is buffered and written periodically using Flush. Programs
// should call Flush before exiting to guarantee all log output is written.
//
// By default, all log statements write to standard error.
// This package provides several flags that modify this behavior.
// As a result, flag.Parse must be called before any logging is done.
//
//	-logtostderr=true
//		Logs are written to standard error instead of to files.
//	-alsologtostderr=false
//		Logs are written to standard error as well as to files.
//	-stderrthreshold=ERROR
//		Log events at or above this severity are logged to standard
//		error as well as to files.
//	-log_dir=""
//		Log files will be written to this directory instead of the
//		default temporary directory.
//
//	Other flags provide aids to debugging.
//
//	-log_backtrace_at=""
//		When set to a file and line number holding a logging statement,
//		such as
//			-log_backtrace_at=gopherflakes.go:234
//		a stack trace will be written to the Info log whenever execution
//		hits that statement. (Unlike with -vmodule, the ".go" must be
//		present.)
//	-v=0
//		Enable V-leveled logging at the specified level.
//	-vmodule=""
//		The syntax of the argument is a comma-separated list of pattern=N,
//		where pattern is a literal file name (minus the ".go" suffix) or
//		"glob" pattern and N is a V level. For instance,
//			-vmodule=gopher*=3
//		sets the V level to 3 in all Go files whose names begin "gopher".
func main() {
    // In this example, we want to show you that all the lines logged
    // end up in the myfile.log. You do NOT need them in your application
    // as all these flags are set up from the command line typically
    flag.Set("logtostderr", "false")     // By default klog logs to stderr, switch that off
    flag.Set("alsologtostderr", "false") // false is default, but this is informative
    flag.Set("stderrthreshold", "FATAL") // stderrthreshold defaults to ERROR, we don't want anything in stderr
    flag.Set("log_file", "myfile.log")   // log to a file

    // parse flags
    flag.Parse()

    // initialize the log object which will intuitively send logs some output defined by the client
    logs.InitLogs()
    logs.AddFlags(flag)
    // defer the flushing of logs until the process exits. Note - can define explicit conditions for this
    defer logs.FlushLogs()

    serviceName := "authenticationservice"
    var logEngine = log.NewLogger(serviceName)
}
```
