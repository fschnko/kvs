package storage

import (
	"testing"
)

func TestStorageGet(t *testing.T) {
	s := New()

	if err := populate(s); err != nil {
		t.Fatalf("populate storage: %v", err)
	}

	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "normal case",
			args: args{
				key: "key_1",
			},
			want: "value_1",
		}, {
			name: "empty key",
			args: args{
				key: "",
			},
			wantErr: ErrKeyEmpty,
		}, {
			name: "key too long (17 bytes)",
			args: args{
				key: "loooooooooooooong",
			},
			wantErr: ErrKeyTooLong,
		}, {
			name: "key not exists",
			args: args{
				key: "dummy key",
			},
			wantErr: ErrKeyNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Get(tt.args.key)
			if err != tt.wantErr {
				t.Errorf("Storage.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorageSet(t *testing.T) {
	s := New()

	if err := populate(s); err != nil {
		t.Fatalf("populate storage: %v", err)
	}

	type args struct {
		key string
		val string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "new pair",
			args: args{
				key: "new_key",
				val: "new_val",
			},
		}, {
			name: "normal case",
			args: args{
				key: "key_1",
				val: "new_val_1",
			},
		}, {
			name: "empty key",
			args: args{
				key: "",
			},
			wantErr: ErrKeyEmpty,
		}, {
			name: "key too long (17 bytes)",
			args: args{
				key: "loooooooooooooong",
			},
			wantErr: ErrKeyTooLong,
		}, {
			name: "val too long (513 bytes)",
			args: args{
				key: "long val",
				val: "loooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooong",
			},
			wantErr: ErrValueTooLong,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Set(tt.args.key, tt.args.val); err != tt.wantErr {
				t.Errorf("Storage.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorageExists(t *testing.T) {
	s := New()

	if err := populate(s); err != nil {
		t.Fatalf("populate storage: %v", err)
	}

	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr error
	}{
		{
			name: "exists case",
			args: args{
				key: "key_1",
			},
			want: true,
		}, {
			name: "not exists case",
			args: args{
				key: "dummy key",
			},
			want: false,
		}, {
			name: "empty key",
			args: args{
				key: "",
			},
			wantErr: ErrKeyEmpty,
		}, {
			name: "key too long (17 bytes)",
			args: args{
				key: "loooooooooooooong",
			},
			wantErr: ErrKeyTooLong,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.Exists(tt.args.key)
			if err != tt.wantErr {
				t.Errorf("Storage.Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorageDelete(t *testing.T) {
	s := New()

	if err := populate(s); err != nil {
		t.Fatalf("populate storage: %v", err)
	}

	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "normal case",
			args: args{
				key: "key_1",
			},
			wantErr: nil,
		}, {
			name: "key not exists",
			args: args{
				key: "dummy key",
			},
			wantErr: ErrKeyNotFound,
		}, {
			name: "empty key",
			args: args{
				key: "",
			},
			wantErr: ErrKeyEmpty,
		}, {
			name: "key too long (17 bytes)",
			args: args{
				key: "loooooooooooooong",
			},
			wantErr: ErrKeyTooLong,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Delete(tt.args.key); err != tt.wantErr {
				t.Errorf("Storage.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func populate(s Storage) error {
	data := map[string]string{
		"key_1": "value_1",
		"key_2": "value_2",
		"key_3": "value_3",
		"key_4": "value_4",
		"key_5": "value_5",
		"key_6": "value_6",
		"key_7": "value_7",
	}

	for key, val := range data {
		if err := s.Set(key, val); err != nil {
			return err
		}
	}
	return nil
}
