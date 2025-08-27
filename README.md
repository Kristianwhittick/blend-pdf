# BlendPDFGo

A tool for merging PDF files with special handling for double-sided scanning.

## Features

- Move single PDF files to output directory
- Merge two PDF files (first file + reversed second file)
- Interactive command-line interface
- Automatic file organization (archive, output, error folders)

## Installation

```bash
# Clone the repository
git clone https://github.com/kris/blendpdfgo.git
cd blendpdfgo

# Build the application
go build
```

## Usage

```bash
# Show help
./blendpdfgo -h

# Run in verbose mode
./blendpdfgo -V

# Watch specific folder
./blendpdfgo /path/to/pdfs

# Watch current directory
./blendpdfgo
```

## Interactive Options

- `S` - Move a single PDF file to the output directory
- `M` - Merge two PDF files (first file + reversed second file)
- `H` - Show help information
- `V` - Toggle verbose mode
- `Q` - Quit the program

## Dependencies

- [pdfcpu](https://github.com/pdfcpu/pdfcpu) - PDF processor