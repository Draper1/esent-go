package types

// CreateInstanceGrbit represents options for JetCreateInstance2
type CreateInstanceGrbit uint32

const (
	// CreateInstanceGrbitNone - Default options.
	CreateInstanceGrbitNone CreateInstanceGrbit = 0
)

// InitGrbit represents options for JetInit2
type InitGrbit uint32

const (
	// InitGrbitNone - Default options.
	InitGrbitNone InitGrbit = 0
)

// TermGrbit represents options for JetTerm2
type TermGrbit uint32

const (
	// TermGrbitNone - Default options.
	TermGrbitNone TermGrbit = 0

	// TermGrbitComplete - Requests that the instance be shut down cleanly.
	// Any optional cleanup work that would ordinarily be done in the background at
	// run time is completed immediately.
	TermGrbitComplete TermGrbit = 1

	// TermGrbitAbrupt - Requests that the instance be shut down as quickly as possible.
	// Any optional work that would ordinarily be done in the background at run time is abandoned.
	TermGrbitAbrupt TermGrbit = 2

	// TermGrbitStopBackup - Interrupts and fails any on-going backup.
	TermGrbitStopBackup TermGrbit = 4
)

// CreateDatabaseGrbit represents options for JetCreateDatabase
type CreateDatabaseGrbit uint32

const (
	// CreateDatabaseGrbitNone - Default options.
	CreateDatabaseGrbitNone CreateDatabaseGrbit = 0

	// CreateDatabaseGrbitOverwriteExisting - By default, if JetCreateDatabase is called and the database already exists,
	// the Api call will fail and the original database will not be overwritten.
	// OverwriteExisting changes this behavior, and the old database will be overwritten with a new one.
	CreateDatabaseGrbitOverwriteExisting CreateDatabaseGrbit = 0x200

	// CreateDatabaseGrbitRecoveryOff - Turns off logging. Setting this bit loses the ability to replay log files
	// and recover the database to a consistent usable state after a crash.
	CreateDatabaseGrbitRecoveryOff CreateDatabaseGrbit = 0x8
)

// BeginTransactionGrbit represents options for JetBeginTransaction
type BeginTransactionGrbit uint32

const (
	// BeginTransactionGrbitNone - Default options.
	BeginTransactionGrbitNone BeginTransactionGrbit = 0

	// BeginTransactionGrbitReadOnly - The transaction will not modify the database.
	// If an update is attempted, that operation will fail.
	BeginTransactionGrbitReadOnly BeginTransactionGrbit = 0x1
)

// CommitTransactionGrbit represents options for JetCommitTransaction
type CommitTransactionGrbit uint32

const (
	// CommitTransactionGrbitNone - Default options.
	CommitTransactionGrbitNone CommitTransactionGrbit = 0

	// CommitTransactionGrbitLazyFlush - The transaction is committed normally but this Api does not wait for
	// the transaction to be flushed to the transaction log file before returning to the caller.
	CommitTransactionGrbitLazyFlush CommitTransactionGrbit = 0x1
)

// BackupGrbit represents options for JetBackupInstance
type BackupGrbit uint32

const (
	// BackupGrbitNone - Default options.
	BackupGrbitNone BackupGrbit = 0

	// BackupGrbitIncremental - Creates an incremental backup as opposed to a full backup.
	// This means that only the log files created since the last full or incremental backup will be backed up.
	BackupGrbitIncremental BackupGrbit = 0x1

	// BackupGrbitAtomic - Creates a full backup of the database. This allows the preservation
	// of an existing backup in the same directory if the new backup fails.
	BackupGrbitAtomic BackupGrbit = 0x4
)

// RestoreGrbit represents options for JetRestore
type RestoreGrbit uint32

const (
	// RestoreGrbitNone - Default options.
	RestoreGrbitNone RestoreGrbit = 0

	// RestoreGrbitStopAtEnd - Stops the restore operation at the end of the current log file.
	RestoreGrbitStopAtEnd RestoreGrbit = 0x1

	// RestoreGrbitStopBeforeEnd - Stops the restore operation before the end of the current log file.
	RestoreGrbitStopBeforeEnd RestoreGrbit = 0x2
)

// BeginExternalBackupGrbit represents options for JetBeginExternalBackupInstance
type BeginExternalBackupGrbit uint32

