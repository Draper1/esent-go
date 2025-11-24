# ESENT Go Documentation

Welcome to the ESENT Go documentation. This library provides a Go interface to the Extensible Storage Engine (ESENT) on Windows, allowing Go applications to interact with ESE databases without relying on C# or ManagedEsent.

## Overview

ESENT Go is a native Go implementation that binds directly to `esent.dll` on Windows. It aims to provide a safe, idiomatic Go API while maintaining the performance and capability of the underlying engine.

## Features

- **Native Go Implementation**: No CGO required (uses `syscall` for DLL loading).
- **Complete Error Codes**: Comprehensive mapping of ESENT error codes to Go errors.
- **Type Safety**: Strong typing for ESENT handles (Session ID, Database ID, Table ID, etc.).
- **Grbits Support**: Complete implementation of ESENT options and flags.
- **Transaction Support**: Full support for ACID transactions.

## Installation

```bash
go get github.com/draper1/esent-go
```

## Basic Usage

Here is a simple example of how to create an instance, session, and database.

```go
package main

import (
    "fmt"
    "github.com/draper1/esent-go/api"
    "github.com/draper1/esent-go/types"
)

func main() {
    // 1. Create API Instance
    jetApi := api.NewJetApi()

    // 2. Create and Initialize ESENT Instance
    instance, err := jetApi.JetCreateInstance2("test_inst", "Test Instance", types.CreateInstanceGrbitNone)
    if err != nil {
        panic(err)
    }
    
    err = jetApi.JetInit2(instance, types.InitGrbitNone)
    if err != nil {
        panic(err)
    }
    defer jetApi.JetTerm2(instance, types.TermGrbitNone)
    
    // 3. Begin Session
    sesid := &types.JET_SESID{}
    err = jetApi.JetBeginSession(instance, sesid, "", "")
    if err != nil {
        panic(err)
    }
    defer jetApi.JetEndSession(sesid, 0)

    // 4. Create Database
    dbid, err := jetApi.JetCreateDatabase(sesid, "test.db", "test.db", types.CreateDatabaseGrbitNone)
    if err != nil {
        panic(err)
    }
    defer jetApi.JetCloseDatabase(sesid, dbid, 0)
    
    fmt.Println("Database created successfully!")
}
```

## Next Steps

- [Architecture Overview](architecture.md)
- [API Reference](api-reference.md)

