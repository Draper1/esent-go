package native

import (
	"fmt"
	"syscall"
)

// ESENT.dll function pointers
var (
	esentDll *syscall.DLL

	// Core functions
	jetCreateInstance2     *syscall.Proc
	jetInit2               *syscall.Proc
	jetTerm2               *syscall.Proc
	jetBeginSession        *syscall.Proc
	jetEndSession          *syscall.Proc
	jetCreateDatabase      *syscall.Proc
	jetAttachDatabase      *syscall.Proc
	jetDetachDatabase      *syscall.Proc
	jetOpenDatabase        *syscall.Proc
	jetCloseDatabase       *syscall.Proc
	jetCreateTable         *syscall.Proc
	jetOpenTable           *syscall.Proc
	jetCloseTable          *syscall.Proc
	jetAddColumn           *syscall.Proc
	jetDeleteColumn        *syscall.Proc
	jetGetColumnInfo       *syscall.Proc
	jetCreateIndex         *syscall.Proc
	jetDeleteIndex         *syscall.Proc
	jetPrepareUpdate       *syscall.Proc
	jetSetColumn           *syscall.Proc
	jetRetrieveColumn      *syscall.Proc
	jetUpdate              *syscall.Proc
	jetDelete              *syscall.Proc
	jetMove                *syscall.Proc
	jetMakeKey             *syscall.Proc
	jetSeek                *syscall.Proc
	jetSetIndexRange       *syscall.Proc
	jetBeginTransaction    *syscall.Proc
	jetBeginTransaction3   *syscall.Proc
	jetCommitTransaction   *syscall.Proc
	jetCommitTransaction2  *syscall.Proc
	jetRollback            *syscall.Proc
	jetBackup              *syscall.Proc
	jetRestore             *syscall.Proc
	jetBeginExternalBackup *syscall.Proc
	jetEndExternalBackup   *syscall.Proc
	jetGetErrorInfo        *syscall.Proc
	jetGetDatabaseInfo     *syscall.Proc
	jetGetTableInfo        *syscall.Proc
	jetGetRecordSize       *syscall.Proc
	jetGetRecordPosition   *syscall.Proc
)

// safeFindProc safely finds a procedure, returning nil if not found
func safeFindProc(dll *syscall.DLL, name string) *syscall.Proc {
	proc, err := dll.FindProc(name)
	if err != nil {
		return nil
	}
	return proc
}

// Initialize loads the ESENT.dll and sets up function pointers
func Initialize() error {
	var err error

	// Load ESENT.dll
	esentDll, err = syscall.LoadDLL("esent.dll")
	if err != nil {
		return err
	}

	// Get function pointers - only load the ones that exist
	jetCreateInstance2 = safeFindProc(esentDll, "JetCreateInstance2")
	jetInit2 = safeFindProc(esentDll, "JetInit2")
	jetTerm2 = safeFindProc(esentDll, "JetTerm2")
	jetBeginSession = safeFindProc(esentDll, "JetBeginSession")
	jetEndSession = safeFindProc(esentDll, "JetEndSession")
	jetCreateDatabase = safeFindProc(esentDll, "JetCreateDatabase")
	jetAttachDatabase = safeFindProc(esentDll, "JetAttachDatabase")
	jetDetachDatabase = safeFindProc(esentDll, "JetDetachDatabase")
	jetOpenDatabase = safeFindProc(esentDll, "JetOpenDatabase")
	jetCloseDatabase = safeFindProc(esentDll, "JetCloseDatabase")
	jetCreateTable = safeFindProc(esentDll, "JetCreateTable")
	jetOpenTable = safeFindProc(esentDll, "JetOpenTable")
	jetCloseTable = safeFindProc(esentDll, "JetCloseTable")
	jetAddColumn = safeFindProc(esentDll, "JetAddColumn")
	jetDeleteColumn = safeFindProc(esentDll, "JetDeleteColumn")
	jetGetColumnInfo = safeFindProc(esentDll, "JetGetColumnInfo")
	jetCreateIndex = safeFindProc(esentDll, "JetCreateIndex")
	jetDeleteIndex = safeFindProc(esentDll, "JetDeleteIndex")
	jetPrepareUpdate = safeFindProc(esentDll, "JetPrepareUpdate")
	jetSetColumn = safeFindProc(esentDll, "JetSetColumn")
	jetRetrieveColumn = safeFindProc(esentDll, "JetRetrieveColumn")
	jetUpdate = safeFindProc(esentDll, "JetUpdate")
	jetDelete = safeFindProc(esentDll, "JetDelete")
	jetMove = safeFindProc(esentDll, "JetMove")
	jetMakeKey = safeFindProc(esentDll, "JetMakeKey")
	jetSeek = safeFindProc(esentDll, "JetSeek")
	jetSetIndexRange = safeFindProc(esentDll, "JetSetIndexRange")
	jetBeginTransaction = safeFindProc(esentDll, "JetBeginTransaction")
	jetBeginTransaction3 = safeFindProc(esentDll, "JetBeginTransaction3")
	jetCommitTransaction = safeFindProc(esentDll, "JetCommitTransaction")
	jetCommitTransaction2 = safeFindProc(esentDll, "JetCommitTransaction2")
	jetRollback = safeFindProc(esentDll, "JetRollback")
	jetBackup = safeFindProc(esentDll, "JetBackup")
	jetRestore = safeFindProc(esentDll, "JetRestore")
	jetBeginExternalBackup = safeFindProc(esentDll, "JetBeginExternalBackup")
	jetEndExternalBackup = safeFindProc(esentDll, "JetEndExternalBackup")
	jetGetErrorInfo = safeFindProc(esentDll, "JetGetErrorInfo")
	jetGetDatabaseInfo = safeFindProc(esentDll, "JetGetDatabaseInfo")
	jetGetTableInfo = safeFindProc(esentDll, "JetGetTableInfo")
	jetGetRecordSize = safeFindProc(esentDll, "JetGetRecordSize")
	jetGetRecordPosition = safeFindProc(esentDll, "JetGetRecordPosition")

	// Check if we have at least the basic functions
	if jetCreateInstance2 == nil || jetInit2 == nil || jetTerm2 == nil {
		return fmt.Errorf("failed to load essential ESENT functions")
	}

	// Log any missing optional functions
	if jetGetErrorInfo == nil {
		fmt.Printf("Warning: JetGetErrorInfo not available in this ESENT version\n")
	}

	return nil
}

// IsInitialized returns true if the native layer has been initialized
func IsInitialized() bool {
	return esentDll != nil
}

// Cleanup releases the ESENT.dll handle
func Cleanup() {
	if esentDll != nil {
		esentDll.Release()
		esentDll = nil
	}
}
