//+build profile

package edb

import "github.com/pkg/profile"
import "fmt"

var PROFILER_ENABLED=true
var PROFILE ProfilerStart

type PkgProfile struct {
}
func (p PkgProfile) Start() ProfilerStart {
  PROFILE = profile.Start(profile.CPUProfile, profile.ProfilePath("."))
  return PROFILE
}
func (p PkgProfile) Stop() {
  p.Stop()
}

var RUN_PROFILER = func() ProfilerStop {
    fmt.Println("RUNNING ENABLED PROFILER")
    return PkgProfile{}
}
