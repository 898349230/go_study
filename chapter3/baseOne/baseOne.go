package baseOne

import (
	"fmt"
)

type BaseOne struct {
	Name string
}

func (baseOne *BaseOne) PrintStr(str string) {
	fmt.Println("BaseOne name: ", baseOne.Name, " ", str)
}
