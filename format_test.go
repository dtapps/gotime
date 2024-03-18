package gotime

import "testing"

func TestZhFormat(t *testing.T) {
	t.Log("DateTimeZhFormat", Current().SetFormat(DateTimeZhFormat))
	t.Log("DateTimeZhShrinkFormat", Current().SetFormat(DateTimeZhShrinkFormat))
	t.Log("DateZhFormat", Current().SetFormat(DateZhFormat))
	t.Log("TimeZhFormat", Current().SetFormat(TimeZhFormat))
	t.Log("TimeZhShrinkFormat", Current().SetFormat(TimeZhShrinkFormat))
}
