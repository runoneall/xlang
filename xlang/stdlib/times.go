package stdlib

import (
	"time"
	"xlang/xlang"
)

var timesModule = map[string]xlang.Object{
	"format_ansic":        &xlang.String{Value: time.ANSIC},
	"format_unix_date":    &xlang.String{Value: time.UnixDate},
	"format_ruby_date":    &xlang.String{Value: time.RubyDate},
	"format_rfc822":       &xlang.String{Value: time.RFC822},
	"format_rfc822z":      &xlang.String{Value: time.RFC822Z},
	"format_rfc850":       &xlang.String{Value: time.RFC850},
	"format_rfc1123":      &xlang.String{Value: time.RFC1123},
	"format_rfc1123z":     &xlang.String{Value: time.RFC1123Z},
	"format_rfc3339":      &xlang.String{Value: time.RFC3339},
	"format_rfc3339_nano": &xlang.String{Value: time.RFC3339Nano},
	"format_kitchen":      &xlang.String{Value: time.Kitchen},
	"format_stamp":        &xlang.String{Value: time.Stamp},
	"format_stamp_milli":  &xlang.String{Value: time.StampMilli},
	"format_stamp_micro":  &xlang.String{Value: time.StampMicro},
	"format_stamp_nano":   &xlang.String{Value: time.StampNano},
	"nanosecond":          &xlang.Int{Value: int64(time.Nanosecond)},
	"microsecond":         &xlang.Int{Value: int64(time.Microsecond)},
	"millisecond":         &xlang.Int{Value: int64(time.Millisecond)},
	"second":              &xlang.Int{Value: int64(time.Second)},
	"minute":              &xlang.Int{Value: int64(time.Minute)},
	"hour":                &xlang.Int{Value: int64(time.Hour)},
	"january":             &xlang.Int{Value: int64(time.January)},
	"february":            &xlang.Int{Value: int64(time.February)},
	"march":               &xlang.Int{Value: int64(time.March)},
	"april":               &xlang.Int{Value: int64(time.April)},
	"may":                 &xlang.Int{Value: int64(time.May)},
	"june":                &xlang.Int{Value: int64(time.June)},
	"july":                &xlang.Int{Value: int64(time.July)},
	"august":              &xlang.Int{Value: int64(time.August)},
	"september":           &xlang.Int{Value: int64(time.September)},
	"october":             &xlang.Int{Value: int64(time.October)},
	"november":            &xlang.Int{Value: int64(time.November)},
	"december":            &xlang.Int{Value: int64(time.December)},
	"sleep": &xlang.UserFunction{
		Name:  "sleep",
		Value: timesSleep,
	}, // sleep(int)
	"parse_duration": &xlang.UserFunction{
		Name:  "parse_duration",
		Value: timesParseDuration,
	}, // parse_duration(str) => int
	"since": &xlang.UserFunction{
		Name:  "since",
		Value: timesSince,
	}, // since(time) => int
	"until": &xlang.UserFunction{
		Name:  "until",
		Value: timesUntil,
	}, // until(time) => int
	"duration_hours": &xlang.UserFunction{
		Name:  "duration_hours",
		Value: timesDurationHours,
	}, // duration_hours(int) => float
	"duration_minutes": &xlang.UserFunction{
		Name:  "duration_minutes",
		Value: timesDurationMinutes,
	}, // duration_minutes(int) => float
	"duration_nanoseconds": &xlang.UserFunction{
		Name:  "duration_nanoseconds",
		Value: timesDurationNanoseconds,
	}, // duration_nanoseconds(int) => int
	"duration_seconds": &xlang.UserFunction{
		Name:  "duration_seconds",
		Value: timesDurationSeconds,
	}, // duration_seconds(int) => float
	"duration_string": &xlang.UserFunction{
		Name:  "duration_string",
		Value: timesDurationString,
	}, // duration_string(int) => string
	"month_string": &xlang.UserFunction{
		Name:  "month_string",
		Value: timesMonthString,
	}, // month_string(int) => string
	"date": &xlang.UserFunction{
		Name:  "date",
		Value: timesDate,
	}, // date(year, month, day, hour, min, sec, nsec) => time
	"now": &xlang.UserFunction{
		Name:  "now",
		Value: timesNow,
	}, // now() => time
	"parse": &xlang.UserFunction{
		Name:  "parse",
		Value: timesParse,
	}, // parse(format, str) => time
	"unix": &xlang.UserFunction{
		Name:  "unix",
		Value: timesUnix,
	}, // unix(sec, nsec) => time
	"add": &xlang.UserFunction{
		Name:  "add",
		Value: timesAdd,
	}, // add(time, int) => time
	"add_date": &xlang.UserFunction{
		Name:  "add_date",
		Value: timesAddDate,
	}, // add_date(time, years, months, days) => time
	"sub": &xlang.UserFunction{
		Name:  "sub",
		Value: timesSub,
	}, // sub(t time, u time) => int
	"after": &xlang.UserFunction{
		Name:  "after",
		Value: timesAfter,
	}, // after(t time, u time) => bool
	"before": &xlang.UserFunction{
		Name:  "before",
		Value: timesBefore,
	}, // before(t time, u time) => bool
	"time_year": &xlang.UserFunction{
		Name:  "time_year",
		Value: timesTimeYear,
	}, // time_year(time) => int
	"time_month": &xlang.UserFunction{
		Name:  "time_month",
		Value: timesTimeMonth,
	}, // time_month(time) => int
	"time_day": &xlang.UserFunction{
		Name:  "time_day",
		Value: timesTimeDay,
	}, // time_day(time) => int
	"time_weekday": &xlang.UserFunction{
		Name:  "time_weekday",
		Value: timesTimeWeekday,
	}, // time_weekday(time) => int
	"time_hour": &xlang.UserFunction{
		Name:  "time_hour",
		Value: timesTimeHour,
	}, // time_hour(time) => int
	"time_minute": &xlang.UserFunction{
		Name:  "time_minute",
		Value: timesTimeMinute,
	}, // time_minute(time) => int
	"time_second": &xlang.UserFunction{
		Name:  "time_second",
		Value: timesTimeSecond,
	}, // time_second(time) => int
	"time_nanosecond": &xlang.UserFunction{
		Name:  "time_nanosecond",
		Value: timesTimeNanosecond,
	}, // time_nanosecond(time) => int
	"time_unix": &xlang.UserFunction{
		Name:  "time_unix",
		Value: timesTimeUnix,
	}, // time_unix(time) => int
	"time_unix_nano": &xlang.UserFunction{
		Name:  "time_unix_nano",
		Value: timesTimeUnixNano,
	}, // time_unix_nano(time) => int
	"time_format": &xlang.UserFunction{
		Name:  "time_format",
		Value: timesTimeFormat,
	}, // time_format(time, format) => string
	"time_location": &xlang.UserFunction{
		Name:  "time_location",
		Value: timesTimeLocation,
	}, // time_location(time) => string
	"time_string": &xlang.UserFunction{
		Name:  "time_string",
		Value: timesTimeString,
	}, // time_string(time) => string
	"is_zero": &xlang.UserFunction{
		Name:  "is_zero",
		Value: timesIsZero,
	}, // is_zero(time) => bool
	"to_local": &xlang.UserFunction{
		Name:  "to_local",
		Value: timesToLocal,
	}, // to_local(time) => time
	"to_utc": &xlang.UserFunction{
		Name:  "to_utc",
		Value: timesToUTC,
	}, // to_utc(time) => time
	"in_location": &xlang.UserFunction{
		Name:  "in_location",
		Value: timesInLocation,
	}, // in_location(time, location) => time
}

