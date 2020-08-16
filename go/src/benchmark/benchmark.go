package benchmark

import (
  "time"
)

func Measure(f func ()) time.Duration {
  start := time.Now()
  f()
  return time.Now().Sub(start)
}
