// Code generated by ent, DO NOT EDIT.

package onemethodservice

const (
	// Label holds the string label denoting the onemethodservice type in the database.
	Label = "one_method_service"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the onemethodservice in the database.
	Table = "one_method_services"
)

// Columns holds all SQL columns for onemethodservice fields.
var Columns = []string{
	FieldID,
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
