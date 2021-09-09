package cache

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/xerrors"
)

var cacheDirectory = ""

var CacheFileNotFoundErr = xerrors.New("cache file is not found error")

const (
	cacheDir  = "labeldb"
	cacheFile = "label.db"
)

func SetDirectory(dir string) string {
	cacheDirectory = dir

	return cacheDirectory
}

func Dir() string {
	if cacheDirectory != "" {
		return cacheDirectory
	}

	dir, err := os.UserCacheDir()
	if err != nil {
		dir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("failed to open user cache directory")
		}
		return filepath.Join(dir, cacheDir)
	}
	return filepath.Join(dir, cacheDir)
}
func Get() string {
	return cacheDirectory
}

func Create() (*os.File, error) {
	filePath := cacheFilePath()

	if !exists(filePath) {
		if err := os.MkdirAll(Dir(), 0700); err != nil {
			return nil, xerrors.Errorf("failed to make directory: %w", err)
		}
	}

	f, err := os.Create(filePath)
	if err != nil {
		return nil, xerrors.Errorf("failed to create: %w", err)
	}
	return f, nil
}

// func Save(v interface{}, path string, key string) error {
// 	if err := os.MkdirAll(filepath.Join(Dir(), key, path), os.FileMode(0700)); err != nil {
// 		return xerrors.Errorf("failed to create directory: %w", err)
// 	}
//
// 	b, err := json.MarshalIndent(v, "", "  ")
// 	if err != nil {
// 		return xerrors.Errorf("failed to marshal json: %w", err)
// 	}
//
// 	f, err := os.Create(filepath.Join(Dir(), key, path, "result.json"))
// 	if err != nil {
// 		return xerrors.Errorf("failed to create file: %w", err)
// 	}
// 	defer f.Close()
// 	n, err := f.Write(b)
// 	if err != nil {
// 		return xerrors.Errorf("failed to write file: %w", err)
// 	}
// 	if len(b) != n {
// 		return xerrors.Errorf("failed to write file length: actual %d, expected %d", n, len(b))
// 	}
//
// 	return nil
// }

func Timestamp() (time.Time, error) {
	fileInfo, err := os.Stat(cacheFilePath())
	if err != nil {
		return time.Time{}, xerrors.Errorf("failed to stat file: %w", err)
	}
	return fileInfo.ModTime(), nil
}

func exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func cacheFilePath() string {
	return filepath.Join(Dir(), cacheFile)
}
