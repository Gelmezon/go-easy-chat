package job

import (
	"fmt"
	"reflect"
	"time"
)

type TestJob struct {
}

func (job *TestJob) Execute(name string) {
	fmt.Println(name, " ", time.Now().Format("2006-01-02 15:04:05"))
}

func (job *TestJob) Register() {
	TypeRegistry["TestJob"] = reflect.TypeOf(job)
}

func init() {
	TypeRegistry["TestJob"] = reflect.TypeOf(TestJob{})
}
