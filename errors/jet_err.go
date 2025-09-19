package errors

// JETErr represents ESENT error codes
type JETErr int32

const (
	// JETErrSuccess - Successful operation.
	JETErrSuccess JETErr = 0

	// JETErrFileCompressed - read/write access is not supported on compressed files
	JETErrFileCompressed JETErr = -4005
	// JETErrFileIOFail - instructs the JET_ABORTRETRYFAILCALLBACK caller to fail the specified I/O
	JETErrFileIOFail JETErr = -4004
	// JETErrFileIORetry - instructs the JET_ABORTRETRYFAILCALLBACK caller to retry the specified I/O
	JETErrFileIORetry JETErr = -4003
	// JETErrFileIOAbort - instructs the JET_ABORTRETRYFAILCALLBACK caller to abort the specified I/O
	JETErrFileIOAbort JETErr = -4002
	// JETErrFileIOBeyondEOF - a read was issued to a location beyond EOF (writes will expand the file)
	JETErrFileIOBeyondEOF JETErr = -4001
	// JETErrFileIOSparse - an I/O was issued to a location that was sparse
	JETErrFileIOSparse JETErr = -4000
	// JETErrLSNotSet - Attempted to retrieve Local Storage from an object which didn't have it set
	JETErrLSNotSet JETErr = -3002
	// JETErrLSAlreadySet - Attempted to set Local Storage for an object which already had it set
	JETErrLSAlreadySet JETErr = -3001
	// JETErrLSCallbackNotSpecified - Attempted to use Local Storage without a callback function being specified
	JETErrLSCallbackNotSpecified JETErr = -3000
	// JETErrInvalidLogDataSequence - Some how the log data provided got out of sequence with the current state of the instance
	JETErrInvalidLogDataSequence JETErr = -2601
	// JETErrTestInjectionNotSupported - Test injection not supported
	JETErrTestInjectionNotSupported JETErr = -2502
	// JETErrTooManyTestInjections - Internal test injection limit hit
	JETErrTooManyTestInjections JETErr = -2501
	// JETErrOSSnapshotInvalidSnapId - invalid JET_OSSNAPID
	JETErrOSSnapshotInvalidSnapId JETErr = -2404
	// JETErrOSSnapshotNotAllowed - OS Shadow copy not allowed (backup or recovery in progress)
	JETErrOSSnapshotNotAllowed JETErr = -2403
	// JETErrOSSnapshotTimeOut - OS Shadow copy ended with time-out
	JETErrOSSnapshotTimeOut JETErr = -2402
	// JETErrOSSnapshotInvalidSequence - OS Shadow copy API used in an invalid sequence
	JETErrOSSnapshotInvalidSequence JETErr = -2401
	// JETErrSpaceHintsInvalid - An element of the JET space hints structure was not correct or actionable.
	JETErrSpaceHintsInvalid JETErr = -2103
	// JETErrCallbackNotResolved - A callback function could not be found
	JETErrCallbackNotResolved JETErr = -2102
	// JETErrCallbackFailed - A callback failed
	JETErrCallbackFailed JETErr = -2101
	// JETErrDatabaseAlreadyRunningMaintenance - The operation did not complete successfully because the database is already running maintenance on specified database
	JETErrDatabaseAlreadyRunningMaintenance JETErr = -2004
	// JETErrFlushMapUnrecoverable - The persisted flush map cannot be reconstructed.
	JETErrFlushMapUnrecoverable JETErr = -1920
	// JETErrFlushMapDatabaseMismatch - The persisted flush map and the database do not match.
	JETErrFlushMapDatabaseMismatch JETErr = -1919
	// JETErrFlushMapVersionUnsupported - The version of the persisted flush map is not supported by this version of the engine.
	JETErrFlushMapVersionUnsupported JETErr = -1918
	// JETErrRollbackError - error during rollback
	JETErrRollbackError JETErr = -1917
	// JETErrOneDatabasePerSession - Just one open user database per session is allowed (JET_paramOneDatabasePerSession)
	JETErrOneDatabasePerSession JETErr = -1916
	// JETErrRecordFormatConversionFailed - Internal error during dynamic record format conversion
	JETErrRecordFormatConversionFailed JETErr = -1915
	// JETErrSessionInUse - Tried to terminate session in use
	JETErrSessionInUse JETErr = -1914
	// JETErrSessionContextNotSetByThisThread - Tried to reset session context, but current thread did not originally set the session context
	JETErrSessionContextNotSetByThisThread JETErr = -1913
	// JETErrSessionContextAlreadySet - Specified session already has a session context set
	JETErrSessionContextAlreadySet JETErr = -1912
	// JETErrEntryPointNotFound - An entry point in a DLL we require could not be found
	JETErrEntryPointNotFound JETErr = -1911
	// JETErrSessionSharingViolation - Multiple threads are using the same session
	JETErrSessionSharingViolation JETErr = -1910
	// JETErrTooManySplits - Infinite split
	JETErrTooManySplits JETErr = -1909
	// JETErrAccessDenied - Access denied
	JETErrAccessDenied JETErr = -1907
	// JETErrInvalidOperation - Invalid operation
	JETErrInvalidOperation JETErr = -1906
	// JETErrLogCorrupted - Logs could not be interpreted
	JETErrLogCorrupted JETErr = -1852
	// JETErrAfterInitialization - Cannot Restore after init.
	JETErrAfterInitialization JETErr = -1850
	// JETErrFileAlreadyExists - File already exists
	JETErrFileAlreadyExists JETErr = -1814
	// JETErrFileInvalidType - Invalid file type
	JETErrFileInvalidType JETErr = -1812
	// JETErrFileNotFound - File not found
	JETErrFileNotFound JETErr = -1811
	// JETErrPermissionDenied - Permission denied
	JETErrPermissionDenied JETErr = -1809
	// JETErrDiskFull - No space left on disk
	JETErrDiskFull JETErr = -1808
	// JETErrTooManyAttachedDatabases - Too many open databases
	JETErrTooManyAttachedDatabases JETErr = -1805
	// JETErrTempFileOpenError - Temp file could not be opened
	JETErrTempFileOpenError JETErr = -1803
	// JETErrInvalidOnSort - Invalid operation on Sort
	JETErrInvalidOnSort JETErr = -1702
	// JETErrTooManySorts - Too many sort processes
	JETErrTooManySorts JETErr = -1701
	// JETErrEncryptionBadItag - Cannot encrypt tagged columns with itag>1
	JETErrEncryptionBadItag JETErr = -1623
	// JETErrDecryptionFailed - Data could not be decrypted
	JETErrDecryptionFailed JETErr = -1622
	// JETErrUpdateMustVersion - No version updates only for uncommitted tables
	JETErrUpdateMustVersion JETErr = -1621
	// JETErrDecompressionFailed - Internal error: data could not be decompressed
	JETErrDecompressionFailed JETErr = -1620
	// JETErrLanguageNotSupported - Windows installation does not support language
	JETErrLanguageNotSupported JETErr = -1619
	// JETErrDataHasChanged - Data has changed, operation aborted
	JETErrDataHasChanged JETErr = -1611
	// JETErrUpdateNotPrepared - No call to JetPrepareUpdate
	JETErrUpdateNotPrepared JETErr = -1609
	// JETErrKeyNotMade - No call to JetMakeKey
	JETErrKeyNotMade JETErr = -1608
	// JETErrAlreadyPrepared - Attempted to update record when record update was already in progress
	JETErrAlreadyPrepared JETErr = -1607
	// JETErrKeyDuplicate - Illegal duplicate key
	JETErrKeyDuplicate JETErr = -1605
	// JETErrRecordPrimaryChanged - Primary key may not change
	JETErrRecordPrimaryChanged JETErr = -1604
	// JETErrNoCurrentRecord - Currency not on a record
	JETErrNoCurrentRecord JETErr = -1603
	// JETErrRecordNoCopy - No working buffer
	JETErrRecordNoCopy JETErr = -1602
	// JETErrRecordNotFound - The key was not found
	JETErrRecordNotFound JETErr = -1601
	// JETErrColumnNoEncryptionKey - Cannot retrieve/set encrypted column without an encryption key
	JETErrColumnNoEncryptionKey JETErr = -1540
	// JETErrColumnCannotBeCompressed - Only JET_coltypLongText and JET_coltypLongBinary columns can be compressed
	JETErrColumnCannotBeCompressed JETErr = -1538
	// JETErrInvalidPlaceholderColumn - Tried to convert column to a primary index placeholder, but column doesn't meet necessary criteria
	JETErrInvalidPlaceholderColumn JETErr = -1530
	// JETErrDerivedColumnCorruption - Invalid column in derived table
	JETErrDerivedColumnCorruption JETErr = -1529
	// JETErrMultiValuedDuplicateAfterTruncation - Duplicate detected on a unique multi-valued column after data was normalized, and normalizing truncated the data before comparison
	JETErrMultiValuedDuplicateAfterTruncation JETErr = -1528
	// JETErrLVCorrupted - Corruption encountered in long-value tree
	JETErrLVCorrupted JETErr = -1526
	// JETErrMultiValuedDuplicate - Duplicate detected on a unique multi-valued column
	JETErrMultiValuedDuplicate JETErr = -1525
	// JETErrDefaultValueTooBig - Default value exceeds maximum size
	JETErrDefaultValueTooBig JETErr = -1524
	// JETErrCannotBeTagged - AutoIncrement and Version cannot be tagged
	JETErrCannotBeTagged JETErr = -1521
	// JETErrColumnInRelationship - Cannot delete, column participates in relationship
	JETErrColumnInRelationship JETErr = -1519
	// JETErrBadItagSequence - Bad itagSequence for tagged column
	JETErrBadItagSequence JETErr = -1518
	// JETErrBadColumnId - Column Id Incorrect
	JETErrBadColumnId JETErr = -1517
	// JETErrKeyIsMade - The key is completely made
	JETErrKeyIsMade JETErr = -1516
	// JETErrNoCurrentIndex - Invalid w/o a current index
	JETErrNoCurrentIndex JETErr = -1515
	// JETErrTaggedNotNULL - No non-NULL tagged columns
	JETErrTaggedNotNULL JETErr = -1514
	// JETErrInvalidColumnType - Invalid column data type
	JETErrInvalidColumnType JETErr = -1511
	// JETErrColumnRedundant - Second autoincrement or version column
	JETErrColumnRedundant JETErr = -1510
	// JETErrMultiValuedColumnMustBeTagged - Attempted to create a multi-valued column, but column was not Tagged
	JETErrMultiValuedColumnMustBeTagged JETErr = -1509
	// JETErrColumnDuplicate - Field is already defined
	JETErrColumnDuplicate JETErr = -1508
	// JETErrColumnNotFound - No such column
	JETErrColumnNotFound JETErr = -1507
	// JETErrColumnTooBig - Field length is greater than maximum
	JETErrColumnTooBig JETErr = -1506
	// JETErrColumnIndexed - Column indexed, cannot delete
	JETErrColumnIndexed JETErr = -1505
	// JETErrNullInvalid - Null not valid
	JETErrNullInvalid JETErr = -1504
	// JETErrColumnDoesNotFit - Field will not fit in record
	JETErrColumnDoesNotFit JETErr = -1503
	// JETErrColumnNoChunk - No such chunk in long value
	JETErrColumnNoChunk JETErr = -1502
	// JETErrColumnLong - Column value is long
	JETErrColumnLong JETErr = -1501
	// JETErrCannotIndexOnEncryptedColumn - Cannot index encrypted column
	JETErrCannotIndexOnEncryptedColumn JETErr = -1440
	// JETErrColumnCannotBeEncrypted - Only JET_coltypLongText and JET_coltypLongBinary columns without default values can be encrypted
	JETErrColumnCannotBeEncrypted JETErr = -1439
	// JETErrInvalidLVChunkSize - Specified LV chunk size is not supported
	JETErrInvalidLVChunkSize JETErr = -1438
	// JETErrIndexTuplesKeyTooSmall - specified key does not meet minimum tuple length
	JETErrIndexTuplesKeyTooSmall JETErr = -1437
	// JETErrIndexTuplesCannotRetrieveFromIndex - cannot call RetrieveColumn() with RetrieveFromIndex on a tuple index
	JETErrIndexTuplesCannotRetrieveFromIndex JETErr = -1436
	// JETErrIndexTuplesInvalidLimits - invalid min/max tuple length or max characters to index specified
	JETErrIndexTuplesInvalidLimits JETErr = -1435
	// JETErrIndexTuplesVarSegMacNotAllowed - tuple index does not allow setting cbVarSegMac
	JETErrIndexTuplesVarSegMacNotAllowed JETErr = -1434
	// JETErrIndexTuplesTextBinaryColumnsOnly - tuple index must be on a text/binary column
	JETErrIndexTuplesTextBinaryColumnsOnly JETErr = -1433
	// JETErrIndexTuplesNonUniqueOnly - tuple index must be a non-unique index
	JETErrIndexTuplesNonUniqueOnly JETErr = -1432
	// JETErrIndexTuplesTooManyColumns - tuple index may only have eleven columns in the index
	JETErrIndexTuplesTooManyColumns JETErr = -1431
	// JETErrIndexTuplesSecondaryIndexOnly - tuple index can only be on a secondary index
	JETErrIndexTuplesSecondaryIndexOnly JETErr = -1430
	// JETErrInvalidIndexId - Illegal index id
	JETErrInvalidIndexId JETErr = -1416
	// JETErrSecondaryIndexCorrupted - Secondary index is corrupt. The database must be defragmented or the affected index must be deleted. If the corrupt index is over Unicode text, a likely cause is a sort-order change.
	JETErrSecondaryIndexCorrupted JETErr = -1414
	// JETErrPrimaryIndexCorrupted - Primary index is corrupt. The database must be defragmented or the table deleted.
	JETErrPrimaryIndexCorrupted JETErr = -1413
	// JETErrIndexBuildCorrupted - Failed to build a secondary index that properly reflects primary index
	JETErrIndexBuildCorrupted JETErr = -1412
	// JETErrMultiValuedIndexViolation - Non-unique inter-record index keys generated for a multivalued index
	JETErrMultiValuedIndexViolation JETErr = -1411
	// JETErrTooManyOpenIndexes - Out of index description blocks
	JETErrTooManyOpenIndexes JETErr = -1410
	// JETErrInvalidCreateIndex - Invalid create index description
	JETErrInvalidCreateIndex JETErr = -1409
	// JETErrIndexInvalidDef - Illegal index definition
	JETErrIndexInvalidDef JETErr = -1406
	// JETErrIndexMustStay - Cannot delete clustered index
	JETErrIndexMustStay JETErr = -1405
	// JETErrIndexNotFound - No such index
	JETErrIndexNotFound JETErr = -1404
	// JETErrIndexDuplicate - Index is already defined
	JETErrIndexDuplicate JETErr = -1403
	// JETErrIndexHasPrimary - Primary index already defined
	JETErrIndexHasPrimary JETErr = -1402
	// JETErrIndexCantBuild - Index build failed
	JETErrIndexCantBuild JETErr = -1401
	// JETErrCannotAddFixedVarColumnToDerivedTable - Template table was created with NoFixedVarColumnsInDerivedTables
	JETErrCannotAddFixedVarColumnToDerivedTable JETErr = -1330
	// JETErrClientRequestToStopJetService - Client has requested stop service
	JETErrClientRequestToStopJetService JETErr = -1329
	// JETErrInvalidSettings - System parameters were set improperly
	JETErrInvalidSettings JETErr = -1328
	// JETErrDDLNotInheritable - Tried to inherit DDL from a table not marked as a template table.
	JETErrDDLNotInheritable JETErr = -1326
	// JETErrCannotNestDDL - Nesting of hierarchical DDL is not currently supported.
	JETErrCannotNestDDL JETErr = -1325
	// JETErrFixedInheritedDDL - On a derived table, DDL operations are prohibited on inherited portion of DDL
	JETErrFixedInheritedDDL JETErr = -1324
	// JETErrFixedDDL - DDL operations prohibited on this table
	JETErrFixedDDL JETErr = -1323
	// JETErrExclusiveTableLockRequired - Must have exclusive lock on table.
	JETErrExclusiveTableLockRequired JETErr = -1322
	// JETErrCannotDeleteTemplateTable - Illegal attempt to delete a template table
	JETErrCannotDeleteTemplateTable JETErr = -1319
	// JETErrCannotDeleteSystemTable - Illegal attempt to delete a system table
	JETErrCannotDeleteSystemTable JETErr = -1318
	// JETErrCannotDeleteTempTable - Use CloseTable instead of DeleteTable to delete temp table
	JETErrCannotDeleteTempTable JETErr = -1317
	// JETErrInvalidObject - Object is invalid for operation
	JETErrInvalidObject JETErr = -1316
	// JETErrObjectDuplicate - Table or object name in use
	JETErrObjectDuplicate JETErr = -1314
	// JETErrTooManyOpenTablesAndCleanupTimedOut - Cannot open any more tables (cleanup attempt failed to complete)
	JETErrTooManyOpenTablesAndCleanupTimedOut JETErr = -1313
	// JETErrIllegalOperation - Oper. not supported on table
	JETErrIllegalOperation JETErr = -1312
	// JETErrTooManyOpenTables - Cannot open any more tables (cleanup already attempted)
	JETErrTooManyOpenTables JETErr = -1311
	// JETErrInvalidTableId - Invalid table id
	JETErrInvalidTableId JETErr = -1310
	// JETErrTableNotEmpty - Table is not empty
	JETErrTableNotEmpty JETErr = -1308
	// JETErrDensityInvalid - Bad file/index density
	JETErrDensityInvalid JETErr = -1307
	// JETErrObjectNotFound - No such table or object
	JETErrObjectNotFound JETErr = -1305
	// JETErrTableInUse - Table is in use, cannot lock
	JETErrTableInUse JETErr = -1304
	// JETErrTableDuplicate - Table already exists
	JETErrTableDuplicate JETErr = -1303
	// JETErrTableLocked - Table is exclusively locked
	JETErrTableLocked JETErr = -1302
	// JETErrTransactionsNotReadyDuringRecovery - Recovery has not seen any Begin0/Commit0 records and so does not know what trxBegin0 to assign to this transaction
	JETErrTransactionsNotReadyDuringRecovery JETErr = -1232
	// JETErrDatabaseAttachedForRecovery - Database is attached but only for recovery.  It must be explicitly attached before it can be opened.
	JETErrDatabaseAttachedForRecovery JETErr = -1231
	// JETErrDatabaseNotReady - Recovery on this database has not yet completed enough to permit access.
	JETErrDatabaseNotReady JETErr = -1230
	// JETErrNoAttachmentsFailedIncrementalReseed - The incremental reseed being performed on the specified database cannot be completed because the min required log contains no attachment info.  A full reseed is required to recover this database.
	JETErrNoAttachmentsFailedIncrementalReseed JETErr = -1229
	// JETErrDatabaseFailedIncrementalReseed - The incremental reseed being performed on the specified database cannot be completed due to a fatal error.  A full reseed is required to recover this database.
	JETErrDatabaseFailedIncrementalReseed JETErr = -1228
	// JETErrDatabaseInvalidIncrementalReseed - The database is not a valid state to perform an incremental reseed.
	JETErrDatabaseInvalidIncrementalReseed JETErr = -1227
	// JETErrDatabaseIncompleteIncrementalReseed - The database cannot be attached because it is currently being rebuilt as part of an incremental reseed.
	JETErrDatabaseIncompleteIncrementalReseed JETErr = -1226
	// JETErrInvalidCreateDbVersion - recovery tried to replay a database creation, but the database was originally created with an incompatible (likely older) version of the database engine
	JETErrInvalidCreateDbVersion JETErr = -1225
	// JETErrDatabaseCorruptedNoRepair - Corrupted db but repair not allowed
	JETErrDatabaseCorruptedNoRepair JETErr = -1224
	// JETErrDatabaseSignInUse - Database with same signature in use
	JETErrDatabaseSignInUse JETErr = -1222
	// JETErrPartiallyAttachedDB - Database is partially attached. Cannot complete attach operation
	JETErrPartiallyAttachedDB JETErr = -1221
	// JETErrCatalogCorrupted - Corruption detected in catalog
	JETErrCatalogCorrupted JETErr = -1220
	// JETErrForceDetachNotAllowed - Force Detach allowed only after normal detach errored out
	JETErrForceDetachNotAllowed JETErr = -1219
	// JETErrDatabaseIdInUse - A database is being assigned an id already in use
	JETErrDatabaseIdInUse JETErr = -1218
	// JETErrDatabaseInvalidPath - Specified path to database file is illegal
	JETErrDatabaseInvalidPath JETErr = -1217
	// JETErrAttachedDatabaseMismatch - An outstanding database attachment has been detected at the start or end of recovery, but database is missing or does not match attachment info
	JETErrAttachedDatabaseMismatch JETErr = -1216
	// JETErrDatabaseSharingViolation - A different database instance is using this database
	JETErrDatabaseSharingViolation JETErr = -1215
	// JETErrTooManyInstances - Cannot start any more database instances
	JETErrTooManyInstances JETErr = -1214
	// JETErrPageSizeMismatch - The database page size does not match the engine
	JETErrPageSizeMismatch JETErr = -1213
	// JETErrDatabase500Format - The database is in an older (500) format
	JETErrDatabase500Format JETErr = -1212
	// JETErrDatabase400Format - The database is in an older (400) format
	JETErrDatabase400Format JETErr = -1211
	// JETErrDatabase200Format - The database is in an older (200) format
	JETErrDatabase200Format JETErr = -1210
	// JETErrInvalidDatabaseVersion - Database engine is incompatible with database
	JETErrInvalidDatabaseVersion JETErr = -1209
	// JETErrCannotDisableVersioning - Cannot disable versioning for this database
	JETErrCannotDisableVersioning JETErr = -1208
	// JETErrDatabaseLocked - Database exclusively locked
	JETErrDatabaseLocked JETErr = -1207
	// JETErrDatabaseCorrupted - Non database file or corrupted db
	JETErrDatabaseCorrupted JETErr = -1206
	// JETErrDatabaseInvalidPages - Invalid number of pages
	JETErrDatabaseInvalidPages JETErr = -1205
	// JETErrDatabaseInvalidName - Invalid database name
	JETErrDatabaseInvalidName JETErr = -1204
	// JETErrDatabaseNotFound - No such database
	JETErrDatabaseNotFound JETErr = -1203
	// JETErrDatabaseInUse - Database in use
	JETErrDatabaseInUse JETErr = -1202
	// JETErrDatabaseDuplicate - Database already exists
	JETErrDatabaseDuplicate JETErr = -1201
	// JETErrDTCCallbackUnexpectedError - Unexpected error code returned from DTC callback
	JETErrDTCCallbackUnexpectedError JETErr = -1162
	// JETErrDTCMissingCallbackOnRecovery - Attempted to recover a distributed transaction but no callback for DTC coordination was specified on initialisation
	JETErrDTCMissingCallbackOnRecovery JETErr = -1161
	// JETErrDTCMissingCallback - Attempted to begin a distributed transaction but no callback for DTC coordination was specified on initialisation
	JETErrDTCMissingCallback JETErr = -1160
	// JETErrCannotNestDistributedTransactions - Attempted to begin a distributed transaction when not at level 0
	JETErrCannotNestDistributedTransactions JETErr = -1154
	// JETErrDistributedTransactionNotYetPreparedToCommit - Attempted to commit a distributed transaction, but PrepareToCommit has not yet been called
	JETErrDistributedTransactionNotYetPreparedToCommit JETErr = -1153
	// JETErrNotInDistributedTransaction - Attempted to PrepareToCommit a non-distributed transaction
	JETErrNotInDistributedTransaction JETErr = -1152
	// JETErrDistributedTransactionAlreadyPreparedToCommit - Attempted a write-operation after a distributed transaction has called PrepareToCommit
	JETErrDistributedTransactionAlreadyPreparedToCommit JETErr = -1151
	// JETErrMustCommitDistributedTransactionToLevel0 - Attempted to PrepareToCommit a distributed transaction to non-zero level
	JETErrMustCommitDistributedTransactionToLevel0 JETErr = -1150
	// JETErrFilteredMoveNotSupported - Attempted to provide a filter to JetSetCursorFilter() in an unsupported scenario.
	JETErrFilteredMoveNotSupported JETErr = -1124
	// JETErrRecoveryVerifyFailure - One or more database pages read from disk during recovery do not match the expected state.
	JETErrRecoveryVerifyFailure JETErr = -1123
	// JETErrFileSystemCorruption - File system operation failed with an error indicating the file system is corrupt.
	JETErrFileSystemCorruption JETErr = -1121
	// JETErrReadLostFlushVerifyFailure - The database page read from disk had a previous write not represented on the page.
	JETErrReadLostFlushVerifyFailure JETErr = -1119
	// JETErrReadPgnoVerifyFailure - The database page read from disk had the wrong page number.
	JETErrReadPgnoVerifyFailure JETErr = -1118
	// JETErrDirtyShutdown - The instance was shutdown successfully but all the attached databases were left in a dirty state by request via JET_bitTermDirty
	JETErrDirtyShutdown JETErr = -1116
	// JETErrInvalidInstance - Invalid instance handle
	JETErrInvalidInstance JETErr = -1115
	// JETErrSesidTableIdMismatch - This session handle can't be used with this table id
	JETErrSesidTableIdMismatch JETErr = -1114
	// JETErrCannotMaterializeForwardOnlySort - The temp table could not be created due to parameters that conflict with JET_bitTTForwardOnly
	JETErrCannotMaterializeForwardOnlySort JETErr = -1113
	// JETErrRecordTooBigForBackwardCompatibility - record would be too big if represented in a database format from a previous version of Jet
	JETErrRecordTooBigForBackwardCompatibility JETErr = -1112
	// JETErrSessionWriteConflict - Attempt to replace the same record by two different cursors in the same session
	JETErrSessionWriteConflict JETErr = -1111
	// JETErrTransReadOnly - Read-only transaction tried to modify the database
	JETErrTransReadOnly JETErr = -1110
	// JETErrRollbackRequired - Must rollback current transaction -- cannot commit or begin a new one
	JETErrRollbackRequired JETErr = -1109
	// JETErrInTransaction - Operation not allowed within a transaction
	JETErrInTransaction JETErr = -1108
	// JETErrWriteConflictPrimaryIndex - Update attempted on uncommitted primary index
	JETErrWriteConflictPrimaryIndex JETErr = -1105
	// JETErrInvalidSesid - Invalid session handle
	JETErrInvalidSesid JETErr = -1104
	// JETErrTransTooDeep - Transactions nested too deeply
	JETErrTransTooDeep JETErr = -1103
	// JETErrWriteConflict - Write lock failed due to outstanding write lock
	JETErrWriteConflict JETErr = -1102
	// JETErrOutOfSessions - Out of sessions
	JETErrOutOfSessions JETErr = -1101
	// JETErrInvalidDbparamId - This JET_dbparam* identifier is not known to the ESE engine.
	JETErrInvalidDbparamId JETErr = -1095
	// JETErrTooManyRecords - There are too many records to enumerate, switch to an API that handles 64-bit numbers
	JETErrTooManyRecords JETErr = -1094
	// JETErrInvalidSesparamId - This JET_sesparam* identifier is not known to the ESE engine.
	JETErrInvalidSesparamId JETErr = -1093
	// JETErrInstanceUnavailableDueToFatalLogDiskFull - This instance cannot be used because it encountered a log-disk-full error performing an operation (likely transaction rollback) that could not tolerate failure
	JETErrInstanceUnavailableDueToFatalLogDiskFull JETErr = -1092
	// JETErrDatabaseUnavailable - This database cannot be used because it encountered a fatal error
	JETErrDatabaseUnavailable JETErr = -1091
	// JETErrInstanceUnavailable - This instance cannot be used because it encountered a fatal error
	JETErrInstanceUnavailable JETErr = -1090
	// JETErrSystemParameterConflict - Global system parameters have already been set, but to a conflicting or disagreeable state to the specified values.
	JETErrSystemParameterConflict JETErr = -1087
	// JETErrInstanceNameInUse - Instance Name already in use
	JETErrInstanceNameInUse JETErr = -1086
	// JETErrTempPathInUse - Temp path already used by another database instance
	JETErrTempPathInUse JETErr = -1085
	// JETErrLogFilePathInUse - Logfile path already used by another database instance
	JETErrLogFilePathInUse JETErr = -1084
	// JETErrSystemPathInUse - System path already used by another database instance
	JETErrSystemPathInUse JETErr = -1083
	// JETErrSystemParamsAlreadySet - Global system parameters have already been set
	JETErrSystemParamsAlreadySet JETErr = -1082
	// JETErrRunningInMultiInstanceMode - Single-instance call with multi-instance mode enabled
	JETErrRunningInMultiInstanceMode JETErr = -1081
	// JETErrRunningInOneInstanceMode - Multi-instance call with single-instance mode enabled
	JETErrRunningInOneInstanceMode JETErr = -1080
	// JETErrOutOfSequentialIndexValues - Sequential index counter has reached maximum value (perform offline defrag to reclaim free/unused SequentialIndex values)
	JETErrOutOfSequentialIndexValues JETErr = -1078
	// JETErrOutOfDbtimeValues - Dbtime counter has reached maximum value (perform offline defrag to reclaim free/unused Dbtime values)
	JETErrOutOfDbtimeValues JETErr = -1077
	// JETErrOutOfAutoincrementValues - Auto-increment counter has reached maximum value (offline defrag WILL NOT be able to reclaim free/unused Auto-increment values).
	JETErrOutOfAutoincrementValues JETErr = -1076
	// JETErrOutOfLongValueIDs - Long-value ID counter has reached maximum value. (perform offline defrag to reclaim free/unused LongValueIDs)
	JETErrOutOfLongValueIDs JETErr = -1075
	// JETErrOutOfObjectIDs - Out of btree ObjectIDs (perform offline defrag to reclaim freed/unused ObjectIds)
	JETErrOutOfObjectIDs JETErr = -1074
	// JETErrTooManyMempoolEntries - Too many mempool entries requested
	JETErrTooManyMempoolEntries JETErr = -1073
	// JETErrRecordNotDeleted - Record has not been deleted
	JETErrRecordNotDeleted JETErr = -1072
	// JETErrCannotIndex - Cannot index escrow column
	JETErrCannotIndex JETErr = -1071
	// JETErrCurrencyStackOutOfMemory - UNUSED: lCSRPerfFUCB * g_lCursorsMax exceeded (XJET only)
	JETErrCurrencyStackOutOfMemory JETErr = -1070
	// JETErrVersionStoreOutOfMemory - Version store out of memory (cleanup already attempted)
	JETErrVersionStoreOutOfMemory JETErr = -1069
	// JETErrVersionStoreOutOfMemoryAndCleanupTimedOut - Version store out of memory (and cleanup attempt failed to complete)
	JETErrVersionStoreOutOfMemoryAndCleanupTimedOut JETErr = -1066
	// JETErrVersionStoreEntryTooBig - Attempted to create a version store entry (RCE) larger than a version bucket
	JETErrVersionStoreEntryTooBig JETErr = -1065
	// JETErrInvalidLCMapStringFlags - Invalid flags for LCMapString()
	JETErrInvalidLCMapStringFlags JETErr = -1064
	// JETErrInvalidCodePage - Invalid or unknown code page
	JETErrInvalidCodePage JETErr = -1063
	// JETErrInvalidLanguageId - Invalid or unknown language id
	JETErrInvalidLanguageId JETErr = -1062
	// JETErrInvalidCountry - Invalid or unknown country/region code
	JETErrInvalidCountry JETErr = -1061
	// JETErrTooManyActiveUsers - Too many active database users
	JETErrTooManyActiveUsers JETErr = -1059
	// JETErrMustRollback - Transaction must rollback because failure of unversioned update
	JETErrMustRollback JETErr = -1057
	// JETErrNotInTransaction - Operation must be within a transaction
	JETErrNotInTransaction JETErr = -1054
	// JETErrNullKeyDisallowed - Null keys are disallowed on index
	JETErrNullKeyDisallowed JETErr = -1053
	// JETErrLinkNotSupported - Link support unavailable
	JETErrLinkNotSupported JETErr = -1052
	// JETErrIndexInUse - Index is in use
	JETErrIndexInUse JETErr = -1051
	// JETErrColumnNotUpdatable - Cannot set column value
	JETErrColumnNotUpdatable JETErr = -1048
	// JETErrInvalidBufferSize - Data buffer doesn't match column size
	JETErrInvalidBufferSize JETErr = -1047
	// JETErrColumnInUse - Column used in an index
	JETErrColumnInUse JETErr = -1046
	// JETErrInvalidBookmark - Invalid bookmark
	JETErrInvalidBookmark JETErr = -1045
	// JETErrInvalidFilename - Filename is invalid
	JETErrInvalidFilename JETErr = -1044
	// JETErrContainerNotEmpty - Container is not empty
	JETErrContainerNotEmpty JETErr = -1043
	// JETErrTooManyColumns - Too many columns defined
	JETErrTooManyColumns JETErr = -1040
	// JETErrBufferTooSmall - Buffer is too small
	JETErrBufferTooSmall JETErr = -1038
	// JETErrSQLLinkNotSupported - SQL Link support unavailable
	JETErrSQLLinkNotSupported JETErr = -1035
	// JETErrQueryNotSupported - Query support unavailable
	JETErrQueryNotSupported JETErr = -1034
	// JETErrFileAccessDenied - Cannot access file, the file is locked or in use
	JETErrFileAccessDenied JETErr = -1032
	// JETErrInitInProgress - Database engine is being initialized
	JETErrInitInProgress JETErr = -1031
	// JETErrAlreadyInitialized - Database engine already initialized
	JETErrAlreadyInitialized JETErr = -1030
	// JETErrNotInitialized - Database engine not initialized
	JETErrNotInitialized JETErr = -1029
	// JETErrInvalidDatabase - Not a database file
	JETErrInvalidDatabase JETErr = -1028
	// JETErrTooManyOpenDatabases - Too many open databases
	JETErrTooManyOpenDatabases JETErr = -1027
	// JETErrRecordTooBig - Record larger than maximum size
	JETErrRecordTooBig JETErr = -1026
	// JETErrInvalidLogDirectory - Invalid log directory
	JETErrInvalidLogDirectory JETErr = -1025
	// JETErrInvalidSystemPath - Invalid system path
	JETErrInvalidSystemPath JETErr = -1024
	// JETErrInvalidPath - Invalid file path
	JETErrInvalidPath JETErr = -1023
	// JETErrDiskIO - Disk IO error
	JETErrDiskIO JETErr = -1022
	// JETErrDiskReadVerificationFailure - The OS returned ERROR_CRC from file IO
	JETErrDiskReadVerificationFailure JETErr = -1021
	// JETErrOutOfFileHandles - Out of file handles
	JETErrOutOfFileHandles JETErr = -1020
	// JETErrPageNotInitialized - Blank database page
	JETErrPageNotInitialized JETErr = -1019
	// JETErrReadVerifyFailure - Checksum error on a database page
	JETErrReadVerifyFailure JETErr = -1018
	// JETErrRecordDeleted - Record has been deleted
	JETErrRecordDeleted JETErr = -1017
	// JETErrTooManyKeys - Too many columns in an index
	JETErrTooManyKeys JETErr = -1016
	// JETErrTooManyIndexes - Too many indexes
	JETErrTooManyIndexes JETErr = -1015
	// JETErrOutOfBuffers - Out of database page buffers
	JETErrOutOfBuffers JETErr = -1014
	// JETErrOutOfCursors - Out of table cursors
	JETErrOutOfCursors JETErr = -1013
	// JETErrOutOfDatabaseSpace - Maximum database size reached
	JETErrOutOfDatabaseSpace JETErr = -1012
	// JETErrOutOfMemory - Out of Memory
	JETErrOutOfMemory JETErr = -1011
	// JETErrInvalidDatabaseId - Invalid database id
	JETErrInvalidDatabaseId JETErr = -1010
	// JETErrDatabaseFileReadOnly - Tried to attach a read-only database file for read/write operations
	JETErrDatabaseFileReadOnly JETErr = -1008
	// JETErrInvalidParameter - Invalid API parameter
	JETErrInvalidParameter JETErr = -1003
	// JETErrInvalidName - Invalid name
	JETErrInvalidName JETErr = -1002
	// JETErrFeatureNotAvailable - API not supported
	JETErrFeatureNotAvailable JETErr = -1001
	// JETErrTermInProgress - Termination in progress
	JETErrTermInProgress JETErr = -1000
	// JETErrInvalidGrbit - Invalid flags parameter
	JETErrInvalidGrbit JETErr = -900
	// JETErrBackupAbortByServer - Backup was aborted by server by calling JetTerm with JET_bitTermStopBackup or by calling JetStopBackup
	JETErrBackupAbortByServer JETErr = -801
	// JETErrEngineFormatVersionSpecifiedTooLowForDatabaseVersion - The specified JET_ENGINEFORMATVERSION is set too low for this database file, the database file has already been upgraded to a higher version.  A higher JET_ENGINEFORMATVERSION value must be set in the param.
	JETErrEngineFormatVersionSpecifiedTooLowForDatabaseVersion JETErr = -623
	// JETErrEngineFormatVersionSpecifiedTooLowForLogVersion - The specified JET_ENGINEFORMATVERSION is set too low for this log stream, the log files have already been upgraded to a higher version.  A higher JET_ENGINEFORMATVERSION value must be set in the param.
	JETErrEngineFormatVersionSpecifiedTooLowForLogVersion JETErr = -622
	// JETErrEngineFormatVersionParamTooLowForRequestedFeature - Thrown by a format feature (not at JetSetSystemParameter) if the client requests a feature that requires a version higher than that set for the JET_paramEngineFormatVersion.
	JETErrEngineFormatVersionParamTooLowForRequestedFeature JETErr = -621
	// JETErrEngineFormatVersionNotYetImplementedTooHigh - The specified JET_ENGINEFORMATVERSION value is too high, higher than this version of ESE knows about.
	JETErrEngineFormatVersionNotYetImplementedTooHigh JETErr = -620
	// JETErrEngineFormatVersionNoLongerSupportedTooLow - The specified JET_ENGINEFORMATVERSION value is too low to be supported by this version of ESE.
	JETErrEngineFormatVersionNoLongerSupportedTooLow JETErr = -619
	// JETErrTransactionTooLong - Too many outstanding generations between JetBeginTransaction and current generation.
	JETErrTransactionTooLong JETErr = -618
	// JETErrSurrogateBackupInProgress - A surrogate backup is in progress.
	JETErrSurrogateBackupInProgress JETErr = -617
	// JETErrLogFileNotCopied - log truncation attempted but not all required logs were copied
	JETErrLogFileNotCopied JETErr = -616
	// JETErrRestoreOfNonBackupDatabase - hard recovery attempted on a database that wasn't a backup database
	JETErrRestoreOfNonBackupDatabase JETErr = -615
	// JETErrCheckpointDepthTooDeep - too many outstanding generations between checkpoint and current generation
	JETErrCheckpointDepthTooDeep JETErr = -614
	// JETErrLogReadVerifyFailure - Checksum error in log file during backup
	JETErrLogReadVerifyFailure JETErr = -612
	// JETErrExistingLogFileIsNotContiguous - Existing log file is not contiguous
	JETErrExistingLogFileIsNotContiguous JETErr = -611
	// JETErrExistingLogFileHasBadSignature - Existing log file has bad signature
	JETErrExistingLogFileHasBadSignature JETErr = -610
	// JETErrUnicodeLanguageValidationFailure - Can not validate the language
	JETErrUnicodeLanguageValidationFailure JETErr = -604
	// JETErrUnicodeNormalizationNotSupported - OS does not provide support for Unicode normalisation (and no normalisation callback was specified)
	JETErrUnicodeNormalizationNotSupported JETErr = -603
	// JETErrUnicodeTranslationFail - Unicode normalization failed
	JETErrUnicodeTranslationFail JETErr = -602
	// JETErrUnicodeTranslationBufferTooSmall - Unicode translation buffer too small
	JETErrUnicodeTranslationBufferTooSmall JETErr = -601
	// JETErrPageInitializedMismatch - Database divergence mismatch. Page was uninitialized on remote node, but initialized on local node.
	JETErrPageInitializedMismatch JETErr = -596
	// JETErrLogSequenceChecksumMismatch - The previous log's accumulated segment checksum doesn't match the next log
	JETErrLogSequenceChecksumMismatch JETErr = -590
	// JETErrCommittedLogFileCorrupt - One or more logs were found to be corrupt during recovery.  These log files are required to maintain durable ACID semantics, but not required to maintain consistency if the JET_bitIgnoreLostLogs bit and JET_paramDeleteOutOfRangeLogs is specified during recovery.
	JETErrCommittedLogFileCorrupt JETErr = -586
	// JETErrRecoveredWithoutUndoDatabasesConsistent - Soft recovery successfully replayed all operations and intended to skip the Undo phase of recovery, but the Undo phase was not required
	JETErrRecoveredWithoutUndoDatabasesConsistent JETErr = -584
	// JETErrSectorSizeNotSupported - The physical sector size reported by the disk subsystem, is unsupported by ESE for a specific file type.
	JETErrSectorSizeNotSupported JETErr = -583
	// JETErrCommittedLogFilesMissing - One or more logs that were committed to this database, are missing.  These log files are required to maintain durable ACID semantics, but not required to maintain consistency if the JET_bitReplayIgnoreLostLogs bit is specified during recovery.
	JETErrCommittedLogFilesMissing JETErr = -582
	// JETErrSoftRecoveryOnSnapshot - Soft recovery on a database from a shadow copy backup set
	JETErrSoftRecoveryOnSnapshot JETErr = -581
	// JETErrDatabasesNotFromSameSnapshot - Databases to be restored are not from the same shadow copy backup
	JETErrDatabasesNotFromSameSnapshot JETErr = -580
	// JETErrRecoveredWithoutUndo - Soft recovery successfully replayed all operations, but the Undo phase of recovery was skipped
	JETErrRecoveredWithoutUndo JETErr = -579
	// JETErrBadRestoreTargetInstance - TargetInstance specified for restore is not found or log files don't match
	JETErrBadRestoreTargetInstance JETErr = -577
	// JETErrMustDisableLoggingForDbUpgrade - Cannot have logging enabled while attempting to upgrade db
	JETErrMustDisableLoggingForDbUpgrade JETErr = -575
	// JETErrLogCorruptDuringHardRecovery - corruption was detected during hard recovery (log was not part of a backup set)
	JETErrLogCorruptDuringHardRecovery JETErr = -574
	// JETErrLogCorruptDuringHardRestore - corruption was detected in a backup set during hard restore
	JETErrLogCorruptDuringHardRestore JETErr = -573
	// JETErrLogTornWriteDuringHardRecovery - torn-write was detected during hard recovery (log was not part of a backup set)
	JETErrLogTornWriteDuringHardRecovery JETErr = -571
	// JETErrLogTornWriteDuringHardRestore - torn-write was detected in a backup set during hard restore
	JETErrLogTornWriteDuringHardRestore JETErr = -570
	// JETErrMissingFileToBackup - Some log or patch files are missing during backup
	JETErrMissingFileToBackup JETErr = -569
	// JETErrDbTimeTooNew - dbtime on page in advance of the dbtimeBefore in record
	JETErrDbTimeTooNew JETErr = -567
	// JETErrDbTimeTooOld - dbtime on page smaller than dbtimeBefore in record
	JETErrDbTimeTooOld JETErr = -566
	// JETErrMissingCurrentLogFiles - Some current log files are missing for continuous restore
	JETErrMissingCurrentLogFiles JETErr = -565
	// JETErrDatabaseIncompleteUpgrade - Attempted to use a database which was only partially converted to the current format -- must restore from backup
	JETErrDatabaseIncompleteUpgrade JETErr = -563
	// JETErrDatabaseAlreadyUpgraded - Attempted to upgrade a database that is already current
	JETErrDatabaseAlreadyUpgraded JETErr = -562
	// JETErrBadBackupDatabaseSize - The backup database size is not in 4k
	JETErrBadBackupDatabaseSize JETErr = -561
	// JETErrMissingFullBackup - The database missed a previous full backup before incremental backup
	JETErrMissingFullBackup JETErr = -560
	// JETErrMissingRestoreLogFiles - Some restore log files are missing
	JETErrMissingRestoreLogFiles JETErr = -557
	// JETErrGivenLogFileIsNotContiguous - Restore log file is not contiguous
	JETErrGivenLogFileIsNotContiguous JETErr = -556
	// JETErrGivenLogFileHasBadSignature - Restore log file has bad signature
	JETErrGivenLogFileHasBadSignature JETErr = -555
	// JETErrStartingRestoreLogTooHigh - The starting log number too high for the restore
	JETErrStartingRestoreLogTooHigh JETErr = -554
	// JETErrEndingRestoreLogTooLow - The starting log number too low for the restore
	JETErrEndingRestoreLogTooLow JETErr = -553
	// JETErrDatabasePatchFileMismatch - Patch file is not generated from this backup
	JETErrDatabasePatchFileMismatch JETErr = -552
	// JETErrConsistentTimeMismatch - Database last consistent time unmatched
	JETErrConsistentTimeMismatch JETErr = -551
	// JETErrDatabaseDirtyShutdown - Database was not shutdown cleanly. Recovery must first be run to properly complete database operations for the previous shutdown.
	JETErrDatabaseDirtyShutdown JETErr = -550
	// JETErrStreamingDataNotLogged - Illegal attempt to replay a streaming file operation where the data wasn't logged. Probably caused by an attempt to roll-forward with circular logging enabled
	JETErrStreamingDataNotLogged JETErr = -549
	// JETErrLogSequenceEndDatabasesConsistent - databases have been recovered, but all possible log generations in the current sequence are used; delete all log files and the checkpoint file and backup the databases before continuing
	JETErrLogSequenceEndDatabasesConsistent JETErr = -548
	// JETErrLogSectorSizeMismatchDatabasesConsistent - databases have been recovered, but the log file sector size (used during recovery) does not match the current volume's sector size
	JETErrLogSectorSizeMismatchDatabasesConsistent JETErr = -547
	// JETErrLogSectorSizeMismatch - the log file sector size does not match the current volume's sector size
	JETErrLogSectorSizeMismatch JETErr = -546
	// JETErrLogFileSizeMismatchDatabasesConsistent - databases have been recovered, but the log file size used during recovery does not match JET_paramLogFileSize
	JETErrLogFileSizeMismatchDatabasesConsistent JETErr = -545
	// JETErrSoftRecoveryOnBackupDatabase - Soft recovery is intended on a backup database. Restore should be used instead
	JETErrSoftRecoveryOnBackupDatabase JETErr = -544
	// JETErrRequiredLogFilesMissing - The required log files for recovery is missing.
	JETErrRequiredLogFilesMissing JETErr = -543
	// JETErrCheckpointFileNotFound - Could not locate checkpoint file
	JETErrCheckpointFileNotFound JETErr = -542
	// JETErrLogFileSizeMismatch - actual log file size does not match JET_paramLogFileSize
	JETErrLogFileSizeMismatch JETErr = -541
	// JETErrDatabaseStreamingFileMismatch - Database and streaming file do not match each other
	JETErrDatabaseStreamingFileMismatch JETErr = -540
	// JETErrDatabaseLogSetMismatch - Database does not belong with the current set of log files
	JETErrDatabaseLogSetMismatch JETErr = -539
	// JETErrPatchFileMissing - Hard restore detected that patch file is missing from backup set
	JETErrPatchFileMissing JETErr = -538
	// JETErrRedoAbruptEnded - Redo abruptly ended due to sudden failure in reading logs from log file
	JETErrRedoAbruptEnded JETErr = -536
	// JETErrBadPatchPage - Patch file page is not valid
	JETErrBadPatchPage JETErr = -535
	// JETErrMissingPatchPage - Patch file page not found during recovery
	JETErrMissingPatchPage JETErr = -534
	// JETErrCheckpointCorrupt - Checkpoint file not found or corrupt
	JETErrCheckpointCorrupt JETErr = -533
	// JETErrBadCheckpointSignature - Bad signature for a checkpoint file
	JETErrBadCheckpointSignature JETErr = -532
	// JETErrBadDbSignature - Bad signature for a db file
	JETErrBadDbSignature JETErr = -531
	// JETErrBadLogSignature - Bad signature for a log file
	JETErrBadLogSignature JETErr = -530
	// JETErrLogDiskFull - Log disk full
	JETErrLogDiskFull JETErr = -529
	// JETErrMissingLogFile - Current log file missing
	JETErrMissingLogFile JETErr = -528
	// JETErrRecoveredWithErrors - Restored with errors
	JETErrRecoveredWithErrors JETErr = -527
	// JETErrInvalidBackup - Cannot perform incremental backup when circular logging enabled
	JETErrInvalidBackup JETErr = -526
	// JETErrMakeBackupDirectoryFail - Could not make backup temp directory
	JETErrMakeBackupDirectoryFail JETErr = -525
	// JETErrDeleteBackupFileFail - Could not delete backup file
	JETErrDeleteBackupFileFail JETErr = -524
	// JETErrBackupNotAllowedYet - Cannot do backup now
	JETErrBackupNotAllowedYet JETErr = -523
	// JETErrInvalidBackupSequence - Backup call out of sequence
	JETErrInvalidBackupSequence JETErr = -521
	// JETErrNoBackup - No backup in progress
	JETErrNoBackup JETErr = -520
	// JETErrLogSequenceEnd - Maximum log file number exceeded
	JETErrLogSequenceEnd JETErr = -519
	// JETErrLogBufferTooSmall - An operation generated a log record which was too large to fit in the log buffer or in a single log file
	JETErrLogBufferTooSmall JETErr = -517
	// JETErrLoggingDisabled - Log is not active
	JETErrLoggingDisabled JETErr = -516
	// JETErrInvalidLogSequence - Timestamp in next log does not match expected
	JETErrInvalidLogSequence JETErr = -515
	// JETErrBadLogVersion - Version of log file is not compatible with Jet version
	JETErrBadLogVersion JETErr = -514
	// JETErrLogGenerationMismatch - Name of logfile does not match internal generation number
	JETErrLogGenerationMismatch JETErr = -513
	// JETErrCannotLogDuringRecoveryRedo - Try to log something during recovery redo
	JETErrCannotLogDuringRecoveryRedo JETErr = -512
	// JETErrLogDisabledDueToRecoveryFailure - Try to log something after recovery failed
	JETErrLogDisabledDueToRecoveryFailure JETErr = -511
	// JETErrLogWriteFail - Failure writing to log file
	JETErrLogWriteFail JETErr = -510
	// JETErrMissingPreviousLogFile - Missing the log file for check point
	JETErrMissingPreviousLogFile JETErr = -509
	// JETErrRestoreInProgress - Restore in progress
	JETErrRestoreInProgress JETErr = -506
	// JETErrBackupInProgress - Backup is active already
	JETErrBackupInProgress JETErr = -505
	// JETErrBackupDirectoryNotEmpty - The backup directory is not empty
	JETErrBackupDirectoryNotEmpty JETErr = -504
	// JETErrNoBackupDirectory - No backup directory given
	JETErrNoBackupDirectory JETErr = -503
	// JETErrLogFileCorrupt - Log file is corrupt
	JETErrLogFileCorrupt JETErr = -501
	// JETErrInvalidLoggedOperation - Logged operation cannot be redone
	JETErrInvalidLoggedOperation JETErr = -500
	// JETErrCompressionIntegrityCheckFailed - A compression integrity check failed. Decompressing data failed the integrity checksum indicating a data corruption in the compress/decompress pipeline.
	JETErrCompressionIntegrityCheckFailed JETErr = -431
	// JETErrStaleColumnReference - Column reference is stale
	JETErrStaleColumnReference JETErr = -427
	// JETErrInvalidColumnReference - Column reference is invalid
	JETErrInvalidColumnReference JETErr = -426
	// JETErrInvalidPreread - Cannot preread long values when current index secondary
	JETErrInvalidPreread JETErr = -424
	// JETErrMustBeSeparateLongValue - Can only preread long value columns that can be separate, e.g. not size constrained so that they are fixed or variable columns
	JETErrMustBeSeparateLongValue JETErr = -423
	// JETErrSeparatedLongValue - Operation not supported on separated long-value
	JETErrSeparatedLongValue JETErr = -421
	// JETErrCannotSeparateIntrinsicLV - illegal attempt to separate an LV which must be intrinsic
	JETErrCannotSeparateIntrinsicLV JETErr = -416
	// JETErrKeyTooBig - Key is too large
	JETErrKeyTooBig JETErr = -408
	// JETErrBadLineCount - Number of lines on the page is too few compared to the line being operated on
	JETErrBadLineCount JETErr = -354
	// JETErrBadEmptyPage - Database corrupted. Searching an unexpectedly empty page.
	JETErrBadEmptyPage JETErr = -351
	// JETErrDatabaseLeakInSpace - Some database pages have become unreachable even from the avail tree, only an offline defragmentation can return the lost space.
	JETErrDatabaseLeakInSpace JETErr = -348
	// JETErrKeyTruncated - key truncated on index that disallows key truncation
	JETErrKeyTruncated JETErr = -346
	// JETErrDbTimeCorrupted - Dbtime on current page is greater than global database dbtime
	JETErrDbTimeCorrupted JETErr = -344
	// JETErrSPOwnExtCorrupted - OwnExt space tree is corrupt
	JETErrSPOwnExtCorrupted JETErr = -343
	// JETErrSPAvailExtCacheOutOfMemory - Out of memory allocating an AvailExt cache node
	JETErrSPAvailExtCacheOutOfMemory JETErr = -342
	// JETErrSPAvailExtCorrupted - AvailExt space tree is corrupt
	JETErrSPAvailExtCorrupted JETErr = -341
	// JETErrSPAvailExtCacheOutOfSync - AvailExt cache doesn't match btree
	JETErrSPAvailExtCacheOutOfSync JETErr = -340
	// JETErrBadParentPageLink - Database corrupted
	JETErrBadParentPageLink JETErr = -338
	// JETErrNTSystemCallFailed - A call to the operating system failed
	JETErrNTSystemCallFailed JETErr = -334
	// JETErrBadBookmark - Bookmark has no corresponding address in database
	JETErrBadBookmark JETErr = -328
	// JETErrBadPageLink - Database corrupted
	JETErrBadPageLink JETErr = -327
	// JETErrKeyBoundary - Reached Key Boundary
	JETErrKeyBoundary JETErr = -324
	// JETErrPageBoundary - Reached Page Boundary
	JETErrPageBoundary JETErr = -323
	// JETErrPreviousVersion - Version already existed. Recovery failure
	JETErrPreviousVersion JETErr = -322
	// JETErrDatabaseBufferDependenciesCorrupted - Buffer dependencies improperly set. Recovery failure
	JETErrDatabaseBufferDependenciesCorrupted JETErr = -255
	// JETErrDeviceFailure - A required hardware device didn't function as expected.
	JETErrDeviceFailure JETErr = -118
	// JETErrDeviceTimeout - Timeout occurred while waiting for a hardware device to respond.
	JETErrDeviceTimeout JETErr = -116
	// JETErrDeviceMisconfigured - A required hardware device was misconfigured externally.
	JETErrDeviceMisconfigured JETErr = -115
	// JETErrDeviceMissing - A required hardware device or functionality was missing.
	JETErrDeviceMissing JETErr = -114
	// JETErrUnloadableOSFunctionality - The desired OS functionality could not be located and loaded / linked.
	JETErrUnloadableOSFunctionality JETErr = -113
	// JETErrDisabledFunctionality - You are running MinESE, that does not have all features compiled in.  This functionality is only supported in a full version of ESE.
	JETErrDisabledFunctionality JETErr = -112
	// JETErrInternalError - Fatal internal error
	JETErrInternalError JETErr = -107
	// JETErrTaskDropped - A requested async task could not be executed
	JETErrTaskDropped JETErr = -106
	// JETErrTooManyIO - System busy due to too many IOs
	JETErrTooManyIO JETErr = -105
	// JETErrOutOfThreads - Could not start thread
	JETErrOutOfThreads JETErr = -103
	// JETErrFileClose - Could not close file
	JETErrFileClose JETErr = -102
	// JETErrRfsNotArmed - Resource Failure Simulator not initialized
	JETErrRfsNotArmed JETErr = -101
	// JETErrRfsFailure - Resource Failure Simulator failure
	JETErrRfsFailure JETErr = -100
)

