package gotime

import (
	"testing"
)

func TestTime(t *testing.T) {
	t.Log("今天此刻 Now：", Current().Now())
	t.Log("测试 SetFormat", Current().SetFormat("20060102150405"))
	t.Log("今天此刻格式化 Format：", Current().Format())
	t.Log("今天此刻日期 ToDateFormat：", Current().ToDateFormat())
	t.Log("今天此刻日期 ToDateFormatTime：", Current().ToDateFormatTime())
	t.Log("今天此刻时间 ToTimeFormat：", Current().ToTimeFormat())
	t.Log("今天此刻时间戳 Timestamp：", Current().Timestamp())
	t.Log("今天此刻时间戳 TimestampWithSecond：", Current().TimestampWithSecond())
	t.Log("今天毫秒级时间戳 TimestampWithMillisecond：", Current().TimestampWithMillisecond())
	t.Log("今天微秒级时间戳 TimestampWithMicrosecond：", Current().TimestampWithMicrosecond())
	t.Log("今天纳秒级时间戳 TimestampWithNanosecond：", Current().TimestampWithNanosecond())

	t.Log("昨天此刻 Now：", Yesterday().Now())
	t.Log("昨天此刻格式化 Format：", Yesterday().Format())
	t.Log("昨天此刻日期 ToDateFormat：", Yesterday().ToDateFormat())
	t.Log("昨天此刻日期 ToDateFormatTime：", Yesterday().ToDateFormatTime())
	t.Log("昨天此刻时间 ToTimeFormat：", Yesterday().ToTimeFormat())
	t.Log("昨天此刻时间戳 Timestamp：", Yesterday().Timestamp())
	t.Log("昨天此刻时间戳 TimestampWithSecond：", Yesterday().TimestampWithSecond())
	t.Log("昨天毫秒级时间戳 TimestampWithMillisecond：", Yesterday().TimestampWithMillisecond())
	t.Log("昨天微秒级时间戳 TimestampWithMicrosecond：", Yesterday().TimestampWithMicrosecond())
	t.Log("昨天纳秒级时间戳 TimestampWithNanosecond：", Yesterday().TimestampWithNanosecond())

	t.Log("明天此刻 Now：", Tomorrow().Now())
	t.Log("明天此刻格式化 Format：", Tomorrow().Format())
	t.Log("明天此刻日期 ToDateFormat：", Tomorrow().ToDateFormat())
	t.Log("明天此刻日期 ToDateFormatTime：", Tomorrow().ToDateFormatTime())
	t.Log("明天此刻时间 ToTimeFormat：", Tomorrow().ToTimeFormat())
	t.Log("明天此刻时间戳 Timestamp：", Tomorrow().Timestamp())
	t.Log("明天此刻时间戳 TimestampWithSecond：", Tomorrow().TimestampWithSecond())
	t.Log("明天毫秒级时间戳 TimestampWithMillisecond：", Tomorrow().TimestampWithMillisecond())
	t.Log("明天微秒级时间戳 TimestampWithMicrosecond：", Tomorrow().TimestampWithMicrosecond())
	t.Log("明天纳秒级时间戳 TimestampWithNanosecond：", Tomorrow().TimestampWithNanosecond())

	t.Log("本世纪开始时间：", Current().StartOfCentury().Format())
	t.Log("本世纪结束时间：", Current().EndOfCentury().Format())
	t.Log("本年代开始时间：", Current().StartOfDecade().Format())
	t.Log("本年代结束时间：", Current().EndOfDecade().Format())
	t.Log("本年开始时间：", Current().StartOfYear().Format())
	t.Log("本年结束时间：", Current().EndOfYear().Format())
	t.Log("本季度开始时间：", Current().StartOfQuarter().Format())
	t.Log("本季度结束时间：", Current().EndOfQuarter().Format())
	t.Log("本月开始时间：", Current().StartOfMonth().Format())
	t.Log("本月结束时间：", Current().EndOfMonth().Format())

	//t.Log("7100秒前的时间：", Current().BeforeSeconds(7100).Format())
	//t.Log("2小时前的时间：", Current().BeforeHour(2).Format())
	//t.Log("7100秒后的时间：", Current().AfterSeconds(7100).Format())
	//t.Log("2小时后的时间：", Current().AfterHour(2).Format())
}

func TestStartOfDay(t *testing.T) {
	t.Log(Current().Format())
	t.Log(Current().StartOfDay().Format())
	t.Log(Current().EndOfDay().Format())
	t.Log(Current().Timestamp())
	t.Log(Current().StartOfDay().Timestamp())
	t.Log(Current().EndOfDay().Timestamp())
	t.Log(Current().BeforeDay(1).Format())
	t.Log(Current().BeforeDay(1).StartOfDay().Format())
	t.Log(Current().BeforeDay(1).EndOfDay().Format())
	t.Log(Current().AfterDay(1).Format())
	t.Log(Current().AfterDay(1).StartOfDay().Format())
	t.Log(Current().AfterDay(1).EndOfDay().Format())
}

func TestDiff(t *testing.T) {
	t.Log(Current().DiffInHourWithAbs(SetCurrentParse("2021-11-26 14:50:00").Time))
	t.Log(Current().DiffInHour(SetCurrentParse("2021-11-26 14:50:00").Time))
	t.Log(Current().DiffInMinutesWithAbs(SetCurrentParse("2021-11-26 14:50:00").Time))
	t.Log(Current().DiffInMinutes(SetCurrentParse("2021-11-26 14:50:00").Time))
	t.Log(SetCurrentParse("2022-03-01T10:03:39+08:00").Format())
	t.Log(SetCurrentParse("2022-03-04T11:12:47+08:00").Format())
}

func TestUnix(t *testing.T) {
	t.Log(SetCurrentUnix(1640067240).Format())
	t.Log(Current().BeforeDay(3 - 2).StartOfDay().Format())
	t.Log(Current().BeforeDay(3 - 1).EndOfDay().Format())
}

func Test2(t *testing.T) {
	t.Log(Current().BeforeDay(1 + 1).Format())
	t.Log(Current().BeforeDay(1).Format())
	t.Log(Current().BeforeHour(24).Format())
	t.Log(Current().Format())
}

func TestGt(t *testing.T) {
	t.Log(SetCurrentParse("2022-07-29 15:05:24").Time)
	t.Log(Current().Lte(SetCurrentParse("2022-07-18 17:05:24").Time))
}

func TestCompare(t *testing.T) {
	t.Log("是否大于", Current().Gt(SetCurrentParse("2022-07-29 14:35:24").Time))
	t.Log("是否小于", Current().Lt(SetCurrentParse("2022-07-29 14:35:24").Time))
}

func TestTaoBao(t *testing.T) {
	var i int64 = 1
	for {
		if i > 3 {
			break
		}
		t.Log(i)
		t.Log(i * 24)
		t.Log((i * 24) - 24)
		t.Log(Current().BeforeHour(i * 24).Format())
		t.Log(Current().BeforeHour((i * 24) - 24).Format())
		i++
	}
}

func TestMT(t *testing.T) {

	var day int64 = 1
	t.Log(day)
	t.Log(Current().BeforeHour(24 * day).Format())
	t.Log(Current().BeforeHour(24 * (day - 1)).Format())

	day = 2
	t.Log(day)
	t.Log(Current().BeforeHour(24 * day).Format())
	t.Log(Current().BeforeHour(24 * (day - 1)).Format())

	day = 3
	t.Log(day)
	t.Log(Current().BeforeHour(24 * day).Format())
	t.Log(Current().BeforeHour(24 * (day - 1)).Format())

}
