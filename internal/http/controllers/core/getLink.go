package core

import "net/http"

type getLinkService interface {
}

// Получение ссылки на страницу просмотра рассылки в реальном времени
func GetLink(s getLinkService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Реализовать.
		w.Write([]byte("Тест GetLink прошел успешно"))
	}
}