const (
	// BeginExternalBackupGrbitNone - Default options.
	BeginExternalBackupGrbitNone BeginExternalBackupGrbit = 0

	// BeginExternalBackupGrbitIncremental - Creates an incremental backup as opposed to a full backup.
	BeginExternalBackupGrbitIncremental BeginExternalBackupGrbit = 0x1
)

// EndExternalBackupGrbit represents options for JetEndExternalBackupInstance
type EndExternalBackupGrbit uint32

const (
	// EndExternalBackupGrbitNone - Default options.
	EndExternalBackupGrbitNone EndExternalBackupGrbit = 0

	// EndExternalBackupGrbitNormal - The client application finished the backup completely, and is ending normally.
	EndExternalBackupGrbitNormal EndExternalBackupGrbit = 0x1

	// EndExternalBackupGrbitAbort - The client application is aborting the backup.
	EndExternalBackupGrbitAbort EndExternalBackupGrbit = 0x2
)

// OpenDatabaseGrbit represents options for JetOpenDatabase
type OpenDatabaseGrbit uint32

const (
	// OpenDatabaseGrbitNone - Default options.
	OpenDatabaseGrbitNone OpenDatabaseGrbit = 0

	// OpenDatabaseGrbitExclusive - Opens the database in exclusive mode.
	OpenDatabaseGrbitExclusive OpenDatabaseGrbit = 0x1

	// OpenDatabaseGrbitReadOnly - Opens the database in read-only mode.
	OpenDatabaseGrbitReadOnly OpenDatabaseGrbit = 0x2
)

// AttachDatabaseGrbit represents options for JetAttachDatabase
type AttachDatabaseGrbit uint32

const (
	// AttachDatabaseGrbitNone - Default options.
	AttachDatabaseGrbitNone AttachDatabaseGrbit = 0

	// AttachDatabaseGrbitReadOnly - Attaches the database in read-only mode.
	AttachDatabaseGrbitReadOnly AttachDatabaseGrbit = 0x1

	// AttachDatabaseGrbitDeleteCorrupted - Deletes corrupted databases during attachment.
	AttachDatabaseGrbitDeleteCorrupted AttachDatabaseGrbit = 0x2
)

// CreateTableGrbit represents options for JetCreateTable
type CreateTableGrbit uint32

const (
	// CreateTableGrbitNone - Default options.
	CreateTableGrbitNone CreateTableGrbit = 0

	// CreateTableGrbitNoFixedVarColumnsInDerivedTables - Prevents fixed and variable columns in derived tables.
	CreateTableGrbitNoFixedVarColumnsInDerivedTables CreateTableGrbit = 0x1

	// CreateTableGrbitNoFixedVarColumnsInDerivedTables2 - Alternative flag for preventing fixed and variable columns.
	CreateTableGrbitNoFixedVarColumnsInDerivedTables2 CreateTableGrbit = 0x2
)

// OpenTableGrbit represents options for JetOpenTable
type OpenTableGrbit uint32

const (
	// OpenTableGrbitNone - Default options.
	OpenTableGrbitNone OpenTableGrbit = 0

	// OpenTableGrbitDenyRead - Denies read access to other sessions.
	OpenTableGrbitDenyRead OpenTableGrbit = 0x1

	// OpenTableGrbitDenyWrite - Denies write access to other sessions.
	OpenTableGrbitDenyWrite OpenTableGrbit = 0x2

	// OpenTableGrbitUpdatable - Opens the table for updates.
	OpenTableGrbitUpdatable OpenTableGrbit = 0x4
)

// SetColumnGrbit represents options for JetSetColumn
type SetColumnGrbit uint32

