package log

import "github.com/bersennaidoo/Loggeez/internal/domain/model"

type Loger interface {
	Append(record model.Record) (uint64, error)
	Read(offset uint64) (model.Record, error)
}
