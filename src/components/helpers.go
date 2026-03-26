package dataparser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetConfigDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(homeDir, ".data-parser")
}

func GetConfigFile() string {
	return filepath.Join(GetConfigDir(), "config.json")
}

func GetDefaultConfig() map[string]interface{} {
	return map[string]interface{}{
		"input_dir":        "/path/to/input/dir",
		"output_dir":       "/path/to/output/dir",
		"file_format":      "json",
		"delimiter":        ",",
		"header_row":       false,
		"skip_rows":        0,
		"max_rows":         0,
		"date_format":      "",
		"time_format":      "",
	}
}

func ValidateConfig(config map[string]interface{}) error {
	if _, ok := config["input_dir"]; !ok {
		return fmt.Errorf("missing 'input_dir' key in config")
	}
	if _, ok := config["output_dir"]; !ok {
		return fmt.Errorf("missing 'output_dir' key in config")
	}
	if _, ok := config["file_format"]; !ok {
		return fmt.Errorf("missing 'file_format' key in config")
	}
	if _, ok := config["delimiter"]; !ok {
		return fmt.Errorf("missing 'delimiter' key in config")
	}
	if _, ok := config["header_row"]; !ok {
		return fmt.Errorf("missing 'header_row' key in config")
	}
	if _, ok := config["skip_rows"]; !ok {
		return fmt.Errorf("missing 'skip_rows' key in config")
	}
	if _, ok := config["max_rows"]; !ok {
		return fmt.Errorf("missing 'max_rows' key in config")
	}
	return nil
}

func GetFileExtension(fileName string) string {
	ext := strings.ToLower(filepath.Ext(fileName))
	ext = strings.TrimPrefix(ext, ".")
	return ext
}