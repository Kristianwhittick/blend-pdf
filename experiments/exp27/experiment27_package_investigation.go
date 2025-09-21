package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("=== Experiment 27: PDFcpu Package Investigation ===")
	
	// Investigate the pdfcpu package structure to find the stream-based functions
	// mentioned by the maintainer
	
	fmt.Println("Looking for pdfcpu source code in Go modules...")
	
	// Get the Go module cache path
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		goPath = filepath.Join(os.Getenv("HOME"), "go")
	}
	
	// Look for pdfcpu in module cache
	modCache := filepath.Join(goPath, "pkg", "mod", "github.com", "pdfcpu")
	
	fmt.Printf("Checking module cache: %s\n", modCache)
	
	// Check if directory exists
	if _, err := os.Stat(modCache); os.IsNotExist(err) {
		fmt.Println("Module cache not found, trying alternative approach...")
		
		// Alternative: use go list to find module location
		fmt.Println("Using go list to find pdfcpu module...")
		
		// For now, let's focus on what we can discover from the API
		fmt.Println("\n=== Available pdfcpu Packages ===")
		fmt.Println("Based on imports that work:")
		fmt.Println("- github.com/pdfcpu/pdfcpu/pkg/api")
		fmt.Println("- github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model")
		fmt.Println("- github.com/pdfcpu/pdfcpu/pkg/pdfcpu")
		
		fmt.Println("\n=== Functions Mentioned by Maintainer ===")
		fmt.Println("1. extract_test.go: func TestExtractPagesLowLevel(t *testing.T)")
		fmt.Println("2. merge_test.go: TestMergeRaw(t *testing.T)")
		fmt.Println("3. trim.go: func Trim(rs io.ReadSeeker, w io.Writer, selectedPages []string, conf *model.Configuration) error")
		
		fmt.Println("\n=== Investigation Strategy ===")
		fmt.Println("1. Try different package imports for Trim function")
		fmt.Println("2. Look for extract and merge functions in different packages")
		fmt.Println("3. Check if functions are in internal packages")
		
		return
	}
	
	// If we found the module cache, search for the mentioned functions
	fmt.Printf("Found pdfcpu module cache: %s\n", modCache)
	
	// Walk through the pdfcpu source to find the functions
	err := filepath.Walk(modCache, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Only look at Go files
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		
		// Skip test files for now (we'll check them separately)
		if strings.HasSuffix(path, "_test.go") {
			return nil
		}
		
		// Parse the Go file
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return nil // Skip files that can't be parsed
		}
		
		// Look for functions named Trim, Extract, or Merge
		ast.Inspect(node, func(n ast.Node) bool {
			if fn, ok := n.(*ast.FuncDecl); ok {
				if fn.Name != nil {
					name := fn.Name.Name
					if strings.Contains(strings.ToLower(name), "trim") ||
						strings.Contains(strings.ToLower(name), "extract") ||
						strings.Contains(strings.ToLower(name), "merge") {
						
						relPath, _ := filepath.Rel(modCache, path)
						fmt.Printf("Found function %s in %s\n", name, relPath)
						
						// Print function signature if it looks relevant
						if name == "Trim" || name == "Extract" || name == "Merge" {
							fmt.Printf("  Package: %s\n", node.Name.Name)
							if fn.Type.Params != nil {
								fmt.Printf("  Parameters: %d\n", len(fn.Type.Params.List))
							}
						}
					}
				}
			}
			return true
		})
		
		return nil
	})
	
	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
	}
	
	fmt.Println("\n=== Next Steps ===")
	fmt.Println("Based on findings, create targeted experiments to test the discovered functions")
}
