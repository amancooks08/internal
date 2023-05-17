package service

import (
	"nickPay/wallet/internal/domain"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	type args struct {
		loginResponse domain.LoginDbResponse
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Generate Token",	
			args: args{
				loginResponse: domain.LoginDbResponse{
					ID: 1,
					Password: "12345678",
				},
			},
			want: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dp",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateToken(tt.args.loginResponse)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
