package GoFor_MFT_Parser

import (
	"os"
	"reflect"
	"sync"
	"testing"
)

func TestDirectoryList_Create(t *testing.T) {
	type args struct {
		inboundBuffer        *chan []byte
		directoryListChannel *chan map[uint64]directory
		waitGroup            *sync.WaitGroup
	}
	tests := []struct {
		name          string
		directoryList unresolvedDirectoryList
		args          args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestMasterFileTableRecord_isThisADirectory(t *testing.T) {
	type args struct {
		mftRecord RawMasterFileTableRecord
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		got     bool
		wantErr bool
	}{
		{
			name:    "this is a directory",
			args:    args{mftRecord: RawMasterFileTableRecord([]byte{70, 73, 76, 69, 48, 0, 3, 0, 150, 29, 38, 147, 1, 0, 0, 0, 7, 0, 2, 0, 56, 0, 3, 0, 192, 3, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 63, 194, 1, 0, 30, 0, 97, 0, 0, 0, 0, 0, 16, 0, 0, 0, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 72, 0, 0, 0, 24, 0, 0, 0, 208, 193, 130, 132, 255, 4, 212, 1, 208, 193, 130, 132, 255, 4, 212, 1, 208, 193, 130, 132, 255, 4, 212, 1, 208, 193, 130, 132, 255, 4, 212, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 136, 122, 89, 107, 0, 0, 0, 0, 48, 0, 0, 0, 120, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 90, 0, 0, 0, 24, 0, 1, 0, 47, 114, 0, 0, 0, 0, 2, 0, 208, 193, 130, 132, 255, 4, 212, 1, 208, 193, 130, 132, 255, 4, 212, 1, 208, 193, 130, 132, 255, 4, 212, 1, 208, 193, 130, 132, 255, 4, 212, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 16, 0, 0, 0, 0, 12, 2, 65, 0, 77, 0, 50, 0, 49, 0, 65, 0, 52, 0, 126, 0, 49, 0, 46, 0, 49, 0, 49, 0, 50, 0, 0, 0, 0, 0, 0, 0, 48, 0, 0, 0, 16, 1, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 242, 0, 0, 0, 24, 0, 1, 0, 47, 114, 0, 0, 0, 0, 2, 0, 208, 193, 130, 132, 255, 4, 212, 1, 208, 193, 130, 132, 255, 4, 212, 1, 208, 193, 130, 132, 255, 4, 212, 1, 208, 193, 130, 132, 255, 4, 212, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 16, 0, 0, 0, 0, 88, 1, 97, 0, 109, 0, 100, 0, 54, 0, 52, 0, 95, 0, 109, 0, 105, 0, 99, 0, 114, 0, 111, 0, 115, 0, 111, 0, 102, 0, 116, 0, 45, 0, 119, 0, 105, 0, 110, 0, 100, 0, 111, 0, 119, 0, 115, 0, 45, 0, 98, 0, 111, 0, 111, 0, 116, 0, 109, 0, 101, 0, 110, 0, 117, 0, 117, 0, 120, 0, 95, 0, 51, 0, 49, 0, 98, 0, 102, 0, 51, 0, 56, 0, 53, 0, 54, 0, 97, 0, 100, 0, 51, 0, 54, 0, 52, 0, 101, 0, 51, 0, 53, 0, 95, 0, 49, 0, 48, 0, 46, 0, 48, 0, 46, 0, 49, 0, 55, 0, 49, 0, 51, 0, 52, 0, 46, 0, 49, 0, 49, 0, 50, 0, 95, 0, 110, 0, 111, 0, 110, 0, 101, 0, 95, 0, 100, 0, 54, 0, 30, 0, 57, 0, 53, 0, 50, 0, 54, 0, 49, 0, 54, 0, 98, 0, 53, 0, 98, 0, 48, 0, 54, 0, 50, 0, 102, 0, 0, 0, 0, 0, 0, 0, 144, 0, 0, 0, 48, 1, 0, 0, 0, 4, 24, 0, 0, 0, 1, 0, 16, 1, 0, 0, 32, 0, 0, 0, 36, 0, 73, 0, 51, 0, 48, 0, 48, 0, 0, 0, 1, 0, 0, 0, 0, 16, 0, 0, 1, 0, 0, 0, 16, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 151, 100, 1, 0, 0, 0, 4, 0, 112, 0, 94, 0, 0, 0, 0, 0, 63, 194, 1, 0, 0, 0, 7, 0, 119, 211, 176, 66, 255, 4, 212, 1, 24, 114, 72, 186, 88, 255, 211, 1, 13, 234, 157, 35, 130, 99, 213, 1, 119, 211, 176, 66, 255, 4, 212, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 138, 9, 0, 0, 0, 0, 0, 32, 6, 4, 0, 23, 0, 0, 128, 14, 1, 66, 0, 111, 0, 111, 0, 116, 0, 77, 0, 101, 0, 110, 0, 117, 0, 85, 0, 88, 0, 46, 0, 100, 0, 108, 0, 108, 0, 0, 0, 151, 100, 1, 0, 0, 0, 4, 0, 112, 0, 90, 0, 0, 0, 0, 0, 63, 194, 1, 0, 0, 0, 7, 0, 119, 211, 176, 66, 255, 4, 212, 1, 24, 114, 72, 186, 88, 255, 211, 1, 13, 234, 157, 35, 130, 99, 213, 1, 119, 211, 176, 66, 255, 4, 212, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 138, 9, 0, 0, 0, 0, 0, 32, 6, 4, 0, 23, 0, 0, 128, 12, 2, 66, 0, 79, 0, 79, 0, 84, 0, 77, 0, 69, 0, 126, 0, 49, 0, 46, 0, 68, 0, 76, 0, 76, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 16, 0, 0, 0, 2, 0, 0, 0, 0, 1, 0, 0, 104, 0, 0, 0, 0, 9, 24, 0, 0, 0, 4, 0, 56, 0, 0, 0, 48, 0, 0, 0, 36, 0, 84, 0, 88, 0, 70, 0, 95, 0, 68, 0, 65, 0, 84, 0, 65, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 5, 0, 1, 0, 0, 0, 1, 0, 0, 0, 37, 225, 0, 0, 0, 0, 0, 0, 0, 142, 0, 0, 21, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 144, 0, 0, 21, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 130, 121, 71, 17, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 30, 0})},
			want:    true,
			wantErr: false,
		},
		{
			name:    "this is not a directory",
			args:    args{mftRecord: RawMasterFileTableRecord([]byte{70, 73, 76, 69, 48, 0, 3, 0, 40, 230, 102, 93, 8, 0, 0, 0, 47, 0, 2, 0, 56, 0, 1, 0, 24, 2, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 23, 34, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 16, 0, 0, 0, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 72, 0, 0, 0, 24, 0, 0, 0, 166, 165, 230, 40, 101, 47, 213, 1, 208, 135, 76, 16, 79, 106, 213, 1, 208, 135, 76, 16, 79, 106, 213, 1, 25, 8, 22, 19, 79, 106, 213, 1, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 161, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 96, 47, 211, 158, 1, 0, 0, 0, 48, 0, 0, 0, 120, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 90, 0, 0, 0, 24, 0, 1, 0, 249, 31, 0, 0, 0, 0, 10, 0, 166, 165, 230, 40, 101, 47, 213, 1, 25, 8, 22, 19, 79, 106, 213, 1, 25, 8, 22, 19, 79, 106, 213, 1, 25, 8, 22, 19, 79, 106, 213, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 0, 0, 0, 0, 0, 0, 0, 12, 2, 79, 0, 78, 0, 69, 0, 68, 0, 82, 0, 73, 0, 126, 0, 51, 0, 46, 0, 80, 0, 78, 0, 71, 0, 0, 0, 0, 0, 0, 0, 48, 0, 0, 0, 184, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 154, 0, 0, 0, 24, 0, 1, 0, 249, 31, 0, 0, 0, 0, 10, 0, 166, 165, 230, 40, 101, 47, 213, 1, 25, 8, 22, 19, 79, 106, 213, 1, 25, 8, 22, 19, 79, 106, 213, 1, 25, 8, 22, 19, 79, 106, 213, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 0, 0, 0, 0, 0, 0, 0, 44, 1, 79, 0, 110, 0, 101, 0, 68, 0, 114, 0, 105, 0, 118, 0, 101, 0, 77, 0, 101, 0, 100, 0, 84, 0, 105, 0, 108, 0, 101, 0, 46, 0, 99, 0, 111, 0, 110, 0, 116, 0, 114, 0, 97, 0, 115, 0, 116, 0, 45, 0, 98, 0, 108, 0, 97, 0, 99, 0, 107, 0, 95, 0, 115, 0, 99, 0, 97, 0, 108, 0, 101, 0, 45, 0, 49, 0, 53, 0, 48, 0, 46, 0, 112, 0, 110, 0, 103, 0, 0, 0, 0, 0, 0, 0, 128, 0, 0, 0, 72, 0, 0, 0, 1, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, 0, 0, 0, 0, 0, 0, 0, 0, 16, 0, 0, 0, 0, 0, 0, 81, 8, 0, 0, 0, 0, 3, 0, 81, 8, 0, 0, 0, 0, 0, 0, 65, 1, 43, 20, 247, 2, 0, 0, 255, 255, 255, 255, 130, 121, 71, 17, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0})},
			want:    false,
			wantErr: false,
		},
		{
			name:    "nil bytes",
			args:    args{mftRecord: nil},
			wantErr: true,
		},
		{
			name:    "not enough bytes",
			args:    args{mftRecord: RawMasterFileTableRecord([]byte{70, 73})},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		var err error
		tt.got, err = tt.args.mftRecord.IsThisADirectory()
		if !reflect.DeepEqual(tt.got, tt.want) || (err != nil) != tt.wantErr {
			t.Errorf("Test %v failed \ngot = %v, \nwant = %v", tt.name, tt.got, tt.want)
		}
	}
}

func TestMftFile_BuildDirectoryTree(t *testing.T) {
	type fields struct {
		FileHandle        *os.File
		MappedDirectories map[uint64]string
		OutputChannel     chan MasterFileTableRecord
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestMftFile_CombineDirectoryInformation(t *testing.T) {
	type fields struct {
		FileHandle        *os.File
		MappedDirectories map[uint64]string
		OutputChannel     chan MasterFileTableRecord
	}
	type args struct {
		directoryListChannel        *chan map[uint64]directory
		waitForDirectoryCombination *sync.WaitGroup
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
