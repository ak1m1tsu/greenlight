package data

import (
	"testing"
)

func TestRuntime_UnmarshalJSON(t *testing.T) {
	type args struct {
		value []byte
	}
	tests := []struct {
		name     string
		r        Runtime
		args     args
		wantErr  bool
		expValue Runtime
	}{
		{
			name: "Negative value",
			r:    Runtime(0),
			args: args{value: []byte("\"-100 mins\"")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.UnmarshalJSON(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Runtime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
