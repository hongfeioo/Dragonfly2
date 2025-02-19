/*
 *     Copyright 2020 The Dragonfly Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package structutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructToMap(t *testing.T) {
	tests := []struct {
		name   string
		s      interface{}
		expect func(*testing.T, map[string]interface{}, error)
	}{
		{
			name: "conver struct to map succeeded",
			s: struct {
				Name string
				Age  float64
			}{
				Name: "foo",
				Age:  18,
			},
			expect: func(t *testing.T, m map[string]interface{}, err error) {
				assert := assert.New(t)
				assert.Equal(m, map[string]interface{}{
					"Name": "foo",
					"Age":  float64(18),
				})
			},
		},
		{
			name: "conver string to map failed",
			s:    "foo",
			expect: func(t *testing.T, m map[string]interface{}, err error) {
				assert := assert.New(t)
				assert.EqualError(err, "json: cannot unmarshal string into Go value of type map[string]interface {}")
			},
		},
		{
			name: "conver number to map failed",
			s:    1,
			expect: func(t *testing.T, m map[string]interface{}, err error) {
				assert := assert.New(t)
				assert.EqualError(err, "json: cannot unmarshal number into Go value of type map[string]interface {}")
			},
		},
		{
			name: "conver nil to map succeeded",
			s:    nil,
			expect: func(t *testing.T, m map[string]interface{}, err error) {
				assert := assert.New(t)
				assert.Equal(m, map[string]interface{}(nil))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m, err := StructToMap(tc.s)
			tc.expect(t, m, err)
		})
	}
}
