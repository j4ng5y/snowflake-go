package snowflake

import (
	"fmt"
	"math/bits"
	"reflect"
	"strconv"
	"time"
)

type Snowflake struct {
	epoch             uint64
	Timestamp         time.Time
	InternalWorkerID  uint64
	InternalProcessID uint64
	IncrimentID       uint64
}

func (S *Snowflake) parseTimestamp(u uint64) time.Time {
	return time.Unix(int64(((u>>22)+S.epoch)/1000), 0).UTC()
}

func (S *Snowflake) parseInternalWorkerID(u uint64) uint64 {
	u = u >> 17
	u = bits.Reverse64(u)
	u = u >> 60
	u = bits.Reverse64(u)
	u = u >> 60
	return u
}

func (S *Snowflake) parseInternalProcessID(u uint64) uint64 {
	u = u >> 12
	u = bits.Reverse64(u)
	u = u >> 60
	u = bits.Reverse64(u)
	u = u >> 60
	return u
}

func (S *Snowflake) parseIncrimentID(u uint64) uint64 {
	u = bits.Reverse64(u)
	u = u >> 53
	u = bits.Reverse64(u)
	u = u >> 53
	return u
}

func (S *Snowflake) unmarshalString(v string) error {
	u, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return err
	}

	S.Timestamp = S.parseTimestamp(u)
	S.InternalWorkerID = S.parseInternalWorkerID(u)
	S.InternalProcessID = S.parseInternalProcessID(u)
	S.IncrimentID = S.parseIncrimentID(u)
	return nil
}

func (S *Snowflake) unmarshalInt(v int) error {
	S.Timestamp = S.parseTimestamp(uint64(v))
	S.InternalWorkerID = S.parseInternalWorkerID(uint64(v))
	S.InternalProcessID = S.parseInternalWorkerID(uint64(v))
	S.IncrimentID = S.parseIncrimentID(uint64(v))
	return nil
}

func (S *Snowflake) unmarshalInt8(v int8) error {
	S.Timestamp = S.parseTimestamp(uint64(v))
	S.InternalWorkerID = S.parseInternalWorkerID(uint64(v))
	S.InternalProcessID = S.parseInternalWorkerID(uint64(v))
	S.IncrimentID = S.parseIncrimentID(uint64(v))
	return nil
}

func (S *Snowflake) unmarshalInt16(v int16) error {
	S.Timestamp = S.parseTimestamp(uint64(v))
	S.InternalWorkerID = S.parseInternalWorkerID(uint64(v))
	S.InternalProcessID = S.parseInternalWorkerID(uint64(v))
	S.IncrimentID = S.parseIncrimentID(uint64(v))
	return nil
}

func (S *Snowflake) unmarshalInt32(v int32) error {
	S.Timestamp = S.parseTimestamp(uint64(v))
	S.InternalWorkerID = S.parseInternalWorkerID(uint64(v))
	S.InternalProcessID = S.parseInternalWorkerID(uint64(v))
	S.IncrimentID = S.parseIncrimentID(uint64(v))
	return nil
}

func (S *Snowflake) unmarshalInt64(v int64) error {
	S.Timestamp = S.parseTimestamp(uint64(v))
	S.InternalWorkerID = S.parseInternalWorkerID(uint64(v))
	S.InternalProcessID = S.parseInternalWorkerID(uint64(v))
	S.IncrimentID = S.parseIncrimentID(uint64(v))
	return nil
}

func (S *Snowflake) unmarshalUint(v uint) error {
	S.Timestamp = S.parseTimestamp(uint64(v))
	S.InternalWorkerID = S.parseInternalWorkerID(uint64(v))
	S.InternalProcessID = S.parseInternalWorkerID(uint64(v))
	S.IncrimentID = S.parseIncrimentID(uint64(v))
	return nil
}

