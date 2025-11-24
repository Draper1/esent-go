# Architecture

ESENT Go is designed with a layered architecture to separate the low-level Windows API calls from the high-level Go interface.

## Layer Structure

### 1. Native Layer (`native/`)
This is the lowest layer of the library. It is responsible for:
- Loading `esent.dll` dynamically using `syscall`.
- Defining function pointers for exported ESENT functions.
- converting Go types to C-compatible types (uintptr, etc.) where necessary for the system calls.

This layer effectively replaces CGO, allowing the library to compile and run on Windows without a C compiler, relying on the system's built-in DLLs.

### 2. Types Layer (`types/`)
This package defines the data structures, constants, and types used throughout the library.
- **Handles**: `JET_SESID`, `JET_DBID`, `JET_TABLEID`, etc. are defined as strong types to prevent mixing up different handle types.
- **Structures**: Go struct representations of ESENT C structures (e.g., `JET_COLUMNDEF`, `JET_INDEXCREATE`).
- **Grbits**: Bitmasks and flags for various operations (e.g., `CreateDatabaseGrbit`).

### 3. API Layer (`api/`)
This is the main entry point for consumers of the library.
- **Interfaces**: Defines the `IJetApi` interface which abstracts the ESENT operations.
- **Implementation**: calls into the `native` package to execute the actual ESENT functions.

## Design Decisions

### No CGO
By avoiding CGO, the library remains purely Go-native (albeit OS-specific). This simplifies the build process and cross-compilation (though the library primarily targets Windows).

### Error Handling
ESENT uses return codes for everything. This library converts non-zero return codes into Go `error` types, specifically wrapping `JET_err` codes so they can be checked programmatically.

### Resource Management
The API is designed to follow Go idioms. However, since ESENT relies heavily on sessions and instances that must be explicitly closed, the library encourages the use of `defer` patterns for `JetTerm`, `JetEndSession`, etc.

