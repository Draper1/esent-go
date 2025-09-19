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
