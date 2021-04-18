package repository

import (
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/niltonkummer/desafio-conductor/app/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func testSeedTransaction() model.Transactions {
	db := DB()
	_ = db.Exec("DELETE FROM transactions; VACUUM;")
	conta := "c13d1ec5-3215-472c-b856-ed4a83ee5c4d"
	list := model.Transactions{
		{ContaID: &conta, Model: model.Model{ID: "c13d1ec5-3215-472c-b856-ed4a83ee5c4d", CreatedAt: time.Date(2021, 04, 19, 2, 14, 32, 0, time.UTC), UpdatedAt: time.Date(2021, 04, 19, 2, 14, 34, 0, time.UTC)}, Descricao: "Netflix", Valor: 32.9},
		{ContaID: &conta, Model: model.Model{ID: "599c1985-486e-4f26-b778-01f9945c7902", CreatedAt: time.Date(2021, 04, 19, 2, 14, 32, 0, time.UTC), UpdatedAt: time.Date(2021, 04, 19, 2, 14, 34, 0, time.UTC)}, Descricao: "iFood", Valor: 58.1},
		{ContaID: &conta, Model: model.Model{ID: "3767072e-5ac2-4d61-ab16-8d133ffe310d", CreatedAt: time.Date(2021, 04, 19, 2, 14, 32, 0, time.UTC), UpdatedAt: time.Date(2021, 04, 19, 2, 14, 34, 0, time.UTC)}, Descricao: "Uber", Valor: 20},
		{ContaID: &conta, Model: model.Model{ID: "17558a06-427f-4ce1-90fd-75c90b077f81", CreatedAt: time.Date(2021, 04, 19, 2, 14, 32, 0, time.UTC), UpdatedAt: time.Date(2021, 04, 19, 2, 14, 34, 0, time.UTC)}, Descricao: "iFood", Valor: 32.9},
		{ContaID: &conta, Model: model.Model{ID: "37c2fabc-49e0-47ed-b91d-2d7cd1d0a55f", CreatedAt: time.Date(2021, 04, 19, 2, 14, 32, 0, time.UTC), UpdatedAt: time.Date(2021, 04, 19, 2, 14, 34, 0, time.UTC)}, Descricao: "iFood", Valor: 25.9},
		{ContaID: &conta, Model: model.Model{ID: "004eaddc-b84b-4683-95f1-a71b44e3e11f", CreatedAt: time.Date(2021, 04, 19, 2, 14, 32, 0, time.UTC), UpdatedAt: time.Date(2021, 04, 19, 2, 14, 34, 0, time.UTC)}, Descricao: "iFood", Valor: 12},
		{ContaID: &conta, Model: model.Model{ID: "a2da1b39-01a2-4225-8f71-e77eb61b8363", CreatedAt: time.Date(2021, 04, 19, 2, 14, 32, 0, time.UTC), UpdatedAt: time.Date(2021, 04, 19, 2, 14, 34, 0, time.UTC)}, Descricao: "iFood", Valor: 34.9},
		{ContaID: &conta, Model: model.Model{ID: "652f98f4-079d-4bce-a87e-9d818d302756", CreatedAt: time.Date(2021, 04, 19, 2, 14, 32, 0, time.UTC), UpdatedAt: time.Date(2021, 04, 19, 2, 14, 34, 0, time.UTC)}, Descricao: "iFood", Valor: 43.9},
		{ContaID: &conta, Model: model.Model{ID: "120643b5-3e23-480b-803b-1b94aa400616", CreatedAt: time.Date(2021, 04, 19, 2, 14, 32, 0, time.UTC), UpdatedAt: time.Date(2021, 04, 19, 2, 14, 34, 0, time.UTC)}, Descricao: "iFood", Valor: 54.9},
		{ContaID: &conta, Model: model.Model{ID: "ef465773-b7bf-429e-a65c-714a2bd9f2bb", CreatedAt: time.Date(2021, 04, 19, 2, 14, 32, 0, time.UTC), UpdatedAt: time.Date(2021, 04, 19, 2, 14, 34, 0, time.UTC)}, Descricao: "iFood", Valor: 13.9},
		{ContaID: &conta, Model: model.Model{ID: "8abf38a0-de75-4189-be17-287821ab5f92", CreatedAt: time.Date(2021, 04, 19, 2, 14, 32, 0, time.UTC), UpdatedAt: time.Date(2021, 04, 19, 2, 14, 34, 0, time.UTC)}, Descricao: "iFood", Valor: 43},
		{ContaID: &conta, Model: model.Model{ID: "01e283bc-d12a-4215-a7bc-4a7c135e6837", CreatedAt: time.Date(2021, 04, 19, 2, 14, 32, 0, time.UTC), UpdatedAt: time.Date(2021, 04, 19, 2, 14, 34, 0, time.UTC)}, Descricao: "iFood", Valor: 60},
		{ContaID: &conta, Model: model.Model{ID: "14bbc2af-f2a1-4693-9638-07d99c11b613", CreatedAt: time.Date(2021, 04, 19, 2, 14, 32, 0, time.UTC), UpdatedAt: time.Date(2021, 04, 19, 2, 14, 34, 0, time.UTC)}, Descricao: "iFood", Valor: 70},
		{ContaID: &conta, Model: model.Model{ID: "c3f6d45f-f605-4513-911e-a276ca8c9c39", CreatedAt: time.Date(2021, 04, 19, 2, 14, 32, 0, time.UTC), UpdatedAt: time.Date(2021, 04, 19, 2, 14, 34, 0, time.UTC)}, Descricao: "iFood", Valor: 120},
		{ContaID: &conta, Model: model.Model{ID: "ffdf3edf-7c2b-4af6-87e5-f856e3cd65ae", CreatedAt: time.Date(2021, 04, 19, 2, 14, 32, 0, time.UTC), UpdatedAt: time.Date(2021, 04, 19, 2, 14, 34, 0, time.UTC)}, Descricao: "iFood", Valor: 32.9},
	}

	if err := db.Model(&model.Transaction{}).Create(list).Error; err != nil {
		log.Fatalf("cannot seed table transactions: %s", err)
	}
	return list
}

func Test_transactionRepo_GetTransactions(t1 *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		accountID uuid.UUID
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantTransactions model.Transactions
		wantErr          bool
	}{
		{
			name: "empty_list",
			fields: fields{
				DB: DB(),
			},
			args: args{
				accountID: uuid.UUID{},
			},
			wantTransactions: model.Transactions{},
			wantErr:          false,
		},
		{
			name: "ok",
			fields: fields{
				DB: DB(),
			},
			args: args{
				accountID: uuid.FromStringOrNil("c13d1ec5-3215-472c-b856-ed4a83ee5c4d"),
			},
			wantTransactions: testSeedTransaction(),
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &transactionRepo{
				DB: tt.fields.DB,
			}
			gotTransactions, err := t.GetTransactions(tt.args.accountID)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTransactions, tt.wantTransactions) {
				t1.Errorf("GetTransactions() gotTransactions = %v, want %v", gotTransactions, tt.wantTransactions)
			}
		})
	}
}