func timesSleep(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	i1, ok := xlang.ToInt64(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	time.Sleep(time.Duration(i1))
	ret = xlang.UndefinedValue

	return
}

func timesParseDuration(args ...xlang.Object) (
	ret xlang.Object,
	err error,
) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	s1, ok := xlang.ToString(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	dur, err := time.ParseDuration(s1)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &xlang.Int{Value: int64(dur)}

	return
}

func timesSince(args ...xlang.Object) (
	ret xlang.Object,
	err error,
) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.Int{Value: int64(time.Since(t1))}

	return
}

func timesUntil(args ...xlang.Object) (
	ret xlang.Object,
	err error,
) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.Int{Value: int64(time.Until(t1))}

	return
}

func timesDurationHours(args ...xlang.Object) (
	ret xlang.Object,
	err error,
) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	i1, ok := xlang.ToInt64(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.Float{Value: time.Duration(i1).Hours()}

	return
}

func timesDurationMinutes(args ...xlang.Object) (
	ret xlang.Object,
	err error,
) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	i1, ok := xlang.ToInt64(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.Float{Value: time.Duration(i1).Minutes()}

	return
}

func timesDurationNanoseconds(args ...xlang.Object) (
	ret xlang.Object,
	err error,
) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	i1, ok := xlang.ToInt64(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.Int{Value: time.Duration(i1).Nanoseconds()}

	return
}

func timesDurationSeconds(args ...xlang.Object) (
	ret xlang.Object,
	err error,
) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	i1, ok := xlang.ToInt64(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.Float{Value: time.Duration(i1).Seconds()}

	return
}

func timesDurationString(args ...xlang.Object) (
	ret xlang.Object,
	err error,
) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	i1, ok := xlang.ToInt64(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.String{Value: time.Duration(i1).String()}

	return
}

func timesMonthString(args ...xlang.Object) (
	ret xlang.Object,
	err error,
) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	i1, ok := xlang.ToInt64(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.String{Value: time.Month(i1).String()}

	return
}

