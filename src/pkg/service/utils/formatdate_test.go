package utils

import (
	//	"inine-track/pkg/service/utils" -- Ciclo desnecessário de importação
	"testing"
	"time"
)

func TestFormateDate(t *testing.T) {
	type args struct {
		data1 string
		data2 string
	}

	layout := "2006-01-02"

	tests := []struct {
		name string
		args args
		err  bool
	}{
		{
			name: "VALID",
			args: args{
				data1: "2025-04-01",
				data2: "2025-04-30",
			},
			err: false,
		},
		{
			name: "INVALID",
			args: args{
				data1: "2025/04/01",
				data2: "2025/04/30",
			},
			err: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2, err := FormateDate(tt.args.data1, tt.args.data2)

			if (err != nil) != tt.err {
				t.Errorf("FormateDate() error = %v, wantErr %v", err, tt.err)
				return
			}

			if !tt.err {
				expected1, _ := time.Parse(layout, tt.args.data1)
				expected2, _ := time.Parse(layout, tt.args.data2)

				if !got1.Equal(expected1) {
					t.Errorf("FormateDate() got1 = %v, want %v", got1, expected1)
				}
				if !got2.Equal(expected2) {
					t.Errorf("FormateDate() got2 = %v, want %v", got2, expected2)
				}
			}
		})
	}
}
