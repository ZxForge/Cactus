package dbstorage

import (
	"cactus/internal/model"
	"context"
	//"github.com/Masterminds/squirrel"
)

func (s *Storage) GetProcesses(ctx context.Context) ([]model.Process, error){
	processes := []model.Process{ 
		{Id: 1, Name: "Рассылка писем", Slug: "email"}, 
		{Id: 2, Name: "Телеграмм уведомления", Slug: "telegram"}, 
		{Id: 3, Name: "Push уведомления", Slug: "push"}, 
		{Id: 4, Name: "SMS-сообщения", Slug: "sms"}, 
		{Id: 5, Name: "1С", Slug: "oneC"}, 
	}
	return processes, nil 
}