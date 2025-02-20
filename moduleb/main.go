package moduleb

import (
	"fmt"

	"github.com/shirakiya/workspace-test/src"
)

func Func1() string {
	return fmt.Sprintf("moduleb.Func1 + %s", src.Func1())
}
