package cmd

import (
	"reflect"
	"testing"
)

func TestParseFromParts(t *testing.T) {
	type args struct {
		parts []string
	}
	tests := []struct {
		name    string
		args    args
		want    Args
		wantErr bool
	}{
		{
			name: "parses project",
			args: args{[]string{"project"}},
			want: Args{"project", ""},
		},
		{
			name: "parses project and branch",
			args: args{[]string{"project", "branch"}},
			want: Args{"project", "branch"},
		},
		{
			name: "parses repository url",
			args: args{[]string{"https://github.com/Spuffynism/pop"}},
			want: Args{"pop", ""},
		},
		{
			name: "parses repository url with trailing slash",
			args: args{[]string{"https://github.com/Spuffynism/pop/"}},
			want: Args{"pop", ""},
		},
		{
			name: "parses repository url with trailing slash in branch",
			args: args{[]string{"https://github.com/Spuffynism/pop/tree/branch/"}},
			want: Args{"pop", "branch"},
		},
		{
			name: "parses repository url with branch in url",
			args: args{[]string{"https://github.com/Spuffynism/pop/tree/branch"}},
			want: Args{"pop", "branch"},
		},
		{
			name: "parses repository url with branch as standalone part",
			args: args{[]string{"https://github.com/Spuffynism/pop", "branch"}},
			want: Args{"pop", "branch"},
		},
		{
			name: "parses branch as standalone part with precedence over branch in url",
			args: args{[]string{"https://github.com/Spuffynism/pop/tree/ignored", "branch"}},
			want: Args{"pop", "branch"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFromParts(tt.args.parts)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFromParts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFromParts() got = %v, want %v", got, tt.want)
			}
		})
	}
}
