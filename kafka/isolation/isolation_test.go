package isolation

import (
	"fmt"
	"github.com/arcosx/Sirius/util"
	"testing"
	"time"
)

func init() {
	util.SetConfig()
}
func TestInitIsolationSet(t *testing.T) {
	InitIsolationSet()
	t.Log("\n" + IsolationSet.String())
}

func TestIsolation_ProduceAsync(t *testing.T) {
	InitIsolationSet()
	isolation, err := IsolationSet.GetIsolation("test")
	if err != nil {
		t.Fatal(err)
	}
	isolation.ProduceAsync("test", []byte("TestIsolation_ProduceAsync"+util.GetTimeNowString()))
	time.Sleep(1 * time.Second)
}

func TestSet_GetIsolation(t *testing.T) {
	InitIsolationSet()
	isolation, err := IsolationSet.GetIsolation("test")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(isolation)
}
