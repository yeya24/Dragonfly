/*
 * Copyright The Dragonfly Authors.
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

package model

import (
	"encoding/json"
	"testing"

	"github.com/go-check/check"
	"gopkg.in/yaml.v2"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type RateSuite struct{}

func init() {
	check.Suite(&RateSuite{})
}

func (suite *RateSuite) TestParseRate(c *check.C) {
	var cases = []struct {
		input    string
		expected Rate
		isWrong  bool
	}{
		{"5m", 5 * MB, false},
		{"5M", 5 * MB, false},
		{"1B", B, false},
		{"10K", 10 * KB, false},
		{"10k", 10 * KB, false},
		{"10G", 10 * GB, false},
		{"10g", 10 * GB, false},
		{"10xx", 0, true},
	}

	for _, cc := range cases {
		output, err := ParseRate(cc.input)
		if !cc.isWrong {
			c.Assert(err, check.IsNil)
			c.Assert(output, check.Equals, cc.expected)
		} else {
			c.Assert(err, check.NotNil)
		}

	}
}

func (suite *RateSuite) TestString(c *check.C) {
	var cases = []struct {
		expected string
		input    Rate
	}{
		{"5M", 5 * MB},
		{"1B", B},
		{"10K", 10 * KB},
		{"1G", GB},
	}

	for _, cc := range cases {
		c.Check(cc.expected, check.Equals, cc.input.String())
	}
}

func (suite *RateSuite) TestMarshalJSON(c *check.C) {
	var cases = []struct {
		input  Rate
		output string
	}{
		{
			5 * MB,
			"\"5M\"",
		},
		{
			1 * GB,
			"\"1G\"",
		},
		{
			1 * B,
			"\"1B\"",
		},
		{
			1 * KB,
			"\"1K\"",
		},
	}

	for _, cc := range cases {
		data, err := json.Marshal(cc.input)
		c.Check(err, check.IsNil)
		c.Check(string(data), check.Equals, cc.output)
	}
}

func (suite *RateSuite) TestUnMarshalJSON(c *check.C) {
	var cases = []struct {
		output Rate
		input  string
	}{
		{
			5 * MB,
			"\"5M\"",
		},
		{
			1 * GB,
			"\"1G\"",
		},
		{
			1 * B,
			"\"1B\"",
		},
		{
			1 * KB,
			"\"1K\"",
		},
	}

	for _, cc := range cases {
		var r Rate
		err := json.Unmarshal([]byte(cc.input), &r)
		c.Check(err, check.IsNil)
		c.Check(r, check.Equals, cc.output)
	}
}

func (suite *RateSuite) TestMarshalYAML(c *check.C) {
	var cases = []struct {
		input  Rate
		output string
	}{
		{
			5 * MB,
			"5M\n",
		},
		{
			1 * GB,
			"1G\n",
		},
		{
			1 * B,
			"1B\n",
		},
		{
			1 * KB,
			"1K\n",
		},
	}

	for _, cc := range cases {
		data, err := yaml.Marshal(cc.input)
		c.Check(err, check.IsNil)
		c.Check(string(data), check.Equals, cc.output)
	}
}

func (suite *RateSuite) TestUnMarshalYAML(c *check.C) {
	var cases = []struct {
		output Rate
		input  string
	}{
		{
			5 * MB,
			"5M\n",
		},
		{
			1 * GB,
			"1G\n",
		},
		{
			1 * B,
			"1B\n",
		},
		{
			1 * KB,
			"1K\n",
		},
	}

	for _, cc := range cases {
		var output Rate
		err := yaml.Unmarshal([]byte(cc.input), &output)
		c.Check(err, check.IsNil)
		c.Check(output, check.Equals, cc.output)
	}
}
