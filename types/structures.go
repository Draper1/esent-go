package types

import (
	"time"
	"unsafe"
)

// JET_SESID represents a session identifier
type JET_SESID struct {
	Value uintptr
}

// JET_DBID represents a database identifier
type JET_DBID struct {
	Value uint32
}

// JET_TABLEID represents a table identifier
type JET_TABLEID struct {
	Value uintptr
}

// JET_COLUMNID represents a column identifier
type JET_COLUMNID struct {
	Value uint32
}

// JET_INSTANCE represents an instance identifier
type JET_INSTANCE struct {
	Value uintptr
}

// JET_ERRINFOBASIC represents basic error information
type JET_ERRINFOBASIC struct {
	ErrorCode      int32
	ErrorInfo      uint32
	ErrorString    string
	SourceFile     string
	SourceLine     uint32
	SourceFunction string
}

// JET_COMMIT_ID represents a commit identifier
type JET_COMMIT_ID struct {
	CommitId [16]byte
	SignLog  [16]byte
}

// JET_RECPOS represents a record position
type JET_RECPOS struct {
	CentriesLT      uint32
	CentriesInRange uint32
	CentriesTotal   uint32
}

// JET_RECSIZE represents record size information
type JET_RECSIZE struct {
	CbData            uint64
	CbLongValueData   uint64
	CNonTaggedColumns uint32
	CTaggedColumns    uint32
	CLongValues       uint32
	CMultiValues      uint32
}

// JET_COLUMNDEF represents a column definition
type JET_COLUMNDEF struct {
	Columnid   JET_COLUMNID
	Coltyp     Coltyp
	Grbit      uint32
	CbMax      uint32
	Cp         uint32
	ColumnName string
}

// NATIVE_COLUMNDEF represents the native column definition structure
type NATIVE_COLUMNDEF struct {
	CbStruct uint32 // Size of the structure
	Columnid uint32 // Column ID
	Coltyp   uint32 // Type of the column
	WCountry uint16 // Reserved. Should be 0.
	Langid   uint16 // Obsolete. Should be 0.
	Cp       uint16 // Code page for text columns
	WCollate uint16 // Reserved. Should be 0.
	CbMax    uint32 // Maximum length of the column
	Grbit    uint32 // Column options
}

// GetNativeColumndef converts a JET_COLUMNDEF to NATIVE_COLUMNDEF
func (c *JET_COLUMNDEF) GetNativeColumndef() NATIVE_COLUMNDEF {
	return NATIVE_COLUMNDEF{
		CbStruct: uint32(unsafe.Sizeof(NATIVE_COLUMNDEF{})),
		Columnid: c.Columnid.Value,
		Coltyp:   uint32(c.Coltyp),
		WCountry: 0,            // Reserved
		Langid:   0,            // Obsolete
		Cp:       uint16(c.Cp), // Convert uint32 to uint16
		WCollate: 0,            // Reserved
		CbMax:    c.CbMax,
		Grbit:    c.Grbit,
	}
}

// JET_INDEXCREATE represents an index creation structure
type JET_INDEXCREATE struct {
	IndexName           string
	IndexKey            string
	Grbit               uint32
	UlDensity           uint32
	PidxUnicode         *JET_UNICODEINDEX
	Rgconditionalcolumn []JET_CONDITIONALCOLUMN
	Rgindexcolumn       []JET_INDEX_COLUMN
	Rgcolumnbase        []JET_COLUMNBASE
	Lcid                uint32
	Pspacehints         *JET_SPACEHINTS
	PTableCreate        *JET_TABLECREATE
}

// JET_UNICODEINDEX represents Unicode index information
type JET_UNICODEINDEX struct {
	Lcid       uint32
	DwMapFlags uint32
}

// JET_CONDITIONALCOLUMN represents a conditional column
type JET_CONDITIONALCOLUMN struct {
	Columnid JET_COLUMNID
	Grbit    uint32
}

// JET_INDEX_COLUMN represents an index column
type JET_INDEX_COLUMN struct {
	Columnid JET_COLUMNID
	Grbit    uint32
}

