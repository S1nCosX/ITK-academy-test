package operationtype

type OperationType int

const (
	DEPOSIT  OperationType = iota
	WITHDRAW OperationType = iota
)

func FromString(str string) (OperationType, error) {
	switch str {
	case "DEPOSIT":
		return DEPOSIT, nil
	case "WITHDRAW":
		return WITHDRAW, nil
	}
	err := InvalidOperationTypeString{}
	return 0, err
}

func ToString(t OperationType) (string, error) {

	switch t {
	case DEPOSIT:
		return "DEPOSIT", nil
	case WITHDRAW:
		return "WITHDRAW", nil
	}
	err := InvalidOperationTypeValue{}
	return "", err
}
