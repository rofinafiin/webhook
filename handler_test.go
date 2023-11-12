package webhook

import (
	"fmt"
	"testing"
)

func TestGetRandomString(t *testing.T) {
	stringssss := GetRandomString(stringreply)
	fmt.Println(stringssss)
}
