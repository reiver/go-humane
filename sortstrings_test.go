package humane_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-humane"
)

func TestSortStrings(t *testing.T) {

	tests := []struct{
		Strings []string
		Expected []string
	}{
		{
			Strings:  []string{},
			Expected: []string{},
		},



		{
			Strings:  []string{"banana"},
			Expected: []string{"banana"},
		},
		{
			Strings:  []string{"Banana"},
			Expected: []string{"Banana"},
		},
		{
			Strings:  []string{"bAnana"},
			Expected: []string{"bAnana"},
		},



		{
			Strings:  []string{"apple","banana","cherry"},
			Expected: []string{"apple","banana","cherry"},
		},
		{
			Strings:  []string{"apple","Banana","CHERRY","dAtE"},
			Expected: []string{"apple","Banana","CHERRY","dAtE"},
		},



		{
			Strings:  []string{"Banana","CHERRY","apple","dAtE"},
			Expected: []string{"apple","Banana","CHERRY","dAtE"},
		},



		{
			Strings:  []string{"A-uppercase","B-uppercase","C-uppercase","a-lowercase","b-lowercase","c-lowercase"},
			Expected: []string{"a-lowercase","A-uppercase","b-lowercase","B-uppercase","c-lowercase","C-uppercase"},
		},



		{
			Strings:  []string{
				"\u0000","\u0001","\u0002","\u0003","\u0004","\u0005","\u0006","\u0007","\u0008","\u0009","\u000A","\u000B","\u000C","\u000D","\u000E","\u000F",
				"\u0010","\u0011","\u0012","\u0013","\u0014","\u0015","\u0016","\u0017","\u0018","\u0019","\u001A","\u001B","\u001C","\u001D","\u001E","\u001F",
				" ",
				"!","\"","#","$","%","&","'","(",")","*","+",",","-",".","/",
				"0","1","2","3","4","5","6","7","8","9",
				":",";","<","=",">","?","@",
				"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z",
				"[","\\","]","^","_","`",
				"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z",
				"{","|","}","~","\u007F",
			},
			Expected: []string{
				"\u0000","\u0001","\u0002","\u0003","\u0004","\u0005","\u0006","\u0007","\u0008","\u0009","\u000A","\u000B","\u000C","\u000D","\u000E","\u000F",
				"\u0010","\u0011","\u0012","\u0013","\u0014","\u0015","\u0016","\u0017","\u0018","\u0019","\u001A","\u001B","\u001C","\u001D","\u001E","\u001F",
				" ",
				"!","\"","#","$","%","&","'","(",")","*","+",",","-",".","/",
				"0","1","2","3","4","5","6","7","8","9",
				"A","a",
				"B","b",
				"C","c",
				"D","d",
				"E","e",
				"F","f",
				"G","g",
				"H","h",
				"I","i",
				"J","j",
				"K","k",
				"L","l",
				"M","m",
				"N","n",
				"O","o",
				"P","p",
				"Q","q",
				"R","r",
				"S","s",
				"T","t",
				"U","u",
				"V","v",
				"W","w",
				"X","x",
				"Y","y",
				"Z","z",
				":",";","<","=",">","?","@",
				"[","\\","]","^","_","`",
				"{","|","}","~","\u007F",
			},
		},









		{
			Strings:  []string{"AbCdE","BcDeF","aBcDe","bCdEf"},
			Expected: []string{"AbCdE","aBcDe","BcDeF","bCdEf"},
		},



		{
			Strings:  []string{"BaNaNa","bAnAn","BaNa","bAn","Ba","b"},
			Expected: []string{"b","Ba","bAn","BaNa","bAnAn","BaNaNa"},
		},



		{
			Strings:  []string{"BaNaNa","bAnAn","BaNa","bAn","Ba","b","bAnAnA","BaNaN","bAnA","BaN","bA","B","banana","BANANA"},
			Expected: []string{"B","b","Ba","bA","BaN","bAn","BaNa","bAnA","BaNaN","bAnAn","BANANA","BaNaNa","bAnAnA","banana"},
		},
	}

	for testNumber, test := range tests {

		var actual []string = make([]string, len(test.Strings))
		copy(actual, test.Strings)
		humane.SortStrings(actual)

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual sorted-strings is not what was expected." , testNumber)
			t.Logf("EXPECTED: (%d) %#v", len(expected), expected)
			t.Logf("ACTUAL:   (%d) %#v", len(actual), actual)
			t.Logf("STRINGS:  (%d) %#v", len(test.Strings), test.Strings)
			continue
		}
	}
}
