package utilLog_test

import (
	"bmachado/Boaz.Api.Go/util"
	"testing"
)

// LogInfoMsg call
// for a valid return value message.
func TestLogInfoMsg(t *testing.T) {
	name := util.LogInfoMsg("aaa")
	if name == "aa" {
		t.Error("Erro teste")
	}
}

// LogInfo call
// for a valid not return fail.
func TestLogInfo(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "aa",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			util.LogInfo(tt.args.obj)
		})
	}
	// Output:
	// Error loading .env file
}
