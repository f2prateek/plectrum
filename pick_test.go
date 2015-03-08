package plectrum

import "github.com/bmizerany/assert"
import "testing"

type test struct {
  keys []string
  expected map[string]interface{}
}

var data = map[string]interface{}{
    "foo": "foo",
    "bar": "bar",
    "qaz": "qaz",
    "qux": "qux",
  }

var tests = []test{
  test{[]string{"foo"}, map[string]interface{}{ "foo": "foo",}},
  test{[]string{"bar", "qaz"}, map[string]interface{}{ "bar": "bar","qaz": "qaz",}},
  test{[]string{"qux", "foo"}, map[string]interface{}{ "qux": "qux","foo": "foo",}},
  test{[]string{"foobarbqazqux"}, map[string]interface{}{}},
  test{[]string{}, map[string]interface{}{}},
}

func TestPick(t *testing.T) {
  for _, test := range tests {
    assert.Equal(t, Pick(data, test.keys), test.expected)
  }
}