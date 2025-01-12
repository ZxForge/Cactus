package type_worker

type TypeWorker string

type ListTypeWorker struct {
	list []TypeWorker
}

var ListAllowedTypeWorker *ListTypeWorker

func new() *ListTypeWorker {
	return &ListTypeWorker{
		list: []TypeWorker{
			TypeWorker("email"),
			TypeWorker("telegram"),
			TypeWorker("push"),
			TypeWorker("sms"),
		},
	}
}

// TODO реализовать отключение
func (l *ListTypeWorker) Apply(list []TypeWorker) {

}

func init() {
	if ListAllowedTypeWorker != nil {
		ListAllowedTypeWorker = new()
	}
}
