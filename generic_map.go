// Copyright (c) 2022 Cisco and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package genericsync

import (
	"sync"
)

type Map[K any, V any] sync.Map

func (m *Map[K, V]) Delete(key K) {
	(*sync.Map)(m).Delete(key)
}

func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	v, ok := (*sync.Map)(m).Load(key)
	if v != nil {
		value = v.(V)
	}
	return value, ok
}

func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := (*sync.Map)(m).LoadAndDelete(key)
	if v != nil {
		value = v.(V)
	}
	return value, loaded
}

func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	v, loaded := (*sync.Map)(m).LoadOrStore(key, value)
	if v != nil {
		value = v.(V)
	}
	return value, loaded
}

func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	(*sync.Map)(m).Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func (m *Map[K, V]) Store(key K, value V) {
	(*sync.Map)(m).Store(key, value)
}