// String returns the string representation of the error code
func (e JETErr) String() string {
	switch e {
	case JETErrSuccess:
		return "Success"
	case JETErrFileCompressed:
		return "FileCompressed"
	case JETErrFileIOFail:
		return "FileIOFail"
	case JETErrFileIORetry:
		return "FileIORetry"
	case JETErrFileIOAbort:
		return "FileIOAbort"
	case JETErrFileIOBeyondEOF:
		return "FileIOBeyondEOF"
	case JETErrFileIOSparse:
		return "FileIOSparse"
	case JETErrLSNotSet:
		return "LSNotSet"
	case JETErrLSAlreadySet:
		return "LSAlreadySet"
	case JETErrLSCallbackNotSpecified:
		return "LSCallbackNotSpecified"
	case JETErrInvalidLogDataSequence:
		return "InvalidLogDataSequence"
	case JETErrTestInjectionNotSupported:
		return "TestInjectionNotSupported"
	case JETErrTooManyTestInjections:
		return "TooManyTestInjections"
	case JETErrOSSnapshotInvalidSnapId:
		return "OSSnapshotInvalidSnapId"
	case JETErrOSSnapshotNotAllowed:
		return "OSSnapshotNotAllowed"
	case JETErrOSSnapshotTimeOut:
		return "OSSnapshotTimeOut"
	case JETErrOSSnapshotInvalidSequence:
		return "OSSnapshotInvalidSequence"
	case JETErrSpaceHintsInvalid:
		return "SpaceHintsInvalid"
	case JETErrCallbackNotResolved:
		return "CallbackNotResolved"
	case JETErrCallbackFailed:
		return "CallbackFailed"
	case JETErrDatabaseAlreadyRunningMaintenance:
		return "DatabaseAlreadyRunningMaintenance"
	case JETErrFlushMapUnrecoverable:
		return "FlushMapUnrecoverable"
	case JETErrFlushMapDatabaseMismatch:
		return "FlushMapDatabaseMismatch"
	case JETErrFlushMapVersionUnsupported:
		return "FlushMapVersionUnsupported"
	case JETErrRollbackError:
		return "RollbackError"
	case JETErrOneDatabasePerSession:
		return "OneDatabasePerSession"
	case JETErrRecordFormatConversionFailed:
		return "RecordFormatConversionFailed"
	case JETErrSessionInUse:
		return "SessionInUse"
	case JETErrSessionContextNotSetByThisThread:
		return "SessionContextNotSetByThisThread"
	case JETErrSessionContextAlreadySet:
		return "SessionContextAlreadySet"
	case JETErrEntryPointNotFound:
		return "EntryPointNotFound"
	case JETErrSessionSharingViolation:
		return "SessionSharingViolation"
	case JETErrTooManySplits:
		return "TooManySplits"
	case JETErrAccessDenied:
		return "AccessDenied"
	case JETErrInvalidOperation:
		return "InvalidOperation"
	case JETErrLogCorrupted:
		return "LogCorrupted"
	case JETErrAfterInitialization:
		return "AfterInitialization"
	case JETErrFileAlreadyExists:
		return "FileAlreadyExists"
	case JETErrFileInvalidType:
		return "FileInvalidType"
	case JETErrFileNotFound:
		return "FileNotFound"
	case JETErrPermissionDenied:
		return "PermissionDenied"
	case JETErrDiskFull:
		return "DiskFull"
	case JETErrTooManyAttachedDatabases:
		return "TooManyAttachedDatabases"
	case JETErrTempFileOpenError:
		return "TempFileOpenError"
	case JETErrInvalidOnSort:
		return "InvalidOnSort"
	case JETErrTooManySorts:
		return "TooManySorts"
	case JETErrEncryptionBadItag:
		return "EncryptionBadItag"
	case JETErrDecryptionFailed:
		return "DecryptionFailed"
	case JETErrUpdateMustVersion:
		return "UpdateMustVersion"
	case JETErrDecompressionFailed:
		return "DecompressionFailed"
	case JETErrLanguageNotSupported:
		return "LanguageNotSupported"
	case JETErrDataHasChanged:
		return "DataHasChanged"
	case JETErrUpdateNotPrepared:
		return "UpdateNotPrepared"
	case JETErrKeyNotMade:
		return "KeyNotMade"
	case JETErrAlreadyPrepared:
		return "AlreadyPrepared"
	case JETErrKeyDuplicate:
		return "KeyDuplicate"
	case JETErrRecordPrimaryChanged:
		return "RecordPrimaryChanged"
	case JETErrNoCurrentRecord:
		return "NoCurrentRecord"
	case JETErrRecordNoCopy:
		return "RecordNoCopy"
	case JETErrRecordNotFound:
		return "RecordNotFound"
	case JETErrColumnNoEncryptionKey:
		return "ColumnNoEncryptionKey"
	case JETErrColumnCannotBeCompressed:
		return "ColumnCannotBeCompressed"
	case JETErrInvalidPlaceholderColumn:
		return "InvalidPlaceholderColumn"
	case JETErrDerivedColumnCorruption:
		return "DerivedColumnCorruption"
	case JETErrMultiValuedDuplicateAfterTruncation:
		return "MultiValuedDuplicateAfterTruncation"
	case JETErrLVCorrupted:
		return "LVCorrupted"
	case JETErrMultiValuedDuplicate:
		return "MultiValuedDuplicate"
	case JETErrDefaultValueTooBig:
		return "DefaultValueTooBig"
	case JETErrCannotBeTagged:
		return "CannotBeTagged"
	case JETErrColumnInRelationship:
		return "ColumnInRelationship"
	case JETErrBadItagSequence:
		return "BadItagSequence"
	case JETErrBadColumnId:
		return "BadColumnId"
	case JETErrKeyIsMade:
		return "KeyIsMade"
	case JETErrNoCurrentIndex:
		return "NoCurrentIndex"
	case JETErrTaggedNotNULL:
		return "TaggedNotNULL"
	case JETErrInvalidColumnType:
		return "InvalidColumnType"
	case JETErrColumnRedundant:
		return "ColumnRedundant"
	case JETErrMultiValuedColumnMustBeTagged:
		return "MultiValuedColumnMustBeTagged"
	case JETErrColumnDuplicate:
		return "ColumnDuplicate"
	case JETErrColumnNotFound:
		return "ColumnNotFound"
	case JETErrColumnTooBig:
		return "ColumnTooBig"
	case JETErrColumnIndexed:
		return "ColumnIndexed"
	case JETErrNullInvalid:
		return "NullInvalid"
	case JETErrColumnDoesNotFit:
		return "ColumnDoesNotFit"
	case JETErrColumnNoChunk:
		return "ColumnNoChunk"
	case JETErrColumnLong:
		return "ColumnLong"
	case JETErrCannotIndexOnEncryptedColumn:
		return "CannotIndexOnEncryptedColumn"
	case JETErrColumnCannotBeEncrypted:
		return "ColumnCannotBeEncrypted"
	case JETErrInvalidLVChunkSize:
		return "InvalidLVChunkSize"
	case JETErrIndexTuplesKeyTooSmall:
		return "IndexTuplesKeyTooSmall"
	case JETErrIndexTuplesCannotRetrieveFromIndex:
		return "IndexTuplesCannotRetrieveFromIndex"
	case JETErrIndexTuplesInvalidLimits:
		return "IndexTuplesInvalidLimits"
	case JETErrIndexTuplesVarSegMacNotAllowed:
		return "IndexTuplesVarSegMacNotAllowed"
	case JETErrIndexTuplesTextBinaryColumnsOnly:
		return "IndexTuplesTextBinaryColumnsOnly"
	case JETErrIndexTuplesNonUniqueOnly:
		return "IndexTuplesNonUniqueOnly"
	case JETErrIndexTuplesTooManyColumns:
		return "IndexTuplesTooManyColumns"
	case JETErrIndexTuplesSecondaryIndexOnly:
		return "IndexTuplesSecondaryIndexOnly"
	case JETErrInvalidIndexId:
		return "InvalidIndexId"
	case JETErrSecondaryIndexCorrupted:
		return "SecondaryIndexCorrupted"
	case JETErrPrimaryIndexCorrupted:
		return "PrimaryIndexCorrupted"
	case JETErrIndexBuildCorrupted:
		return "IndexBuildCorrupted"
	case JETErrMultiValuedIndexViolation:
		return "MultiValuedIndexViolation"
	case JETErrTooManyOpenIndexes:
		return "TooManyOpenIndexes"
	case JETErrInvalidCreateIndex:
		return "InvalidCreateIndex"
	case JETErrIndexInvalidDef:
		return "IndexInvalidDef"
	case JETErrIndexMustStay:
		return "IndexMustStay"
	case JETErrIndexNotFound:
		return "IndexNotFound"
	case JETErrIndexDuplicate:
		return "IndexDuplicate"
	case JETErrIndexHasPrimary:
		return "IndexHasPrimary"
	case JETErrIndexCantBuild:
		return "IndexCantBuild"
	case JETErrCannotAddFixedVarColumnToDerivedTable:
		return "CannotAddFixedVarColumnToDerivedTable"
	case JETErrClientRequestToStopJetService:
		return "ClientRequestToStopJetService"
	case JETErrInvalidSettings:
		return "InvalidSettings"
	case JETErrDDLNotInheritable:
		return "DDLNotInheritable"
	case JETErrCannotNestDDL:
		return "CannotNestDDL"
	case JETErrFixedInheritedDDL:
		return "FixedInheritedDDL"
	case JETErrFixedDDL:
		return "FixedDDL"
	case JETErrExclusiveTableLockRequired:
		return "ExclusiveTableLockRequired"
	case JETErrCannotDeleteTemplateTable:
		return "CannotDeleteTemplateTable"
	case JETErrCannotDeleteSystemTable:
		return "CannotDeleteSystemTable"
	case JETErrCannotDeleteTempTable:
		return "CannotDeleteTempTable"
	case JETErrInvalidObject:
		return "InvalidObject"
	case JETErrObjectDuplicate:
		return "ObjectDuplicate"
	case JETErrTooManyOpenTablesAndCleanupTimedOut:
		return "TooManyOpenTablesAndCleanupTimedOut"
	case JETErrIllegalOperation:
		return "IllegalOperation"
	case JETErrTooManyOpenTables:
		return "TooManyOpenTables"
	case JETErrInvalidTableId:
		return "InvalidTableId"
	case JETErrTableNotEmpty:
		return "TableNotEmpty"
	case JETErrDensityInvalid:
		return "DensityInvalid"
	case JETErrObjectNotFound:
		return "ObjectNotFound"
	case JETErrTableInUse:
		return "TableInUse"
	case JETErrTableDuplicate:
		return "TableDuplicate"
	case JETErrTableLocked:
		return "TableLocked"
	case JETErrTransactionsNotReadyDuringRecovery:
		return "TransactionsNotReadyDuringRecovery"
	case JETErrDatabaseAttachedForRecovery:
		return "DatabaseAttachedForRecovery"
	case JETErrDatabaseNotReady:
		return "DatabaseNotReady"
	case JETErrNoAttachmentsFailedIncrementalReseed:
		return "NoAttachmentsFailedIncrementalReseed"
	case JETErrDatabaseFailedIncrementalReseed:
		return "DatabaseFailedIncrementalReseed"
	case JETErrDatabaseInvalidIncrementalReseed:
		return "DatabaseInvalidIncrementalReseed"
	case JETErrDatabaseIncompleteIncrementalReseed:
		return "DatabaseIncompleteIncrementalReseed"
	case JETErrInvalidCreateDbVersion:
		return "InvalidCreateDbVersion"
	case JETErrDatabaseCorruptedNoRepair:
		return "DatabaseCorruptedNoRepair"
	case JETErrDatabaseSignInUse:
		return "DatabaseSignInUse"
	case JETErrPartiallyAttachedDB:
		return "PartiallyAttachedDB"
	case JETErrCatalogCorrupted:
		return "CatalogCorrupted"
	case JETErrForceDetachNotAllowed:
		return "ForceDetachNotAllowed"
	case JETErrDatabaseIdInUse:
		return "DatabaseIdInUse"
	case JETErrDatabaseInvalidPath:
		return "DatabaseInvalidPath"
	case JETErrAttachedDatabaseMismatch:
		return "AttachedDatabaseMismatch"
	case JETErrDatabaseSharingViolation:
		return "DatabaseSharingViolation"
	case JETErrTooManyInstances:
		return "TooManyInstances"
	case JETErrPageSizeMismatch:
		return "PageSizeMismatch"
	case JETErrDatabase500Format:
		return "Database500Format"
	case JETErrDatabase400Format:
		return "Database400Format"
	case JETErrDatabase200Format:
		return "Database200Format"
	case JETErrInvalidDatabaseVersion:
		return "InvalidDatabaseVersion"
	case JETErrCannotDisableVersioning:
		return "CannotDisableVersioning"
	case JETErrDatabaseLocked:
		return "DatabaseLocked"
	case JETErrDatabaseCorrupted:
		return "DatabaseCorrupted"
	case JETErrDatabaseInvalidPages:
		return "DatabaseInvalidPages"
	case JETErrDatabaseInvalidName:
		return "DatabaseInvalidName"
	case JETErrDatabaseNotFound:
		return "DatabaseNotFound"
	case JETErrDatabaseInUse:
		return "DatabaseInUse"
	case JETErrDatabaseDuplicate:
		return "DatabaseDuplicate"
	case JETErrDTCCallbackUnexpectedError:
		return "DTCCallbackUnexpectedError"
	case JETErrDTCMissingCallbackOnRecovery:
		return "DTCMissingCallbackOnRecovery"
	case JETErrDTCMissingCallback:
		return "DTCMissingCallback"
	case JETErrCannotNestDistributedTransactions:
		return "CannotNestDistributedTransactions"
	case JETErrDistributedTransactionNotYetPreparedToCommit:
		return "DistributedTransactionNotYetPreparedToCommit"
	case JETErrNotInDistributedTransaction:
		return "NotInDistributedTransaction"
	case JETErrDistributedTransactionAlreadyPreparedToCommit:
		return "DistributedTransactionAlreadyPreparedToCommit"
	case JETErrMustCommitDistributedTransactionToLevel0:
		return "MustCommitDistributedTransactionToLevel0"
	case JETErrFilteredMoveNotSupported:
		return "FilteredMoveNotSupported"
	case JETErrRecoveryVerifyFailure:
		return "RecoveryVerifyFailure"
	case JETErrFileSystemCorruption:
		return "FileSystemCorruption"
	case JETErrReadLostFlushVerifyFailure:
		return "ReadLostFlushVerifyFailure"
	case JETErrReadPgnoVerifyFailure:
		return "ReadPgnoVerifyFailure"
	case JETErrDirtyShutdown:
		return "DirtyShutdown"
	case JETErrInvalidInstance:
		return "InvalidInstance"
	case JETErrSesidTableIdMismatch:
		return "SesidTableIdMismatch"
	case JETErrCannotMaterializeForwardOnlySort:
		return "CannotMaterializeForwardOnlySort"
	case JETErrRecordTooBigForBackwardCompatibility:
		return "RecordTooBigForBackwardCompatibility"
	case JETErrSessionWriteConflict:
		return "SessionWriteConflict"
	case JETErrTransReadOnly:
		return "TransReadOnly"
	case JETErrRollbackRequired:
		return "RollbackRequired"
	case JETErrInTransaction:
		return "InTransaction"
	case JETErrWriteConflictPrimaryIndex:
		return "WriteConflictPrimaryIndex"
	case JETErrInvalidSesid:
		return "InvalidSesid"
	case JETErrTransTooDeep:
		return "TransTooDeep"
	case JETErrWriteConflict:
		return "WriteConflict"
	case JETErrOutOfSessions:
		return "OutOfSessions"
	case JETErrInvalidDbparamId:
		return "InvalidDbparamId"
	case JETErrTooManyRecords:
		return "TooManyRecords"
	case JETErrInvalidSesparamId:
		return "InvalidSesparamId"
	case JETErrInstanceUnavailableDueToFatalLogDiskFull:
		return "InstanceUnavailableDueToFatalLogDiskFull"
	case JETErrDatabaseUnavailable:
		return "DatabaseUnavailable"
	case JETErrInstanceUnavailable:
		return "InstanceUnavailable"
	case JETErrSystemParameterConflict:
		return "SystemParameterConflict"
	case JETErrInstanceNameInUse:
		return "InstanceNameInUse"
	case JETErrTempPathInUse:
		return "TempPathInUse"
	case JETErrLogFilePathInUse:
		return "LogFilePathInUse"
	case JETErrSystemPathInUse:
		return "SystemPathInUse"
	case JETErrSystemParamsAlreadySet:
		return "SystemParamsAlreadySet"
	case JETErrRunningInMultiInstanceMode:
		return "RunningInMultiInstanceMode"
	case JETErrRunningInOneInstanceMode:
		return "RunningInOneInstanceMode"
	case JETErrOutOfSequentialIndexValues:
		return "OutOfSequentialIndexValues"
	case JETErrOutOfDbtimeValues:
		return "OutOfDbtimeValues"
	case JETErrOutOfAutoincrementValues:
		return "OutOfAutoincrementValues"
	case JETErrOutOfLongValueIDs:
		return "OutOfLongValueIDs"
	case JETErrOutOfObjectIDs:
		return "OutOfObjectIDs"
	case JETErrTooManyMempoolEntries:
		return "TooManyMempoolEntries"
	case JETErrRecordNotDeleted:
		return "RecordNotDeleted"
	case JETErrCannotIndex:
		return "CannotIndex"
	case JETErrCurrencyStackOutOfMemory:
		return "CurrencyStackOutOfMemory"
	case JETErrVersionStoreOutOfMemory:
		return "VersionStoreOutOfMemory"
	case JETErrVersionStoreOutOfMemoryAndCleanupTimedOut:
		return "VersionStoreOutOfMemoryAndCleanupTimedOut"
	case JETErrVersionStoreEntryTooBig:
		return "VersionStoreEntryTooBig"
	case JETErrInvalidLCMapStringFlags:
		return "InvalidLCMapStringFlags"
	case JETErrInvalidCodePage:
		return "InvalidCodePage"
	case JETErrInvalidLanguageId:
		return "InvalidLanguageId"
	case JETErrInvalidCountry:
		return "InvalidCountry"
	case JETErrTooManyActiveUsers:
		return "TooManyActiveUsers"
	case JETErrMustRollback:
		return "MustRollback"
	case JETErrNotInTransaction:
		return "NotInTransaction"
	case JETErrNullKeyDisallowed:
		return "NullKeyDisallowed"
	case JETErrLinkNotSupported:
		return "LinkNotSupported"
	case JETErrIndexInUse:
		return "IndexInUse"
	case JETErrColumnNotUpdatable:
		return "ColumnNotUpdatable"
	case JETErrInvalidBufferSize:
		return "InvalidBufferSize"
	case JETErrColumnInUse:
		return "ColumnInUse"
	case JETErrInvalidBookmark:
		return "InvalidBookmark"
	case JETErrInvalidFilename:
		return "InvalidFilename"
	case JETErrContainerNotEmpty:
		return "ContainerNotEmpty"
	case JETErrTooManyColumns:
		return "TooManyColumns"
	case JETErrBufferTooSmall:
		return "BufferTooSmall"
	case JETErrSQLLinkNotSupported:
		return "SQLLinkNotSupported"
	case JETErrQueryNotSupported:
		return "QueryNotSupported"
	case JETErrFileAccessDenied:
		return "FileAccessDenied"
	case JETErrInitInProgress:
		return "InitInProgress"
	case JETErrAlreadyInitialized:
		return "AlreadyInitialized"
	case JETErrNotInitialized:
		return "NotInitialized"
	case JETErrInvalidDatabase:
		return "InvalidDatabase"
	case JETErrTooManyOpenDatabases:
		return "TooManyOpenDatabases"
	case JETErrRecordTooBig:
		return "RecordTooBig"
	case JETErrInvalidLogDirectory:
		return "InvalidLogDirectory"
	case JETErrInvalidSystemPath:
		return "InvalidSystemPath"
	case JETErrInvalidPath:
		return "InvalidPath"
	case JETErrDiskIO:
		return "DiskIO"
	case JETErrDiskReadVerificationFailure:
		return "DiskReadVerificationFailure"
	case JETErrOutOfFileHandles:
		return "OutOfFileHandles"
	case JETErrPageNotInitialized:
		return "PageNotInitialized"
	case JETErrReadVerifyFailure:
		return "ReadVerifyFailure"
	case JETErrRecordDeleted:
		return "RecordDeleted"
	case JETErrTooManyKeys:
		return "TooManyKeys"
	case JETErrTooManyIndexes:
		return "TooManyIndexes"
	case JETErrOutOfBuffers:
		return "OutOfBuffers"
	case JETErrOutOfCursors:
		return "OutOfCursors"
	case JETErrOutOfDatabaseSpace:
		return "OutOfDatabaseSpace"
	case JETErrOutOfMemory:
		return "OutOfMemory"
	case JETErrInvalidDatabaseId:
		return "InvalidDatabaseId"
	case JETErrDatabaseFileReadOnly:
		return "DatabaseFileReadOnly"
	case JETErrInvalidParameter:
		return "InvalidParameter"
	case JETErrInvalidName:
		return "InvalidName"
	case JETErrFeatureNotAvailable:
		return "FeatureNotAvailable"
	case JETErrTermInProgress:
		return "TermInProgress"
	case JETErrInvalidGrbit:
		return "InvalidGrbit"
	case JETErrBackupAbortByServer:
		return "BackupAbortByServer"
	case JETErrEngineFormatVersionSpecifiedTooLowForDatabaseVersion:
		return "EngineFormatVersionSpecifiedTooLowForDatabaseVersion"
	case JETErrEngineFormatVersionSpecifiedTooLowForLogVersion:
		return "EngineFormatVersionSpecifiedTooLowForLogVersion"
	case JETErrEngineFormatVersionParamTooLowForRequestedFeature:
		return "EngineFormatVersionParamTooLowForRequestedFeature"
	case JETErrEngineFormatVersionNotYetImplementedTooHigh:
		return "EngineFormatVersionNotYetImplementedTooHigh"
	case JETErrEngineFormatVersionNoLongerSupportedTooLow:
		return "EngineFormatVersionNoLongerSupportedTooLow"
	case JETErrTransactionTooLong:
		return "TransactionTooLong"
	case JETErrSurrogateBackupInProgress:
		return "SurrogateBackupInProgress"
	case JETErrLogFileNotCopied:
		return "LogFileNotCopied"
	case JETErrRestoreOfNonBackupDatabase:
		return "RestoreOfNonBackupDatabase"
	case JETErrCheckpointDepthTooDeep:
		return "CheckpointDepthTooDeep"
	case JETErrLogReadVerifyFailure:
		return "LogReadVerifyFailure"
	case JETErrExistingLogFileIsNotContiguous:
		return "ExistingLogFileIsNotContiguous"
	case JETErrExistingLogFileHasBadSignature:
		return "ExistingLogFileHasBadSignature"
	case JETErrUnicodeLanguageValidationFailure:
		return "UnicodeLanguageValidationFailure"
	case JETErrUnicodeNormalizationNotSupported:
		return "UnicodeNormalizationNotSupported"
	case JETErrUnicodeTranslationFail:
		return "UnicodeTranslationFail"
	case JETErrUnicodeTranslationBufferTooSmall:
		return "UnicodeTranslationBufferTooSmall"
	case JETErrPageInitializedMismatch:
		return "PageInitializedMismatch"
	case JETErrLogSequenceChecksumMismatch:
		return "LogSequenceChecksumMismatch"
	case JETErrCommittedLogFileCorrupt:
		return "CommittedLogFileCorrupt"
	case JETErrRecoveredWithoutUndoDatabasesConsistent:
		return "RecoveredWithoutUndoDatabasesConsistent"
	case JETErrSectorSizeNotSupported:
		return "SectorSizeNotSupported"
	case JETErrCommittedLogFilesMissing:
		return "CommittedLogFilesMissing"
	case JETErrSoftRecoveryOnSnapshot:
		return "SoftRecoveryOnSnapshot"
	case JETErrDatabasesNotFromSameSnapshot:
		return "DatabasesNotFromSameSnapshot"
	case JETErrRecoveredWithoutUndo:
		return "RecoveredWithoutUndo"
	case JETErrBadRestoreTargetInstance:
		return "BadRestoreTargetInstance"
	case JETErrMustDisableLoggingForDbUpgrade:
		return "MustDisableLoggingForDbUpgrade"
	case JETErrLogCorruptDuringHardRecovery:
		return "LogCorruptDuringHardRecovery"
	case JETErrLogCorruptDuringHardRestore:
		return "LogCorruptDuringHardRestore"
	case JETErrLogTornWriteDuringHardRecovery:
		return "LogTornWriteDuringHardRecovery"
	case JETErrLogTornWriteDuringHardRestore:
		return "LogTornWriteDuringHardRestore"
	case JETErrMissingFileToBackup:
		return "MissingFileToBackup"
	case JETErrDbTimeTooNew:
		return "DbTimeTooNew"
	case JETErrDbTimeTooOld:
		return "DbTimeTooOld"
	case JETErrMissingCurrentLogFiles:
		return "MissingCurrentLogFiles"
	case JETErrDatabaseIncompleteUpgrade:
		return "DatabaseIncompleteUpgrade"
	case JETErrDatabaseAlreadyUpgraded:
		return "DatabaseAlreadyUpgraded"
	case JETErrBadBackupDatabaseSize:
		return "BadBackupDatabaseSize"
	case JETErrMissingFullBackup:
		return "MissingFullBackup"
	case JETErrMissingRestoreLogFiles:
		return "MissingRestoreLogFiles"
	case JETErrGivenLogFileIsNotContiguous:
		return "GivenLogFileIsNotContiguous"
	case JETErrGivenLogFileHasBadSignature:
		return "GivenLogFileHasBadSignature"
	case JETErrStartingRestoreLogTooHigh:
		return "StartingRestoreLogTooHigh"
	case JETErrEndingRestoreLogTooLow:
		return "EndingRestoreLogTooLow"
	case JETErrDatabasePatchFileMismatch:
		return "DatabasePatchFileMismatch"
	case JETErrConsistentTimeMismatch:
		return "ConsistentTimeMismatch"
	case JETErrDatabaseDirtyShutdown:
		return "DatabaseDirtyShutdown"
	case JETErrStreamingDataNotLogged:
		return "StreamingDataNotLogged"
	case JETErrLogSequenceEndDatabasesConsistent:
		return "LogSequenceEndDatabasesConsistent"
	case JETErrLogSectorSizeMismatchDatabasesConsistent:
		return "LogSectorSizeMismatchDatabasesConsistent"
	case JETErrLogSectorSizeMismatch:
		return "LogSectorSizeMismatch"
	case JETErrLogFileSizeMismatchDatabasesConsistent:
		return "LogFileSizeMismatchDatabasesConsistent"
	case JETErrSoftRecoveryOnBackupDatabase:
		return "SoftRecoveryOnBackupDatabase"
	case JETErrRequiredLogFilesMissing:
		return "RequiredLogFilesMissing"
	case JETErrCheckpointFileNotFound:
		return "CheckpointFileNotFound"
	case JETErrLogFileSizeMismatch:
		return "LogFileSizeMismatch"
	case JETErrDatabaseStreamingFileMismatch:
		return "DatabaseStreamingFileMismatch"
	case JETErrDatabaseLogSetMismatch:
		return "DatabaseLogSetMismatch"
	case JETErrPatchFileMissing:
		return "PatchFileMissing"
	case JETErrRedoAbruptEnded:
		return "RedoAbruptEnded"
	case JETErrBadPatchPage:
		return "BadPatchPage"
	case JETErrMissingPatchPage:
		return "MissingPatchPage"
	case JETErrCheckpointCorrupt:
		return "CheckpointCorrupt"
	case JETErrBadCheckpointSignature:
		return "BadCheckpointSignature"
	case JETErrBadDbSignature:
		return "BadDbSignature"
	case JETErrBadLogSignature:
		return "BadLogSignature"
	case JETErrLogDiskFull:
		return "LogDiskFull"
	case JETErrMissingLogFile:
		return "MissingLogFile"
	case JETErrRecoveredWithErrors:
		return "RecoveredWithErrors"
	case JETErrInvalidBackup:
		return "InvalidBackup"
	case JETErrMakeBackupDirectoryFail:
		return "MakeBackupDirectoryFail"
	case JETErrDeleteBackupFileFail:
		return "DeleteBackupFileFail"
	case JETErrBackupNotAllowedYet:
		return "BackupNotAllowedYet"
	case JETErrInvalidBackupSequence:
		return "InvalidBackupSequence"
	case JETErrNoBackup:
		return "NoBackup"
	case JETErrLogSequenceEnd:
		return "LogSequenceEnd"
	case JETErrLogBufferTooSmall:
		return "LogBufferTooSmall"
	case JETErrLoggingDisabled:
		return "LoggingDisabled"
	case JETErrInvalidLogSequence:
		return "InvalidLogSequence"
	case JETErrBadLogVersion:
		return "BadLogVersion"
	case JETErrLogGenerationMismatch:
		return "LogGenerationMismatch"
	case JETErrCannotLogDuringRecoveryRedo:
		return "CannotLogDuringRecoveryRedo"
	case JETErrLogDisabledDueToRecoveryFailure:
		return "LogDisabledDueToRecoveryFailure"
	case JETErrLogWriteFail:
		return "LogWriteFail"
	case JETErrMissingPreviousLogFile:
		return "MissingPreviousLogFile"
	case JETErrRestoreInProgress:
		return "RestoreInProgress"
	case JETErrBackupInProgress:
		return "BackupInProgress"
	case JETErrBackupDirectoryNotEmpty:
		return "BackupDirectoryNotEmpty"
	case JETErrNoBackupDirectory:
		return "NoBackupDirectory"
	case JETErrLogFileCorrupt:
		return "LogFileCorrupt"
	case JETErrInvalidLoggedOperation:
		return "InvalidLoggedOperation"
	case JETErrCompressionIntegrityCheckFailed:
		return "CompressionIntegrityCheckFailed"
	case JETErrStaleColumnReference:
		return "StaleColumnReference"
	case JETErrInvalidColumnReference:
		return "InvalidColumnReference"
	case JETErrInvalidPreread:
		return "InvalidPreread"
	case JETErrMustBeSeparateLongValue:
		return "MustBeSeparateLongValue"
	case JETErrSeparatedLongValue:
		return "SeparatedLongValue"
	case JETErrCannotSeparateIntrinsicLV:
		return "CannotSeparateIntrinsicLV"
	case JETErrKeyTooBig:
		return "KeyTooBig"
	case JETErrBadLineCount:
		return "BadLineCount"
	case JETErrBadEmptyPage:
		return "BadEmptyPage"
	case JETErrDatabaseLeakInSpace:
		return "DatabaseLeakInSpace"
	case JETErrKeyTruncated:
		return "KeyTruncated"
	case JETErrDbTimeCorrupted:
		return "DbTimeCorrupted"
	case JETErrSPOwnExtCorrupted:
		return "SPOwnExtCorrupted"
	case JETErrSPAvailExtCacheOutOfMemory:
		return "SPAvailExtCacheOutOfMemory"
	case JETErrSPAvailExtCorrupted:
		return "SPAvailExtCorrupted"
	case JETErrSPAvailExtCacheOutOfSync:
		return "SPAvailExtCacheOutOfSync"
	case JETErrBadParentPageLink:
		return "BadParentPageLink"
	case JETErrNTSystemCallFailed:
		return "NTSystemCallFailed"
	case JETErrBadBookmark:
		return "BadBookmark"
	case JETErrBadPageLink:
		return "BadPageLink"
	case JETErrKeyBoundary:
		return "KeyBoundary"
	case JETErrPageBoundary:
		return "PageBoundary"
	case JETErrPreviousVersion:
		return "PreviousVersion"
	case JETErrDatabaseBufferDependenciesCorrupted:
		return "DatabaseBufferDependenciesCorrupted"
	case JETErrDeviceFailure:
		return "DeviceFailure"
	case JETErrDeviceTimeout:
		return "DeviceTimeout"
	case JETErrDeviceMisconfigured:
		return "DeviceMisconfigured"
	case JETErrDeviceMissing:
		return "DeviceMissing"
	case JETErrUnloadableOSFunctionality:
		return "UnloadableOSFunctionality"
	case JETErrDisabledFunctionality:
		return "DisabledFunctionality"
	case JETErrInternalError:
		return "InternalError"
	case JETErrTaskDropped:
		return "TaskDropped"
	case JETErrTooManyIO:
		return "TooManyIO"
	case JETErrOutOfThreads:
		return "OutOfThreads"
	case JETErrFileClose:
		return "FileClose"
	case JETErrRfsNotArmed:
		return "RfsNotArmed"
	case JETErrRfsFailure:
		return "RfsFailure"
	default:
		return "Unknown"
	}
}

// Error implements the error interface
func (e JETErr) Error() string {
	return e.String()
}