package utils

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// GetBackupNameByTime 根据时间获取备份新命名
func GetBackupNameByTime(name string, layout string) string {
	dir := filepath.Dir(name)
	filename := filepath.Base(name)
	ext := filepath.Ext(filename)
	prefix := strings.TrimSuffix(filename, ext)
	nowTm := time.Now()

	return filepath.Join(dir, fmt.Sprintf("%s_%s%s", prefix, nowTm.Format(layout), ext))
}

func TimeFromFileName(name, prefix, ext, layout string) (time.Time, error) {
	// 前缀
	if !strings.HasPrefix(name, prefix) {
		return time.Time{}, errors.New("invalid file name not prefix")
	}
	// 后缀
	if !strings.HasSuffix(name, ext) {
		return time.Time{}, errors.New("invalid file name not suffix")
	}

	ts := name[len(prefix) : len(name)-len(ext)]
	return time.Parse(layout, ts)
}

// CompressFileByGzip 压缩文件为gzip
func CompressFileByGzip(src, dst string) (err error) {
	defer func() {
		if err != nil {
			// 如果中途出错 删除目标文件
			os.Remove(dst)
		}
	}()

	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	gzFile, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, fileInfo.Mode())
	if err != nil {
		return err
	}
	defer gzFile.Close()

	gz := gzip.NewWriter(gzFile)
	// 拷贝到gz写入器
	if _, err = io.Copy(gz, file); err != nil {
		return err
	}
	if err = gz.Flush(); err != nil {
		return err
	}
	if err = gz.Close(); err != nil {
		return err
	}
	if err = gzFile.Close(); err != nil {
		return err
	}
	if err = file.Close(); err != nil {
		return err
	}
	if err = os.Remove(src); err != nil {
		return err
	}

	return nil
}
