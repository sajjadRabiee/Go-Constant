package constant

import (
	"errors"
	"fmt"
)

type constInterface[v any] interface {
	fmt.Stringer
	GetMembers() []v
	GetDefault() v
}

func ToConst[v constInterface[v]](value string) (v, error) {
	var constInterfaceInstance v
	for _, member := range constInterfaceInstance.GetMembers() {
		if member.String() == value {
			return member, nil
		}
	}

	return constInterfaceInstance.GetDefault(), errors.New("value doesn't exist")
}
