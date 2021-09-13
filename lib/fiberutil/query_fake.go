// +build fortest

package util

import (
	"github.com/gofiber/fiber/v2"
)

var queries = make(map[string]string)

func CtxSetQuery(key string, value string) {
	queries[key] = value
}

func CtxQuery(ctx *fiber.Ctx, key string, defaultValue ...string) string {
	if value, ok := queries[key]; ok {
		return value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}