// JET_COLUMNBASE represents a column base
type JET_COLUMNBASE struct {
	Columnid   JET_COLUMNID
	Coltyp     Coltyp
	Grbit      uint32
	CbMax      uint32
	Cp         uint32
	ColumnName string
}

// JET_SPACEHINTS represents space hints
type JET_SPACEHINTS struct {
	UlInitialPages uint32
	UlGrowth       uint32
	UlMinimumPages uint32
	UlMaximumPages uint32
	Grbit          uint32
}

// JET_TABLECREATE represents table creation information
type JET_TABLECREATE struct {
	TableName       string
	TableNameLength uint32
	Grbit           uint32
	UlDensity       uint32
	Rgcolumncreate  []JET_COLUMNCREATE
	Rgindexcreate   []JET_INDEXCREATE
	Rgcallback      []JET_CALLBACK
	Pspacehints     *JET_SPACEHINTS
	Tableid         JET_TABLEID
}

// JET_COLUMNCREATE represents column creation information
type JET_COLUMNCREATE struct {
	Columnid    JET_COLUMNID
	Coltyp      Coltyp
	Grbit       uint32
	UlSeed      uint32
	UlIncrement uint32
	UlDefault   uint32
	UlMax       uint32
	UlMin       uint32
	CbMax       uint32
	Cp          uint32
	ColumnName  string
}

// JET_CALLBACK represents a callback function
type JET_CALLBACK struct {
	CallbackType uint32
	Callback     interface{}
}

// JET_LOGTIME represents a log time
type JET_LOGTIME struct {
	Time time.Time
}

// JET_LGPOS represents a log position
type JET_LGPOS struct {
	LGeneration uint32
	ISector     uint32
	IB          uint16
}

// JET_SIGNATURE represents a database signature
type JET_SIGNATURE struct {
	LogtimeCreate JET_LOGTIME
	LGeneration   uint32
	UlRandom      uint32
}

// JET_SETINFO represents set column information
type JET_SETINFO struct {
	ItagSequence uint32
	IbLongValue  uint32
}

// JET_RETINFO represents retrieve column information
type JET_RETINFO struct {
	ItagSequence       uint32
	IbLongValue        uint32
	ColumnidNextTagged uint32
}

