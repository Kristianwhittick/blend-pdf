package main

import (
	"fmt"
	"os"
	"os/exec"
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
		fmt.Println("  09 - Memory context API test")
		fmt.Println("  10 - Memory page extraction (simple) test")
		fmt.Println("  11 - Memory page extraction test")
		fmt.Println("  12 - Memory context merging test")
		fmt.Println("  13 - Memory bytes processing test")
		fmt.Println("  14 - API exploration test")
		fmt.Println("  15 - Working memory approach test")
		fmt.Println("  16 - Hybrid memory approach test")
		fmt.Println("  17 - Final memory approach test")
		fmt.Println("  18 - CollectFile API availability test")
		fmt.Println("  19 - CollectFile strategy analysis")
		return
	}

	experiment := os.Args[1]
	
	var cmd *exec.Cmd
	switch experiment {
	case "01":
		cmd = exec.Command("go", "run", "experiments/exp01/experiment01_pagecount.go")
	case "02":
		cmd = exec.Command("go", "run", "experiments/exp02/experiment02_validate.go")
	case "03":
		cmd = exec.Command("go", "run", "experiments/exp03/experiment03_extract.go")
	case "04":
		cmd = exec.Command("go", "run", "experiments/exp04/experiment04_extract_multi.go")
	case "05":
		cmd = exec.Command("go", "run", "experiments/exp05/experiment05_reverse.go")
	case "06":
		cmd = exec.Command("go", "run", "experiments/exp06/experiment06_merge.go")
	case "07":
		cmd = exec.Command("go", "run", "experiments/exp07/experiment07_page_merge.go")
	case "08":
		cmd = exec.Command("go", "run", "experiments/exp08/experiment08_interleaved.go")
	case "09":
		cmd = exec.Command("go", "run", "experiments/exp09/experiment09_memory_context.go")
	case "10":
		cmd = exec.Command("go", "run", "experiments/exp10/experiment10_memory_extract_simple.go")
	case "11":
		cmd = exec.Command("go", "run", "experiments/exp11/experiment11_memory_extract.go")
	case "12":
		cmd = exec.Command("go", "run", "experiments/exp12/experiment12_memory_merge.go")
	case "13":
		cmd = exec.Command("go", "run", "experiments/exp13/experiment13_memory_bytes.go")
	case "14":
		cmd = exec.Command("go", "run", "experiments/exp14/experiment14_api_exploration.go")
	case "15":
		cmd = exec.Command("go", "run", "experiments/exp15/experiment15_working_memory.go")
	case "16":
		cmd = exec.Command("go", "run", "experiments/exp16/experiment16_hybrid_memory.go")
	case "17":
		cmd = exec.Command("go", "run", "experiments/exp17/experiment17_final_memory_approach.go")
	case "18":
		cmd = exec.Command("go", "run", "experiments/exp18/experiment18_collect.go")
	case "19":
		cmd = exec.Command("go", "run", "experiments/exp19/experiment19_collect_interleaved.go")
	case "20":
		cmd = exec.Command("go", "run", "experiments/exp20/experiment20_zip_merge_basic.go")
	case "21":
		cmd = exec.Command("go", "run", "experiments/exp21/experiment21_collect_zip_merge.go")
	case "22":
		cmd = exec.Command("go", "run", "experiments/exp22/experiment22_complete_zip_flow.go")
	default:
		fmt.Printf("Unknown experiment: %s\n", experiment)
		return
	}
	
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running experiment: %v\n", err)
	}
}
