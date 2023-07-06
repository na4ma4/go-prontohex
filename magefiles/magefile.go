//go:build mage
// +build mage

package main

import (
	"context"
	"log"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = Test

func Test(ctx context.Context) error {
	log.Println("Testing...")
	mg.CtxDeps(ctx, ModDownload)
	args := []string{
		"test",
	}

	if os.Getenv("MAGEFILE_DEBUG") == "1" {
		args = append(args, "-test.v=1")
	}

	if err := sh.RunV("go", append(args, "./...")...); err != nil {
		return err
	}

	return nil
}

func Clean(ctx context.Context) error {
	log.Println("Cleaning...")
	if err := os.RemoveAll(".makefiles"); err != nil {
		return err
	}
	if err := os.RemoveAll("artifacts/lint"); err != nil {
		return err
	}
	return nil
}

// Manage your deps, or running package managers.
func ModDownload(ctx context.Context) error {
	log.Println("Downloading Modules...")
	cmd := exec.Command("go", "mod", "download")
	return cmd.Run()
}
