// Copyright 2020 Jordan Gregory. All rights reserved.
// Use of this source code is governed by the license
// that can be found in the LICENSE.txt file.

// Package snowflake implements methods to generate
// and unmarshal unique identifiers in the style of
// Twitters snowflake network ID tool.
package snowflake

import (
	"fmt"
	"math/bits"
	"reflect"
	"strconv"
	"time"
)

// Snowflake is the core structure to held data used to create
// and unmarshal new snowflake identifiers.
type Snowflake struct {
	// epoch is used only to make the timestamp match custom
	// timestamps that use a custom epoch, such as Discord.
	epoch uint64

	// Timestamp is used to generate the last 42 bits of the
	// snowflake.
	Timestamp time.Time

	// InternalWorkerID is used to identify the worker that
	// is generating this snowflake. This field makes up
	// the 17th - 21st bits of the snowflake.
	InternalWorkerID uint64

	// InternalProcessID is used to identify the process of
	// the worker that is generating this snowflake. This
	// field makes up the 12th - 16th bits of the snowflake.
	InternalProcessID uint64

	// IncrimentID is incremented each time a worker/process
	// combo creates another identifier. This field makes up
	// the first 11 bits of the snowflake.
	IncrimentID uint64
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

// Unmarshal takes a value and inserts the relevant values into
// the parent instance of Snowflake.
//
// Arguments:
//     value (interface{}): The value to unmarshal.
//
// Returns:
//     (error): An error if one exists, nil otherwise.
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

// String returns the string value of this snowflake.
//
// Argumenets:
//     None
//
// Returns:
//     (string): The string value of the snowflake.
func (S Snowflake) String() string {
	s := (uint64(S.Timestamp.Unix()) - S.epoch) << 22
	s = S.InternalWorkerID << 17
	s = S.InternalProcessID << 12
	s = S.IncrimentID << 0
	return fmt.Sprintf("%d", s)
}

// Int returns the value of this snowflake as in int type.
//
// Argumenets:
//     None
//
// Returns:
//     (int): The int value of the snowflake.
func (S Snowflake) Int() int {
	s := (uint64(S.Timestamp.Unix()) - S.epoch) << 22
	s = S.InternalWorkerID << 17
	s = S.InternalProcessID << 12
	s = S.IncrimentID << 0
	return int(s)
}

// Int64 returns the value of this snowflake as an int64 type.
//
// Argumenets:
//     None
//
// Returns:
//     (int64): The int64 value of the snowflake.
func (S Snowflake) Int64() int64 {
	s := (uint64(S.Timestamp.Unix()) - S.epoch) << 22
	s = S.InternalWorkerID << 17
	s = S.InternalProcessID << 12
	s = S.IncrimentID << 0
	return int64(s)
}

// Uint returns the value of this snowflake as a uint type.
//
// Argumenets:
//     None
//
// Returns:
//     (uint): The uint value of the snowflake.
func (S Snowflake) Uint() uint {
	s := (uint64(S.Timestamp.Unix()) - S.epoch) << 22
	s = S.InternalWorkerID << 17
	s = S.InternalProcessID << 12
	s = S.IncrimentID << 0
	return uint(s)
}

// Uint64 returns the value of this snowflake as a uint64 type.
//
// Argumenets:
//     None
//
// Returns:
//     (uint64): The uint64 value of the snowflake.
func (S Snowflake) Uint64() uint64 {
	s := (uint64(S.Timestamp.Unix()) - S.epoch) << 22
	s = S.InternalWorkerID << 17
	s = S.InternalProcessID << 12
	s = S.IncrimentID << 0
	return s
}

// Option is the interface for all Snowflake Options
// to follow in order to modify a new Snowflake via the
// NewWithOptions function.
type Option func(*Snowflake)

// New generates a new Snowflake with default values:
//     * The standard UNIX epoch
//     * The current time
//     * A WorkerID of 0
//     * A ProcessID of 0
//     * An IncrimentID of 0
//
// Arguments:
//     None
//
// Returns:
//     (*Snowflake): A pointer to the new instance of Snowflake.
func New() *Snowflake {
	return &Snowflake{
		Timestamp: time.Now(),
	}
}

// WithEpoch is a function to set a custom epoch for generating
// or parsing a Snowflake with a custom Epoch (such as Discord.)
//
// Arguments:
//     e (uint64): The value of the epoch
//
// Returns:
//     (Option): A Snowflake Option instance to be consumed
//               by NewWithOptions.
func WithEpoch(e uint64) Option {
	return func(S *Snowflake) {
		S.epoch = e
	}
}

// WithTime is a function to set a custom time for generating
// or parsing a Snowflake.
//
// Arguments:
//     t (time.Time): The value of the custom time
//
// Returns:
//     (Option): A Snowflake Option instance to be consumed
//               by NewWithOptions.
func WithTime(t time.Time) Option {
	return func(S *Snowflake) {
		S.Timestamp = t
	}
}

// WithWorkerID is a function to set a custom WorkerID for
// generating a Snowflake with a specific WorkerID.
//
// This Option has no bearing on how a WorkerID would be parsed
// by the Unmarshal function.
//
// Arguments:
//     id (uint64): The value of the WorkerID
//
// Returns:
//     (Option): A Snowflake Option instance to be consumed
//               by NewWithOptions.
func WithWorkerID(id uint64) Option {
	return func(S *Snowflake) {
		S.InternalWorkerID = id
	}
}

// WithProcessID is a function to set a custom ProcessID for
// generating a Snowflake with a specific ProcessID.
//
// This Option has no bearing on how a ProcessID would be parsed
// by the Unmarshal function.
//
// Arguments:
//     id (uint64): The value of the ProcessID
//
// Returns:
//     (Option): A Snowflake Option instance to be consumed
//               by NewWithOptions.
func WithProcessID(id uint64) Option {
	return func(S *Snowflake) {
		S.InternalProcessID = id
	}
}

// WithIncrimentID is a function to set a custom IncrimentID for
// generating a Snowflake with a specific IncrimentID.
//
// This Option has no bearing on how an IncrimentID would be parsed
// by the Unmarshal function.
//
// Arguments:
//     id (uint64): The value of the IncrimentID
//
// Returns:
//     (Option): A Snowflake Option instance to be consumed
//               by NewWithOptions.
func WithIncrimentID(id uint64) Option {
	return func(S *Snowflake) {
		S.IncrimentID = id
	}
}

// NewWithOptions generates a new Snowflake with custom
// values, or uses the default value if an option is not
// specified. Default values are:
//     * The standard UNIX epoch
//     * The current time
//     * A WorkerID of 0
//     * A ProcessID of 0
//     * An IncrimentID of 0
//
// Arguments:
//     None
//
// Returns:
//     (*Snowflake): A pointer to the new instance of Snowflake.
func NewWithOptions(opts ...Option) *Snowflake {
	S := &Snowflake{
		Timestamp: time.Now(),
	}

	for _, opt := range opts {
		opt(S)
	}

	return S
}
