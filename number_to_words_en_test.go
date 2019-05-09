package numbertowords

import (
	"testing"
)

func TestIntegerEN(t *testing.T) {
	t.Run("case 0 without decimal point", func(*testing.T) {
		str := "0"

		got, _ := integerEN(str, en)

		want := "zero"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 1-9 without decimal point", func(*testing.T) {
		str := "5"

		got, _ := integerEN(str, en)

		want := "five"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10 - 99 without decimal point", func(*testing.T) {
		str := "11"

		got, _ := integerEN(str, en)

		want := "eleven"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10 - 99 without decimal point", func(*testing.T) {
		str := "15"

		got, _ := integerEN(str, en)

		want := "fifteen"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10-99", func(*testing.T) {
		str := "88"

		got, _ := integerEN(str, en)

		want := "eighty eight"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10-99", func(*testing.T) {
		str := "44"

		got, _ := integerEN(str, en)

		want := "fourty four"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 100-999 without decimal point", func(*testing.T) {
		str := "111"

		got, _ := integerEN(str, en)

		want := "onehundred eleven"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 100-999 without decimal point", func(*testing.T) {
		str := "555"

		got, _ := integerEN(str, en)

		want := "fivehundred fifty five"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 100-999 without decimal point", func(*testing.T) {
		str := "515"

		got, _ := integerEN(str, en)

		want := "fivehundred fifteen"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 1000-9999 without decimal point", func(*testing.T) {
		str := "4414"

		got, _ := integerEN(str, en)

		want := "fourthousand fourhundred fourteen"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10000-99999 without decimal point (1)", func(*testing.T) {
		str := "15512"

		got, _ := integerEN(str, en)

		want := "fifteenthousand fivehundred twelve"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10000-99999 without decimal point (2)", func(*testing.T) {
		str := "52152"

		got, _ := integerEN(str, en)

		want := "fifty twothousand onehundred fifty two"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10000-99999 without decimal point (3)", func(*testing.T) {
		str := "15155"

		got, _ := integerEN(str, en)

		want := "fifteenthousand onehundred fifty five"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 100000-999999 without decimal point", func(*testing.T) {
		str := "141512"

		got, _ := integerEN(str, en)

		want := "onehundred fourty onethousand fivehundred twelve"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 1000000-9999999 without decimal point", func(*testing.T) {
		str := "1111111"

		got, _ := integerEN(str, en)

		want := "onemillion onehundred eleventhousand onehundred eleven"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10000000-99999999 without decimal point", func(*testing.T) {
		str := "22222222"

		got, _ := integerEN(str, en)

		want := "twenty twomillion twohundred twenty twothousand twohundred twenty two"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10000000-99999999 without decimal point", func(*testing.T) {
		str := "88888888"

		got, _ := integerEN(str, en)

		want := "eighty eightmillion eighthundred eighty eightthousand eighthundred eighty eight"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})
}

func TestDecimalEN(t *testing.T) {
	t.Run("case no decimal", func(*testing.T) {
		str := "00"

		got, _ := decimalEN(str, en)

		want := ""

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 0-9", func(*testing.T) {
		str := "01"

		got, _ := decimalEN(str, en)

		want := "point zero one"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10-19", func(*testing.T) {
		str := "10"

		got, _ := decimalEN(str, en)

		want := "point one"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10-19", func(*testing.T) {
		str := "15"

		got, _ := decimalEN(str, en)

		want := "point one five"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 20-29", func(*testing.T) {
		str := "20"

		got, _ := decimalEN(str, en)

		want := "point two"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 20-29", func(*testing.T) {
		str := "25"

		got, _ := decimalEN(str, en)

		want := "point two five"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})
}

func TestConvertEN(t *testing.T) {
	t.Run("case 0-9 (1)", func(*testing.T) {
		number := 0.00

		got, _ := Convert(number, en)

		want := "zero baht"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 0-9 (2)", func(*testing.T) {
		number := 0.25

		got, _ := Convert(number, en)

		want := "zero point two five baht"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 0-9 (3)", func(*testing.T) {
		number := 9.99

		got, _ := Convert(number, en)

		want := "nine point nine nine baht"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10-99 (1)", func(*testing.T) {
		number := 10.00

		got, _ := Convert(number, en)

		want := "ten baht"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10-99 (2)", func(*testing.T) {
		number := 22.22

		got, _ := Convert(number, en)

		want := "twenty two point two two baht"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10-99 (3)", func(*testing.T) {
		number := 88.88

		got, _ := Convert(number, en)

		want := "eighty eight point eight eight baht"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 100-999 (3)", func(*testing.T) {
		number := 111.11

		got, _ := Convert(number, en)

		want := "onehundred eleven point one one baht"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})
}
