package core

import (
	"net/http"
	"os"
)

type getFilesService interface {
	GetFile(hash string) (*os.File, error)
}

// Получение ссылки на страницу просмотра рассылки в реальном времени
func GetFiles(s getFilesService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// var req request.GetFilesRequest
		// _ = json.NewDecoder(r.Body).Decode(&req)

		// // Переписать валидатор, так как сейчас мы пишем в структуре валидацию и приходится туда сюда прыгать.
		// errors := validation.ValidateStructure(&req)
		// if errors != nil {
		// 	w.Header().Set("Content-Type", "application/json")
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	jsonErrors, _ := json.Marshal(response.CreateResponseError("Ошибка валидации, проверьте отправляемые поля", errors))
		// 	w.Write(jsonErrors)
		// 	return
		// }

		// s.GetFile()

		// w.Header().Set("Content-Type", "application/octet-stream")
		// w.Header().Set("Content-Disposition", "attachment; filename=file.txt")

		// http.ServeFile()

		w.Write([]byte("Тест GetFiles прошел успешно"))
	}
}
