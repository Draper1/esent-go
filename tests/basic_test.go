package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/andyd/esent-go/api"
	"github.com/andyd/esent-go/errors"
	"github.com/andyd/esent-go/types"
)

func TestBasicTypes(t *testing.T) {
	// Test column types
	if types.ColtypNil != 0 {
		t.Errorf("Expected ColtypNil to be 0, got %d", types.ColtypNil)
	}

	if types.ColtypBit != 1 {
		t.Errorf("Expected ColtypBit to be 1, got %d", types.ColtypBit)
	}

	if types.ColtypText != 10 {
		t.Errorf("Expected ColtypText to be 10, got %d", types.ColtypText)
	}
}

func TestErrorCodes(t *testing.T) {
	// Test error codes
	if errors.JETErrSuccess != 0 {
		t.Errorf("Expected JETErrSuccess to be 0, got %d", errors.JETErrSuccess)
	}

	if errors.JETErrFileClose != -102 {
		t.Errorf("Expected JETErrFileClose to be -102, got %d", errors.JETErrFileClose)
	}

	// Test error string conversion
	if errors.JETErrSuccess.String() != "Success" {
		t.Errorf("Expected JETErrSuccess.String() to return 'Success', got '%s'", errors.JETErrSuccess.String())
	}
}

func TestGrbits(t *testing.T) {
	// Test grbits
	if types.CreateInstanceGrbitNone != 0 {
		t.Errorf("Expected CreateInstanceGrbitNone to be 0, got %d", types.CreateInstanceGrbitNone)
	}

	if types.TermGrbitComplete != 1 {
		t.Errorf("Expected TermGrbitComplete to be 1, got %d", types.TermGrbitComplete)
	}

	if types.CreateDatabaseGrbitOverwriteExisting != 0x200 {
		t.Errorf("Expected CreateDatabaseGrbitOverwriteExisting to be 0x200, got %d", types.CreateDatabaseGrbitOverwriteExisting)
	}
}

func TestStructures(t *testing.T) {
	// Test basic structure creation
	sesid := &types.JET_SESID{Value: 0x12345678}
	if sesid.Value != 0x12345678 {
		t.Errorf("Expected sesid.Value to be 0x12345678, got %x", sesid.Value)
	}

	dbid := &types.JET_DBID{Value: 0x87654321}
	if dbid.Value != 0x87654321 {
		t.Errorf("Expected dbid.Value to be 0x87654321, got %x", dbid.Value)
	}

	// Create a proper column ID for the column definition
	columnid := &types.JET_COLUMNID{Value: 0x11111111}

	columndef := &types.JET_COLUMNDEF{
		Columnid:   *columnid,
		Coltyp:     types.ColtypText,
		Grbit:      0,
		CbMax:      255,
		Cp:         0,
		ColumnName: "TestColumn",
	}

	if columndef.ColumnName != "TestColumn" {
		t.Errorf("Expected ColumnName to be 'TestColumn', got '%s'", columndef.ColumnName)
	}

	if columndef.Coltyp != types.ColtypText {
		t.Errorf("Expected Coltyp to be ColtypText, got %d", columndef.Coltyp)
	}
}

