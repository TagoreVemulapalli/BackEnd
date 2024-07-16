package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler Suite")
}

var _ = Describe("Handlers", func() {
	var (
		e   *echo.Echo
		rec *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		e = echo.New()
		rec = httptest.NewRecorder()
	})

	Describe("GET /users", func() {
		It("should return 200 OK", func() {
			req := httptest.NewRequest(http.MethodGet, "/users", nil)
			c := e.NewContext(req, rec)

			err := GetUsers(c)
			Expect(err).To(BeNil())
			Expect(rec.Code).To(Equal(http.StatusOK))
		})

		It("should fail with 404 Not Found", func() {
			req := httptest.NewRequest(http.MethodGet, "/nonexistent", nil)
			c := e.NewContext(req, rec)

			err := GetUsers(c) // Intentional failure: getUsers should not handle /nonexistent
			Expect(err).ToNot(BeNil())
			Expect(rec.Code).To(Equal(http.StatusNotFound))
		})
	})

	Describe("POST /users", func() {
		It("should create a user and return 201", func() {
			req := httptest.NewRequest(http.MethodPost, "/users", nil)
			c := e.NewContext(req, rec)

			err := CreateUser(c)
			Expect(err).To(BeNil())
			Expect(rec.Code).To(Equal(http.StatusCreated))
		})

		It("should fail to create a user with invalid data", func() {
			req := httptest.NewRequest(http.MethodPost, "/users", nil)
			c := e.NewContext(req, rec)

			err := CreateUser(c)
			Expect(err).ToNot(BeNil())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))
		})
	})

	Describe("PUT /users/:id", func() {
		It("should update a user and return 200", func() {
			req := httptest.NewRequest(http.MethodPut, "/users/1", nil)
			c := e.NewContext(req, rec)

			err := UpdateUser(c)
			Expect(err).To(BeNil())
			Expect(rec.Code).To(Equal(http.StatusOK))
		})

		It("should fail to update a user with invalid ID", func() {
			req := httptest.NewRequest(http.MethodPut, "/users/invalid-id", nil)
			c := e.NewContext(req, rec)

			err := UpdateUser(c) // Intentional failure
			Expect(err).ToNot(BeNil())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))
		})
	})

	Describe("DELETE /users/:id", func() {
		It("should delete a user and return 204", func() {
			req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
			c := e.NewContext(req, rec)

			err := DeleteUser(c)
			Expect(err).To(BeNil())
			Expect(rec.Code).To(Equal(http.StatusNoContent))
		})

		It("should fail to delete a user with invalid ID", func() {
			req := httptest.NewRequest(http.MethodDelete, "/users/invalid-id", nil)
			c := e.NewContext(req, rec)

			err := DeleteUser(c) // Intentional failure
			Expect(err).ToNot(BeNil())
			Expect(rec.Code).To(Equal(http.StatusBadRequest))
		})
	})
})