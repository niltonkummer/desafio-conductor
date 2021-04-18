package app

import (
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func SetupRoutes(a *app) http.Handler {

	r := chi.NewRouter()

	tokenAuth := jwtauth.New("HS256", []byte("MY_SECRET"), nil)

	r.Use(jwtauth.Verifier(tokenAuth))
	r.Use(jwtauth.Authenticator)

	r.Use(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _, err := jwtauth.FromContext(r.Context())
			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			handler.ServeHTTP(w, r)
		})
	})

	r.Get("/", ListAccounts(a))

	r.Get("/{id}", GetAccount(a))

	r.Get("/{id}/transacoes", GetTransactions(a))

	r.Get("/{id}/transacoes.pdf", GeneratePDF(a))

	return r
}

// ListAccounts Lista Contas godoc
// @Summary Lista Contas
// @Description Retorna uma lista de contas
// @Produce  json
// @Success 200 {object} []model.Account
// @Header 200 {string} Bearer "token"
// @Router /contas [get]
func ListAccounts(a *app) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		list, err := a.account.ListAccounts()
		if err != nil {
			render.Status(request, http.StatusInternalServerError)
			render.Respond(writer, request, map[string]string{
				"error": err.Error(),
			})
			return
		}
		render.Status(request, http.StatusOK)
		render.Respond(writer, request, list)
	}
}

// GetAccount Busca Conta godoc
// @Summary Busca Conta por ID
// @Description Retorna uma conta por id
// @Param id path string true "ID Conta"
// @Produce json
// @Success 200 {object} model.Account
// @Header 200 {string} Bearer "token"
// @Router /contas/{id} [get]
func GetAccount(a *app) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		id := chi.URLParam(request, "id")
		if !govalidator.IsUUIDv4(id) {
			render.Status(request, http.StatusBadRequest)
			render.Respond(writer, request, map[string]string{
				"error": "invalid parameter",
			})
			return
		}

		account, err := a.account.GetAccount(uuid.FromStringOrNil(id))
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				render.Status(request, http.StatusInternalServerError)
			} else {
				render.Status(request, http.StatusBadRequest)
			}
			render.Respond(writer, request, map[string]string{
				"error": err.Error(),
			})
			return
		}
		render.Status(request, http.StatusOK)
		render.Respond(writer, request, account)
	}
}

// GetTransactions Busca transações godoc
// @Summary Busca as transações de uma conta
// @Description Retorna uma lista de transações de uma determinada conta
// @Param id path string true "ID Conta"
// @Produce json
// @Success 200 {object} model.Transactions
// @Header 200 {string} Bearer "token"
// @Router /contas/{id}/transacoes [get]
func GetTransactions(a *app) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		id := chi.URLParam(request, "id")
		if !govalidator.IsUUIDv4(id) {
			render.Status(request, http.StatusBadRequest)
			render.Respond(writer, request, map[string]string{
				"error": "invalid parameter",
			})
			return
		}

		transactions, err := a.transaction.GetTransactions(uuid.FromStringOrNil(id))
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				render.Status(request, http.StatusInternalServerError)
			} else {
				render.Status(request, http.StatusBadRequest)
			}
			render.Respond(writer, request, map[string]string{
				"error": err.Error(),
			})
			return
		}
		render.Status(request, http.StatusOK)
		render.Respond(writer, request, transactions)
	}
}

// GeneratePDF Gerar Fatura godoc
// @Summary Gera a fatura em PDF
// @Description Retorna as transações de uma conta em formato PDF
// @Param id path string true "ID Conta"
// @Produce application/json
// @Success 200
// @Header 200 {string} Bearer "token"
// @Router /contas/{id}/transacoes.pdf [get]
func GeneratePDF(a *app) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		id := chi.URLParam(request, "id")
		if !govalidator.IsUUIDv4(id) {
			render.Status(request, http.StatusBadRequest)
			render.Respond(writer, request, map[string]string{
				"error": "invalid parameter",
			})
			return
		}

		data, err := a.transaction.GeneratePDF(uuid.FromStringOrNil(id))
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				render.Status(request, http.StatusInternalServerError)
			} else {
				render.Status(request, http.StatusBadRequest)
			}
			render.Respond(writer, request, map[string]string{
				"error": err.Error(),
			})
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/pdf")
		writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%v", "fatura.pdf"))
		_, _ = writer.Write(data)

		render.Status(request, http.StatusOK)
		render.Respond(writer, request, data)
	}
}
