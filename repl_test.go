package main

import (
	"testing"
)

func TestFirstInput(t *testing.T){

  cases := []struct{
    input     string
    expected  string
  }{
    {
      input : "CHARMANDER",
      expected : "",
    },
    {
      input : "Pickachu",
      expected: "",
    },
    {
      input: "Map",
      expected:"map",
    },
  }

  for _, ca := range cases {
    actual, _ := cleanInput(ca.input)
    expectedWord := ca.expected

    if actual != expectedWord {
      t.Errorf("Failed %v is not equal to %v", actual, expectedWord)
      continue
    }
  }

}

//func TestCleanInput(t *testing.T) {
//
//	cases := []struct {
//		input    string
//		expected []string
//	}{
//		{
//
//			input:    " hello world ",
//			expected: []string{"hello", "world"},
//		},
//    {
//
//			input:    " nice world ",
//			expected: []string{"nice", "world"},
//		},
//    {
//
//			input:    " there is a  world ",
//			expected: []string{"there","is","a", "world"},
//		},
//    {
//
//			input:    "which is nice",
//			expected: []string{"which","is", "nice"},
//		},
//
//   {
//
//			input:    "1234 12",
//			expected: []string{"1234", "12"},
//		},
//
//   {
//
//			input:    " my goodness me",
//			expected: []string{"my","goodness", "me"},
//		},
//
//
//
//	}
//
//	for _, c := range cases {
//		actual := cleanInput(c.input)
//    if len(actual) != len(c.expected){
//      t.Errorf("length of %v is not equal to %v", actual, c.expected)
//      continue
//    }
//
//		for i := range actual {
//			word := actual[i]
//			expectedWord := c.expected[i]
//			if word != expectedWord {
//				t.Errorf("failed: expecting %s, got %s", expectedWord, actual)
//        continue
//			}
//		}
//
//	}
//}
