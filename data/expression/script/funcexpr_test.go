package script

import (
	"bytes"
	"fmt"
	"github.com/project-flogo/core/data/resolve"
	"testing"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestBuiltinFuncExpr(t *testing.T) {

	expr, err := factory.NewExpr(`len("test")`)
	assert.Nil(t, err)

	v, err := expr.Eval(nil)
	assert.Nil(t, err)

	assert.Equal(t, 4, v)
}

func TestFuncExprNoSpace(t *testing.T) {
	expr, err := factory.NewExpr(`tstring.concat("a","b")`)
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, "ab", v)
}

func TestFuncExprNested(t *testing.T) {

	expr, err := factory.NewExpr(`tstring.concat("This", "is",tstring.concat("my","first"),"gocc",tstring.concat("lexer","and","parser"),tstring.concat("go","program","!!!"))`)
	assert.Nil(t, err)

	v, err := expr.Eval(nil)
	assert.Nil(t, err)

	assert.Equal(t, "Thisismyfirstgocclexerandparsergoprogram!!!", v.(string))
}

func TestFuncExprNestedMultiSpace(t *testing.T) {

	expr, err := factory.NewExpr(`tstring.concat("This",   " is" , " Flogo")`)
	assert.Nil(t, err)

	v, err := expr.Eval(nil)
	assert.Nil(t, err)

	assert.Equal(t, "This is Flogo", v.(string))
}

func TestFunctionWithRef(t *testing.T) {

	scope := data.NewSimpleScope(map[string]interface{}{"queryParams": map[string]interface{}{"id": "flogo"}}, nil)
	factory := NewExprFactory(resolve.GetBasicResolver())
	testcases := make(map[string]interface{})
	testcases[`tstring.concat("This", " is ", $.queryParams.id)`] = "This is flogo"

	for k, v := range testcases {
		vv, err := factory.NewExpr(k)
		assert.Nil(t, err)
		result, err := vv.Eval(scope)
		assert.Nil(t, err)
		if !assert.ObjectsAreEqual(v, result) {
			assert.Fail(t, fmt.Sprintf("test function [%s] failed, expected [%+v] but actual [%+v]", k, v, result))
		}
	}

}

func init() {
	function.Register(&fnConcat{})
}

type fnConcat struct {
}

func (fnConcat) Name() string {
	return "tstring.concat"
}

func (fnConcat) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, true
}

func (fnConcat) Eval(params ...interface{}) (interface{}, error) {
	if len(params) >= 2 {
		var buffer bytes.Buffer

		for _, v := range params {
			buffer.WriteString(v.(string))
		}
		return buffer.String(), nil
	}

	return "", fmt.Errorf("fnConcat function must have at least two arguments")
}

func TestFuncExprSingleQuote(t *testing.T) {
	expr, err := factory.NewExpr("tstring.concat('abc','def')")
	assert.Nil(t, err)
	v, err := expr.Eval(nil)
	assert.Nil(t, err)
	assert.Equal(t, "abcdef", v)
}

func init() {
	function.Register(&tLength{})
}

type tLength struct {
}

func (tLength) Name() string {
	return "tstring.length"
}

func (tLength) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, false
}

func (tLength) Eval(params ...interface{}) (interface{}, error) {
	p := params[0].(string)
	return len(p), nil
}
