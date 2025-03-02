// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ottlfuncs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func Test_deleteKey(t *testing.T) {
	input := pcommon.NewMap()
	input.PutStr("test", "hello world")
	input.PutInt("test2", 3)
	input.PutBool("test3", true)

	target := &ottl.StandardGetSetter[pcommon.Map]{
		Getter: func(ctx pcommon.Map) (interface{}, error) {
			return ctx, nil
		},
	}

	tests := []struct {
		name   string
		target ottl.Getter[pcommon.Map]
		key    string
		want   func(pcommon.Map)
	}{
		{
			name:   "delete test",
			target: target,
			key:    "test",
			want: func(expectedMap pcommon.Map) {
				expectedMap.PutBool("test3", true)
				expectedMap.PutInt("test2", 3)
			},
		},
		{
			name:   "delete test2",
			target: target,
			key:    "test2",
			want: func(expectedMap pcommon.Map) {
				expectedMap.PutStr("test", "hello world")
				expectedMap.PutBool("test3", true)
			},
		},
		{
			name:   "delete nothing",
			target: target,
			key:    "not a valid key",
			want: func(expectedMap pcommon.Map) {
				expectedMap.PutStr("test", "hello world")
				expectedMap.PutInt("test2", 3)
				expectedMap.PutBool("test3", true)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scenarioMap := pcommon.NewMap()
			input.CopyTo(scenarioMap)

			exprFunc, err := DeleteKey(tt.target, tt.key)
			assert.NoError(t, err)

			_, err = exprFunc(scenarioMap)
			assert.Nil(t, err)

			expected := pcommon.NewMap()
			tt.want(expected)

			assert.Equal(t, expected, scenarioMap)
		})
	}
}

func Test_deleteKey_bad_input(t *testing.T) {
	input := pcommon.NewValueStr("not a map")
	target := &ottl.StandardGetSetter[interface{}]{
		Getter: func(ctx interface{}) (interface{}, error) {
			return ctx, nil
		},
		Setter: func(ctx interface{}, val interface{}) error {
			t.Errorf("nothing should be set in this scenario")
			return nil
		},
	}

	key := "anything"

	exprFunc, err := DeleteKey[interface{}](target, key)
	assert.NoError(t, err)
	result, err := exprFunc(input)
	assert.NoError(t, err)
	assert.Nil(t, result)
	assert.Equal(t, pcommon.NewValueStr("not a map"), input)
}

func Test_deleteKey_get_nil(t *testing.T) {
	target := &ottl.StandardGetSetter[interface{}]{
		Getter: func(ctx interface{}) (interface{}, error) {
			return ctx, nil
		},
		Setter: func(ctx interface{}, val interface{}) error {
			t.Errorf("nothing should be set in this scenario")
			return nil
		},
	}

	key := "anything"

	exprFunc, err := DeleteKey[interface{}](target, key)
	assert.NoError(t, err)
	result, err := exprFunc(nil)
	assert.NoError(t, err)
	assert.Nil(t, result)
}
