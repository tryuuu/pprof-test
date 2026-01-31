package bad

import (
	"time"
)

func Do() {
	endTime := time.Now().Add(5 * time.Second)
	data := make([]string, 0)
	for time.Now().Before(endTime) {
		for i := 0; i < 100000; i++ {
			data = append(data, "pprof-test")
		}
	}
}
