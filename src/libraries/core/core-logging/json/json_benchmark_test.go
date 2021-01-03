/*
Copyright 2019 The Blackspace Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package logs

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"go.uber.org/zap"
)

var record = struct {
	Error   error                  `json:"err"`
	Level   int                    `json:"v"`
	Message string                 `json:"msg"`
	Time    time.Time              `json:"ts"`
	Fields  map[string]interface{} `json:"fields"`
}{
	Error:   fmt.Errorf("test for error:%s", "default"),
	Level:   2,
	Message: "test",
	Time:    time.Unix(0, 123),
	Fields: map[string]interface{}{
		"str":     "foo",
		"int64-1": int64(1),
		"int64-2": int64(1),
		"float64": float64(1.0),
		"string1": "\n",
		"string2": "💩",
		"string3": "🤔",
		"string4": "🙊",
		"bool":    true,
		"request": struct {
			Method  string `json:"method"`
			Timeout int    `json:"timeout"`
			secret  string `json:"secret"`
		}{
			Method:  "GET",
			Timeout: 10,
			secret:  "pony",
		},
	},
}

func BenchmarkInfoLoggerInfo(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			jLogger := NewJSONLogger(nil, nil)
			jLogger.Info("test",
				"str", "foo",
				"int64-1", int64(1),
				"int64-2", int64(1),
				"float64", float64(1.0),
				"string1", "\n",
				"string2", "💩",
				"string3", "🤔",
				"string4", "🙊",
				"bool", true,
				"request", struct {
					Method  string `json:"method"`
					Timeout int    `json:"timeout"`
					secret  string `json:"secret"`
				}{
					Method:  "GET",
					Timeout: 10,
					secret:  "pony",
				},
			)
		}
	})
}

func BenchmarkInfoMLoggerInfo(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			jLogger := NewJSONLogger(nil, nil)
			jLogger.InfoM("test",
				zap.String("test0", "response0"),
				zap.String("test1", "response1"),
				zap.String("test2", "response2"),
				zap.String("test3", "response3"),
			)
		}
	})
}

func BenchmarkInfoLoggerInfoStandardJSON(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			json.Marshal(record)
		}
	})
}

func BenchmarkZapLoggerError(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			jLogger := NewJSONLogger(nil, nil)
			jLogger.Error(fmt.Errorf("test for error:%s", "default"),
				"test",
				"str", "foo",
				"int64-1", int64(1),
				"int64-2", int64(1),
				"float64", float64(1.0),
				"string1", "\n",
				"string2", "💩",
				"string3", "🤔",
				"string4", "🙊",
				"bool", true,
				"request", struct {
					Method  string `json:"method"`
					Timeout int    `json:"timeout"`
					secret  string `json:"secret"`
				}{
					Method:  "GET",
					Timeout: 10,
					secret:  "pony",
				},
			)
		}
	})
}

func BenchmarkErrorMLoggerInfo(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			jLogger := NewJSONLogger(nil, nil)
			jLogger.ErrorM(errors.New("test0"), "error occured")
		}
	})
}

func BenchmarkZapLoggerErrorStandardJSON(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			json.Marshal(record)
		}
	})
}

func BenchmarkZapLoggerV(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			jLogger := NewJSONLogger(nil, nil)
			jLogger.V(1).Info("test",
				"str", "foo",
				"int64-1", int64(1),
				"int64-2", int64(1),
				"float64", float64(1.0),
				"string1", "\n",
				"string2", "💩",
				"string3", "🤔",
				"string4", "🙊",
				"bool", true,
				"request", struct {
					Method  string `json:"method"`
					Timeout int    `json:"timeout"`
					secret  string `json:"secret"`
				}{
					Method:  "GET",
					Timeout: 10,
					secret:  "pony",
				},
			)
		}
	})
}
