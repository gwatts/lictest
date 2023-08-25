package bar

import (
	"encoding/json"
	"fmt"
	"github.com/jmespath/go-jmespath"
)

func Bar() {
	var jsondata = []byte(`{"foo": {"bar": {"baz": [0, 1, 2, 3, 4]}}}`) // your data
	var data interface{}
	json.Unmarshal(jsondata, &data)
	result, _ := jmespath.Search("foo.bar.baz[2]", data)
	fmt.Println(result)
}