const (
	// SetColumnGrbitNone - Default options.
	SetColumnGrbitNone SetColumnGrbit = 0

	// SetColumnGrbitAppendLV - Appends to long value columns.
	SetColumnGrbitAppendLV SetColumnGrbit = 0x1

	// SetColumnGrbitOverwriteLV - Overwrites long value columns.
	SetColumnGrbitOverwriteLV SetColumnGrbit = 0x2

	// SetColumnGrbitSizeLV - Sets the size of long value columns.
	SetColumnGrbitSizeLV SetColumnGrbit = 0x4

	// SetColumnGrbitZeroLength - Sets zero-length values.
	SetColumnGrbitZeroLength SetColumnGrbit = 0x8

	// SetColumnGrbitSeparateLV - Separates long value columns.
	SetColumnGrbitSeparateLV SetColumnGrbit = 0x10

	// SetColumnGrbitUniqueMultiValues - Ensures unique multi-values.
	SetColumnGrbitUniqueMultiValues SetColumnGrbit = 0x20

	// SetColumnGrbitUniqueNormalizedMultiValues - Ensures unique normalized multi-values.
	SetColumnGrbitUniqueNormalizedMultiValues SetColumnGrbit = 0x40

	// SetColumnGrbitUncompressed - Stores data uncompressed.
	SetColumnGrbitUncompressed SetColumnGrbit = 0x80
)

// RetrieveColumnGrbit represents options for JetRetrieveColumn
type RetrieveColumnGrbit uint32

const (
	// RetrieveColumnGrbitNone - Default options.
	RetrieveColumnGrbitNone RetrieveColumnGrbit = 0

	// RetrieveColumnGrbitRetrieveCopy - Retrieves a copy of the column.
	RetrieveColumnGrbitRetrieveCopy RetrieveColumnGrbit = 0x1

	// RetrieveColumnGrbitRetrieveFromIndex - Retrieves from index.
	RetrieveColumnGrbitRetrieveFromIndex RetrieveColumnGrbit = 0x2

	// RetrieveColumnGrbitRetrieveFromPrimaryBookmark - Retrieves from primary bookmark.
	RetrieveColumnGrbitRetrieveFromPrimaryBookmark RetrieveColumnGrbit = 0x4

	// RetrieveColumnGrbitRetrieveTag - Retrieves tag information.
	RetrieveColumnGrbitRetrieveTag RetrieveColumnGrbit = 0x8

	// RetrieveColumnGrbitRetrieveNull - Retrieves null values.
	RetrieveColumnGrbitRetrieveNull RetrieveColumnGrbit = 0x10

	// RetrieveColumnGrbitRetrieveIgnoreDefault - Ignores default values.
	RetrieveColumnGrbitRetrieveIgnoreDefault RetrieveColumnGrbit = 0x20

	// RetrieveColumnGrbitRetrieveLongId - Retrieves long ID.
	RetrieveColumnGrbitRetrieveLongId RetrieveColumnGrbit = 0x40

	// RetrieveColumnGrbitRetrieveLongIdMax - Retrieves maximum long ID.
	RetrieveColumnGrbitRetrieveLongIdMax RetrieveColumnGrbit = 0x80
)

// CreateIndexGrbit represents options for JetCreateIndex
type CreateIndexGrbit uint32

const (
	// CreateIndexGrbitNone - Default options.
	CreateIndexGrbitNone CreateIndexGrbit = 0

	// CreateIndexGrbitUnique - Creates a unique index.
	CreateIndexGrbitUnique CreateIndexGrbit = 0x1

	// CreateIndexGrbitPrimary - Creates a primary index.
	CreateIndexGrbitPrimary CreateIndexGrbit = 0x2

	// CreateIndexGrbitDisallowNull - Disallows null values in the index.
	CreateIndexGrbitDisallowNull CreateIndexGrbit = 0x4

	// CreateIndexGrbitIgnoreNull - Ignores null values in the index.
	CreateIndexGrbitIgnoreNull CreateIndexGrbit = 0x8

	// CreateIndexGrbitIgnoreAnyNull - Ignores any null values in the index.
	CreateIndexGrbitIgnoreAnyNull CreateIndexGrbit = 0x10

	// CreateIndexGrbitIgnoreFirstNull - Ignores the first null value in the index.
	CreateIndexGrbitIgnoreFirstNull CreateIndexGrbit = 0x20

	// CreateIndexGrbitAllowNull - Allows null values in the index.
	CreateIndexGrbitAllowNull CreateIndexGrbit = 0x40

	// CreateIndexGrbitAllowNull2 - Alternative flag for allowing null values.
	CreateIndexGrbitAllowNull2 CreateIndexGrbit = 0x80
)

// MoveGrbit represents options for JetMove
type MoveGrbit uint32

