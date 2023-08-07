package repo

import (
	"sync"

	"github.com/bersennaidoo/Loggeez/internal/domain/model"
)

type Log struct {
	mu      sync.Mutex
	records []model.Record
}

func NewLog() *Log {
	return &Log{}
}

func (c *Log) Append(record model.Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)

	return record.Offset, nil
}

func (c *Log) Read(offset uint64) (model.Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if offset >= uint64(len(c.records)) {
		return model.Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}
