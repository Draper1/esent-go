package native

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/draper1/esent-go/errors"
	"github.com/draper1/esent-go/types"
)

// JetCreateInstance2 creates a new ESENT instance
func JetCreateInstance2(instanceName, displayName string, grbit types.CreateInstanceGrbit) (*types.JET_INSTANCE, error) {
	if !IsInitialized() || jetCreateInstance2 == nil {
		return nil, errors.JETErrInternalError
	}

	var instance types.JET_INSTANCE
	instanceNamePtr, _ := syscall.BytePtrFromString(instanceName)
	displayNamePtr, _ := syscall.BytePtrFromString(displayName)

	ret, _, _ := jetCreateInstance2.Call(
		uintptr(unsafe.Pointer(&instance)),
		uintptr(unsafe.Pointer(instanceNamePtr)),
		uintptr(unsafe.Pointer(displayNamePtr)),
		uintptr(grbit),
	)

	if ret == 0 {
		return &instance, nil
	}

	return nil, errors.JETErr(int32(ret))
}

// JetInit2 initializes an ESENT instance
func JetInit2(instance *types.JET_INSTANCE, grbit types.InitGrbit) error {
	if !IsInitialized() || jetInit2 == nil {
		return errors.JETErrInternalError
	}

	ret, _, _ := jetInit2.Call(
		uintptr(unsafe.Pointer(&instance.Value)), // Pass reference to instance value
		uintptr(grbit),
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetTerm2 terminates an ESENT instance
func JetTerm2(instance *types.JET_INSTANCE, grbit types.TermGrbit) error {
	if !IsInitialized() || jetTerm2 == nil {
		return errors.JETErrInternalError
	}

	ret, _, _ := jetTerm2.Call(
		instance.Value, // Pass the instance value, not pointer
		uintptr(grbit),
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetBeginSession begins a new session
func JetBeginSession(instance *types.JET_INSTANCE, sesid *types.JET_SESID, username, password string) error {
	if !IsInitialized() || jetBeginSession == nil {
		return errors.JETErrInternalError
	}

	usernamePtr, _ := syscall.BytePtrFromString(username)
	passwordPtr, _ := syscall.BytePtrFromString(password)

	ret, _, _ := jetBeginSession.Call(
		instance.Value,                        // Pass instance value
		uintptr(unsafe.Pointer(&sesid.Value)), // Pass reference to sesid value
		uintptr(unsafe.Pointer(usernamePtr)),
		uintptr(unsafe.Pointer(passwordPtr)),
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetEndSession ends a session
func JetEndSession(sesid *types.JET_SESID, grbit uint32) error {
	if !IsInitialized() || jetEndSession == nil {
		return errors.JETErrInternalError
	}

	ret, _, _ := jetEndSession.Call(
		sesid.Value, // Pass sesid value
		uintptr(grbit),
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetCreateDatabase creates a new database
func JetCreateDatabase(sesid *types.JET_SESID, databaseName, physicalName string, grbit types.CreateDatabaseGrbit) (*types.JET_DBID, error) {
	if !IsInitialized() || jetCreateDatabase == nil {
		return nil, errors.JETErrInternalError
	}

	var dbid types.JET_DBID
	databaseNamePtr, _ := syscall.BytePtrFromString(databaseName)
	physicalNamePtr, _ := syscall.BytePtrFromString(physicalName)

	ret, _, _ := jetCreateDatabase.Call(
		sesid.Value, // Pass sesid value
		uintptr(unsafe.Pointer(databaseNamePtr)),
		uintptr(unsafe.Pointer(physicalNamePtr)),
		uintptr(unsafe.Pointer(&dbid.Value)), // Pass reference to dbid value
		uintptr(grbit),
	)

	if ret == 0 {
		return &dbid, nil
	}

	return nil, errors.JETErr(int32(ret))
}

// JetAttachDatabase attaches an existing database
func JetAttachDatabase(sesid *types.JET_SESID, databaseName string, grbit uint32) error {
	if !IsInitialized() || jetAttachDatabase == nil {
		return errors.JETErrInternalError
	}

	databaseNamePtr, _ := syscall.BytePtrFromString(databaseName)

	ret, _, _ := jetAttachDatabase.Call(
		uintptr(unsafe.Pointer(sesid)),
		uintptr(unsafe.Pointer(databaseNamePtr)),
		uintptr(grbit),
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetDetachDatabase detaches a database
func JetDetachDatabase(sesid *types.JET_SESID, databaseName string) error {
	if !IsInitialized() || jetDetachDatabase == nil {
		return errors.JETErrInternalError
	}

	databaseNamePtr, _ := syscall.BytePtrFromString(databaseName)

	ret, _, _ := jetDetachDatabase.Call(
		uintptr(unsafe.Pointer(sesid)),
		uintptr(unsafe.Pointer(databaseNamePtr)),
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetOpenDatabase opens a database
func JetOpenDatabase(sesid *types.JET_SESID, databaseName string, grbit uint32) (*types.JET_DBID, error) {
	if !IsInitialized() || jetOpenDatabase == nil {
		return nil, errors.JETErrInternalError
	}

	var dbid types.JET_DBID
	databaseNamePtr, _ := syscall.BytePtrFromString(databaseName)

	ret, _, _ := jetOpenDatabase.Call(
		uintptr(unsafe.Pointer(sesid)),
		uintptr(unsafe.Pointer(databaseNamePtr)),
		uintptr(unsafe.Pointer(&dbid)),
		uintptr(grbit),
	)

	if ret == 0 {
		return &dbid, nil
	}

	return nil, errors.JETErr(int32(ret))
}

// JetCloseDatabase closes a database
func JetCloseDatabase(sesid *types.JET_SESID, dbid *types.JET_DBID, grbit uint32) error {
	if !IsInitialized() || jetCloseDatabase == nil {
		return errors.JETErrInternalError
	}

	ret, _, _ := jetCloseDatabase.Call(
		sesid.Value,         // Pass sesid value
		uintptr(dbid.Value), // Pass dbid value
		uintptr(grbit),
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetCreateTable creates a new table
func JetCreateTable(sesid *types.JET_SESID, dbid *types.JET_DBID, tableName string, cPages, density uint32) (*types.JET_TABLEID, error) {
	if !IsInitialized() || jetCreateTable == nil {
		return nil, errors.JETErrInternalError
	}

	var tableid types.JET_TABLEID
	tableNamePtr, _ := syscall.BytePtrFromString(tableName)

	ret, _, _ := jetCreateTable.Call(
		sesid.Value,         // Pass sesid value
		uintptr(dbid.Value), // Pass dbid value
		uintptr(unsafe.Pointer(tableNamePtr)),
		uintptr(cPages),
		uintptr(density),                        // Pass density value
		uintptr(unsafe.Pointer(&tableid.Value)), // Pass reference to tableid value
	)

	if ret == 0 {
		return &tableid, nil
	}

	return nil, errors.JETErr(int32(ret))
}

// JetOpenTable opens an existing table
func JetOpenTable(sesid *types.JET_SESID, dbid *types.JET_DBID, tableName string, grbit uint32) (*types.JET_TABLEID, error) {
	if !IsInitialized() || jetOpenTable == nil {
		return nil, errors.JETErrInternalError
	}

	var tableid types.JET_TABLEID
	tableNamePtr, _ := syscall.BytePtrFromString(tableName)

	ret, _, _ := jetOpenTable.Call(
		uintptr(unsafe.Pointer(sesid)),
		uintptr(unsafe.Pointer(dbid)),
		uintptr(unsafe.Pointer(tableNamePtr)),
		uintptr(unsafe.Pointer(&tableid)),
		uintptr(grbit),
	)

	if ret == 0 {
		return &tableid, nil
	}

	return nil, errors.JETErr(int32(ret))
}

// JetCloseTable closes a table
func JetCloseTable(sesid *types.JET_SESID, tableid *types.JET_TABLEID) error {
	if !IsInitialized() || jetCloseTable == nil {
		return errors.JETErrInternalError
	}

	ret, _, _ := jetCloseTable.Call(
		sesid.Value,   // Pass sesid value
		tableid.Value, // Pass tableid value
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetBeginTransaction begins a transaction
func JetBeginTransaction(sesid *types.JET_SESID, grbit uint32) error {
	if !IsInitialized() || jetBeginTransaction == nil {
		return errors.JETErrInternalError
	}

	ret, _, _ := jetBeginTransaction.Call(
		uintptr(unsafe.Pointer(sesid)),
		uintptr(grbit),
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetCommitTransaction commits a transaction
func JetCommitTransaction(sesid *types.JET_SESID, grbit uint32) error {
	if !IsInitialized() || jetCommitTransaction == nil {
		return errors.JETErrInternalError
	}

	ret, _, _ := jetCommitTransaction.Call(
		uintptr(unsafe.Pointer(sesid)),
		uintptr(grbit),
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetRollback rolls back a transaction
func JetRollback(sesid *types.JET_SESID, grbit uint32) error {
	if !IsInitialized() || jetRollback == nil {
		return errors.JETErrInternalError
	}

	ret, _, _ := jetRollback.Call(
		uintptr(unsafe.Pointer(sesid)),
		uintptr(grbit),
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetBeginExternalBackup begins an external backup
func JetBeginExternalBackup(instance *types.JET_INSTANCE, grbit types.BeginExternalBackupGrbit) error {
	if !IsInitialized() || jetBeginExternalBackup == nil {
		return errors.JETErrInternalError
	}

	ret, _, _ := jetBeginExternalBackup.Call(
		uintptr(unsafe.Pointer(instance)),
		uintptr(grbit),
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetEndExternalBackup ends an external backup
func JetEndExternalBackup(instance *types.JET_INSTANCE, grbit types.EndExternalBackupGrbit) error {
	if !IsInitialized() || jetEndExternalBackup == nil {
		return errors.JETErrInternalError
	}

	ret, _, _ := jetEndExternalBackup.Call(
		uintptr(unsafe.Pointer(instance)),
		uintptr(grbit),
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetGetErrorInfo gets detailed error information
func JetGetErrorInfo(errorCode int32, languageId uint32) (*types.JET_ERRINFOBASIC, error) {
	if !IsInitialized() || jetGetErrorInfo == nil {
		return nil, errors.JETErrInternalError
	}

	var errInfo types.JET_ERRINFOBASIC

	ret, _, _ := jetGetErrorInfo.Call(
		uintptr(errorCode),
		uintptr(languageId),
		uintptr(unsafe.Pointer(&errInfo)),
		uintptr(unsafe.Sizeof(errInfo)),
	)

	if ret == 0 {
		return &errInfo, nil
	}

	return nil, errors.JETErr(int32(ret))
}

// JetGetDatabaseInfo gets database information
func JetGetDatabaseInfo(sesid *types.JET_SESID, dbid *types.JET_DBID, infoLevel uint32, result []byte, maxResultSize uint32, actualResultSize *uint32) error {
	if !IsInitialized() || jetGetDatabaseInfo == nil {
		return errors.JETErrInternalError
	}

	var resultPtr uintptr
	if result != nil {
		resultPtr = uintptr(unsafe.Pointer(&result[0]))
	}

	var actualResultSizePtr uintptr
	if actualResultSize != nil {
		actualResultSizePtr = uintptr(unsafe.Pointer(actualResultSize))
	}

	ret, _, _ := jetGetDatabaseInfo.Call(
		uintptr(unsafe.Pointer(sesid)),
		uintptr(unsafe.Pointer(dbid)),
		uintptr(infoLevel),
		resultPtr,
		uintptr(maxResultSize),
		actualResultSizePtr,
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetGetTableInfo gets table information
func JetGetTableInfo(sesid *types.JET_SESID, tableid *types.JET_TABLEID, infoLevel uint32, result []byte, maxResultSize uint32, actualResultSize *uint32) error {
	if !IsInitialized() || jetGetTableInfo == nil {
		return errors.JETErrInternalError
	}

	var resultPtr uintptr
	if result != nil {
		resultPtr = uintptr(unsafe.Pointer(&result[0]))
	}

	var actualResultSizePtr uintptr
	if actualResultSize != nil {
		actualResultSizePtr = uintptr(unsafe.Pointer(actualResultSize))
	}

	ret, _, _ := jetGetTableInfo.Call(
		uintptr(unsafe.Pointer(sesid)),
		uintptr(unsafe.Pointer(tableid)),
		uintptr(infoLevel),
		resultPtr,
		uintptr(maxResultSize),
		actualResultSizePtr,
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetGetRecordSize gets record size information
func JetGetRecordSize(sesid *types.JET_SESID, tableid *types.JET_TABLEID, grbit uint32) (*types.JET_RECSIZE, error) {
	if !IsInitialized() || jetGetRecordSize == nil {
		return nil, errors.JETErrInternalError
	}

	var recsize types.JET_RECSIZE

	ret, _, _ := jetGetRecordSize.Call(
		uintptr(unsafe.Pointer(sesid)),
		uintptr(unsafe.Pointer(tableid)),
		uintptr(grbit),
		uintptr(unsafe.Pointer(&recsize)),
	)

	if ret == 0 {
		return &recsize, nil
	}

	return nil, errors.JETErr(int32(ret))
}

// JetGetRecordPosition gets record position information
func JetGetRecordPosition(sesid *types.JET_SESID, tableid *types.JET_TABLEID, grbit uint32) (*types.JET_RECPOS, error) {
	if !IsInitialized() || jetGetRecordPosition == nil {
		return nil, errors.JETErrInternalError
	}

	var recpos types.JET_RECPOS

	ret, _, _ := jetGetRecordPosition.Call(
		uintptr(unsafe.Pointer(sesid)),
		uintptr(unsafe.Pointer(tableid)),
		uintptr(grbit),
		uintptr(unsafe.Pointer(&recpos)),
	)

	if ret == 0 {
		return &recpos, nil
	}

	return nil, errors.JETErr(int32(ret))
}

// JetAddColumn adds a column to a table
func JetAddColumn(sesid *types.JET_SESID, tableid *types.JET_TABLEID, columnName string, columndef *types.JET_COLUMNDEF, defaultValue []byte, defaultValueSize uint32) (*types.JET_COLUMNID, error) {
	if !IsInitialized() || jetAddColumn == nil {
		return nil, errors.JETErrInternalError
	}

	var columnid types.JET_COLUMNID
	columnNamePtr, _ := syscall.BytePtrFromString(columnName)

	// Convert to native struct
	nativeColumndef := columndef.GetNativeColumndef()

	var defaultValuePtr uintptr
	if defaultValue != nil && len(defaultValue) > 0 {
		defaultValuePtr = uintptr(unsafe.Pointer(&defaultValue[0]))
	}

	ret, _, _ := jetAddColumn.Call(
		sesid.Value,                               // Pass sesid value
		tableid.Value,                             // Pass tableid value
		uintptr(unsafe.Pointer(columnNamePtr)),    // Pass column name pointer
		uintptr(unsafe.Pointer(&nativeColumndef)), // Pass native column definition pointer
		defaultValuePtr,                           // Pass default value pointer
		uintptr(defaultValueSize),                 // Pass default value size
		uintptr(unsafe.Pointer(&columnid.Value)),  // Pass reference to columnid value
	)

	if ret == 0 {
		return &columnid, nil
	}

	return nil, errors.JETErr(int32(ret))
}

// JetDeleteColumn deletes a column from a table
func JetDeleteColumn(sesid *types.JET_SESID, tableid *types.JET_TABLEID, columnName string) error {
	if !IsInitialized() || jetDeleteColumn == nil {
		return errors.JETErrInternalError
	}

	columnNamePtr, _ := syscall.BytePtrFromString(columnName)

	ret, _, _ := jetDeleteColumn.Call(
		sesid.Value,                            // Pass sesid value
		tableid.Value,                          // Pass tableid value
		uintptr(unsafe.Pointer(columnNamePtr)), // Pass column name pointer
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetGetColumnInfo gets information about a column
func JetGetColumnInfo(sesid *types.JET_SESID, tableid *types.JET_TABLEID, columnName string, grbit uint32) (*types.JET_COLUMNDEF, error) {
	if !IsInitialized() || jetGetColumnInfo == nil {
		return nil, errors.JETErrInternalError
	}

	var columndef types.JET_COLUMNDEF
	columnNamePtr, _ := syscall.BytePtrFromString(columnName)

	ret, _, _ := jetGetColumnInfo.Call(
		sesid.Value,                            // Pass sesid value
		tableid.Value,                          // Pass tableid value
		uintptr(unsafe.Pointer(columnNamePtr)), // Pass column name pointer
		uintptr(grbit),                         // Pass grbit
		uintptr(unsafe.Pointer(&columndef)),    // Pass reference to column definition
	)

	if ret == 0 {
		return &columndef, nil
	}

	return nil, errors.JETErr(int32(ret))
}

// JetCreateIndex creates an index on a table
func JetCreateIndex(sesid *types.JET_SESID, tableid *types.JET_TABLEID, indexName string, grbit uint32, indexKey string, indexKeyLength uint32, density uint32) error {
	if !IsInitialized() || jetCreateIndex == nil {
		return errors.JETErrInternalError
	}

	indexNamePtr, _ := syscall.BytePtrFromString(indexName)
	indexKeyPtr, _ := syscall.BytePtrFromString(indexKey)

	ret, _, _ := jetCreateIndex.Call(
		sesid.Value,                           // Pass sesid value
		tableid.Value,                         // Pass tableid value
		uintptr(unsafe.Pointer(indexNamePtr)), // Pass index name pointer
		uintptr(grbit),                        // Pass grbit
		uintptr(unsafe.Pointer(indexKeyPtr)),  // Pass index key pointer
		uintptr(indexKeyLength),               // Pass index key length
		uintptr(density),                      // Pass density
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetDeleteIndex deletes an index from a table
func JetDeleteIndex(sesid *types.JET_SESID, tableid *types.JET_TABLEID, indexName string) error {
	if !IsInitialized() || jetDeleteIndex == nil {
		return errors.JETErrInternalError
	}

	indexNamePtr, _ := syscall.BytePtrFromString(indexName)

	ret, _, _ := jetDeleteIndex.Call(
		sesid.Value,                           // Pass sesid value
		tableid.Value,                         // Pass tableid value
		uintptr(unsafe.Pointer(indexNamePtr)), // Pass index name pointer
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetPrepareUpdate prepares a record for update
func JetPrepareUpdate(sesid *types.JET_SESID, tableid *types.JET_TABLEID, prep uint32) error {
	if !IsInitialized() || jetPrepareUpdate == nil {
		return errors.JETErrInternalError
	}

	ret, _, _ := jetPrepareUpdate.Call(
		sesid.Value,   // Pass sesid value
		tableid.Value, // Pass tableid value
		uintptr(prep), // Pass prep parameter
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetSetColumn sets a column value
func JetSetColumn(sesid *types.JET_SESID, tableid *types.JET_TABLEID, columnid *types.JET_COLUMNID, data []byte, dataSize uint32, grbit uint32, setinfo *types.JET_SETINFO) error {
	if !IsInitialized() || jetSetColumn == nil {
		return errors.JETErrInternalError
	}

	var dataPtr uintptr
	if data != nil && len(data) > 0 {
		dataPtr = uintptr(unsafe.Pointer(&data[0]))
	}

	var setinfoPtr uintptr
	if setinfo != nil {
		setinfoPtr = uintptr(unsafe.Pointer(setinfo))
	}

	ret, _, _ := jetSetColumn.Call(
		sesid.Value,             // Pass sesid value
		tableid.Value,           // Pass tableid value
		uintptr(columnid.Value), // Pass columnid value
		dataPtr,                 // Pass data pointer
		uintptr(dataSize),       // Pass data size
		uintptr(grbit),          // Pass grbit
		setinfoPtr,              // Pass setinfo pointer
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetRetrieveColumn retrieves a column value
func JetRetrieveColumn(sesid *types.JET_SESID, tableid *types.JET_TABLEID, columnid *types.JET_COLUMNID, data []byte, dataSize uint32, actualDataSize *uint32, grbit uint32, retinfo *types.JET_RETINFO) error {
	if !IsInitialized() || jetRetrieveColumn == nil {
		return errors.JETErrInternalError
	}

	var dataPtr uintptr
	if data != nil && len(data) > 0 {
		dataPtr = uintptr(unsafe.Pointer(&data[0]))
	}

	var actualDataSizePtr uintptr
	if actualDataSize != nil {
		actualDataSizePtr = uintptr(unsafe.Pointer(actualDataSize))
	}

	var retinfoPtr uintptr
	if retinfo != nil {
		retinfoPtr = uintptr(unsafe.Pointer(retinfo))
	}

	ret, _, _ := jetRetrieveColumn.Call(
		sesid.Value,             // Pass sesid value
		tableid.Value,           // Pass tableid value
		uintptr(columnid.Value), // Pass columnid value
		dataPtr,                 // Pass data pointer
		uintptr(dataSize),       // Pass data size
		actualDataSizePtr,       // Pass actual data size pointer
		uintptr(grbit),          // Pass grbit
		retinfoPtr,              // Pass retinfo pointer
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetUpdate updates a record
func JetUpdate(sesid *types.JET_SESID, tableid *types.JET_TABLEID, data []byte, dataSize uint32, actualDataSize *uint32) error {
	if !IsInitialized() || jetUpdate == nil {
		return errors.JETErrInternalError
	}

	var dataPtr uintptr
	if data != nil && len(data) > 0 {
		dataPtr = uintptr(unsafe.Pointer(&data[0]))
	}

	var actualDataSizePtr uintptr
	if actualDataSize != nil {
		actualDataSizePtr = uintptr(unsafe.Pointer(actualDataSize))
	}

	ret, _, _ := jetUpdate.Call(
		sesid.Value,       // Pass sesid value
		tableid.Value,     // Pass tableid value
		dataPtr,           // Pass data pointer
		uintptr(dataSize), // Pass data size
		actualDataSizePtr, // Pass actual data size pointer
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetDelete deletes a record
func JetDelete(sesid *types.JET_SESID, tableid *types.JET_TABLEID) error {
	if !IsInitialized() || jetDelete == nil {
		return errors.JETErrInternalError
	}

	ret, _, _ := jetDelete.Call(
		sesid.Value,   // Pass sesid value
		tableid.Value, // Pass tableid value
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetMove moves the cursor
func JetMove(sesid *types.JET_SESID, tableid *types.JET_TABLEID, cRow int32, grbit uint32) error {
	if !IsInitialized() || jetMove == nil {
		return errors.JETErrInternalError
	}

	ret, _, _ := jetMove.Call(
		sesid.Value,    // Pass sesid value
		tableid.Value,  // Pass tableid value
		uintptr(cRow),  // Pass cRow parameter
		uintptr(grbit), // Pass grbit
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetMakeKey makes a key for seeking
func JetMakeKey(sesid *types.JET_SESID, tableid *types.JET_TABLEID, data []byte, dataSize uint32, grbit uint32) error {
	if !IsInitialized() || jetMakeKey == nil {
		return errors.JETErrInternalError
	}

	var dataPtr uintptr
	if data != nil && len(data) > 0 {
		dataPtr = uintptr(unsafe.Pointer(&data[0]))
	}

	ret, _, _ := jetMakeKey.Call(
		sesid.Value,       // Pass sesid value
		tableid.Value,     // Pass tableid value
		dataPtr,           // Pass data pointer
		uintptr(dataSize), // Pass data size
		uintptr(grbit),    // Pass grbit
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetSeek seeks to a key
func JetSeek(sesid *types.JET_SESID, tableid *types.JET_TABLEID, grbit uint32) error {
	if !IsInitialized() || jetSeek == nil {
		return errors.JETErrInternalError
	}

	ret, _, _ := jetSeek.Call(
		sesid.Value,    // Pass sesid value
		tableid.Value,  // Pass tableid value
		uintptr(grbit), // Pass grbit
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetSetIndexRange sets an index range
func JetSetIndexRange(sesid *types.JET_SESID, tableid *types.JET_TABLEID, grbit uint32) error {
	if !IsInitialized() || jetSetIndexRange == nil {
		return errors.JETErrInternalError
	}

	ret, _, _ := jetSetIndexRange.Call(
		sesid.Value,    // Pass sesid value
		tableid.Value,  // Pass tableid value
		uintptr(grbit), // Pass grbit
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetRestore performs a restore operation
func JetRestore(backupPath string, grbit uint32) error {
	if !IsInitialized() || jetRestore == nil {
		return errors.JETErrInternalError
	}

	backupPathPtr, _ := syscall.BytePtrFromString(backupPath)

	ret, _, _ := jetRestore.Call(
		uintptr(unsafe.Pointer(backupPathPtr)), // Pass backup path pointer
		uintptr(grbit),                         // Pass grbit
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetBeginTransaction3 begins a transaction with user ID
func JetBeginTransaction3(sesid *types.JET_SESID, userTransactionId int64, grbit uint32) error {
	if !IsInitialized() || jetBeginTransaction3 == nil {
		return errors.JETErrInternalError
	}

	ret, _, _ := jetBeginTransaction3.Call(
		sesid.Value,                // Pass sesid value
		uintptr(userTransactionId), // Pass user transaction ID
		uintptr(grbit),             // Pass grbit
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetCommitTransaction2 commits a transaction with durable commit
func JetCommitTransaction2(sesid *types.JET_SESID, grbit uint32, durableCommit time.Duration) (*types.JET_COMMIT_ID, error) {
	if !IsInitialized() || jetCommitTransaction2 == nil {
		return nil, errors.JETErrInternalError
	}

	var commitId types.JET_COMMIT_ID

	ret, _, _ := jetCommitTransaction2.Call(
		sesid.Value,                           // Pass sesid value
		uintptr(grbit),                        // Pass grbit
		uintptr(durableCommit.Milliseconds()), // Pass durable commit timeout
		uintptr(unsafe.Pointer(&commitId)),    // Pass reference to commit ID
	)

	if ret == 0 {
		return &commitId, nil
	}

	return nil, errors.JETErr(int32(ret))
}

// JetBackup performs a backup operation
func JetBackup(backupPath string, grbit uint32, statusCallback interface{}) error {
	if !IsInitialized() || jetBackup == nil {
		return errors.JETErrInternalError
	}

	backupPathPtr, _ := syscall.BytePtrFromString(backupPath)

	var callbackPtr uintptr
	if statusCallback != nil {
		// For now, we'll pass 0 as callback - this would need proper callback handling
		callbackPtr = 0
	}

	ret, _, _ := jetBackup.Call(
		uintptr(unsafe.Pointer(backupPathPtr)), // Pass backup path pointer
		uintptr(grbit),                         // Pass grbit
		callbackPtr,                            // Pass callback pointer
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetSetSystemParameter sets a system parameter for an ESENT instance
func JetSetSystemParameter(instance *types.JET_INSTANCE, sesid *types.JET_SESID, param types.JET_param, paramValue uintptr, paramString string) error {
	if !IsInitialized() || jetSetSystemParameter == nil {
		return errors.JETErrInternalError
	}

	var instancePtr uintptr
	if instance != nil {
		instancePtr = instance.Value
	}

	var sesidPtr uintptr
	if sesid != nil {
		sesidPtr = sesid.Value
	}

	var paramStringPtr uintptr
	if paramString != "" {
		paramStringBytes, _ := syscall.BytePtrFromString(paramString)
		paramStringPtr = uintptr(unsafe.Pointer(paramStringBytes))
	}

	ret, _, _ := jetSetSystemParameter.Call(
		instancePtr,    // Pass instance value
		sesidPtr,       // Pass sesid value
		uintptr(param), // Pass parameter ID
		paramValue,     // Pass parameter value
		paramStringPtr, // Pass parameter string
	)

	if ret == 0 {
		return nil
	}

	return errors.JETErr(int32(ret))
}

// JetGetDatabaseFileInfo gets information about a database file (int32 version)
func JetGetDatabaseFileInfo(databaseName string, infoLevel types.JET_DbInfo) (int32, error) {
	if !IsInitialized() {
		return 0, errors.JETErrInternalError
	}

	if jetGetDatabaseFileInfo == nil {
		return 0, errors.JETErrInternalError
	}

	databaseNamePtr, _ := syscall.BytePtrFromString(databaseName)
	var result int32

	ret, _, _ := jetGetDatabaseFileInfo.Call(
		uintptr(unsafe.Pointer(databaseNamePtr)), // Pass database name pointer
		uintptr(unsafe.Pointer(&result)),         // Pass result pointer (out parameter)
		uintptr(unsafe.Sizeof(result)),           // Pass size of result
		uintptr(infoLevel),                       // Pass info level
	)

	if ret == 0 {
		return result, nil
	}

	return 0, errors.JETErr(int32(ret))
}

// JetGetDatabaseFileInfoLong gets information about a database file (int64 version)
func JetGetDatabaseFileInfoLong(databaseName string, infoLevel types.JET_DbInfo) (int64, error) {
	if !IsInitialized() || jetGetDatabaseFileInfo == nil {
		return 0, errors.JETErrInternalError
	}

	databaseNamePtr, _ := syscall.BytePtrFromString(databaseName)
	var result int64

	ret, _, _ := jetGetDatabaseFileInfo.Call(
		uintptr(unsafe.Pointer(databaseNamePtr)), // Pass database name pointer
		uintptr(unsafe.Pointer(&result)),         // Pass result pointer
		uintptr(unsafe.Sizeof(result)),           // Pass size of result
		uintptr(infoLevel),                       // Pass info level
	)

	if ret == 0 {
		return result, nil
	}

	return 0, errors.JETErr(int32(ret))
}

// JetGetDatabaseFileInfoMisc gets miscellaneous information about a database file
func JetGetDatabaseFileInfoMisc(databaseName string, infoLevel types.JET_DbInfo) (*types.JET_DBINFOMISC, error) {
	if !IsInitialized() || jetGetDatabaseFileInfo == nil {
		return nil, errors.JETErrInternalError
	}

	databaseNamePtr, _ := syscall.BytePtrFromString(databaseName)
	var result types.JET_DBINFOMISC

	ret, _, _ := jetGetDatabaseFileInfo.Call(
		uintptr(unsafe.Pointer(databaseNamePtr)), // Pass database name pointer
		uintptr(unsafe.Pointer(&result)),         // Pass result pointer
		uintptr(unsafe.Sizeof(result)),           // Pass size of result
		uintptr(infoLevel),                       // Pass info level
	)

	if ret == 0 {
		return &result, nil
	}

	return nil, errors.JETErr(int32(ret))
}
