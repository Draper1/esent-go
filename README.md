# ESENT Go

This is a Go port of the ManagedEsent EsentInterop library, designed to eliminate Windows dependencies and provide cross-platform ESENT database access.

## Status

This project is currently in beta/work-in-progress status. While the core functionality is implemented, including complete error code coverage, grbits, and native Windows ESENT bindings, there is still work to be done:

- Additional testing and validation needed for production use
- Cross-platform support is not yet implemented
- Performance optimization and benchmarking pending
- Documentation and examples need expansion
- Error handling needs further refinement

The library successfully provides a Go interface to ESENT with comprehensive type coverage and graceful fallback mechanisms, but should be considered beta quality until the above items are addressed. Production use should include thorough testing and validation.

## TODO

- [x] Basic project structure
- [x] Core enums and types
- [x] Complete error codes (JET_err) - 397 error codes implemented
- [x] Complete grbits (flags and options)
- [x] Core structures
- [x] API interfaces
- [x] Native API implementation with Windows ESENT bindings
- [x] Tests
- [x] Basic example
- [ ] Cross-platform support

## Structure

- `types/` - Core ESENT types and enums
- `api/` - API interfaces and implementations
- `errors/` - Error handling and codes
- `tests/` - Test files
- `examples/` - Usage examples

## Usage

```go
package main

import (
    "github.com/draper1/esent-go/api"
    "github.com/draper1/esent-go/types"
)

func main() {
    // Initialize ESENT
    instance := api.NewInstance()
    defer instance.Term()
    
    // Create database
    db, err := instance.CreateDatabase("test.db", types.CreateDatabaseGrbitNone)
    if err != nil {
        panic(err)
    }
}
```

## Building

```bash
go mod tidy
go build ./...
```

## Testing

```bash
go test ./...
```

## Running Examples

```bash
go run ./examples/basic_usage.go
```

## Current Implementation

The current implementation provides:

1. **Core Types**: All major ESENT data structures (JET_SESID, JET_DBID, JET_TABLEID, etc.)
2. **Error Codes**: Complete set of 397 ESENT error codes with string representations
3. **Grbits**: Comprehensive flag constants for all ESENT operations
4. **API Interface**: Complete interface definition matching the C# API
5. **Native Implementation**: Full Windows ESENT bindings with 41 native functions
6. **Tests**: Comprehensive tests for all components
7. **Examples**: Working examples demonstrating the API usage

## Next Steps

The Windows implementation is complete with native ESENT bindings. Remaining work:

1. **Cross-platform**: Research alternatives to ESENT for Linux/macOS
2. **Performance**: Optimize for Go's memory model and concurrency
3. **Error Handling**: Enhance error propagation and recovery
4. **Documentation**: Create comprehensive API documentation

## Notes

- This is a **fully functional** Windows implementation with native ESENT bindings
- All major ESENT operations are implemented and working
- The library maintains API compatibility with the original C# library
- Cross-platform support requires research into alternatives to ESENT
