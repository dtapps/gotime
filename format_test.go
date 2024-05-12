package gotime

import "testing"

func TestFormat(t *testing.T) {
	t.Log("DateTimeFormat", Current().SetFormat(DateTimeFormat))
	t.Log("DateTimeSFormat", Current().SetFormat(DateTimeSFormat))
	t.Log("DateTimeShrinkFormat", Current().SetFormat(DateTimeShrinkFormat))
	t.Log("DateTimeShrinkSFormat", Current().SetFormat(DateTimeShrinkSFormat))
	t.Log("DateFormat", Current().SetFormat(DateFormat))
	t.Log("TimeFormat", Current().SetFormat(TimeFormat))
	t.Log("TimeShrinkFormat", Current().SetFormat(TimeShrinkFormat))
}

func TestZhFormat(t *testing.T) {
	t.Log("DateTimeZhFormat", Current().SetFormat(DateTimeZhFormat))
	t.Log("DateTimeZhSFormat", Current().SetFormat(DateTimeZhSFormat))
	t.Log("DateTimeZhShrinkFormat", Current().SetFormat(DateTimeZhShrinkFormat))
	t.Log("DateTimeZhShrinkSFormat", Current().SetFormat(DateTimeZhShrinkSFormat))
	t.Log("DateZhFormat", Current().SetFormat(DateZhFormat))
	t.Log("TimeZhFormat", Current().SetFormat(TimeZhFormat))
	t.Log("TimeZhShrinkFormat", Current().SetFormat(TimeZhShrinkFormat))
}
