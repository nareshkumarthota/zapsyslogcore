package zapsyslogcore

import "fmt"

func init() {
	fmt.Println("am getting ininit call")
}

// this code is taken from github.com/tchap/zapext. Modified it accoriding to the needs.

// SysLogCore root struct for zap core user friendly syslog.
