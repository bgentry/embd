// Package all conviniently loads all the inbuilt/supported host drivers.
package all

import (
	_ "github.com/bgentry/embd/host/bbb"
	_ "github.com/bgentry/embd/host/rpi"
)
