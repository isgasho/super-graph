package json

import (
	"bytes"
	"testing"
)

var (
	input1 = `
	{ 
	"data": {
	"test": { "__twitter_id": "ABCD" },
	"users": [
		{
			"id": 1,
			"full_name": "Sidney Stroman",
			"email": "user0@demo.com",
			"__twitter_id": "2048666903444506956",
			"embed": {
				"id": 8,
				"full_name": "Caroll Orn Sr.",
				"email": "joannarau@hegmann.io",
				"__twitter_id": "ABC123"
			}
		},
		{
			"id": 2,
			"full_name": "Jerry Dickinson",
			"email": "user1@demo.com",
			"__twitter_id": [{ "name": "hello" }, { "name": "world"}]
		},
		{
			"id": 3,
			"full_name": "Kenna Cassin",
			"email": "user2@demo.com",
			"__twitter_id": { "name": "hello", "address": { "work": "1 infinity loop" } }
		},
		{
			"id": 4,
			"full_name": "Mr. Pat Parisian",
			"email": "__twitter_id",
			"__twitter_id": 1234567890
		},
		{
			"id": 5,
			"full_name": "Bette Ebert",
			"email": "janeenrath@goyette.com",
			"__twitter_id": 1.23E
		},
		{
			"id": 6,
			"full_name": "Everett Kiehn",
			"email": "michael@bartoletti.com",
			"__twitter_id": true
		},
		{
			"id": 7,
			"full_name": "Katrina Cronin",
			"email": "loretaklocko@framivolkman.org",
			"__twitter_id": false
		},
		{
			"id": 8,
			"full_name": "Caroll Orn Sr.",
			"email": "joannarau@hegmann.io",
			"__twitter_id": "2048666903444506956"
		},
		{
			"id": 9,
			"full_name": "Gwendolyn Ziemann",
			"email": "renaytoy@rutherford.co",
			"__twitter_id": ["hello", "world"]
		},
		{
			"id": 10,
			"full_name": "Mrs. Rosann Fritsch",
			"email": "holliemosciski@thiel.org",
			"__twitter_id": "2048666903444506956"
		},
		{
			"id": 11,
			"full_name": "Arden Koss",
			"email": "cristobalankunding@howewelch.org",
			"__twitter_id": "2048666903444506956"
		},
		{
			"id": 12,
			"full_name": "Brenton Bauch PhD",
			"email": "renee@miller.co",
			"__twitter_id": 1
		},
		{
			"id": 13,
			"full_name": "Daine Gleichner",
			"email": "andrea@gmail.com",
			"__twitter_id": "",
			"id__twitter_id": "NOOO",
			"work_email": "andrea@nienow.co"
		}
	]}
	}`

	input2 = `
	[{
		"id": 1,
		"full_name": "Sidney Stroman",
		"email": "user0@demo.com",
		"__twitter_id": "2048666903444506956",
		"embed": {
			"id": 8,
			"full_name": "Caroll Orn Sr.",
			"email": "joannarau@hegmann.io",
			"__twitter_id": "ABC123"
		}
	},
	{
		"m": 1,
		"id": 2,
		"full_name": "Jerry Dickinson",
		"email": "user1@demo.com",
		"__twitter_id": [{ "name": "hello" }, { "name": "world"}]
	}]`

	input3 = `
	{ 
		"data": {
			"test": { "__twitter_id": "ABCD" },
			"users": [{"id":1,"embed":{"id":8}},{"id":2},{"id":3},{"id":4},{"id":5},{"id":6},{"id":7},{"id":8},{"id":9},{"id":10},{"id":11},{"id":12},{"id":13}]
		}
	}`

	input4 = `
	[{
		"id": 1,
		"full_name": "Sidney Stroman",
		"email": "user0@demo.com",
		"__twitter_id": "2048666903444506956",
		"embed": {
			"id": 8,
			"full_name": "Caroll Orn Sr.",
			"email": "joannarau@hegmann.io",
			"__twitter_id": "ABC123"
		}
	},
	{
		"m": 1,
		"id": 2,
		"full_name": "Jerry Dickinson",
		"email": "user1@demo.com",
		"__twitter_id": [{ "name": "hello" }, { "name": "world"}]
	}]`
)

