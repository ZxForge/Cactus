package model

type PipelineStatus struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}