// JET_DBINFOMISC represents miscellaneous database information
type JET_DBINFOMISC struct {
	UlVersion          uint32
	UlUpdate           uint32
	SignDb             JET_SIGNATURE
	Dbstate            uint32
	LgposConsistent    JET_LGPOS
	LogtimeConsistent  JET_LOGTIME
	LogtimeAttach      JET_LOGTIME
	LgposAttach        JET_LGPOS
	LogtimeDetach      JET_LOGTIME
	LgposDetach        JET_LGPOS
	LgposLogRecovery   JET_LGPOS
	LogtimeLogRecovery JET_LOGTIME
	UlConsistent       uint32
	UlAttach           uint32
	UlDetach           uint32
	UlLogRecovery      uint32
	UlLogRecovery2     uint32
	UlLogRecovery3     uint32
	UlLogRecovery4     uint32
	UlLogRecovery5     uint32
	UlLogRecovery6     uint32
	UlLogRecovery7     uint32
	UlLogRecovery8     uint32
	UlLogRecovery9     uint32
	UlLogRecovery10    uint32
	UlLogRecovery11    uint32
	UlLogRecovery12    uint32
	UlLogRecovery13    uint32
	UlLogRecovery14    uint32
	UlLogRecovery15    uint32
	UlLogRecovery16    uint32
	UlLogRecovery17    uint32
	UlLogRecovery18    uint32
	UlLogRecovery19    uint32
	UlLogRecovery20    uint32
	UlLogRecovery21    uint32
	UlLogRecovery22    uint32
	UlLogRecovery23    uint32
	UlLogRecovery24    uint32
	UlLogRecovery25    uint32
	UlLogRecovery26    uint32
	UlLogRecovery27    uint32
	UlLogRecovery28    uint32
	UlLogRecovery29    uint32
	UlLogRecovery30    uint32
	UlLogRecovery31    uint32
	UlLogRecovery32    uint32
	UlLogRecovery33    uint32
	UlLogRecovery34    uint32
	UlLogRecovery35    uint32
	UlLogRecovery36    uint32
	UlLogRecovery37    uint32
	UlLogRecovery38    uint32
	UlLogRecovery39    uint32
	UlLogRecovery40    uint32
	UlLogRecovery41    uint32
	UlLogRecovery42    uint32
	UlLogRecovery43    uint32
	UlLogRecovery44    uint32
	UlLogRecovery45    uint32
	UlLogRecovery46    uint32
	UlLogRecovery47    uint32
	UlLogRecovery48    uint32
	UlLogRecovery49    uint32
	UlLogRecovery50    uint32
	UlLogRecovery51    uint32
	UlLogRecovery52    uint32
	UlLogRecovery53    uint32
	UlLogRecovery54    uint32
	UlLogRecovery55    uint32
	UlLogRecovery56    uint32
	UlLogRecovery57    uint32
	UlLogRecovery58    uint32
	UlLogRecovery59    uint32
	UlLogRecovery60    uint32
	UlLogRecovery61    uint32
	UlLogRecovery62    uint32
	UlLogRecovery63    uint32
	UlLogRecovery64    uint32
	UlLogRecovery65    uint32
	UlLogRecovery66    uint32
	UlLogRecovery67    uint32
	UlLogRecovery68    uint32
	UlLogRecovery69    uint32
	UlLogRecovery70    uint32
	UlLogRecovery71    uint32
	UlLogRecovery72    uint32
	UlLogRecovery73    uint32
	UlLogRecovery74    uint32
	UlLogRecovery75    uint32
	UlLogRecovery76    uint32
	UlLogRecovery77    uint32
	UlLogRecovery78    uint32
	UlLogRecovery79    uint32
	UlLogRecovery80    uint32
	UlLogRecovery81    uint32
	UlLogRecovery82    uint32
	UlLogRecovery83    uint32
	UlLogRecovery84    uint32
	UlLogRecovery85    uint32
	UlLogRecovery86    uint32
	UlLogRecovery87    uint32
	UlLogRecovery88    uint32
	UlLogRecovery89    uint32
	UlLogRecovery90    uint32
	UlLogRecovery91    uint32
	UlLogRecovery92    uint32
	UlLogRecovery93    uint32
	UlLogRecovery94    uint32
	UlLogRecovery95    uint32
	UlLogRecovery96    uint32
	UlLogRecovery97    uint32
	UlLogRecovery98    uint32
	UlLogRecovery99    uint32
	UlLogRecovery100   uint32
}

// JET_param represents ESENT system parameters
type JET_param uint32

