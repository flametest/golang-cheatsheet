// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldMoney holds the string denoting the money field in the database.
	FieldMoney = "money"
	// FieldMeta holds the string denoting the meta field in the database.
	FieldMeta = "meta"

	// Table holds the table name of the user in the database.
	Table = "User"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldMoney,
	FieldMeta,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultMoney holds the default value on creation for the money field.
	DefaultMoney int64
)