func timesDate(args ...xlang.Object) (
	ret xlang.Object,
	err error,
) {
	if len(args) < 7 || len(args) > 8 {
		err = xlang.ErrWrongNumArguments
		return
	}

	i1, ok := xlang.ToInt(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}
	i2, ok := xlang.ToInt(args[1])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}
	i3, ok := xlang.ToInt(args[2])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}
	i4, ok := xlang.ToInt(args[3])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}
	i5, ok := xlang.ToInt(args[4])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "fifth",
			Expected: "int(compatible)",
			Found:    args[4].TypeName(),
		}
		return
	}
	i6, ok := xlang.ToInt(args[5])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "sixth",
			Expected: "int(compatible)",
			Found:    args[5].TypeName(),
		}
		return
	}
	i7, ok := xlang.ToInt(args[6])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "seventh",
			Expected: "int(compatible)",
			Found:    args[6].TypeName(),
		}
		return
	}

	var loc *time.Location
	if len(args) == 8 {
		i8, ok := xlang.ToString(args[7])
		if !ok {
			err = xlang.ErrInvalidArgumentType{
				Name:     "eighth",
				Expected: "string(compatible)",
				Found:    args[7].TypeName(),
			}
			return
		}
		loc, err = time.LoadLocation(i8)
		if err != nil {
			ret = wrapError(err)
			return
		}
	} else {
		loc = time.Now().Location()
	}

	ret = &xlang.Time{
		Value: time.Date(i1,
			time.Month(i2), i3, i4, i5, i6, i7, loc),
	}

	return
}

func timesNow(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 0 {
		err = xlang.ErrWrongNumArguments
		return
	}

	ret = &xlang.Time{Value: time.Now()}

	return
}

func timesParse(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 2 {
		err = xlang.ErrWrongNumArguments
		return
	}

	s1, ok := xlang.ToString(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := xlang.ToString(args[1])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	parsed, err := time.Parse(s1, s2)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &xlang.Time{Value: parsed}

	return
}

func timesUnix(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 2 {
		err = xlang.ErrWrongNumArguments
		return
	}

	i1, ok := xlang.ToInt64(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := xlang.ToInt64(args[1])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &xlang.Time{Value: time.Unix(i1, i2)}

	return
}

func timesAdd(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 2 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := xlang.ToInt64(args[1])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &xlang.Time{Value: t1.Add(time.Duration(i2))}

	return
}

func timesSub(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 2 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := xlang.ToTime(args[1])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &xlang.Int{Value: int64(t1.Sub(t2))}

	return
}

func timesAddDate(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 4 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := xlang.ToInt(args[1])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	i3, ok := xlang.ToInt(args[2])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	i4, ok := xlang.ToInt(args[3])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}

	ret = &xlang.Time{Value: t1.AddDate(i2, i3, i4)}

	return
}

func timesAfter(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 2 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := xlang.ToTime(args[1])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	if t1.After(t2) {
		ret = xlang.TrueValue
	} else {
		ret = xlang.FalseValue
	}

	return
}

func timesBefore(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 2 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := xlang.ToTime(args[1])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.Before(t2) {
		ret = xlang.TrueValue
	} else {
		ret = xlang.FalseValue
	}

	return
}

func timesTimeYear(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.Int{Value: int64(t1.Year())}

	return
}

func timesTimeMonth(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.Int{Value: int64(t1.Month())}

	return
}

func timesTimeDay(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.Int{Value: int64(t1.Day())}

	return
}

func timesTimeWeekday(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.Int{Value: int64(t1.Weekday())}

	return
}

func timesTimeHour(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.Int{Value: int64(t1.Hour())}

	return
}

func timesTimeMinute(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.Int{Value: int64(t1.Minute())}

	return
}

func timesTimeSecond(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.Int{Value: int64(t1.Second())}

	return
}

func timesTimeNanosecond(args ...xlang.Object) (
	ret xlang.Object,
	err error,
) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.Int{Value: int64(t1.Nanosecond())}

	return
}

func timesTimeUnix(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.Int{Value: t1.Unix()}

	return
}

func timesTimeUnixNano(args ...xlang.Object) (
	ret xlang.Object,
	err error,
) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.Int{Value: t1.UnixNano()}

	return
}

func timesTimeFormat(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 2 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := xlang.ToString(args[1])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	s := t1.Format(s2)
	if len(s) > xlang.MaxStringLen {

		return nil, xlang.ErrStringLimit
	}

	ret = &xlang.String{Value: s}

	return
}

func timesIsZero(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.IsZero() {
		ret = xlang.TrueValue
	} else {
		ret = xlang.FalseValue
	}

	return
}

func timesToLocal(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.Time{Value: t1.Local()}

	return
}

func timesToUTC(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.Time{Value: t1.UTC()}

	return
}

func timesTimeLocation(args ...xlang.Object) (
	ret xlang.Object,
	err error,
) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.String{Value: t1.Location().String()}

	return
}

func timesInLocation(args ...xlang.Object) (
	ret xlang.Object,
	err error,
) {
	if len(args) != 2 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := xlang.ToString(args[1])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	location, err := time.LoadLocation(s2)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &xlang.Time{Value: t1.In(location)}

	return
}

func timesTimeString(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 1 {
		err = xlang.ErrWrongNumArguments
		return
	}

	t1, ok := xlang.ToTime(args[0])
	if !ok {
		err = xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &xlang.String{Value: t1.String()}

	return
}
