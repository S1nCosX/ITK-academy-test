package operationtype

type InvalidOperationTypeString struct {
}

func (InvalidOperationTypeString) Error() string {
	return "Invalid string value for operationType"
}

type InvalidOperationTypeValue struct {
}

func (InvalidOperationTypeValue) Error() string {
	return "Invalid int value of operationType"
}
