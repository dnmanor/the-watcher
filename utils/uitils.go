package utils

import (
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func ShouldRebuild(path string, op fsnotify.Op) bool {
	base := filepath.Base(path)

	//Mac OS dep
	if base == ".DS_Store" {
		return false
	}

	// Vim temporary file
	if base == "4913" {
		return false
	}

	// Vim backups
	if strings.HasSuffix(base, "~") {
		return false
	}

	return true

}
