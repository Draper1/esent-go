package types

// Coltyp represents ESENT column types
type Coltyp uint32

const (
	// ColtypNil - Null column type. Invalid for column creation.
	ColtypNil Coltyp = 0

	// ColtypBit - True, False or NULL.
	ColtypBit Coltyp = 1

	// ColtypUnsignedByte - 1-byte integer, unsigned.
	ColtypUnsignedByte Coltyp = 2

	// ColtypShort - 2-byte integer, signed.
	ColtypShort Coltyp = 3

	// ColtypLong - 4-byte integer, signed.
	ColtypLong Coltyp = 4

	// ColtypCurrency - 8-byte integer, signed.
	ColtypCurrency Coltyp = 5

	// ColtypIEEESingle - 4-byte IEEE single-precision.
	ColtypIEEESingle Coltyp = 6

	// ColtypIEEEDouble - 8-byte IEEE double-precision.
	ColtypIEEEDouble Coltyp = 7

	// ColtypDateTime - Integral date, fractional time.
	ColtypDateTime Coltyp = 8

	// ColtypBinary - Binary data, up to 255 bytes.
	ColtypBinary Coltyp = 9

	// ColtypText - Text data, up to 255 bytes.
	ColtypText Coltyp = 10

	// ColtypLongBinary - Binary data, up to 2GB.
	ColtypLongBinary Coltyp = 11

	// ColtypLongText - Text data, up to 2GB.
	ColtypLongText Coltyp = 12
)

// String returns the string representation of the column type
func (c Coltyp) String() string {
	switch c {
	case ColtypNil:
		return "Nil"
	case ColtypBit:
		return "Bit"
	case ColtypUnsignedByte:
		return "UnsignedByte"
	case ColtypShort:
		return "Short"
	case ColtypLong:
		return "Long"
	case ColtypCurrency:
		return "Currency"
	case ColtypIEEESingle:
		return "IEEESingle"
	case ColtypIEEEDouble:
		return "IEEEDouble"
	case ColtypDateTime:
		return "DateTime"
	case ColtypBinary:
		return "Binary"
	case ColtypText:
		return "Text"
	case ColtypLongBinary:
		return "LongBinary"
	case ColtypLongText:
		return "LongText"
	default:
		return "Unknown"
	}
}
