// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build unit

package validate

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestURL(t *testing.T) {
	tests := []struct {
		name    string
		val     interface{}
		wantErr bool
	}{
		{
			name:    "Valid URL",
			val:     "http://test.com/",
			wantErr: false,
		},
		{
			name:    "invalid value",
			val:     1,
			wantErr: true,
		},
		{
			name:    "missing trailing slash",
			val:     "http://test.com",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		val := tt.val
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := URL(val); (err != nil) != wantErr {
				t.Errorf("URL() error = %v, wantErr %v", err, wantErr)
			}
		})
	}
}

func TestOptionalObjectID(t *testing.T) {
	tests := []struct {
		name    string
		val     interface{}
		wantErr bool
	}{
		{
			name:    "Empty value",
			val:     "",
			wantErr: false,
		},
		{
			name:    "nil value",
			val:     nil,
			wantErr: false,
		},
		{
			name:    "Valid ObjectID",
			val:     "5e9f088b4797476aa0a5d56a",
			wantErr: false,
		},
		{
			name:    "Short ObjectID",
			val:     "5e9f088b4797476aa0a5d56",
			wantErr: true,
		},
		{
			name:    "Invalid ObjectID",
			val:     "5e9f088b4797476aa0a5d56z",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		val := tt.val
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := OptionalObjectID(val); (err != nil) != wantErr {
				t.Errorf("OptionalObjectID() error = %v, wantErr %v", err, wantErr)
			}
		})
	}
}

func TestObjectID(t *testing.T) {
	tests := []struct {
		name    string
		val     string
		wantErr bool
	}{
		{
			name:    "Empty value",
			val:     "",
			wantErr: true,
		},
		{
			name:    "Valid ObjectID",
			val:     "5e9f088b4797476aa0a5d56a",
			wantErr: false,
		},
		{
			name:    "Short ObjectID",
			val:     "5e9f088b4797476aa0a5d56",
			wantErr: true,
		},
		{
			name:    "Invalid ObjectID",
			val:     "5e9f088b4797476aa0a5d56z",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		val := tt.val
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := ObjectID(val); (err != nil) != wantErr {
				t.Errorf("OptionalObjectID() error = %v, wantErr %v", err, wantErr)
			}
		})
	}
}

func TestCredentials(t *testing.T) {
	t.Run("no credentials", func(t *testing.T) {
		err := Credentials()
		if err == nil {
			t.Fatal("Credentials() expected an error\n")
		}
	})
	t.Run("with credentials", func(t *testing.T) {
		// this function depends on the global config (globals are bad I know)
		// the easiest way we have to test it is via ENV vars
		viper.AutomaticEnv()
		_ = os.Setenv("PUBLIC_API_KEY", "test")
		_ = os.Setenv("PRIVATE_API_KEY", "test")

		err := Credentials()
		if err != nil {
			t.Fatalf("Credentials() unexpected error %v\n", err)
		}
	})
}

func TestFlagInSlice(t *testing.T) {
	t.Parallel()
	type args struct {
		value       string
		flag        string
		validValues []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "value is present",
			args: args{
				value:       "test",
				flag:        "flag",
				validValues: []string{"test", "not-test"},
			},
			wantErr: false,
		},
		{
			name: "value is present",
			args: args{
				value:       "test",
				flag:        "flag",
				validValues: []string{"not-test"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		value := tt.args.value
		flag := tt.args.flag
		validValues := tt.args.validValues
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := FlagInSlice(value, flag, validValues); (err != nil) != wantErr {
				t.Errorf("FlagInSlice() error = %v, wantErr %v", err, wantErr)
			}
		})
	}
}

func TestOptionalURL(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		val     interface{}
		wantErr bool
	}{
		{
			name:    "valid",
			val:     "http://test.com/",
			wantErr: false,
		},
		{
			name:    "empty",
			val:     "",
			wantErr: false,
		},
		{
			name:    "nil",
			val:     nil,
			wantErr: false,
		},
		{
			name:    "invalid value",
			val:     1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		val := tt.val
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := OptionalURL(val); (err != nil) != wantErr {
				t.Errorf("OptionalURL() error = %v, wantErr %v", err, wantErr)
			}
		})
	}
}

func TestPath(t *testing.T) {
	f, err := ioutil.TempFile("", "sample")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())

	if err := Path(f.Name()); err != nil {
		t.Errorf("ValidatePath() error = %v", err)
	}

	if err := Path("invalid"); err == nil {
		t.Error("ValidatePath() expected error but got none")
	}
}

func TestClusterName(t *testing.T) {
	tests := []struct {
		name    string
		val     interface{}
		wantErr bool
	}{
		{
			name:    "valid (single word)",
			val:     "Cluster0",
			wantErr: false,
		},
		{
			name:    "valid (dashed)",
			val:     "Cluster-0",
			wantErr: false,
		},
		{
			name:    "invalid (space)",
			val:     "Cluster 0",
			wantErr: true,
		},
		{
			name:    "invalid (underscore)",
			val:     "Cluster_0",
			wantErr: true,
		},
		{
			name:    "invalid (spacial char)",
			val:     "Cluster-ñ",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		val := tt.val
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			if err := ClusterName(val); (err != nil) != wantErr {
				t.Errorf("ClusterName() error = %v, wantErr %v", err, wantErr)
			}
		})
	}
}

func TestDBUsername(t *testing.T) {
	tests := []struct {
		name    string
		val     interface{}
		wantErr bool
	}{
		{
			name:    "valid (single word)",
			val:     "admin",
			wantErr: false,
		},
		{
			name:    "valid (dashed)",
			val:     "admin-test",
			wantErr: false,
		},
		{
			name:    "valid (underscore)",
			val:     "admin_test",
			wantErr: false,
		},
		{
			name:    "invalid (space)",
			val:     "admin test",
			wantErr: true,
		},
		{
			name:    "invalid (spacial char)",
			val:     "admin-ñ",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		val := tt.val
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			if err := DBUsername(val); (err != nil) != wantErr {
				t.Errorf("DBUsername() error = %v, wantErr %v", err, wantErr)
			}
		})
	}
}
