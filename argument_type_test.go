package marker

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetArgumentTypeInfo(t *testing.T) {
	var rawValue []byte
	var interfaceValue interface{}
	testCases := []struct {
		Type              reflect.Type
		ShouldReturnError bool
		ExpectedType      ArgumentType
		ExpectedItemType  *ArgumentTypeInfo
	}{

		{
			Type:              reflect.TypeOf(make([]interface{}, 0)),
			ShouldReturnError: false,
			ExpectedType:      SliceType,
			ExpectedItemType:  &ArgumentTypeInfo{ActualType: AnyType},
		},
		{
			Type:              reflect.TypeOf(&rawValue),
			ShouldReturnError: false,
			ExpectedType:      RawType,
		},
		{
			Type:              reflect.TypeOf(&interfaceValue),
			ShouldReturnError: false,
			ExpectedType:      AnyType,
		},
		{
			Type:              reflect.TypeOf(true),
			ShouldReturnError: false,
			ExpectedType:      BoolType,
		},
		{
			Type:              reflect.TypeOf(uint8(0)),
			ShouldReturnError: false,
			ExpectedType:      UnsignedIntegerType,
		},
		{
			Type:              reflect.TypeOf(uint16(0)),
			ShouldReturnError: false,
			ExpectedType:      UnsignedIntegerType,
		},
		{
			Type:              reflect.TypeOf(uint(0)),
			ShouldReturnError: false,
			ExpectedType:      UnsignedIntegerType,
		},
		{
			Type:              reflect.TypeOf(uint32(0)),
			ShouldReturnError: false,
			ExpectedType:      UnsignedIntegerType,
		},
		{
			Type:              reflect.TypeOf(uint64(0)),
			ShouldReturnError: false,
			ExpectedType:      UnsignedIntegerType,
		},
		{
			Type:              reflect.TypeOf(int8(0)),
			ShouldReturnError: false,
			ExpectedType:      SignedIntegerType,
		},
		{
			Type:              reflect.TypeOf(int16(0)),
			ShouldReturnError: false,
			ExpectedType:      SignedIntegerType,
		},
		{
			Type:              reflect.TypeOf(0),
			ShouldReturnError: false,
			ExpectedType:      SignedIntegerType,
		},
		{
			Type:              reflect.TypeOf(int32(0)),
			ShouldReturnError: false,
			ExpectedType:      SignedIntegerType,
		},
		{
			Type:              reflect.TypeOf(int64(0)),
			ShouldReturnError: false,
			ExpectedType:      SignedIntegerType,
		},
		{
			Type:              reflect.TypeOf("test"),
			ShouldReturnError: false,
			ExpectedType:      StringType,
		},
		{
			Type:              reflect.TypeOf(make([]bool, 0)),
			ShouldReturnError: false,
			ExpectedType:      SliceType,
			ExpectedItemType:  &ArgumentTypeInfo{ActualType: BoolType},
		},
		{
			Type:              reflect.TypeOf(make([]interface{}, 0)),
			ShouldReturnError: false,
			ExpectedType:      SliceType,
			ExpectedItemType:  &ArgumentTypeInfo{ActualType: AnyType},
		},
		{
			Type:              reflect.TypeOf(make(map[string]int)),
			ShouldReturnError: false,
			ExpectedType:      MapType,
			ExpectedItemType:  &ArgumentTypeInfo{ActualType: SignedIntegerType},
		},
		{
			Type:              reflect.TypeOf(make(map[string]interface{})),
			ShouldReturnError: false,
			ExpectedType:      MapType,
			ExpectedItemType:  &ArgumentTypeInfo{ActualType: AnyType},
		},
		{
			Type: reflect.TypeOf(&struct {
			}{}),
			ShouldReturnError: true,
			ExpectedType:      InvalidType,
		},
	}

	for index, testCase := range testCases {
		typeInfo, err := ArgumentTypeInfoFromType(testCase.Type)
		if testCase.ShouldReturnError && err == nil {
			t.Errorf("%d. test case must have an error.", index)
		}

		if typeInfo.ActualType != testCase.ExpectedType {
			t.Errorf("actual type is not equal to expected, got %q; want %q", typeInfo.ActualType, testCase.ExpectedType)
		}

		if testCase.ExpectedItemType != nil && typeInfo.ItemType.ActualType != testCase.ExpectedItemType.ActualType {
			t.Errorf("item type is not equal to expected, got %q; want %q", typeInfo.ItemType.ActualType, testCase.ExpectedItemType.ActualType)
		}
	}
}

