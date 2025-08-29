// Copyright 2025 Kristian Whittick
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ui

import (
	"path/filepath"
	"sort"
)

// FileOpsBridge bridges the UI with existing file operations
type FileOpsBridge struct {
	watchDir   string
	archiveDir string
	outputDir  string
	errorDir   string
	
	// Function pointers to existing operations
	findPDFFilesFunc      func(string) ([]string, error)
	countPDFFilesFunc     func(string) int
	getFileSizeFunc       func(string) string
	processSingleFileFunc func() error
	processMergeFilesFunc func() error
}

// NewFileOpsBridge creates a new bridge with function pointers
func NewFileOpsBridge(watchDir, archiveDir, outputDir, errorDir string) *FileOpsBridge {
	return &FileOpsBridge{
		watchDir:   watchDir,
		archiveDir: archiveDir,
		outputDir:  outputDir,
		errorDir:   errorDir,
	}
}

// SetFunctions sets the function pointers to existing operations
func (b *FileOpsBridge) SetFunctions(
	findPDFFiles func(string) ([]string, error),
	countPDFFiles func(string) int,
	getFileSize func(string) string,
	processSingleFile func() error,
	processMergeFiles func() error,
) {
	b.findPDFFilesFunc = findPDFFiles
	b.countPDFFilesFunc = countPDFFiles
	b.getFileSizeFunc = getFileSize
	b.processSingleFileFunc = processSingleFile
	b.processMergeFilesFunc = processMergeFiles
}

// FindPDFFiles implements FileOperations interface
func (b *FileOpsBridge) FindPDFFiles(dir string) ([]string, error) {
	if b.findPDFFilesFunc != nil {
		files, err := b.findPDFFilesFunc(dir)
		if err != nil {
			return nil, err
		}
		
		// Sort files alphabetically
		sort.Strings(files)
		return files, nil
	}
	return []string{}, nil
}

// CountPDFFiles implements FileOperations interface
func (b *FileOpsBridge) CountPDFFiles(dir string) int {
	if b.countPDFFilesFunc != nil {
		return b.countPDFFilesFunc(dir)
	}
	return 0
}

// GetHumanReadableSize implements FileOperations interface
func (b *FileOpsBridge) GetHumanReadableSize(filename string) string {
	if b.getFileSizeFunc != nil {
		// If filename is not absolute, make it relative to watch dir
		if !filepath.IsAbs(filename) {
			filename = filepath.Join(b.watchDir, filename)
		}
		return b.getFileSizeFunc(filename)
	}
	return "0B"
}

// ProcessSingleFile implements FileOperations interface
func (b *FileOpsBridge) ProcessSingleFile() error {
	if b.processSingleFileFunc != nil {
		return b.processSingleFileFunc()
	}
	return nil
}

// ProcessMergeFiles implements FileOperations interface
func (b *FileOpsBridge) ProcessMergeFiles() error {
	if b.processMergeFilesFunc != nil {
		return b.processMergeFilesFunc()
	}
	return nil
}