func (S *Snowflake) unmarshalUint8(v uint8) error {
	S.Timestamp = S.parseTimestamp(uint64(v))
	S.InternalWorkerID = S.parseInternalWorkerID(uint64(v))
	S.InternalProcessID = S.parseInternalWorkerID(uint64(v))
	S.IncrimentID = S.parseIncrimentID(uint64(v))
	return nil
}

func (S *Snowflake) unmarshalUint16(v uint16) error {
	S.Timestamp = S.parseTimestamp(uint64(v))
	S.InternalWorkerID = S.parseInternalWorkerID(uint64(v))
	S.InternalProcessID = S.parseInternalWorkerID(uint64(v))
	S.IncrimentID = S.parseIncrimentID(uint64(v))
	return nil
}

func (S *Snowflake) unmarshalUint32(v uint32) error {
	S.Timestamp = S.parseTimestamp(uint64(v))
	S.InternalWorkerID = S.parseInternalWorkerID(uint64(v))
	S.InternalProcessID = S.parseInternalWorkerID(uint64(v))
	S.IncrimentID = S.parseIncrimentID(uint64(v))
	return nil
}

func (S *Snowflake) unmarshalUint64(v uint64) error {
	S.Timestamp = S.parseTimestamp(uint64(v))
	S.InternalWorkerID = S.parseInternalWorkerID(uint64(v))
	S.InternalProcessID = S.parseInternalWorkerID(uint64(v))
	S.IncrimentID = S.parseIncrimentID(uint64(v))
	return nil
}

// Unmarshal takes a value and inserts the relevant values into the Snowflake struct
func (S *Snowflake) Unmarshal(value interface{}) error {
	switch reflect.TypeOf(value).String() {
	case "string":
		return S.unmarshalString(value.(string))
	case "int":
		return S.unmarshalInt(value.(int))
	case "int8":
		return S.unmarshalInt8(value.(int8))
	case "int16":
		return S.unmarshalInt16(value.(int16))
	case "int32":
		return S.unmarshalInt32(value.(int32))
	case "int64":
		return S.unmarshalInt64(value.(int64))
	case "uint":
		return S.unmarshalUint(value.(uint))
	case "uint8":
		return S.unmarshalUint8(value.(uint8))
	case "uint16":
		return S.unmarshalUint16(value.(uint16))
	case "uint32":
		return S.unmarshalUint32(value.(uint32))
	case "uint64":
		return S.unmarshalUint64(value.(uint64))
	default:
		return fmt.Errorf("the type '%s', of the value provided, '%v', is not supported by this operation", reflect.TypeOf(value).String(), value)
	}
}

func (S Snowflake) String() string {
	s := (uint64(S.Timestamp.Unix()) - S.epoch) << 22
	s = S.InternalWorkerID << 17
	s = S.InternalProcessID << 12
	s = S.IncrimentID << 0
	return fmt.Sprintf("%d", s)
}

func (S Snowflake) Int() int {
	s := (uint64(S.Timestamp.Unix()) - S.epoch) << 22
	s = S.InternalWorkerID << 17
	s = S.InternalProcessID << 12
	s = S.IncrimentID << 0
	return int(s)
}

func (S Snowflake) Int64() int64 {
	s := (uint64(S.Timestamp.Unix()) - S.epoch) << 22
	s = S.InternalWorkerID << 17
	s = S.InternalProcessID << 12
	s = S.IncrimentID << 0
	return int64(s)
}

func (S Snowflake) Uint() uint {
	s := (uint64(S.Timestamp.Unix()) - S.epoch) << 22
	s = S.InternalWorkerID << 17
	s = S.InternalProcessID << 12
	s = S.IncrimentID << 0
	return uint(s)
}

func (S Snowflake) Uint64() uint64 {
	s := (uint64(S.Timestamp.Unix()) - S.epoch) << 22
	s = S.InternalWorkerID << 17
	s = S.InternalProcessID << 12
	s = S.IncrimentID << 0
	return s
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
