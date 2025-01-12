package plugin

type Plugin interface {
	New() Plugin // TODO Переписать на состояние REQUEST, создаю каждый раз только из за того что отличить REQEST'ы не могу друг от друга, нужен REQUEST ID.
	GetSchema() any
}
