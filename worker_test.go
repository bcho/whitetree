package whitetree_test

import (
	"testing"

	"github.com/bcho/whitetree"
)

func TestWorkGroupSpawnWithPanic(t *testing.T) {
	panicParser := func(whitetree.TaskContext) (whitetree.HandlerId, error) {
		panic(1)
	}

	w := whitetree.NewWorkerGroup(
		1,
		&whitetree.ParserPackage{"panic parser": panicParser},
		&whitetree.HandlerPackage{},
	)

	go w.Spawn(whitetree.NewTaskContext([]byte("test ctx")))
	err := w.Wait()
	if err == nil {
		t.Errorf("should recover from panic")
	}
}