const (
	// MoveGrbitNone - Default options.
	MoveGrbitNone MoveGrbit = 0

	// MoveGrbitMoveKeyNE - Moves to the next key that is not equal.
	MoveGrbitMoveKeyNE MoveGrbit = 0x1

	// MoveGrbitMoveKeyGT - Moves to the next key that is greater than.
	MoveGrbitMoveKeyGT MoveGrbit = 0x2

	// MoveGrbitMoveKeyGE - Moves to the next key that is greater than or equal.
	MoveGrbitMoveKeyGE MoveGrbit = 0x4

	// MoveGrbitMoveKeyLT - Moves to the next key that is less than.
	MoveGrbitMoveKeyLT MoveGrbit = 0x8

	// MoveGrbitMoveKeyLE - Moves to the next key that is less than or equal.
	MoveGrbitMoveKeyLE MoveGrbit = 0x10
)

// SeekGrbit represents options for JetSeek
type SeekGrbit uint32

const (
	// SeekGrbitNone - Default options.
	SeekGrbitNone SeekGrbit = 0

	// SeekGrbitSeekEQ - Seeks to an exact match.
	SeekGrbitSeekEQ SeekGrbit = 0x1

	// SeekGrbitSeekGE - Seeks to a value greater than or equal.
	SeekGrbitSeekGE SeekGrbit = 0x2

	// SeekGrbitSeekGT - Seeks to a value greater than.
	SeekGrbitSeekGT SeekGrbit = 0x4

	// SeekGrbitSeekLE - Seeks to a value less than or equal.
	SeekGrbitSeekLE SeekGrbit = 0x8

	// SeekGrbitSeekLT - Seeks to a value less than.
	SeekGrbitSeekLT SeekGrbit = 0x10

	// SeekGrbitSeekNE - Seeks to a value not equal.
	SeekGrbitSeekNE SeekGrbit = 0x20

	// SeekGrbitSetIndexRange - Sets the index range after seeking.
	SeekGrbitSetIndexRange SeekGrbit = 0x40
)

// MakeKeyGrbit represents options for JetMakeKey
type MakeKeyGrbit uint32

const (
	// MakeKeyGrbitNone - Default options.
	MakeKeyGrbitNone MakeKeyGrbit = 0

	// MakeKeyGrbitNewKey - Creates a new key.
	MakeKeyGrbitNewKey MakeKeyGrbit = 0x1

	// MakeKeyGrbitStrLimit - Limits the key to string boundaries.
	MakeKeyGrbitStrLimit MakeKeyGrbit = 0x2

	// MakeKeyGrbitSubStrLimit - Limits the key to substring boundaries.
	MakeKeyGrbitSubStrLimit MakeKeyGrbit = 0x4

	// MakeKeyGrbitNormalizedKey - Creates a normalized key.
	MakeKeyGrbitNormalizedKey MakeKeyGrbit = 0x8

	// MakeKeyGrbitKeyDataZeroLength - Creates a key with zero-length data.
	MakeKeyGrbitKeyDataZeroLength MakeKeyGrbit = 0x10
)

// SetIndexRangeGrbit represents options for JetSetIndexRange
type SetIndexRangeGrbit uint32

const (
	// SetIndexRangeGrbitNone - Default options.
	SetIndexRangeGrbitNone SetIndexRangeGrbit = 0

	// SetIndexRangeGrbitRangeInclusive - Makes the range inclusive.
	SetIndexRangeGrbitRangeInclusive SetIndexRangeGrbit = 0x1

	// SetIndexRangeGrbitRangeUpperLimit - Sets the upper limit of the range.
	SetIndexRangeGrbitRangeUpperLimit SetIndexRangeGrbit = 0x2

	// SetIndexRangeGrbitRangeInstantDuration - Sets instant duration for the range.
	SetIndexRangeGrbitRangeInstantDuration SetIndexRangeGrbit = 0x4

	// SetIndexRangeGrbitRangeRemove - Removes the current range.
	SetIndexRangeGrbitRangeRemove SetIndexRangeGrbit = 0x8
)

// String methods for debugging
func (g CreateInstanceGrbit) String() string {
	switch g {
	case CreateInstanceGrbitNone:
		return "None"
	default:
		return "Unknown"
	}
}

func (g InitGrbit) String() string {
	switch g {
	case InitGrbitNone:
		return "None"
	default:
		return "Unknown"
	}
}

