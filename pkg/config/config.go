/*
Copyright 2023 The Alibaba Cloud Serverless Authors.
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

package config

import "time"

type Config struct {
	ClientAddr           string
	GcInterval           time.Duration
	IdleDurationBeforeGC time.Duration
}

var DefaultConfig = Config{
	ClientAddr:           "127.0.0.1:50051",
	// default 
	// GcInterval:           10 * time.Second,
	// IdleDurationBeforeGC: 5 * time.Minute,

	// 51 score: 51.1475
	// GcInterval:           15 * time.Second,
	// IdleDurationBeforeGC: 5 * time.Minute,

	// 51.2059
	// 51.4524
	// GcInterval:           20 * time.Second,
	// IdleDurationBeforeGC: 6 * time.Minute,

	GcInterval:           20 * time.Second,
	IdleDurationBeforeGC: 300 * time.Second,
}
