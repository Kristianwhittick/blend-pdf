# PDFcpu API Experiments

## Overview
Testing individual pdfcpu API functions to understand their behavior and requirements.

## Test Files
- **Input**: Doc_A.pdf (3 pages: A1, A2, A3)
- **Input**: Doc_B.pdf (3 pages: B3, B2, B1)
- **Output**: test_XX.pdf files in output/ directory

## API Testing Procedures

### Running Individual API Tests
```bash
cd /home/kris/scan/blendpdfgo
go run experiments/run_experiments.go 01    # Test page counting
go run experiments/run_experiments.go 02    # Test PDF validation
go run experiments/run_experiments.go 03    # Test single page extraction
go run experiments/run_experiments.go 04    # Test multiple page extraction
go run experiments/run_experiments.go 05    # Test reverse page extraction
go run experiments/run_experiments.go 06    # Test simple merge
go run experiments/run_experiments.go 07    # Test individual page merge
go run experiments/run_experiments.go 08    # Test complete interleaved pattern
go run experiments/run_experiments.go 17    # Test CollectFile API availability
go run experiments/run_experiments.go 18    # Test CollectFile strategy
```

### Expected API Test Results
- **Test 01**: Both PDFs should have 3 pages
- **Test 02**: Both PDFs should validate successfully
- **Test 03**: Extract A1 to single page PDF
- **Test 04**: Extract A1, A2 to two-page PDF
- **Test 05**: Extract pages in reverse order
- **Test 06**: Simple concatenation: A1, A2, A3, M, 9, f
- **Test 07**: Partial interleaved: A1, f, A2, 9
- **Test 08**: Full interleaved: A1, f, A2, 9, A3, M ✅
- **Test 17**: CollectFile function availability and signature confirmation ✅
- **Test 18**: CollectFile strategy analysis and implementation approach ✅

## Notes
- Each test creates a standalone Go program
- Results verified using pdftotext
- Detailed experiment results documented in api_knowledge.md
- Memory processing research documented in memory_processing_research.md