func (g TermGrbit) String() string {
	switch g {
	case TermGrbitNone:
		return "None"
	case TermGrbitComplete:
		return "Complete"
	case TermGrbitAbrupt:
		return "Abrupt"
	case TermGrbitStopBackup:
		return "StopBackup"
	default:
		return "Unknown"
	}
}

func (g CreateDatabaseGrbit) String() string {
	switch g {
	case CreateDatabaseGrbitNone:
		return "None"
	case CreateDatabaseGrbitOverwriteExisting:
		return "OverwriteExisting"
	case CreateDatabaseGrbitRecoveryOff:
		return "RecoveryOff"
	default:
		return "Unknown"
	}
}

func (g BeginTransactionGrbit) String() string {
	switch g {
	case BeginTransactionGrbitNone:
		return "None"
	case BeginTransactionGrbitReadOnly:
		return "ReadOnly"
	default:
		return "Unknown"
	}
}

func (g CommitTransactionGrbit) String() string {
	switch g {
	case CommitTransactionGrbitNone:
		return "None"
	case CommitTransactionGrbitLazyFlush:
		return "LazyFlush"
	default:
		return "Unknown"
	}
}

func (g BackupGrbit) String() string {
	switch g {
	case BackupGrbitNone:
		return "None"
	case BackupGrbitIncremental:
		return "Incremental"
	case BackupGrbitAtomic:
		return "Atomic"
	default:
		return "Unknown"
	}
}

func (g RestoreGrbit) String() string {
	switch g {
	case RestoreGrbitNone:
		return "None"
	case RestoreGrbitStopAtEnd:
		return "StopAtEnd"
	case RestoreGrbitStopBeforeEnd:
		return "StopBeforeEnd"
	default:
		return "Unknown"
	}
}

func (g BeginExternalBackupGrbit) String() string {
	switch g {
	case BeginExternalBackupGrbitNone:
		return "None"
	case BeginExternalBackupGrbitIncremental:
		return "Incremental"
	default:
		return "Unknown"
	}
}

func (g EndExternalBackupGrbit) String() string {
	switch g {
	case EndExternalBackupGrbitNone:
		return "None"
	case EndExternalBackupGrbitNormal:
		return "Normal"
	case EndExternalBackupGrbitAbort:
		return "Abort"
	default:
		return "Unknown"
	}
}

func (g SetIndexRangeGrbit) String() string {
	switch g {
	case SetIndexRangeGrbitNone:
		return "None"
	case SetIndexRangeGrbitRangeInclusive:
		return "RangeInclusive"
	case SetIndexRangeGrbitRangeUpperLimit:
		return "RangeUpperLimit"
	case SetIndexRangeGrbitRangeInstantDuration:
		return "RangeInstantDuration"
	case SetIndexRangeGrbitRangeRemove:
		return "RangeRemove"
	default:
		return "Unknown"
	}
}

func (g OpenDatabaseGrbit) String() string {
	switch g {
	case OpenDatabaseGrbitNone:
		return "None"
	case OpenDatabaseGrbitExclusive:
		return "Exclusive"
	case OpenDatabaseGrbitReadOnly:
		return "ReadOnly"
	default:
		return "Unknown"
	}
}

func (g AttachDatabaseGrbit) String() string {
	switch g {
	case AttachDatabaseGrbitNone:
		return "None"
	case AttachDatabaseGrbitReadOnly:
		return "ReadOnly"
	case AttachDatabaseGrbitDeleteCorrupted:
		return "DeleteCorrupted"
	default:
		return "Unknown"
	}
}

func (g CreateTableGrbit) String() string {
	switch g {
	case CreateTableGrbitNone:
		return "None"
	case CreateTableGrbitNoFixedVarColumnsInDerivedTables:
		return "NoFixedVarColumnsInDerivedTables"
	case CreateTableGrbitNoFixedVarColumnsInDerivedTables2:
		return "NoFixedVarColumnsInDerivedTables2"
	default:
		return "Unknown"
	}
}

func (g OpenTableGrbit) String() string {
	switch g {
	case OpenTableGrbitNone:
		return "None"
	case OpenTableGrbitDenyRead:
		return "DenyRead"
	case OpenTableGrbitDenyWrite:
		return "DenyWrite"
	case OpenTableGrbitUpdatable:
		return "Updatable"
	default:
		return "Unknown"
	}
}

