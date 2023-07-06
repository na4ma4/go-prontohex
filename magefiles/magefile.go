//go:build mage

package main

import (
	"context"

	"github.com/magefile/mage/mg"
	//mage:import na4ma4
	"github.com/na4ma4/mage"
)

var Default = Test

func Lint(ctx context.Context) error {
	mg.CtxDeps(ctx, mage.GoLint)
	return nil
}

func Test(ctx context.Context) error {
	// chgBack := mage.ChangeLogger(func(format string, v ...any) {
	// 	fmt.Printf(fmt.Sprintf("Intercepted Logs: %s\n", format), v...)
	// })
	// defer chgBack()
	mg.CtxDeps(ctx, mage.GoTest)
	mg.CtxDeps(ctx, mage.GoLint)
	return nil
}

func Clean(ctx context.Context) error {
	mg.CtxDeps(ctx, mage.Clean)
	return nil
}

// Manage your deps, or running package managers.
func ModDownload(ctx context.Context) error {
	mg.CtxDeps(ctx, mage.GoModDownload)
	return nil
}