func TestArgumentTypeInfo_ParseString(t *testing.T) {
	typeInfo, err := ArgumentTypeInfoFromType(reflect.TypeOf("anyTest"))
	assert.Nil(t, err)
	assert.Equal(t, StringType, typeInfo.ActualType)

	strValue := ""

	scanner := NewScanner(" anyTestString ")
	scanner.Peek()

	err = typeInfo.parseString(scanner, reflect.ValueOf(&strValue))
	assert.Nil(t, err)
	assert.Equal(t, "anyTestString", strValue)

	scanner = NewScanner("\"anyTestString\"")
	scanner.Peek()

	err = typeInfo.parseString(scanner, reflect.ValueOf(&strValue))
	assert.Nil(t, err)
	assert.Equal(t, "anyTestString", strValue)

	scanner = NewScanner("`anyTestString`")
	scanner.Peek()

	err = typeInfo.parseString(scanner, reflect.ValueOf(&strValue))
	assert.Nil(t, err)
	assert.Equal(t, "anyTestString", strValue)

	scanner = NewScanner(" anyTestString ")
	scanner.Peek()

	err = typeInfo.Parse(scanner, reflect.ValueOf(&strValue))
	assert.Nil(t, err)
	assert.Equal(t, "anyTestString", strValue)

	scanner = NewScanner("\"anyTestString\"")
	scanner.Peek()

	err = typeInfo.Parse(scanner, reflect.ValueOf(&strValue))
	assert.Nil(t, err)
	assert.Equal(t, "anyTestString", strValue)

	scanner = NewScanner("`anyTestString`")
	scanner.Peek()

	err = typeInfo.Parse(scanner, reflect.ValueOf(&strValue))
	assert.Nil(t, err)
	assert.Equal(t, "anyTestString", strValue)
}

func TestArgumentTypeInfo_ParseBoolean(t *testing.T) {
	typeInfo, err := ArgumentTypeInfoFromType(reflect.TypeOf(true))
	assert.Nil(t, err)
	assert.Equal(t, BoolType, typeInfo.ActualType)

	boolValue := false

	scanner := NewScanner(" true ")
	scanner.Peek()

	err = typeInfo.parseBoolean(scanner, reflect.ValueOf(&boolValue))
	assert.Nil(t, err)
	assert.True(t, boolValue)

	scanner = NewScanner(" false ")
	scanner.Peek()

	err = typeInfo.parseBoolean(scanner, reflect.ValueOf(&boolValue))
	assert.Nil(t, err)
	assert.False(t, boolValue)

	scanner = NewScanner(" true ")
	scanner.Peek()

	err = typeInfo.Parse(scanner, reflect.ValueOf(&boolValue))
	assert.Nil(t, err)
	assert.True(t, boolValue)

	scanner = NewScanner(" false ")
	scanner.Peek()

	err = typeInfo.Parse(scanner, reflect.ValueOf(&boolValue))
	assert.Nil(t, err)
	assert.False(t, boolValue)
}

func TestArgumentTypeInfo_ParseInteger(t *testing.T) {
	typeInfo, err := ArgumentTypeInfoFromType(reflect.TypeOf(0))
	assert.Nil(t, err)
	assert.Equal(t, SignedIntegerType, typeInfo.ActualType)

	signedIntegerValue := 0

	scanner := NewScanner(" -091215 ")
	scanner.Peek()

	err = typeInfo.parseInteger(scanner, reflect.ValueOf(&signedIntegerValue))
	assert.Nil(t, err)
	assert.Equal(t, -91215, signedIntegerValue)

	scanner = NewScanner(" -070519 ")
	scanner.Peek()

	err = typeInfo.parseInteger(scanner, reflect.ValueOf(&signedIntegerValue))
	assert.Nil(t, err)
	assert.Equal(t, -70519, signedIntegerValue)

	typeInfo, err = ArgumentTypeInfoFromType(reflect.TypeOf(uint(0)))
	assert.Nil(t, err)
	assert.Equal(t, UnsignedIntegerType, typeInfo.ActualType)

	scanner = NewScanner(" -091215 ")
	scanner.Peek()

	err = typeInfo.Parse(scanner, reflect.ValueOf(&signedIntegerValue))
	assert.Nil(t, err)
	assert.Equal(t, -91215, signedIntegerValue)

	scanner = NewScanner(" -070519 ")
	scanner.Peek()

	err = typeInfo.Parse(scanner, reflect.ValueOf(&signedIntegerValue))
	assert.Nil(t, err)
	assert.Equal(t, -70519, signedIntegerValue)

	typeInfo, err = ArgumentTypeInfoFromType(reflect.TypeOf(uint(0)))
	assert.Nil(t, err)
	assert.Equal(t, UnsignedIntegerType, typeInfo.ActualType)

	unsignedIntegerValue := uint(0)

	scanner = NewScanner(" 091215 ")
	scanner.Peek()

	err = typeInfo.parseInteger(scanner, reflect.ValueOf(&unsignedIntegerValue))
	assert.Nil(t, err)
	assert.Equal(t, uint(91215), unsignedIntegerValue)

	scanner = NewScanner(" 070519 ")
	scanner.Peek()

	err = typeInfo.parseInteger(scanner, reflect.ValueOf(&unsignedIntegerValue))
	assert.Nil(t, err)
	assert.Equal(t, uint(70519), unsignedIntegerValue)

	scanner = NewScanner(" 091215 ")
	scanner.Peek()

	err = typeInfo.Parse(scanner, reflect.ValueOf(&unsignedIntegerValue))
	assert.Nil(t, err)
	assert.Equal(t, uint(91215), unsignedIntegerValue)

	scanner = NewScanner(" 070519 ")
	scanner.Peek()

	err = typeInfo.Parse(scanner, reflect.ValueOf(&unsignedIntegerValue))
	assert.Nil(t, err)
	assert.Equal(t, uint(70519), unsignedIntegerValue)
}