func (g SetColumnGrbit) String() string {
	switch g {
	case SetColumnGrbitNone:
		return "None"
	case SetColumnGrbitAppendLV:
		return "AppendLV"
	case SetColumnGrbitOverwriteLV:
		return "OverwriteLV"
	case SetColumnGrbitSizeLV:
		return "SizeLV"
	case SetColumnGrbitZeroLength:
		return "ZeroLength"
	case SetColumnGrbitSeparateLV:
		return "SeparateLV"
	case SetColumnGrbitUniqueMultiValues:
		return "UniqueMultiValues"
	case SetColumnGrbitUniqueNormalizedMultiValues:
		return "UniqueNormalizedMultiValues"
	case SetColumnGrbitUncompressed:
		return "Uncompressed"
	default:
		return "Unknown"
	}
}

func (g RetrieveColumnGrbit) String() string {
	switch g {
	case RetrieveColumnGrbitNone:
		return "None"
	case RetrieveColumnGrbitRetrieveCopy:
		return "RetrieveCopy"
	case RetrieveColumnGrbitRetrieveFromIndex:
		return "RetrieveFromIndex"
	case RetrieveColumnGrbitRetrieveFromPrimaryBookmark:
		return "RetrieveFromPrimaryBookmark"
	case RetrieveColumnGrbitRetrieveTag:
		return "RetrieveTag"
	case RetrieveColumnGrbitRetrieveNull:
		return "RetrieveNull"
	case RetrieveColumnGrbitRetrieveIgnoreDefault:
		return "RetrieveIgnoreDefault"
	case RetrieveColumnGrbitRetrieveLongId:
		return "RetrieveLongId"
	case RetrieveColumnGrbitRetrieveLongIdMax:
		return "RetrieveLongIdMax"
	default:
		return "Unknown"
	}
}

func (g CreateIndexGrbit) String() string {
	switch g {
	case CreateIndexGrbitNone:
		return "None"
	case CreateIndexGrbitUnique:
		return "Unique"
	case CreateIndexGrbitPrimary:
		return "Primary"
	case CreateIndexGrbitDisallowNull:
		return "DisallowNull"
	case CreateIndexGrbitIgnoreNull:
		return "IgnoreNull"
	case CreateIndexGrbitIgnoreAnyNull:
		return "IgnoreAnyNull"
	case CreateIndexGrbitIgnoreFirstNull:
		return "IgnoreFirstNull"
	case CreateIndexGrbitAllowNull:
		return "AllowNull"
	case CreateIndexGrbitAllowNull2:
		return "AllowNull2"
	default:
		return "Unknown"
	}
}

func (g MoveGrbit) String() string {
	switch g {
	case MoveGrbitNone:
		return "None"
	case MoveGrbitMoveKeyNE:
		return "MoveKeyNE"
	case MoveGrbitMoveKeyGT:
		return "MoveKeyGT"
	case MoveGrbitMoveKeyGE:
		return "MoveKeyGE"
	case MoveGrbitMoveKeyLT:
		return "MoveKeyLT"
	case MoveGrbitMoveKeyLE:
		return "MoveKeyLE"
	default:
		return "Unknown"
	}
}

func (g SeekGrbit) String() string {
	switch g {
	case SeekGrbitNone:
		return "None"
	case SeekGrbitSeekEQ:
		return "SeekEQ"
	case SeekGrbitSeekGE:
		return "SeekGE"
	case SeekGrbitSeekGT:
		return "SeekGT"
	case SeekGrbitSeekLE:
		return "SeekLE"
	case SeekGrbitSeekLT:
		return "SeekLT"
	case SeekGrbitSeekNE:
		return "SeekNE"
	case SeekGrbitSetIndexRange:
		return "SetIndexRange"
	default:
		return "Unknown"
	}
}

func (g MakeKeyGrbit) String() string {
	switch g {
	case MakeKeyGrbitNone:
		return "None"
	case MakeKeyGrbitNewKey:
		return "NewKey"
	case MakeKeyGrbitStrLimit:
		return "StrLimit"
	case MakeKeyGrbitSubStrLimit:
		return "SubStrLimit"
	case MakeKeyGrbitNormalizedKey:
		return "NormalizedKey"
	case MakeKeyGrbitKeyDataZeroLength:
		return "KeyDataZeroLength"
	default:
		return "Unknown"
	}
}
