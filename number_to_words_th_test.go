package numbertowords

import (
	"testing"
)

func TestIntegerTH(t *testing.T) {
	t.Run("case 0 without decimal point", func(*testing.T) {
		str := "0"

		got, _ := readerTH(str, th)

		want := "ศูนย์"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 1-9 without decimal point", func(*testing.T) {
		str := "321"

		got, _ := readerTH(str, th)

		want := "สามร้อยยี่สิบเอ็ด"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 1-9 without decimal point", func(*testing.T) {
		str := "31"

		got, _ := readerTH(str, th)

		want := "สามสิบเอ็ด"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 1-9 without decimal point", func(*testing.T) {
		str := "5"

		got, _ := readerTH(str, th)

		want := "ห้า"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10 - 99 without decimal point", func(*testing.T) {
		str := "15"

		got, _ := readerTH(str, th)

		want := "สิบห้า"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10 - 99 without decimal point", func(*testing.T) {
		str := "22"

		got, _ := readerTH(str, th)

		want := "ยี่สิบสอง"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 100-999 without decimal point", func(*testing.T) {
		str := "555"

		got, _ := readerTH(str, th)

		want := "ห้าร้อยห้าสิบห้า"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 100-999 without decimal point", func(*testing.T) {
		str := "515"

		got, _ := readerTH(str, th)

		want := "ห้าร้อยสิบห้า"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 1000-9999 without decimal point", func(*testing.T) {
		str := "4414"

		got, _ := readerTH(str, th)

		want := "สี่พันสี่ร้อยสิบสี่"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10000-99999 without decimal point (1)", func(*testing.T) {
		str := "15512"

		got, _ := readerTH(str, th)

		want := "หนึ่งหมื่นห้าพันห้าร้อยสิบสอง"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10000-99999 without decimal point (2)", func(*testing.T) {
		str := "52152"

		got, _ := readerTH(str, th)

		want := "ห้าหมื่นสองพันหนึ่งร้อยห้าสิบสอง"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10000-99999 without decimal point (3)", func(*testing.T) {
		str := "15155"

		got, _ := readerTH(str, th)

		want := "หนึ่งหมื่นห้าพันหนึ่งร้อยห้าสิบห้า"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 100000-999999 without decimal point", func(*testing.T) {
		str := "141512"

		got, _ := readerTH(str, th)

		want := "หนึ่งแสนสี่หมื่นหนึ่งพันห้าร้อยสิบสอง"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 1000000-9999999 without decimal point", func(*testing.T) {
		str := "1111111"

		got, _ := readerTH(str, th)

		want := "หนึ่งล้านหนึ่งแสนหนึ่งหมื่นหนึ่งพันหนึ่งร้อยสิบเอ็ด"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10000000-99999999 without decimal point", func(*testing.T) {
		str := "22222222"

		got, _ := readerTH(str, th)

		want := "ยี่สิบสองล้านสองแสนสองหมื่นสองพันสองร้อยยี่สิบสอง"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})
}

func TestDecimalTH(t *testing.T) {
	t.Run("case no decimal", func(*testing.T) {
		str := "00"

		got, _ := decimalTH(str, th)

		want := "ถ้วน"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 0-9", func(*testing.T) {
		str := "01"

		got, _ := decimalTH(str, th)

		want := "หนึ่งสตางค์"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10-19", func(*testing.T) {
		str := "10"

		got, _ := decimalTH(str, th)

		want := "สิบสตางค์"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10-19", func(*testing.T) {
		str := "15"

		got, _ := decimalTH(str, th)

		want := "สิบห้าสตางค์"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 20-29", func(*testing.T) {
		str := "20"

		got, _ := decimalTH(str, th)

		want := "ยี่สิบสตางค์"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 20-29", func(*testing.T) {
		str := "25"

		got, _ := decimalTH(str, th)

		want := "ยี่สิบห้าสตางค์"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})
}

func TestConvertTH(t *testing.T) {
	t.Run("case 0-9 (1)", func(*testing.T) {
		number := 0.00

		got, _ := Convert(number, th)

		want := "ศูนย์บาทถ้วน"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 0-9 (2)", func(*testing.T) {
		number := 0.25

		got, _ := Convert(number, th)

		want := "ยี่สิบห้าสตางค์"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 0-9 (3)", func(*testing.T) {
		number := 9.99

		got, _ := Convert(number, th)

		want := "เก้าบาทเก้าสิบเก้าสตางค์"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10-99 (1)", func(*testing.T) {
		number := 10.00

		got, _ := Convert(number, th)

		want := "สิบบาทถ้วน"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10-99 (2)", func(*testing.T) {
		number := 22.22

		got, _ := Convert(number, th)

		want := "ยี่สิบสองบาทยี่สิบสองสตางค์"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 10-99 (3)", func(*testing.T) {
		number := 88.88

		got, _ := Convert(number, th)

		want := "แปดสิบแปดบาทแปดสิบแปดสตางค์"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})

	t.Run("case 100-999 (3)", func(*testing.T) {
		number := 111.11

		got, _ := Convert(number, th)

		want := "หนึ่งร้อยสิบเอ็ดบาทสิบเอ็ดสตางค์"

		if got != want {
			t.Errorf(`Expect "%v" got "%v"`, want, got)
		}
	})
}
