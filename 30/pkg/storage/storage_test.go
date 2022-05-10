package storage

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		constr string
	}
	tests := []struct {
		name    string
		args    args
		want    *Storage
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.constr)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_NewTask(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		t Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.NewTask(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_Tasks(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		taskID   int
		authorID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.Tasks(tt.args.taskID, tt.args.authorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tasks() got = %v, want %v", got, tt.want)
			}
		})
	}
}
