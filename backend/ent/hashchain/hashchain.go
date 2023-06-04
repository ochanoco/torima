// Code generated by ent, DO NOT EDIT.

package hashchain

const (
	// Label holds the string label denoting the hashchain type in the database.
	Label = "hash_chain"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldSignature holds the string denoting the signature field in the database.
	FieldSignature = "signature"
	// EdgeLog holds the string denoting the log edge name in mutations.
	EdgeLog = "log"
	// Table holds the table name of the hashchain in the database.
	Table = "hash_chains"
	// LogTable is the table that holds the log relation/edge.
	LogTable = "hash_chains"
	// LogInverseTable is the table name for the ServiceLog entity.
	// It exists in this package in order to avoid circular dependency with the "servicelog" package.
	LogInverseTable = "service_logs"
	// LogColumn is the table column denoting the log relation/edge.
	LogColumn = "service_log_hashchains"
)

// Columns holds all SQL columns for hashchain fields.
var Columns = []string{
	FieldID,
	FieldHash,
	FieldSignature,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "hash_chains"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"service_log_hashchains",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}