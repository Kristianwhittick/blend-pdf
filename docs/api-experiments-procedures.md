# PDFcpu API Experiments

## Overview
Testing individual pdfcpu API functions to understand their behaviour and requirements.

## Test Files
- **Input**: Doc_A.pdf (3 pages: A1, A2, A3)
- **Input**: Doc_B.pdf (3 pages: B3, B2, B1)
- **Output**: test_XX.pdf files in output/ directory

## API Testing Procedures

### Running Individual API Tests
```bash
cd /home/kris/workspace/scan/blend-pdf
go run experiments/run_experiments.go 01    # Test page counting
go run experiments/run_experiments.go 02    # Test PDF validation
go run experiments/run_experiments.go 03    # Test single page extraction
go run experiments/run_experiments.go 04    # Test multiple page extraction
go run experiments/run_experiments.go 05    # Test reverse page extraction
go run experiments/run_experiments.go 06    # Test simple merge
go run experiments/run_experiments.go 07    # Test individual page merge
go run experiments/run_experiments.go 08    # Test complete interleaved pattern
go run experiments/run_experiments.go 09    # Test memory context API
go run experiments/run_experiments.go 10    # Test memory page extraction (simple)
go run experiments/run_experiments.go 11    # Test memory page extraction
go run experiments/run_experiments.go 12    # Test memory context merging
go run experiments/run_experiments.go 13    # Test memory bytes processing
go run experiments/run_experiments.go 14    # Test API exploration
go run experiments/run_experiments.go 15    # Test working memory approach
go run experiments/run_experiments.go 16    # Test hybrid memory approach
go run experiments/run_experiments.go 17    # Test final memory approach
go run experiments/run_experiments.go 18    # Test CollectFile API availability
go run experiments/run_experiments.go 19    # Test CollectFile strategy
go run experiments/run_experiments.go 20    # Test basic zip merge
go run experiments/run_experiments.go 21    # Test CollectFile + zip merge
go run experiments/run_experiments.go 22    # Test complete zip flow validation
go run experiments/run_experiments.go 23    # Test low-level extract pages API
go run experiments/run_experiments.go 24    # Test stream-based Trim function
go run experiments/run_experiments.go 25    # Test raw merge function
go run experiments/run_experiments.go 26    # Test complete in-memory workflow
go run experiments/run_experiments.go 27    # Test package investigation
go run experiments/run_experiments.go 28    # Test import path testing
go run experiments/run_experiments.go 29    # Test discovered stream-based APIs
```

### Experiment Structure
- **Experiments 01-22**: Located in individual folders (`exp01/`, `exp02/`, etc.) to prevent Go linting warnings
- **Experiments 23-29**: Located in root experiments directory as single files
- **Runner**: `experiments/run_experiments.go` handles both folder and file structures

### Expected API Test Results
- **Test 01**: Both PDFs should have 3 pages
- **Test 02**: Both PDFs should validate successfully
- **Test 03**: Extract A1 to single page PDF
- **Test 04**: Extract A1, A2 to two-page PDF
- **Test 05**: Extract pages in reverse order
- **Test 06**: Simple concatenation: A1, A2, A3, M, 9, f
- **Test 07**: Partial interleaved: A1, f, A2, 9
- **Test 08**: Full interleaved: A1, f, A2, 9, A3, M ✅
- **Test 18**: CollectFile function availability and signature confirmation ✅
- **Test 19**: CollectFile strategy analysis and implementation approach ✅
- **Test 20**: Basic zip merge functionality - interleaved pattern A1, M, A2, 9, A3, f ✅
- **Test 21**: Complete solution - CollectFile reversal + zip merge = A1, f, A2, 9, A3, M ✅
- **Test 22**: Full workflow validation with content verification ✅

## Notes
- Each experiment is in its own subfolder (exp01/, exp02/, etc.)
- All experiments run through the unified runner: `go run experiments/run_experiments.go <number>`
- Results verified using pdftotext
- Detailed experiment results documented in api-knowledge.md
- Memory processing research documented in memory-processing-research.md
