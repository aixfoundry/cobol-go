# COBOL Parser for Go

A Go COBOL parser built on [ANTLR v4](https://github.com/antlr/antlr4) and the [COBOL85 grammar](https://github.com/antlr/grammars-v4/tree/master/cobol85). Parses COBOL source into a Protobuf AST with support for copybook expansion, REPLACE preprocessing, and symbol table resolution.

## Source Format Support

| Format | Description | Indicator Column |
|--------|-------------|-----------------|
| **FIXED** | Standard ANSI/IBM reference format (80-column) | 7 |
| **TANDEM** | HP Tandem format | 1 |
| **VARIABLE** | Variable-length format | 7 |
| **FREE** | COBOL 2002/2014 free format (no fixed columns) | 1 |

### Format Details

**FIXED** (default):
```
Cols 1-6   : sequence area
Col  7     : indicator field (*, /, -, D, $, space)
Cols 8-12  : Area A
Cols 13-72 : Area B
Cols 73-80 : comment area
```

**FREE**:
```
Col  1     : indicator field
Cols 2+    : source text (no fixed column boundaries)
```

## Project Structure

```
├── asg/            AST generation, visitor, symbol table
│   ├── conv/       COBOL parse tree → Protobuf AST conversion
│   ├── symbol/     Symbol table building and name resolution
│   └── visitor/    AST visitors (data, environment, procedure)
├── cmd/            CLI tools
│   ├── tree/       Parse tree output (.tree, .error)
│   ├── preprocess/ Copybook expansion + REPLACE (.preprocessed)
│   ├── dobf/       Name obfuscation (stdin → JSON)
│   └── func/       Paragraph/function extraction (stdin → JSON)
├── constant/       COBOL keywords and character constants
├── copybook/       Copybook file resolution
├── document/       Preprocessor, REPLACE handling, line combining
├── format/         Source format and dialect definitions
├── gen/            ANTLR-generated parser code
├── line/           Line parsing, linked-line processing, continuations
├── options/        Functional options pattern
├── pb/             Protobuf definitions
└── test/           Integration tests against NIST corpus
```

## Quick Start

### Build

```bash
make build        # Builds all CLI tools to bin/
go build ./...    # Verify all packages compile
```

### CLI Tools

All tools accept `-format` flag (`FIXED`, `TANDEM`, `VARIABLE`, `FREE`):

**tree** — Parse COBOL files and output parse trees:
```bash
bin/tree -format=FREE -path=/path/to/sources -suffix=.cbl -copyPath=/path/to/copybooks
# Produces: file.cbl.tree (parse tree), file.cbl.error (syntax errors)
```

**preprocess** — Expand copybooks and REPLACE directives:
```bash
bin/preprocess -format=FIXED -path=/path/to/sources
# Produces: file.cbl.preprocessed
```

**dobf** — Obfuscate identifiers (stdin → JSON):
```bash
cat program.cbl | bin/dobf
# Outputs: {"text":"...","vars":{"VAR_1":"WS-COUNT"},...}
```

**func** — Extract paragraph/function structure (stdin → JSON):
```bash
cat program.cbl | bin/func
# Outputs: [{"name":"0100-INIT","text":"...","tokens":[...]}]
```

## API Usage

```go
import (
    "github.com/aixfoundry/cobol-go/asg"
    "github.com/aixfoundry/cobol-go/format"
    "github.com/aixfoundry/cobol-go/options"
)

func main() {
    opts := options.NewOptions().
        SetFormat(format.FREE).
        AddCopyBookDirectory("./copybooks")

    // Parse and get full AST
    program, err := asg.AnalyzeFile("program.cbl", opts)
    if err != nil {
        panic(err)
    }

    // Access parsed compilation units
    for _, cu := range program.GetCompilationUnits() {
        for _, pu := range cu.GetProgramUnits() {
            // Work with program units...
        }
    }

    // Build symbol table for name resolution
    table := symbol.Build(program)
    entries := table.DataEntries["WS-COUNT"]
}
```

### Functional Options

```go
opts := options.NewOptions().
    SetFormat(format.FIXED).           // Source format
    SetDialect(format.OSVS).           // COBOL dialect
    AddCopyBookDirectory("./lib").     // Copybook search paths
    AddCopyBookFile("./macros.cpy").   // Individual copybooks
    AddCopyBookExtension("cpy")        // File extension to search
```

## Code Generation

ANTLR grammar (requires Java):
```bash
go generate ./...
```

Protobuf stubs:
```bash
make proto
```

## References

- [ANTLR v4 COBOL85 Grammar](https://github.com/antlr/grammars-v4/tree/master/cobol85)
- [ANTLR v4](https://github.com/antlr/antlr4)
- [ProLeap COBOL Parser (Java)](https://github.com/uwol/proleap-cobol-parser)
- [ProLeap COBOL Transform](https://github.com/proleap/proleap-cobol)

## License

Apache-2.0 — see [LICENSE](LICENSE)
