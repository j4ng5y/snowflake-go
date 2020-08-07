package snowflake

import "time"

type Snowflake struct {
	epoch             uint64
	Timestamp         time.Time
	InternalWorkerID  uint64
	InternalProcessID uint64
	IncrimentID       uint64
}

func (S *Snowflake) ParseString(s string) error {
	return nil
}

func (S *Snowflake) ParseInt(n int) error {
	return nil
}

func (S *Snowflake) ParseInt8(n int8) error {
	return nil
}

func (S *Snowflake) ParseInt16(n int16) error {
	return nil
}

func (S *Snowflake) ParseInt32(n int32) error {
	return nil
}

func (S *Snowflake) ParseInt64(n int64) error {
	return nil
}

func (S *Snowflake) ParseUint(n uint) error {
	return nil
}

func (S *Snowflake) ParseUint8(n uint8) error {
	return nil
}

func (S *Snowflake) ParseUint16(n uint16) error {
	return nil
}

func (S *Snowflake) ParseUint32(n uint32) error {
	return nil
}

func (S *Snowflake) ParseUint64(n uint64) error {
	return nil
}

func (S Snowflake) Int() int {
	return int(0)
}

func (S Snowflake) Int8() int8 {
	return int8(0)
}

func (S Snowflake) Int16() int16 {
	return int16(0)
}

func (S Snowflake) Int32() int32 {
	return int32(0)
}

func (S Snowflake) Int64() int64 {
	return int64(0)
}

func (S Snowflake) Uint() uint {
	return uint(0)
}

func (S Snowflake) Uint8() uint8 {
	return uint8(0)
}

func (S Snowflake) Uint16() uint16 {
	return uint16(0)
}

func (S Snowflake) Uint32() uint32 {
	return uint32(0)
}

func (S Snowflake) Uint64() uint64 {
	return uint64(0)
}

type SnowflakeOption func(*Snowflake)

func New() *Snowflake {
	return &Snowflake{
		Timestamp: time.Now(),
	}
}

func WithEpoch(e uint64) SnowflakeOption {
	return func(S *Snowflake) {
		S.epoch = e
	}
}

func WithSpecificTime(t time.Time) SnowflakeOption {
	return func(S *Snowflake) {
		S.Timestamp = t
	}
}

func WithSpecificWorkerID(id uint64) SnowflakeOption {
	return func(S *Snowflake) {
		S.InternalWorkerID = id
	}
}

func WithSpecificProcessID(id uint64) SnowflakeOption {
	return func(S *Snowflake) {
		S.InternalProcessID = id
	}
}

func WithSpecificIncrimentID(id uint64) SnowflakeOption {
	return func(S *Snowflake) {
		S.IncrimentID = id
	}
}

func NewWithOptions(opts ...SnowflakeOption) *Snowflake {
	S := &Snowflake{}

	for _, opt := range opts {
		opt(S)
	}

	return S
}
