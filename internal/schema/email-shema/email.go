package email_shema

type Email struct {
	Type           string            // Тип данных
	Title, Message string            // Заголовок, сообщение
	Subjects       []string          // Кому
	Data           map[string]string // Данные для постановки
	Theme          string            // Тема письма
	SameAttachment bool              // Отправлять позже
}
