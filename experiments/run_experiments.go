package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run experiments/run_experiments.go <experiment_number>")
		fmt.Println("Available experiments:")
		fmt.Println("  01 - Page count API test")
		fmt.Println("  02 - PDF validation API test")
		fmt.Println("  03 - Extract single page test")
		fmt.Println("  04 - Extract multiple pages test")
		fmt.Println("  05 - Extract pages in reverse order test")
		fmt.Println("  06 - Simple merge test")
		fmt.Println("  07 - Individual page merge test")
		fmt.Println("  08 - Complete interleaved pattern test")
		fmt.Println("  17 - CollectFile API availability test")
		fmt.Println("  18 - CollectFile strategy analysis")
		return
	}

	experiment := os.Args[1]
	
	switch experiment {
	case "01":
		experiment01PageCount()
	case "02":
		experiment02Validate()
	case "03":
		experiment03Extract()
	case "04":
		experiment04ExtractMulti()
	case "05":
		experiment05Reverse()
	case "06":
		experiment06Merge()
	case "07":
		experiment07PageMerge()
	case "08":
		experiment08Interleaved()
	case "17":
		experiment17Collect()
	case "18":
		experiment18CollectInterleaved()
	default:
		fmt.Printf("Unknown experiment: %s\n", experiment)
		fmt.Println("Use 'go run experiments/run_experiments.go' to see available experiments")
	}
}
