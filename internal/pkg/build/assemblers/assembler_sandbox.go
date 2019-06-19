// Copyright (c) 2018-2019, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package assemblers

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/sylabs/singularity/internal/pkg/client/cache"
	"github.com/sylabs/singularity/internal/pkg/sylog"
	"github.com/sylabs/singularity/pkg/build/types"
)

// SandboxAssembler stores data required to assemble the image, for instannce the
// image cache
type SandboxAssembler struct {
	ImgCache *cache.ImgCache
}

// Assemble creates a Sandbox image from a Bundle
func (a *SandboxAssembler) Assemble(b *types.Bundle, path string) (err error) {
	sylog.Infof("Creating sandbox directory...")

	// move bundle rootfs to sandboxdir as final sandbox
	sylog.Debugf("Moving sandbox from %v to %v", b.Rootfs(), path)
	if _, err := os.Stat(path); err == nil {
		os.RemoveAll(path)
	}

	var stderr bytes.Buffer
	cmd := exec.Command("mv", b.Rootfs(), path)
	// If the assemble is associated to a specific image cache, we make
	// sure that the environment variable points to it for the process
	// that will create.
	if a.ImgCache != nil {
		cmd.Env = append(os.Environ(), cache.DirEnv+"="+a.ImgCache.BaseDir)
	}
	cmd.Stderr = &stderr
	if err = cmd.Run(); err != nil {
		return fmt.Errorf("Sandbox Assemble Failed: %v: %v", err, stderr.String())
	}

	return nil
}
