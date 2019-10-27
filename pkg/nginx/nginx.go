package nginx

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/goodwithtech/deckoder/extractor"

	"github.com/goodwithtech/deckoder/analyzer"

	"archive/tar"
	"path/filepath"

	deckodertypes "github.com/goodwithtech/deckoder/types"
)

var ErrNoConf = errors.New("no nginx.conf files")
var accessLogPrefix = "access_log"

func ScanImage(imageName string, option deckodertypes.DockerOption) (err error) {
	ctx := context.Background()
	files, err := analyzer.Analyze(ctx, imageName, confFilterFunc(), option)
	if err != nil {
		return fmt.Errorf("failed to analyze image: %w", err)
	}
	return checkConfFile(files)
}

func checkConfFile(files extractor.FileMap) error {
	var existFile bool
	for filePath, fileData := range files {
		if checkFilePath(filePath) {
			existFile = true
			if err := checkLogFormat(fileData.Body); err != nil {
				return fmt.Errorf("%s: %w", filePath, err)
			}
		}
	}
	if !existFile {
		return ErrNoConf
	}
	return nil
}

// split by space : configuration line
func splitBySpace(command string) []string {
	splitted := strings.Split(command, " ")
	cmds := []string{}
	for _, cmd := range splitted {
		trimmed := strings.TrimSpace(cmd)
		if trimmed != "" {
			cmds = append(cmds, trimmed)
		}
	}
	return cmds
}

// is ltsv format?
func checkLogFormat(body []byte) error {
	scanner := bufio.NewScanner(bytes.NewBuffer(body))
	for scanner.Scan() {
		line := scanner.Text()
		cmds := splitBySpace(line)
		// detect format :  "access_log /path/to/log format;"
		if len(cmds) >= 3 && cmds[0] == accessLogPrefix {
			// use ltsv log format
			if !strings.Contains(cmds[2], "ltsv") {
				return fmt.Errorf(`Expect log format contains "ltsv" but %q`, cmds[2])
			}
		}
	}
	return nil
}

// is nginx configuration files?
func checkFilePath(path string) bool {
	return strings.Contains(path, "nginx") && strings.HasSuffix(path, "conf")
}

// for deckoder func
func confFilterFunc() deckodertypes.FilterFunc {
	return func(h *tar.Header) (bool, error) {
		filePath := filepath.Clean(h.Name)
		if checkFilePath(filePath) {
			return true, nil
		}
		return false, nil
	}
}
