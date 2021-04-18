package repository

import (
	"reflect"
	"testing"
	"time"

	"github.com/niltonkummer/desafio-conductor/app/setup"
	"gorm.io/driver/sqlite"

	"github.com/niltonkummer/desafio-conductor/app/model"
	uuid "github.com/satori/go.uuid"

	"gorm.io/gorm"
)

var DB = func() *gorm.DB {
	db, _ := setup.SetupDB(sqlite.Open("../../db/test_data.sqlite"))
	return db
}

func Test_accountRepo_GetAccount(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Account
		wantErr bool
	}{
		{
			name: "found",
			fields: fields{
				DB: DB(),
			},
			args: args{
				id: uuid.FromStringOrNil("c13d1ec5-3215-472c-b856-ed4a83ee5c4d"),
			},
			want: &model.Account{
				Model: model.Model{
					ID:        "c13d1ec5-3215-472c-b856-ed4a83ee5c4d",
					CreatedAt: time.Date(2021, 4, 19, 1, 12, 35, 0, time.UTC),
					UpdatedAt: time.Date(2021, 4, 19, 1, 12, 39, 0, time.UTC),
					DeletedAt: gorm.DeletedAt{},
				},
				Status: "ATIVO",
			},
			wantErr: false,
		},
		{
			name: "not_found_or_db_error",
			fields: fields{
				DB: DB(),
			},
			args: args{
				id: uuid.FromStringOrNil("90087d5a-f3b9-49b4-bd1f-6ff11ad973e5"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty_uuid",
			fields: fields{
				DB: DB(),
			},
			args: args{
				id: uuid.FromStringOrNil(""),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &accountRepo{
				DB: tt.fields.DB,
			}
			got, err := a.GetAccount(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountRepo_ListAccounts(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    model.Accounts
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				DB: DB(),
			},
			want: model.Accounts{
				{
					Model: model.Model{
						ID:        "c13d1ec5-3215-472c-b856-ed4a83ee5c4d",
						CreatedAt: time.Date(2021, 4, 19, 1, 12, 35, 0, time.UTC),
						UpdatedAt: time.Date(2021, 4, 19, 1, 12, 39, 0, time.UTC),
						DeletedAt: gorm.DeletedAt{},
					},
					Status: "ATIVO",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &accountRepo{
				DB: tt.fields.DB,
			}
			got, err := a.ListAccounts()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAccounts() got = %v, want %v", got, tt.want)
			}
		})
	}
}
