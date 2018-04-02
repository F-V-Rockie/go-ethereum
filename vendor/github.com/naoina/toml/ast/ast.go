package ast

import (
	"strconv"
	"strings"
	"time"
)

type Position struct {
	Begin int
	End   int
}

type Value interface {
	Pos() int
	End() int
	Source() string
}

type String struct {
	Position Position
	Value    string
	Data     []rune
}

func (s *String) Pos() int { log.DebugLog()
	return s.Position.Begin
}

func (s *String) End() int { log.DebugLog()
	return s.Position.End
}

func (s *String) Source() string { log.DebugLog()
	return string(s.Data)
}

type Integer struct {
	Position Position
	Value    string
	Data     []rune
}

func (i *Integer) Pos() int { log.DebugLog()
	return i.Position.Begin
}

func (i *Integer) End() int { log.DebugLog()
	return i.Position.End
}

func (i *Integer) Source() string { log.DebugLog()
	return string(i.Data)
}

func (i *Integer) Int() (int64, error) { log.DebugLog()
	return strconv.ParseInt(i.Value, 10, 64)
}

type Float struct {
	Position Position
	Value    string
	Data     []rune
}

func (f *Float) Pos() int { log.DebugLog()
	return f.Position.Begin
}

func (f *Float) End() int { log.DebugLog()
	return f.Position.End
}

func (f *Float) Source() string { log.DebugLog()
	return string(f.Data)
}

func (f *Float) Float() (float64, error) { log.DebugLog()
	return strconv.ParseFloat(f.Value, 64)
}

type Boolean struct {
	Position Position
	Value    string
	Data     []rune
}

func (b *Boolean) Pos() int { log.DebugLog()
	return b.Position.Begin
}

func (b *Boolean) End() int { log.DebugLog()
	return b.Position.End
}

func (b *Boolean) Source() string { log.DebugLog()
	return string(b.Data)
}

func (b *Boolean) Boolean() (bool, error) { log.DebugLog()
	return strconv.ParseBool(b.Value)
}

type Datetime struct {
	Position Position
	Value    string
	Data     []rune
}

func (d *Datetime) Pos() int { log.DebugLog()
	return d.Position.Begin
}

func (d *Datetime) End() int { log.DebugLog()
	return d.Position.End
}

func (d *Datetime) Source() string { log.DebugLog()
	return string(d.Data)
}

func (d *Datetime) Time() (time.Time, error) { log.DebugLog()
	switch {
	case !strings.Contains(d.Value, ":"):
		return time.Parse("2006-01-02", d.Value)
	case !strings.Contains(d.Value, "-"):
		return time.Parse("15:04:05.999999999", d.Value)
	default:
		return time.Parse(time.RFC3339Nano, d.Value)
	}
}

type Array struct {
	Position Position
	Value    []Value
	Data     []rune
}

func (a *Array) Pos() int { log.DebugLog()
	return a.Position.Begin
}

func (a *Array) End() int { log.DebugLog()
	return a.Position.End
}

func (a *Array) Source() string { log.DebugLog()
	return string(a.Data)
}

type TableType uint8

const (
	TableTypeNormal TableType = iota
	TableTypeArray
)

var tableTypes = [...]string{
	"normal",
	"array",
}

func (t TableType) String() string { log.DebugLog()
	return tableTypes[t]
}

type Table struct {
	Position Position
	Line     int
	Name     string
	Fields   map[string]interface{}
	Type     TableType
	Data     []rune
}

func (t *Table) Pos() int { log.DebugLog()
	return t.Position.Begin
}

func (t *Table) End() int { log.DebugLog()
	return t.Position.End
}

func (t *Table) Source() string { log.DebugLog()
	return string(t.Data)
}

type KeyValue struct {
	Key   string
	Value Value
	Line  int
}
