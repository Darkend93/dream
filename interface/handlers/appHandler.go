package handlers

type AppHandler struct {
	Dream interface{ DreamHandler }
}
