package utils_test

import (
	"inine-track/pkg/database"
	"inine-track/pkg/service/utils"
	"log"
	"testing"
)

func TestGetProject(t *testing.T) {
	err := database.ConnectDB()

	if err != nil {
		log.Fatal(err.Error())
	}

	type args struct {
		idProject int64
	}
	tests := []struct {
		name string
		args args
		err  bool
	}{
		{
			name: "VALID",
			args: args{
				1,
			},
			err: false,
		}, {
			name: "INVALID",
			args: args{
				999999999,
			},
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := utils.GetProject(tt.args.idProject); (err != nil) != tt.err {
				t.Errorf("GetProject() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}
