package humane

import (
	"testing"

	"github.com/reiver/go-ascii"
)

func TestMapRune(t *testing.T) {

	tests := []struct{
		Rune rune
		Expected rune
	}{
		{
			Rune: ascii.NUL,
			Expected: 0x00,
		},
		{
			Rune: ascii.SOH,
			Expected: 0x01,
		},
		{
			Rune: ascii.STX,
			Expected: 0x02,
		},
		{
			Rune: ascii.ETX,
			Expected: 0x03,
		},
		{
			Rune: ascii.EOT,
			Expected: 0x04,
		},
		{
			Rune: ascii.ENQ,
			Expected: 0x05,
		},
		{
			Rune: ascii.ACK,
			Expected: 0x06,
		},
		{
			Rune: ascii.BEL,
			Expected: 0x07,
		},
		{
			Rune: ascii.BS,
			Expected: 0x08,
		},
		{
			Rune: ascii.HT,
			Expected: 0x09,
		},
		{
			Rune: ascii.LF,
			Expected: 0x0A,
		},
		{
			Rune: ascii.VT,
			Expected: 0x0B,
		},
		{
			Rune: ascii.FF,
			Expected: 0x0C,
		},
		{
			Rune: ascii.CR,
			Expected: 0x0D,
		},
		{
			Rune: ascii.SO,
			Expected: 0x0E,
		},
		{
			Rune: ascii.SI,
			Expected: 0x0F,
		},
		{
			Rune: ascii.DLE,
			Expected: 0x10,
		},
		{
			Rune: ascii.DC1, // == ascii.XON
			Expected: 0x11,
		},
		{
			Rune: ascii.DC2,
			Expected: 0x12,
		},
		{
			Rune: ascii.DC3, // == ascii.XOFF
			Expected: 0x13,
		},
		{
			Rune: ascii.DC4,
			Expected: 0x14,
		},
		{
			Rune: ascii.NAK,
			Expected: 0x15,
		},
		{
			Rune: ascii.SYN,
			Expected: 0x16,
		},
		{
			Rune: ascii.ETB,
			Expected: 0x17,
		},
		{
			Rune: ascii.CAN,
			Expected: 0x18,
		},
		{
			Rune: ascii.EM,
			Expected: 0x19,
		},
		{
			Rune: ascii.SUB,
			Expected: 0x1A,
		},
		{
			Rune: ascii.ESC,
			Expected: 0x1B,
		},
		{
			Rune: ascii.FS,
			Expected: 0x1C,
		},
		{
			Rune: ascii.GS,
			Expected: 0x1D,
		},
		{
			Rune: ascii.RS,
			Expected: 0x1E,
		},
		{
			Rune: ascii.US,
			Expected: 0x1F,
		},
		{
			Rune: ascii.SP,
			Expected: 0x20,
		},
		{
			Rune: '!',
			Expected: 0x21,
		},
		{
			Rune: '"',
			Expected: 0x22,
		},
		{
			Rune: '#',
			Expected: 0x23,
		},
		{
			Rune: '$',
			Expected: 0x24,
		},
		{
			Rune: '%',
			Expected: 0x25,
		},
		{
			Rune: '&',
			Expected: 0x26,
		},
		{
			Rune: '\'',
			Expected: 0x27,
		},
		{
			Rune: '(',
			Expected: 0x28,
		},
		{
			Rune: ')',
			Expected: 0x29,
		},
		{
			Rune: '*',
			Expected: 0x2A,
		},
		{
			Rune: '+',
			Expected: 0x2B,
		},
		{
			Rune: ',',
			Expected: 0x2C,
		},
		{
			Rune: '-',
			Expected: 0x2D,
		},
		{
			Rune: '.',
			Expected: 0x2E,
		},
		{
			Rune: '/',
			Expected: 0x2F,
		},
		{
			Rune: '0',
			Expected: 0x30,
		},
		{
			Rune: '1',
			Expected: 0x31,
		},
		{
			Rune: '2',
			Expected: 0x32,
		},
		{
			Rune: '3',
			Expected: 0x33,
		},
		{
			Rune: '4',
			Expected: 0x34,
		},
		{
			Rune: '5',
			Expected: 0x35,
		},
		{
			Rune: '6',
			Expected: 0x36,
		},
		{
			Rune: '7',
			Expected: 0x37,
		},
		{
			Rune: '8',
			Expected: 0x38,
		},
		{
			Rune: '9',
			Expected: 0x39,
		},









		{
			Rune: 'A',
			Expected: 0x3A,
		},
		{
			Rune: 'a',
			Expected: 0x3B,
		},
		{
			Rune: 'B',
			Expected: 0x3C,
		},
		{
			Rune: 'b',
			Expected: 0x3D,
		},
		{
			Rune: 'C',
			Expected: 0x3E,
		},
		{
			Rune: 'c',
			Expected: 0x3F,
		},
		{
			Rune: 'D',
			Expected: 0x40,
		},
		{
			Rune: 'd',
			Expected: 0x41,
		},
		{
			Rune: 'E',
			Expected: 0x42,
		},
		{
			Rune: 'e',
			Expected: 0x43,
		},
		{
			Rune: 'F',
			Expected: 0x44,
		},
		{
			Rune: 'f',
			Expected: 0x45,
		},
		{
			Rune: 'G',
			Expected: 0x46,
		},
		{
			Rune: 'g',
			Expected: 0x47,
		},
		{
			Rune: 'H',
			Expected: 0x48,
		},
		{
			Rune: 'h',
			Expected: 0x49,
		},
		{
			Rune: 'I',
			Expected: 0x4A,
		},
		{
			Rune: 'i',
			Expected: 0x4B,
		},
		{
			Rune: 'J',
			Expected: 0x4C,
		},
		{
			Rune: 'j',
			Expected: 0x4D,
		},
		{
			Rune: 'K',
			Expected: 0x4E,
		},
		{
			Rune: 'k',
			Expected: 0x4F,
		},
		{
			Rune: 'L',
			Expected: 0x50,
		},
		{
			Rune: 'l',
			Expected: 0x51,
		},
		{
			Rune: 'M',
			Expected: 0x52,
		},
		{
			Rune: 'm',
			Expected: 0x53,
		},
		{
			Rune: 'N',
			Expected: 0x54,
		},
		{
			Rune: 'n',
			Expected: 0x55,
		},
		{
			Rune: 'O',
			Expected: 0x56,
		},
		{
			Rune: 'o',
			Expected: 0x57,
		},
		{
			Rune: 'P',
			Expected: 0x58,
		},
		{
			Rune: 'p',
			Expected: 0x59,
		},
		{
			Rune: 'Q',
			Expected: 0x5A,
		},
		{
			Rune: 'q',
			Expected: 0x5B,
		},
		{
			Rune: 'R',
			Expected: 0x5C,
		},
		{
			Rune: 'r',
			Expected: 0x5D,
		},
		{
			Rune: 'S',
			Expected: 0x5E,
		},
		{
			Rune: 's',
			Expected: 0x5F,
		},
		{
			Rune: 'T',
			Expected: 0x60,
		},
		{
			Rune: 't',
			Expected: 0x61,
		},
		{
			Rune: 'U',
			Expected: 0x62,
		},
		{
			Rune: 'u',
			Expected: 0x63,
		},
		{
			Rune: 'V',
			Expected: 0x64,
		},
		{
			Rune: 'v',
			Expected: 0x65,
		},
		{
			Rune: 'W',
			Expected: 0x66,
		},
		{
			Rune: 'w',
			Expected: 0x67,
		},
		{
			Rune: 'X',
			Expected: 0x68,
		},
		{
			Rune: 'x',
			Expected: 0x69,
		},
		{
			Rune: 'Y',
			Expected: 0x6A,
		},
		{
			Rune: 'y',
			Expected: 0x6B,
		},
		{
			Rune: 'Z',
			Expected: 0x6C,
		},
		{
			Rune: 'z',
			Expected: 0x6D,
		},









		{
			Rune: ':',
			Expected: 0x6E,
		},
		{
			Rune: ';',
			Expected: 0x6F,
		},
		{
			Rune: '<',
			Expected: 0x70,
		},
		{
			Rune: '=',
			Expected: 0x71,
		},
		{
			Rune: '>',
			Expected: 0x72,
		},
		{
			Rune: '?',
			Expected: 0x73,
		},
		{
			Rune: '@',
			Expected: 0x74,
		},









		{
			Rune: '[',
			Expected: 0x75,
		},
		{
			Rune: '\\',
			Expected: 0x76,
		},
		{
			Rune: ']',
			Expected: 0x77,
		},
		{
			Rune: '^',
			Expected: 0x78,
		},
		{
			Rune: '_',
			Expected: 0x79,
		},
		{
			Rune: '`',
			Expected: 0x7A,
		},









		{
			Rune: '{',
			Expected: 0x7B,
		},
		{
			Rune: '|',
			Expected: 0x7C,
		},
		{
			Rune: '}',
			Expected: 0x7D,
		},
		{
			Rune: '~',
			Expected: 0x7E,
		},
		{
			Rune: ascii.DEL,
			Expected: 0x7F,
		},
		{
			Rune:     0x80,
			Expected: 0x80,
		},
		{
			Rune:     0x81,
			Expected: 0x81,
		},
		{
			Rune:     0x82,
			Expected: 0x82,
		},
		{
			Rune:     0x83,
			Expected: 0x83,
		},
		{
			Rune:     0x84,
			Expected: 0x84,
		},
		{
			Rune:     0x85,
			Expected: 0x85,
		},
		{
			Rune:     0x86,
			Expected: 0x86,
		},
		{
			Rune:     0x87,
			Expected: 0x87,
		},
		{
			Rune:     0x88,
			Expected: 0x88,
		},
		{
			Rune:     0x89,
			Expected: 0x89,
		},
		{
			Rune:     0x8A,
			Expected: 0x8A,
		},
		{
			Rune:     0x8B,
			Expected: 0x8B,
		},
		{
			Rune:     0x8C,
			Expected: 0x8C,
		},
		{
			Rune:     0x8D,
			Expected: 0x8D,
		},
		{
			Rune:     0x8E,
			Expected: 0x8E,
		},
		{
			Rune:     0x8F,
			Expected: 0x8F,
		},
		{
			Rune:     0x90,
			Expected: 0x90,
		},
		{
			Rune:     0x91,
			Expected: 0x91,
		},
		{
			Rune:     0x92,
			Expected: 0x92,
		},
		{
			Rune:     0x93,
			Expected: 0x93,
		},
		{
			Rune:     0x94,
			Expected: 0x94,
		},
		{
			Rune:     0x95,
			Expected: 0x95,
		},
		{
			Rune:     0x96,
			Expected: 0x96,
		},
		{
			Rune:     0x97,
			Expected: 0x97,
		},
		{
			Rune:     0x98,
			Expected: 0x98,
		},
		{
			Rune:     0x99,
			Expected: 0x99,
		},
		{
			Rune:     0x9A,
			Expected: 0x9A,
		},
		{
			Rune:     0x9B,
			Expected: 0x9B,
		},
		{
			Rune:     0x9C,
			Expected: 0x9C,
		},
		{
			Rune:     0x9D,
			Expected: 0x9D,
		},
		{
			Rune:     0x9E,
			Expected: 0x9E,
		},
		{
			Rune:     0x9F,
			Expected: 0x9F,
		},
		{
			Rune:     0xA0,
			Expected: 0xA0,
		},
		{
			Rune:     0xA1,
			Expected: 0xA1,
		},
		{
			Rune:     0xA2,
			Expected: 0xA2,
		},
		{
			Rune:     0xA3,
			Expected: 0xA3,
		},
		{
			Rune:     0xA4,
			Expected: 0xA4,
		},
		{
			Rune:     0xA5,
			Expected: 0xA5,
		},
		{
			Rune:     0xA6,
			Expected: 0xA6,
		},
		{
			Rune:     0xA7,
			Expected: 0xA7,
		},
		{
			Rune:     0xA8,
			Expected: 0xA8,
		},
		{
			Rune:     0xA9,
			Expected: 0xA9,
		},
		{
			Rune:     0xAA,
			Expected: 0xAA,
		},
		{
			Rune:     0xAB,
			Expected: 0xAB,
		},
		{
			Rune:     0xAC,
			Expected: 0xAC,
		},
		{
			Rune:     0xAD,
			Expected: 0xAD,
		},
		{
			Rune:     0xAE,
			Expected: 0xAE,
		},
		{
			Rune:     0xAF,
			Expected: 0xAF,
		},
		{
			Rune:     0xB0,
			Expected: 0xB0,
		},
		{
			Rune:     0xB1,
			Expected: 0xB1,
		},
		{
			Rune:     0xB2,
			Expected: 0xB2,
		},
		{
			Rune:     0xB3,
			Expected: 0xB3,
		},
		{
			Rune:     0xB4,
			Expected: 0xB4,
		},
		{
			Rune:     0xB5,
			Expected: 0xB5,
		},
		{
			Rune:     0xB6,
			Expected: 0xB6,
		},
		{
			Rune:     0xB7,
			Expected: 0xB7,
		},
		{
			Rune:     0xB8,
			Expected: 0xB8,
		},
		{
			Rune:     0xB9,
			Expected: 0xB9,
		},
		{
			Rune:     0xBA,
			Expected: 0xBA,
		},
		{
			Rune:     0xBB,
			Expected: 0xBB,
		},
		{
			Rune:     0xBC,
			Expected: 0xBC,
		},
		{
			Rune:     0xBD,
			Expected: 0xBD,
		},
		{
			Rune:     0xBE,
			Expected: 0xBE,
		},
		{
			Rune:     0xBF,
			Expected: 0xBF,
		},
		{
			Rune:     0xC0,
			Expected: 0xC0,
		},
		{
			Rune:     0xC1,
			Expected: 0xC1,
		},
		{
			Rune:     0xC2,
			Expected: 0xC2,
		},
		{
			Rune:     0xC3,
			Expected: 0xC3,
		},
		{
			Rune:     0xC4,
			Expected: 0xC4,
		},
		{
			Rune:     0xC5,
			Expected: 0xC5,
		},
		{
			Rune:     0xC6,
			Expected: 0xC6,
		},
		{
			Rune:     0xC7,
			Expected: 0xC7,
		},
		{
			Rune:     0xC8,
			Expected: 0xC8,
		},
		{
			Rune:     0xC9,
			Expected: 0xC9,
		},
		{
			Rune:     0xCA,
			Expected: 0xCA,
		},
		{
			Rune:     0xCB,
			Expected: 0xCB,
		},
		{
			Rune:     0xCC,
			Expected: 0xCC,
		},
		{
			Rune:     0xCD,
			Expected: 0xCD,
		},
		{
			Rune:     0xCE,
			Expected: 0xCE,
		},
		{
			Rune:     0xCF,
			Expected: 0xCF,
		},
		{
			Rune:     0xD0,
			Expected: 0xD0,
		},
		{
			Rune:     0xD1,
			Expected: 0xD1,
		},
		{
			Rune:     0xD2,
			Expected: 0xD2,
		},
		{
			Rune:     0xD3,
			Expected: 0xD3,
		},
		{
			Rune:     0xD4,
			Expected: 0xD4,
		},
		{
			Rune:     0xD5,
			Expected: 0xD5,
		},
		{
			Rune:     0xD6,
			Expected: 0xD6,
		},
		{
			Rune:     0xD7,
			Expected: 0xD7,
		},
		{
			Rune:     0xD8,
			Expected: 0xD8,
		},
		{
			Rune:     0xD9,
			Expected: 0xD9,
		},
		{
			Rune:     0xDA,
			Expected: 0xDA,
		},
		{
			Rune:     0xDB,
			Expected: 0xDB,
		},
		{
			Rune:     0xDC,
			Expected: 0xDC,
		},
		{
			Rune:     0xDD,
			Expected: 0xDD,
		},
		{
			Rune:     0xDE,
			Expected: 0xDE,
		},
		{
			Rune:     0xDF,
			Expected: 0xDF,
		},
		{
			Rune:     0xE0,
			Expected: 0xE0,
		},
		{
			Rune:     0xE1,
			Expected: 0xE1,
		},
		{
			Rune:     0xE2,
			Expected: 0xE2,
		},
		{
			Rune:     0xE3,
			Expected: 0xE3,
		},
		{
			Rune:     0xE4,
			Expected: 0xE4,
		},
		{
			Rune:     0xE5,
			Expected: 0xE5,
		},
		{
			Rune:     0xE6,
			Expected: 0xE6,
		},
		{
			Rune:     0xE7,
			Expected: 0xE7,
		},
		{
			Rune:     0xE8,
			Expected: 0xE8,
		},
		{
			Rune:     0xE9,
			Expected: 0xE9,
		},
		{
			Rune:     0xEA,
			Expected: 0xEA,
		},
		{
			Rune:     0xEB,
			Expected: 0xEB,
		},
		{
			Rune:     0xEC,
			Expected: 0xEC,
		},
		{
			Rune:     0xED,
			Expected: 0xED,
		},
		{
			Rune:     0xEE,
			Expected: 0xEE,
		},
		{
			Rune:     0xEF,
			Expected: 0xEF,
		},
		{
			Rune:     0xF0,
			Expected: 0xF0,
		},
		{
			Rune:     0xF1,
			Expected: 0xF1,
		},
		{
			Rune:     0xF2,
			Expected: 0xF2,
		},
		{
			Rune:     0xF3,
			Expected: 0xF3,
		},
		{
			Rune:     0xF4,
			Expected: 0xF4,
		},
		{
			Rune:     0xF5,
			Expected: 0xF5,
		},
		{
			Rune:     0xF6,
			Expected: 0xF6,
		},
		{
			Rune:     0xF7,
			Expected: 0xF7,
		},
		{
			Rune:     0xF8,
			Expected: 0xF8,
		},
		{
			Rune:     0xF9,
			Expected: 0xF9,
		},
		{
			Rune:     0xFA,
			Expected: 0xFA,
		},
		{
			Rune:     0xFB,
			Expected: 0xFB,
		},
		{
			Rune:     0xFC,
			Expected: 0xFC,
		},
		{
			Rune:     0xFD,
			Expected: 0xFD,
		},
		{
			Rune:     0xFE,
			Expected: 0xFE,
		},
		{
			Rune:     0xFF,
			Expected: 0xFF,
		},



		{
			Rune:     '۰',
			Expected: 0x06F0,
		},
		{
			Rune:     '۱',
			Expected: 0x06F1,
		},
		{
			Rune:     '۲',
			Expected: 0x06F2,
		},
		{
			Rune:     '۳',
			Expected: 0x06F3,
		},
		{
			Rune:     '۴',
			Expected: 0x06F4,
		},
		{
			Rune:     '۵',
			Expected: 0x06F5,
		},
		{
			Rune:     '۶',
			Expected: 0x06F6,
		},
		{
			Rune:     '۷',
			Expected: 0x06F7,
		},
		{
			Rune:     '۸',
			Expected: 0x06F8,
		},
		{
			Rune:     '۹',
			Expected: 0x06F9,
		},



		{
			Rune:     '👾',
			Expected: 0x1F47E,
		},



		{
			Rune:     '😈',
			Expected: 0x1F608,
		},



		{
			Rune:     '🙂',
			Expected: 0x1F642,
		},
	}

	for testNumber, test := range tests {

		actual := mapRune(test.Rune)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual rune-mapping is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q (%U)", expected, expected)
			t.Logf("ACTUAL  : %q (%U)", actual, actual)
			t.Logf("RUNE:     %q (%U)", test.Rune, test.Rune)
			continue
		}
	}
}