const (
	// SystemPath - This parameter indicates the relative or absolute file system path of the
	// folder that will contain the checkpoint file for the instance. The path
	// must be terminated with a backslash character, which indicates that the
	// target path is a folder.
	JET_paramSystemPath JET_param = 0

	// TempPath - This parameter indicates the relative or absolute file system path of
	// the folder or file that will contain the temporary database for the instance.
	// If the path is to a folder that will contain the temporary database then it
	// must be terminated with a backslash character.
	JET_paramTempPath JET_param = 1

	// LogFilePath - This parameter indicates the relative or absolute file system path of the
	// folder that will contain the transaction logs for the instance. The path must
	// be terminated with a backslash character, which indicates that the target path
	// is a folder.
	JET_paramLogFilePath JET_param = 2

	// BaseName - This parameter sets the three letter prefix used for many of the files used by
	// the database engine. For example, the checkpoint file is called EDB.CHK by
	// default because EDB is the default base name.
	JET_paramBaseName JET_param = 3

	// EventSource - This parameter supplies an application specific string that will be added to
	// any event log messages that are emitted by the database engine. This allows
	// easy correlation of event log messages with the source application. By default
	// the host application executable name will be used.
	JET_paramEventSource JET_param = 4

	// MaxSessions - This parameter reserves the requested number of session resources for use by an
	// instance. A session resource directly corresponds to a JET_SESID data type.
	// This setting will affect how many sessions can be used at the same time.
	JET_paramMaxSessions JET_param = 5

	// MaxOpenTables - This parameter reserves the requested number of B+ Tree resources for use by
	// an instance. This setting will affect how many tables can be used at the same time.
	JET_paramMaxOpenTables JET_param = 6

	// MaxCursors - This parameter reserves the requested number of cursor resources for use by an
	// instance. A cursor resource directly corresponds to a JET_TABLEID data type.
	// This setting will affect how many cursors can be used at the same time. A cursor
	// resource cannot be shared by different sessions so this parameter must be set to
	// a large enough value so that each session can use as many cursors as are required.
	JET_paramMaxCursors JET_param = 8

	// MaxVerPages - This parameter reserves the requested number of version store pages for use by an instance.
	JET_paramMaxVerPages JET_param = 9

	// MaxTemporaryTables - This parameter reserves the requested number of temporary table resources for use
	// by an instance. This setting will affect how many temporary tables can be used at
	// the same time. If this system parameter is set to zero then no temporary database
	// will be created and any activity that requires use of the temporary database will
	// fail. This setting can be useful to avoid the I/O required to create the temporary
	// database if it is known that it will not be used.
	JET_paramMaxTemporaryTables JET_param = 10

	// LogFileSize - This parameter will configure the size of the transaction log files. Each
	// transaction log file is a fixed size. The size is equal to the setting of
	// this system parameter in units of 1024 bytes.
	JET_paramLogFileSize JET_param = 11

	// LogBuffers - This parameter will configure the amount of memory used to cache log records
	// before they are written to the transaction log file. The unit for this
	// parameter is the sector size of the volume that holds the transaction log files.
	// The sector size is almost always 512 bytes, so it is safe to assume that size
	// for the unit. This parameter has an impact on performance. When the database
	// engine is under heavy update load, this buffer can become full very rapidly.
	// A larger cache size for the transaction log file is critical for good update
	// performance under such a high load condition. The default is known to be too small
	// for this case.
	// Do not set this parameter to a number of buffers that is larger (in bytes) than
	// half the size of a transaction log file.
	JET_paramLogBuffers JET_param = 12

	// CircularLog - This parameter configures how transaction log files are managed by the database
	// engine. When circular logging is off, all transaction log files that are generated
	// are retained on disk until they are no longer needed because a full backup of the
	// database has been performed. When circular logging is on, only transaction log files
	// that are younger than the current checkpoint are retained on disk. The benefit of
	// this mode is that backups are not required to retire old transaction log files.
	JET_paramCircularLog JET_param = 17

	// DbExtensionSize - This parameter controls the amount of space that is added to a database file each
	// time it needs to grow to accommodate more data. The size is in database pages.
	JET_paramDbExtensionSize JET_param = 18

	// PageTempDBMin - This parameter controls the initial size of the temporary database. The size is in
	// database pages. A size of zero indicates that the default size of an ordinary
	// database should be used. It is often desirable for small applications to configure
	// the temporary database to be as small as possible. Setting this parameter to
	// SystemParameters.PageTempDBSmallest will achieve the smallest temporary database possible.
	JET_paramPageTempDBMin JET_param = 19

	// CacheSizeMax - This parameter configures the maximum size of the database page cache. The size
	// is in database pages. If this parameter is left to its default value, then the
	// maximum size of the cache will be set to the size of physical memory when JetInit
	// is called.
	JET_paramCacheSizeMax JET_param = 23

	// CheckpointDepthMax - This parameter controls how aggressively database pages are flushed from the
	// database page cache to minimize the amount of time it will take to recover from a
	// crash. The parameter is a threshold in bytes for about how many transaction log
	// files will need to be replayed after a crash. If circular logging is enabled using
	// JET_param.CircularLog then this parameter will also control the approximate amount
	// of transaction log files that will be retained on disk.
	JET_paramCheckpointDepthMax JET_param = 24

	// LrukCorrInterval - This parameter controls the correlation interval of ESE's LRU-K page replacement
	// algorithm.
	JET_paramLrukCorrInterval JET_param = 25

	// LrukTimeout - This parameter controls the timeout interval of ESE's LRU-K page replacement
	// algorithm.
	JET_paramLrukTimeout JET_param = 28

	// OutstandingIOMax - This parameter controls how many database file I/Os can be queued
	// per-disk in the host operating system at one time.  A larger value
	// for this parameter can significantly help the performance of a large
	// database application.
	JET_paramOutstandingIOMax JET_param = 30

	// StartFlushThreshold - This parameter controls when the database page cache begins evicting pages from the
	// cache to make room for pages that are not cached. When the number of page buffers in the cache
	// drops below this threshold then a background process will be started to replenish that pool
	// of available buffers. This threshold is always relative to the maximum cache size as set by
	// JET_paramCacheSizeMax. This threshold must also always be less than the stop threshold as
	// set by JET_paramStopFlushThreshold.
	JET_paramStartFlushThreshold JET_param = 31

	// StopFlushThreshold - This parameter controls when the database page cache ends evicting pages from the cache to make
	// room for pages that are not cached. When the number of page buffers in the cache rises above
	// this threshold then the background process that was started to replenish that pool of available
	// buffers is stopped. This threshold is always relative to the maximum cache size as set by
	// JET_paramCacheSizeMax. This threshold must also always be greater than the start threshold
	// as set by JET_paramStartFlushThreshold.
	JET_paramStopFlushThreshold JET_param = 32

	// Recovery - This parameter is the master switch that controls crash recovery for an instance.
	// If this parameter is set to "On" then ARIES style recovery will be used to bring all
	// databases in the instance to a consistent state in the event of a process or machine
	// crash. If this parameter is set to "Off" then all databases in the instance will be
	// managed without the benefit of crash recovery. That is to say, that if the instance
	// is not shut down cleanly using JetTerm prior to the process exiting or machine shutdown
	// then the contents of all databases in that instance will be corrupted.
	JET_paramRecovery JET_param = 34

	// EnableOnlineDefrag - This parameter controls the behavior of online defragmentation when initiated using
	// JetDefragment and JetDefragment2.
	JET_paramEnableOnlineDefrag JET_param = 35

	// CacheSize - This parameter can be used to control the size of the database page cache at run time.
	// Ordinarily, the cache will automatically tune its size as a function of database and
	// machine activity levels. If the application sets this parameter to zero, then the cache
	// will tune its own size in this manner. However, if the application sets this parameter
	// to a non-zero value then the cache will adjust itself to that target size.
	JET_paramCacheSize JET_param = 41

	// EnableIndexChecking - When this parameter is true, every database is checked at JetAttachDatabase time for
	// indexes over Unicode key columns that were built using an older version of the NLS
	// library in the operating system. This must be done because the database engine persists
	// the sort keys generated by LCMapStringW and the value of these sort keys change from release to release.
	// If a primary index is detected to be in this state then JetAttachDatabase will always fail with
	// JET_err.PrimaryIndexCorrupted.
	// If any secondary indexes are detected to be in this state then there are two possible outcomes.
	// If AttachDatabaseGrbit.DeleteCorruptIndexes was passed to JetAttachDatabase then these indexes
	// will be deleted and JET_wrnCorruptIndexDeleted will be returned from JetAttachDatabase. These
	// indexes will need to be recreated by your application. If AttachDatabaseGrbit.DeleteCorruptIndexes
	// was not passed to JetAttachDatabase then the call will fail with JET_errSecondaryIndexCorrupted.
	JET_paramEnableIndexChecking JET_param = 45

	// EventSourceKey - This parameter can be used to control which event log the database engine uses for its event log
	// messages. By default, all event log messages will go to the Application event log. If the registry
	// key name for another event log is configured then the event log messages will go there instead.
	JET_paramEventSourceKey JET_param = 49

	// NoInformationEvent - When this parameter is true, informational event log messages that would ordinarily be generated by
	// the database engine will be suppressed.
	JET_paramNoInformationEvent JET_param = 50

	// EventLoggingLevel - Configures the detail level of eventlog messages that are emitted
	// to the eventlog by the database engine. Higher numbers will result
	// in more detailed eventlog messages.
	JET_paramEventLoggingLevel JET_param = 51

	// DeleteOutOfRangeLogs - Delete the log files that are not matching (generation wise) during soft recovery.
	JET_paramDeleteOutOfRangeLogs JET_param = 52

	// EnableIndexCleanup - After Windows 7, it was discovered that JET_paramEnableIndexCleanup had some implementation limitations, reducing its effectiveness.
	// Rather than update it to work with locale names, the functionality is removed altogether.
	// Unfortunately JET_paramEnableIndexCleanup can not be ignored altogether. JET_paramEnableIndexChecking defaults to false, so if
	// JET_paramEnableIndexCleanup were to be removed entirely, then by default there were would be no checks for NLS changes!
	// The current behavious (when enabled) is to track the language sort versions for the indices, and when the sort version for that
	// particular locale changes, the engine knows which indices are now invalid. For example, if the sort version for only "de-de" changes,
	// then the "de-de" indices are invalid, but the "en-us" indices will be fine.
	// Post-Windows 8:
	// JET_paramEnableIndexChecking accepts JET_INDEXCHECKING (which is an enum). The values of '0' and '1' have the same meaning as before,
	// but '2' is JET_IndexCheckingDeferToOpenTable, which means that the NLS up-to-date-ness is NOT checked when the database is attached.
	// It is deferred to JetOpenTable(), which may now fail with JET_errPrimaryIndexCorrupted or JET_errSecondaryIndexCorrupted (which
	// are NOT actual corruptions, but instead reflect an NLS sort change).
	// IN SUMMARY:
	// New code should explicitly set both IndexChecking and IndexCleanup to the same value.
	JET_paramEnableIndexCleanup JET_param = 54

	// CacheSizeMin - This parameter configures the minimum size of the database page cache. The size is in database pages.
	JET_paramCacheSizeMin JET_param = 60

	// PreferredVerPages - This parameter represents a threshold relative to JET_param.MaxVerPages that controls
	// the discretionary use of version pages by the database engine. If the size of the version store exceeds
	// this threshold then any information that is only used for optional background tasks, such as reclaiming
	// deleted space in the database, is instead sacrificed to preserve room for transactional information.
	JET_paramPreferredVerPages JET_param = 63

	// DatabasePageSize - This parameter configures the page size for the database. The page
	// size is the smallest unit of space allocation possible for a database
	// file. The database page size is also very important because it sets
	// the upper limit on the size of an individual record in the database.
	// Only one database page size is supported per process at this time.
	// This means that if you are in a single process that contains different
	// applications that use the database engine then they must all agree on
	// a database page size.
	JET_paramDatabasePageSize JET_param = 64

	// ErrorToString - This parameter can be used to convert a JET_ERR into a string.
	// This should only be used with JetGetSystemParameter.
	JET_paramErrorToString JET_param = 70

	// RuntimeCallback - Configures the engine with a JET_CALLBACK delegate.
	// This callback may be called for the following reasons:
	// JET_cbtyp.FreeCursorLS, JET_cbtyp.FreeTableLS
	// or JET_cbtyp.Null. See JetSetLS
	// for more information. This parameter cannot currently be retrieved.
	JET_paramRuntimeCallback JET_param = 73

	// CleanupMismatchedLogFiles - This parameter controls the outcome of JetInit when the database
	// engine is configured to start using transaction log files on disk
	// that are of a different size than what is configured. Normally,
	// JetInit will successfully recover the databases
	// but will fail with JET_err.LogFileSizeMismatchDatabasesConsistent
	// to indicate that the log file size is misconfigured. However, when
	// this parameter is set to true then the database engine will silently
	// delete all the old log files, start a new set of transaction log files
	// using the configured log file size. This parameter is useful when the
	// application wishes to transparently change its transaction log file
	// size yet still work transparently in upgrade and restore scenarios.
	JET_paramCleanupMismatchedLogFiles JET_param = 77

	// ExceptionAction - This parameter controls what happens when an exception is thrown by the
	// database engine or code that is called by the database engine. When set
	// to JET_ExceptionMsgBox, any exception will be thrown to the Windows unhandled
	// exception filter. This will result in the exception being handled as an
	// application failure. The intent is to prevent application code from erroneously
	// trying to catch and ignore an exception generated by the database engine.
	// This cannot be allowed because database corruption could occur. If the application
	// wishes to properly handle these exceptions then the protection can be disabled
	// by setting this parameter to JET_ExceptionNone.
	JET_paramExceptionAction JET_param = 98

	// CreatePathIfNotExist - When this parameter is set to true then any folder that is missing in a file system path in use by
	// the database engine will be silently created. Otherwise, the operation that uses the missing file system
	// path will fail with JET_err.InvalidPath.
	JET_paramCreatePathIfNotExist JET_param = 100

	// OneDatabasePerSession - When this parameter is true then only one database is allowed to
	// be opened using JetOpenDatabase by a given session at one time.
	// The temporary database is excluded from this restriction.
	JET_paramOneDatabasePerSession JET_param = 102

	// MaxInstances - This parameter controls the maximum number of instances that can be created in a single process.
	JET_paramMaxInstances JET_param = 104

	// VersionStoreTaskQueueMax - This parameter controls the number of background cleanup work items that
	// can be queued to the database engine thread pool at any one time.
	JET_paramVersionStoreTaskQueueMax JET_param = 105

	// DisablePerfmon - This parameter controls whether perfmon counters should be enabled or not.
	// By default, perfmon counters are enabled, but there is memory overhead for enabling
	// them.
	JET_paramDisablePerfmon JET_param = 107
)

