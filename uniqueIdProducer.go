package go_plugins

import (
	"time"
)

func UniqueId(puber string) string {
	return puber + time.Now().Format(time.StampMilli)
}
