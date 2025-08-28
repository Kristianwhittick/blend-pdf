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
go run experiments/experiment01_pagecount.go    # Test page counting
go run experiments/experiment02_validate.go     # Test PDF validation
go run experiments/experiment03_extract.go      # Test single page extraction
go run experiments/experiment04_extract_multi.go # Test multiple page extraction
go run experiments/experiment05_reverse.go      # Test reverse page extraction
go run experiments/experiment06_merge.go        # Test simple merge
go run experiments/experiment07_page_merge.go   # Test individual page merge
go run experiments/experiment08_interleaved.go  # Test complete interleaved pattern
go run experiments/experiment17_collect.go      # Test CollectFile API availability
go run experiments/experiment18_collect_interleaved.go # Test CollectFile strategy
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
