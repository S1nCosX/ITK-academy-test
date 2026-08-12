package operationtype

type InvalidOperationTypeString struct {
}

func (InvalidOperationTypeString) Error() string {
	return "Invalid string value for operationType"
}
