package main

import (
    "os"
    "fmt"
    //"log"
    "io/ioutil"
    //"strconv"
    //"github.com/davecgh/go-spew/spew"
    log "github.com/sirupsen/logrus"
)


func main() {
    // load config
    if len(os.Args) == 1 {
        fmt.Fprintln(os.Stderr, "usage: fuse [config]")
        fmt.Fprintln(os.Stderr, "error: no config specified")
        os.Exit(1)
    }
    bytes, err := ioutil.ReadFile(os.Args[len(os.Args)-1])

    if os.Args[1] == "-v" {
        log.SetLevel(log.DebugLevel)
    }

    // parse config
    result, err := Parse(string(bytes))
    if err != nil {
        fmt.Fprintln(os.Stderr, "error during parsing config file:", err)
        os.Exit(1)
    }

    // prepare notifier
    notifer := NewNotifer()
    for name, alerter := range result.Alerters {
        notifer.AddAlerter(name, alerter)
    }

    // prepare monitors and create fuse
    fuse := NewFuse()
    for _, monitor := range result.Monitors {
        fuse.AddMonitor(monitor)
    }

    // start monitor's gorutines and wait
    fuse.RunWith(notifer)
}
