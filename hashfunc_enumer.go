// Code generated by "enumer -type HashFunc"; DO NOT EDIT.

//
package xmssmt

import (
	"fmt"
)

const _HashFuncName = "SHA2SHAKE"

var _HashFuncIndex = [...]uint8{0, 4, 9}

func (i HashFunc) String() string {
	if i >= HashFunc(len(_HashFuncIndex)-1) {
		return fmt.Sprintf("HashFunc(%d)", i)
	}
	return _HashFuncName[_HashFuncIndex[i]:_HashFuncIndex[i+1]]
}

var _HashFuncValues = []HashFunc{0, 1}

var _HashFuncNameToValueMap = map[string]HashFunc{
	_HashFuncName[0:4]: 0,
	_HashFuncName[4:9]: 1,
}

// HashFuncString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func HashFuncString(s string) (HashFunc, error) {
	if val, ok := _HashFuncNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to HashFunc values", s)
}

// HashFuncValues returns all values of the enum
func HashFuncValues() []HashFunc {
	return _HashFuncValues
}

// IsAHashFunc returns "true" if the value is listed in the enum definition. "false" otherwise
func (i HashFunc) IsAHashFunc() bool {
	for _, v := range _HashFuncValues {
		if i == v {
			return true
		}
	}
	return false
}