// JET_DbInfo represents info levels for retrieving database information
type JET_DbInfo uint32

const (
	// Filename - Returns the path to the database file (string)
	JET_DbInfoFilename JET_DbInfo = 0

	// LCID - Returns the locale identifier (LCID) associated with this database (Int32)
	JET_DbInfoLCID JET_DbInfo = 3

	// Options - Returns a OpenDatabaseGrbit. This indicates whether the
	// database is opened in exclusive mode. If the database is in exclusive mode then
	// OpenDatabaseGrbit.Exclusive will be returned, otherwise zero is
	// returned. Other database grbit options for JetAttachDatabase and JetOpenDatabase
	// are not returned.
	JET_DbInfoOptions JET_DbInfo = 6

	// Transactions - Returns a number one greater than the maximum level to which transactions can be
	// nested. If JetBeginTransaction is called (in a nesting fashion, that is, on the
	// same session, without a commit or rollback) as many times as this value, on the
	// last call JET_err.TransTooDeep will be returned (Int32).
	JET_DbInfoTransactions JET_DbInfo = 7

	// Version - Returns the major version of the database engine (Int32)
	JET_DbInfoVersion JET_DbInfo = 8

	// Filesize - Returns the filesize of the database, in pages (Int32)
	JET_DbInfoFilesize JET_DbInfo = 10

	// SpaceOwned - Returns the owned space of the database, in pages (Int32)
	JET_DbInfoSpaceOwned JET_DbInfo = 11

	// SpaceAvailable - Returns the available space in the database, in pages (Int32)
	JET_DbInfoSpaceAvailable JET_DbInfo = 12

	// Misc - Returns a JET_DBINFOMISC object
	JET_DbInfoMisc JET_DbInfo = 14

	// DBInUse - Returns a boolean indicating whether the database is attached (boolean)
	JET_DbInfoDBInUse JET_DbInfo = 15

	// PageSize - Returns the page size of the database (Int32)
	JET_DbInfoPageSize JET_DbInfo = 17

	// FileType - Returns the type of the database (JET_filetype)
	JET_DbInfoFileType JET_DbInfo = 19
)