func TestGet(t *testing.T) {
	values := Get([]byte(input1), [][]byte{
		[]byte("__twitter_id"),
		[]byte("work_email"),
	})

	expected := []Field{
		{[]byte("__twitter_id"), []byte(`"ABCD"`)},
		{[]byte("__twitter_id"), []byte(`"2048666903444506956"`)},
		{[]byte("__twitter_id"), []byte(`"ABC123"`)},
		{[]byte("__twitter_id"),
			[]byte(`[{ "name": "hello" }, { "name": "world"}]`)},
		{[]byte("__twitter_id"),
			[]byte(`{ "name": "hello", "address": { "work": "1 infinity loop" } }`),
		},
		{[]byte("__twitter_id"), []byte(`1234567890`)},
		{[]byte("__twitter_id"), []byte(`1.23E`)},
		{[]byte("__twitter_id"), []byte(`true`)},
		{[]byte("__twitter_id"), []byte(`false`)},
		{[]byte("__twitter_id"), []byte(`"2048666903444506956"`)},
		{[]byte("__twitter_id"), []byte(`["hello", "world"]`)},
		{[]byte("__twitter_id"), []byte(`"2048666903444506956"`)},
		{[]byte("__twitter_id"), []byte(`"2048666903444506956"`)},
		{[]byte("__twitter_id"), []byte(`1`)},
		{[]byte("__twitter_id"), []byte(`""`)},
		{[]byte("work_email"), []byte(`"andrea@nienow.co"`)},
	}

	if len(values) != len(expected) {
		t.Fatal("len(values) != len(expected)")
	}

	for i := range expected {
		if bytes.Equal(values[i].Key, expected[i].Key) == false {
			t.Error(string(values[i].Key), " != ", string(expected[i].Key))
		}

		if bytes.Equal(values[i].Value, expected[i].Value) == false {
			t.Error(string(values[i].Value), " != ", string(expected[i].Value))
		}
	}
}

func TestFilter(t *testing.T) {
	var b bytes.Buffer
	Filter(&b, []byte(input2), []string{"id", "full_name", "embed"})

	expected := `[{"id": 1,"full_name": "Sidney Stroman","embed": {"id": 8,"full_name": "Caroll Orn Sr.","email": "joannarau@hegmann.io","__twitter_id": "ABC123"}},{"id": 2,"full_name": "Jerry Dickinson"}]`

	if b.String() != expected {
		t.Error("Does not match expected json")
	}
}

func TestStrip(t *testing.T) {
	value := Strip([]byte(input3), []string{"data", "users"})

	expected := []byte(`[{"id":1,"embed":{"id":8}},{"id":2},{"id":3},{"id":4},{"id":5},{"id":6},{"id":7},{"id":8},{"id":9},{"id":10},{"id":11},{"id":12},{"id":13}]`)

	if bytes.Equal(value, expected) == false {
		t.Log(value)
		t.Error("Does not match expected json")
	}
}

func TestReplace(t *testing.T) {
	var buf bytes.Buffer

	from := []Field{
		{[]byte("__twitter_id"), []byte(`[{ "name": "hello" }, { "name": "world"}]`)},
		{[]byte("__twitter_id"), []byte(`"ABC123"`)},
	}

	to := []Field{
		{[]byte("__twitter_id"), []byte(`"1234567890"`)},
		{[]byte("some_list"), []byte(`[{"id":1,"embed":{"id":8}},{"id":2},{"id":3},{"id":4},{"id":5},{"id":6},{"id":7},{"id":8},{"id":9},{"id":10},{"id":11},{"id":12},{"id":13}]`)},
	}

	expected := `[{
		"id": 1,
		"full_name": "Sidney Stroman",
		"email": "user0@demo.com",
		"__twitter_id": "2048666903444506956",
		"embed": {
			"id": 8,
			"full_name": "Caroll Orn Sr.",
			"email": "joannarau@hegmann.io",
			"some_list":[{"id":1,"embed":{"id":8}},{"id":2},{"id":3},{"id":4},{"id":5},{"id":6},{"id":7},{"id":8},{"id":9},{"id":10},{"id":11},{"id":12},{"id":13}]
		}
	},
	{
		"m": 1,
		"id": 2,
		"full_name": "Jerry Dickinson",
		"email": "user1@demo.com",
		"__twitter_id":"1234567890"
	}]`

	err := Replace(&buf, []byte(input4), from, to)
	if err != nil {
		t.Fatal(err)
	}

	if buf.String() != expected {
		t.Error("Does not match expected json")
	}
}

func BenchmarkGet(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		Get([]byte(input1), [][]byte{[]byte("__twitter_id")})
	}
}

func BenchmarkFilter(b *testing.B) {
	var buf bytes.Buffer

	keys := []string{"id", "full_name", "embed", "email", "__twitter_id"}
	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		err := Filter(&buf, []byte(input2), keys)
		if err != nil {
			b.Fatal(err)
		}
		buf.Reset()
	}
}

func BenchmarkStrip(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		Strip([]byte(input3), []string{"data", "users"})
	}
}

func BenchmarkReplace(b *testing.B) {
	var buf bytes.Buffer

	from := []Field{
		{[]byte("__twitter_id"), []byte(`[{ "name": "hello" }, { "name": "world"}]`)},
		{[]byte("__twitter_id"), []byte(`"ABC123"`)},
	}

	to := []Field{
		{[]byte("__twitter_id"), []byte(`"1234567890"`)},
		{[]byte("some_list"), []byte(`[{"id":1,"embed":{"id":8}},{"id":2},{"id":3},{"id":4},{"id":5},{"id":6},{"id":7},{"id":8},{"id":9},{"id":10},{"id":11},{"id":12},{"id":13}]`)},
	}

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		err := Replace(&buf, []byte(input4), from, to)
		if err != nil {
			b.Fatal(err)
		}
		buf.Reset()
	}
}
