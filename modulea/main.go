package modulea

import (
	"fmt"

	"github.com/shirakiya/workspace-test/internal"
)

func Func1() string {
	return fmt.Sprintf("modulea.Func1 + %s", internal.Func1())
}
