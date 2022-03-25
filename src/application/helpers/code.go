package helpers

import "strconv"

type Code struct{}

func NewCode() *Code {
	return &Code{}
}

func (c Code) Generate(count uint32, prefix string, length uint8) string {
	code := strconv.FormatInt(int64(count), 10)

	for i := 0; i < int(length); i++ {
		if len(code) < int(length) {
			code = "0" + code
		}
	}

	return prefix + code
}
