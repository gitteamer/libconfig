/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2018 Aliaksandr Valialkin
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 * Author: Aliaksandr Valialkin <valyala@gmail.com>
 */
package libconfig

import "math/big"

var handyPool ParserPool

// GetString returns string value for the field identified by keys path
// in JSON data.
//
// Array indexes may be represented as decimal numbers in keys.
//
// An empty string is returned on error. Use Parser for proper error handling.
//
// Parser is faster for obtaining multiple fields from JSON.
func GetString(data []byte, keys ...string) string {
	p := handyPool.Get()
	v, err := p.ParseBytes(data)
	if err != nil {
		handyPool.Put(p)
		return ""
	}
	sb := v.GetStringBytes(keys...)
	str := string(sb)
	handyPool.Put(p)
	return str
}

// GetBytes returns string value for the field identified by keys path
// in JSON data.
//
// Array indexes may be represented as decimal numbers in keys.
//
// nil is returned on error. Use Parser for proper error handling.
//
// Parser is faster for obtaining multiple fields from JSON.
func GetBytes(data []byte, keys ...string) []byte {
	p := handyPool.Get()
	v, err := p.ParseBytes(data)
	if err != nil {
		handyPool.Put(p)
		return nil
	}
	sb := v.GetStringBytes(keys...)

	// Make a copy of sb, since sb belongs to p.
	var b []byte
	if sb != nil {
		b = append(b, sb...)
	}

	handyPool.Put(p)
	return b
}

// GetInt returns int value for the field identified by keys path
// in JSON data.
//
// Array indexes may be represented as decimal numbers in keys.
//
// 0 is returned on error. Use Parser for proper error handling.
//
// Parser is faster for obtaining multiple fields from JSON.
func GetInt(data []byte, keys ...string) int {
	p := handyPool.Get()
	v, err := p.ParseBytes(data)
	if err != nil {
		handyPool.Put(p)
		return 0
	}
	n := v.GetInt(keys...)
	handyPool.Put(p)
	return n
}

func GetHex(data []byte, keys ...string) string {
	p := handyPool.Get()
	v, err := p.ParseBytes(data)
	if err != nil {
		handyPool.Put(p)
		return ""
	}
	n := v.GetHex(keys...)
	handyPool.Put(p)
	return n
}

func GetBigint(data []byte, keys ...string) *big.Int {
	p := handyPool.Get()
	v, err := p.ParseBytes(data)
	if err != nil {
		handyPool.Put(p)
		return big.NewInt(0)
	}
	n := v.GetBigint(keys...)
	handyPool.Put(p)
	return n
}

// GetFloat64 returns float64 value for the field identified by keys path
// in JSON data.
//
// Array indexes may be represented as decimal numbers in keys.
//
// 0 is returned on error. Use Parser for proper error handling.
//
// Parser is faster for obtaining multiple fields from JSON.
func GetFloat64(data []byte, keys ...string) float64 {
	p := handyPool.Get()
	v, err := p.ParseBytes(data)
	if err != nil {
		handyPool.Put(p)
		return 0
	}
	f := v.GetFloat64(keys...)
	handyPool.Put(p)
	return f
}

// GetBool returns boolean value for the field identified by keys path
// in JSON data.
//
// Array indexes may be represented as decimal numbers in keys.
//
// False is returned on error. Use Parser for proper error handling.
//
// Parser is faster for obtaining multiple fields from JSON.
func GetBool(data []byte, keys ...string) bool {
	p := handyPool.Get()
	v, err := p.ParseBytes(data)
	if err != nil {
		handyPool.Put(p)
		return false
	}
	b := v.GetBool(keys...)
	handyPool.Put(p)
	return b
}

// Exists returns true if the field identified by keys path exists in JSON data.
//
// Array indexes may be represented as decimal numbers in keys.
//
// False is returned on error. Use Parser for proper error handling.
//
// Parser is faster when multiple fields must be checked in the JSON.
func Exists(data []byte, keys ...string) bool {
	p := handyPool.Get()
	v, err := p.ParseBytes(data)
	if err != nil {
		handyPool.Put(p)
		return false
	}
	ok := v.Exists(keys...)
	handyPool.Put(p)
	return ok
}

// Parse parses json string s.
//
// The function is slower than the Parser.Parse for re-used Parser.
func Parse(s string) (*Value, error) {
	var p Parser
	return p.Parse(s)
}

// MustParse parses json string s.
//
// The function panics if s cannot be parsed.
// The function is slower than the Parser.Parse for re-used Parser.
func MustParse(s string) *Value {
	v, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return v
}

// ParseBytes parses b containing json.
//
// The function is slower than the Parser.ParseBytes for re-used Parser.
func ParseBytes(b []byte) (*Value, error) {
	var p Parser
	return p.ParseBytes(b)
}

// MustParseBytes parses b containing json.
//
// The function panics if b cannot be parsed.
// The function is slower than the Parser.ParseBytes for re-used Parser.
func MustParseBytes(b []byte) *Value {
	v, err := ParseBytes(b)
	if err != nil {
		panic(err)
	}
	return v
}
