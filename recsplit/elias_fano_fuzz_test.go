//go:build gofuzzbeta
// +build gofuzzbeta

/*
   Copyright 2021 Erigon contributors

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

package recsplit

import (
	"testing"
)

// gotip test -trimpath -v -fuzz=FuzzEliasFano -fuzztime=10s ./recsplit

func FuzzEliasFano(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte) {
		if len(in)%2 == 1 {
			t.Skip()
		}
		if len(in) == 0 {
			t.Skip()
		}
		var ef DoubleEliasFano
		// Treat each byte of the sequence as difference between previous value and the next
		numBuckets := len(in) / 2
		cumKeys := make([]uint64, numBuckets+1)
		position := make([]uint64, numBuckets+1)
		for i, b := range in[:numBuckets] {
			cumKeys[i+1] = cumKeys[i] + uint64(b)
		}
		for i, b := range in[numBuckets:] {
			position[i+1] = position[i] + uint64(b)
		}
		ef.Build(cumKeys, position)
		// Try to read from ef
		for bucket := 0; bucket < numBuckets; bucket++ {
			cumKey, bitPos := ef.Get2(uint64(bucket))
			if cumKey != cumKeys[bucket] {
				t.Fatalf("bucket %d: cumKey from EF = %d, expected %d", bucket, cumKey, cumKeys[bucket])
			}
			if bitPos != position[bucket] {
				t.Fatalf("bucket %d: position from EF = %d, expected %d", bucket, bitPos, position[bucket])
			}
		}
		for bucket := 0; bucket < numBuckets; bucket++ {
			cumKey, cumKeysNext, bitPos := ef.Get3(uint64(bucket))
			if cumKey != cumKeys[bucket] {
				t.Fatalf("bucket %d: cumKey from EF = %d, expected %d", bucket, cumKey, cumKeys[bucket])
			}
			if bitPos != position[bucket] {
				t.Fatalf("bucket %d: position from EF = %d, expected %d", bucket, bitPos, position[bucket])
			}
			if cumKeysNext != cumKeys[bucket+1] {
				t.Fatalf("bucket %d: cumKeysNext from EF = %d, expected %d", bucket, cumKeysNext, cumKeys[bucket+1])
			}
		}
	})
}
