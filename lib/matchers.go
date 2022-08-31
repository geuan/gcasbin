package lib

import (
	"fmt"
	"strings"
)

var ADMINS=[]string{"admin","root"}

// 自定义matcher
func init() {
	E.AddFunction("methodMatch", func(arguments ...interface{}) (i interface{}, e error) {
		if len(arguments) == 2 {
			k1, k2 := arguments[0].(string), arguments[1].(string)
			return MethodMatch(k1, k2), nil
		}
		return nil, fmt.Errorf("methodMatch error.")
	})

	E.AddFunction("isSuperAdmin", func(arguments ...interface{}) (i interface{}, e error) {
		if len(arguments) == 1 {
			k1 := arguments[0].(string)
			return IsSuperAdmin(k1),nil
		}
		return nil, fmt.Errorf("isSuperAdmin error")
	})

}

func MethodMatch(key1, key2 string) bool {
	ks := strings.Split(key2, " ")
	for _, k := range ks {
		if k == key1 {
			return true
		}
	}
	return false
}

func IsSuperAdmin(userName string) bool  {
	for _,user := range ADMINS {
		if user == userName {
			return true
		}
	}
	return false
}
