package api

import (
	"time"

	"github.com/draper1/esent-go/errors"
	"github.com/draper1/esent-go/types"
)

// IJetApi defines the interface for ESENT operations
type IJetApi interface {
	// Instance management
	JetCreateInstance2(instanceName string, displayName string, grbit types.CreateInstanceGrbit) (*types.JET_INSTANCE, error)
	JetInit2(instance *types.JET_INSTANCE, grbit types.InitGrbit) error
	JetTerm2(instance *types.JET_INSTANCE, grbit types.TermGrbit) error

	// Session management
	JetBeginSession(instance *types.JET_INSTANCE, sesid *types.JET_SESID, username string, password string) error
	JetEndSession(sesid *types.JET_SESID, grbit uint32) error

	// Database operations
	JetCreateDatabase(sesid *types.JET_SESID, databaseName string, physicalName string, grbit types.CreateDatabaseGrbit) (*types.JET_DBID, error)
	JetAttachDatabase(sesid *types.JET_SESID, databaseName string, grbit uint32) error
	JetDetachDatabase(sesid *types.JET_SESID, databaseName string) error
	JetOpenDatabase(sesid *types.JET_SESID, databaseName string, grbit uint32) (*types.JET_DBID, error)
	JetCloseDatabase(sesid *types.JET_SESID, dbid *types.JET_DBID, grbit uint32) error

	// Table operations
	JetCreateTable(sesid *types.JET_SESID, dbid *types.JET_DBID, tableName string, cPages uint32, density uint32) (*types.JET_TABLEID, error)
	JetOpenTable(sesid *types.JET_SESID, dbid *types.JET_DBID, tableName string, grbit uint32) (*types.JET_TABLEID, error)
	JetCloseTable(sesid *types.JET_SESID, tableid *types.JET_TABLEID) error

	// Column operations
	JetAddColumn(sesid *types.JET_SESID, tableid *types.JET_TABLEID, columnName string, columndef *types.JET_COLUMNDEF, defaultValue []byte, defaultValueSize uint32) (*types.JET_COLUMNID, error)
	JetDeleteColumn(sesid *types.JET_SESID, tableid *types.JET_TABLEID, columnName string) error
	JetGetColumnInfo(sesid *types.JET_SESID, tableid *types.JET_TABLEID, columnName string, grbit uint32) (*types.JET_COLUMNDEF, error)

	// Index operations
	JetCreateIndex(sesid *types.JET_SESID, tableid *types.JET_TABLEID, indexName string, grbit uint32, indexKey string, indexKeyLength uint32, density uint32) error
	JetDeleteIndex(sesid *types.JET_SESID, tableid *types.JET_TABLEID, indexName string) error

	// Record operations
	JetPrepareUpdate(sesid *types.JET_SESID, tableid *types.JET_TABLEID, prep uint32) error
	JetSetColumn(sesid *types.JET_SESID, tableid *types.JET_TABLEID, columnid *types.JET_COLUMNID, data []byte, dataSize uint32, grbit uint32, setinfo *types.JET_SETINFO) error
	JetRetrieveColumn(sesid *types.JET_SESID, tableid *types.JET_TABLEID, columnid *types.JET_COLUMNID, data []byte, dataSize uint32, actualDataSize *uint32, grbit uint32, retinfo *types.JET_RETINFO) error
	JetUpdate(sesid *types.JET_SESID, tableid *types.JET_TABLEID, data []byte, dataSize uint32, actualDataSize *uint32) error
	JetDelete(sesid *types.JET_SESID, tableid *types.JET_TABLEID) error

	// Cursor operations
	JetMove(sesid *types.JET_SESID, tableid *types.JET_TABLEID, cRow int32, grbit uint32) error
	JetMakeKey(sesid *types.JET_SESID, tableid *types.JET_TABLEID, data []byte, dataSize uint32, grbit uint32) error
	JetSeek(sesid *types.JET_SESID, tableid *types.JET_TABLEID, grbit uint32) error
	JetSetIndexRange(sesid *types.JET_SESID, tableid *types.JET_TABLEID, grbit uint32) error

	// Transaction operations
	JetBeginTransaction(sesid *types.JET_SESID, grbit uint32) error
	JetBeginTransaction3(sesid *types.JET_SESID, userTransactionId int64, grbit uint32) error
	JetCommitTransaction(sesid *types.JET_SESID, grbit uint32) error
	JetCommitTransaction2(sesid *types.JET_SESID, grbit uint32, durableCommit time.Duration) (*types.JET_COMMIT_ID, error)
	JetRollback(sesid *types.JET_SESID, grbit uint32) error

	// Backup and recovery
	JetBackup(backupPath string, grbit uint32, statusCallback interface{}) error
	JetRestore(backupPath string, grbit uint32) error

	// Error handling
	JetGetErrorInfo(error errors.JETErr, grbit uint32) (*types.JET_ERRINFOBASIC, error)

	// Database information
	JetGetDatabaseInfo(sesid *types.JET_SESID, dbid *types.JET_DBID, infoLevel uint32) (interface{}, error)
	JetGetTableInfo(sesid *types.JET_SESID, tableid *types.JET_TABLEID, infoLevel uint32) (interface{}, error)

	// Utility functions
	JetGetRecordSize(sesid *types.JET_SESID, tableid *types.JET_TABLEID) (*types.JET_RECSIZE, error)
	JetGetRecordPosition(sesid *types.JET_SESID, tableid *types.JET_TABLEID) (*types.JET_RECPOS, error)
}