func TestApiInterface(t *testing.T) {
	t.Log("Starting TestApiInterface...")

	// Test API creation
	jetApi := api.NewJetApi()
	if jetApi == nil {
		t.Fatal("Expected NewJetApi() to return non-nil")
	}
	t.Log("✓ API created")

	// Test instance creation
	t.Log("Creating instance...")
	instance, err := jetApi.JetCreateInstance2("test", "Test Instance", types.CreateInstanceGrbitNone)
	if err != nil {
		t.Errorf("Expected no error from JetCreateInstance2, got %v", err)
	}
	if instance == nil {
		t.Fatal("Expected instance to be non-nil")
	}
	t.Log("✓ Instance created")

	// Initialize instance
	t.Log("Initializing instance...")
	err = jetApi.JetInit2(instance, types.InitGrbitNone)
	if err != nil {
		t.Errorf("Expected no error from JetInit2, got %v", err)
	}
	t.Log("✓ Instance initialized")

	// Test session creation
	t.Log("Creating session...")
	sesid := &types.JET_SESID{}
	err = jetApi.JetBeginSession(instance, sesid, "", "")
	if err != nil {
		t.Errorf("Expected no error from JetBeginSession, got %v", err)
	}
	t.Log("✓ Session created")

	// Test database creation with unique name
	t.Log("Creating database...")
	dbName := fmt.Sprintf("test_%d.db", time.Now().UnixNano())
	t.Logf("Database name: %s", dbName)
	dbid, err := jetApi.JetCreateDatabase(sesid, dbName, dbName, types.CreateDatabaseGrbitNone)
	if err != nil {
		t.Errorf("Expected no error from JetCreateDatabase, got %v", err)
	}
	if dbid == nil {
		t.Fatal("Expected dbid to be non-nil")
	}
	t.Log("✓ Database created")

	// Test table creation
	t.Log("Creating table...")
	tableid, err := jetApi.JetCreateTable(sesid, dbid, "TestTable", 0, 0)
	if err != nil {
		t.Errorf("Expected no error from JetCreateTable, got %v", err)
	}
	if tableid == nil {
		t.Fatal("Expected tableid to be non-nil")
	}
	t.Log("✓ Table created")

	// Test column creation - column ID should be empty, will be filled by JetAddColumn
	columndef := &types.JET_COLUMNDEF{
		Columnid:   types.JET_COLUMNID{Value: 0}, // Initialize to empty, will be filled by JetAddColumn
		Coltyp:     types.ColtypText,
		Grbit:      0,
		CbMax:      255,
		Cp:         0,
		ColumnName: "TestColumn",
	}

	createdColumnid, err := jetApi.JetAddColumn(sesid, tableid, "TestColumn", columndef, nil, 0)
	if err != nil {
		t.Errorf("Expected no error from JetAddColumn, got %v", err)
	}
	if createdColumnid == nil {
		t.Fatal("Expected columnid to be non-nil")
	}

	// Test cleanup
	err = jetApi.JetCloseTable(sesid, tableid)
	if err != nil {
		t.Errorf("Expected no error from JetCloseTable, got %v", err)
	}

	err = jetApi.JetCloseDatabase(sesid, dbid, 0)
	if err != nil {
		t.Errorf("Expected no error from JetCloseDatabase, got %v", err)
	}

	err = jetApi.JetEndSession(sesid, 0)
	if err != nil {
		t.Errorf("Expected no error from JetEndSession, got %v", err)
	}

	err = jetApi.JetTerm2(instance, types.TermGrbitNone)
	if err != nil {
		t.Errorf("Expected no error from JetTerm2, got %v", err)
	}

	// Clean up database files
	cleanupDatabaseFiles(t)
}

// cleanupDatabaseFiles removes all database files from the tests directory and project root
func cleanupDatabaseFiles(t *testing.T) {
	t.Log("Cleaning up database files...")

	// Get the current directory (tests folder)
	testsDir, err := os.Getwd()
	if err != nil {
		t.Logf("Warning: Could not get current directory: %v", err)
		return
	}

	// Get the parent directory (project root)
	projectRoot := filepath.Dir(testsDir)

	// Directories to clean up
	dirsToClean := []string{testsDir, projectRoot}

	// Database file extensions to clean up
	extensions := []string{".db", ".jfm", ".log", ".chk", ".jrs"}

	// Clean up files with each extension in each directory
	for _, dir := range dirsToClean {
		for _, ext := range extensions {
			pattern := filepath.Join(dir, "*"+ext)
			matches, err := filepath.Glob(pattern)
			if err != nil {
				t.Logf("Warning: Could not glob pattern %s: %v", pattern, err)
				continue
			}

			for _, match := range matches {
				err := os.Remove(match)
				if err != nil {
					t.Logf("Warning: Could not remove file %s: %v", match, err)
				} else {
					relPath, _ := filepath.Rel(projectRoot, match)
					t.Logf("Removed: %s", relPath)
				}
			}
		}
	}

	t.Log("Database cleanup completed")
}
