// +build !fortest

package fiberutil

import "github.com/gofiber/fiber/v2"

var queries = make(map[string]string)

func CtxSetQuery(key string, value string) {
	queries[key] = value
}

func CtxQuery(ctx *fiber.Ctx, key string, defaultValue ...string) string {
	return ctx.Query(key, defaultValue...)
}