func TestArgumentTypeInfo_ParseMap(t *testing.T) {
	m := make(map[string]any)
	typeInfo, err := ArgumentTypeInfoFromType(reflect.TypeOf(&m))
	assert.Nil(t, err)
	assert.Equal(t, MapType, typeInfo.ActualType)
	assert.Equal(t, AnyType, typeInfo.ItemType.ActualType)

	scanner := NewScanner(" {anyKey1:123,anyKey2:true,anyKey3:\"anyValue1\",anyKey4:`anyValue2`} ")
	scanner.Peek()

	err = typeInfo.parseMap(scanner, reflect.ValueOf(&m))
	assert.Nil(t, err)
	assert.Contains(t, m, "anyKey1")
	assert.Equal(t, 123, m["anyKey1"])
	assert.Contains(t, m, "anyKey2")
	assert.Equal(t, true, m["anyKey2"])
	assert.Contains(t, m, "anyKey3")
	assert.Equal(t, "anyValue1", m["anyKey3"])
	assert.Contains(t, m, "anyKey4")
	assert.Equal(t, "anyValue2", m["anyKey4"])

	scanner = NewScanner(" {anyKey1:123,anyKey2:true,anyKey3:\"anyValue1\",anyKey4:`anyValue2`} ")
	scanner.Peek()

	err = typeInfo.Parse(scanner, reflect.ValueOf(&m))
	assert.Nil(t, err)
	assert.Contains(t, m, "anyKey1")
	assert.Equal(t, 123, m["anyKey1"])
	assert.Contains(t, m, "anyKey2")
	assert.Equal(t, true, m["anyKey2"])
	assert.Contains(t, m, "anyKey3")
	assert.Equal(t, "anyValue1", m["anyKey3"])
	assert.Contains(t, m, "anyKey4")
	assert.Equal(t, "anyValue2", m["anyKey4"])
}

func TestArgumentTypeInfo_ParseSlice(t *testing.T) {
	s := make([]int, 0)
	typeInfo, err := ArgumentTypeInfoFromType(reflect.TypeOf(&s))
	assert.Nil(t, err)
	assert.Equal(t, SliceType, typeInfo.ActualType)
	assert.Equal(t, SignedIntegerType, typeInfo.ItemType.ActualType)

	scanner := NewScanner(" 1;2;3;4;5 ")
	scanner.Peek()

	err = typeInfo.parseSlice(scanner, reflect.ValueOf(&s))
	assert.Nil(t, err)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, s)

	scanner = NewScanner(" 1;2;3;4;5 ")
	scanner.Peek()

	err = typeInfo.Parse(scanner, reflect.ValueOf(&s))
	assert.Nil(t, err)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, s)

	scanner = NewScanner(" {1,2,3,4,5} ")
	scanner.Peek()

	err = typeInfo.parseSlice(scanner, reflect.ValueOf(&s))
	assert.Nil(t, err)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, s)

	scanner = NewScanner(" {1,2,3,4,5} ")
	scanner.Peek()

	err = typeInfo.Parse(scanner, reflect.ValueOf(&s))
	assert.Nil(t, err)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, s)
}
