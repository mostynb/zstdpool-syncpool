// Copyright 2020 Mostyn Bramley-Moore.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package syncpool

import (
	"runtime"
	"sync"

	"github.com/klauspost/compress/zstd"
)

// EncoderWrapper is a wrapper that embeds a *zstd.Encoder, and is safe for
// use in a sync.Pool.
type EncoderWrapper struct {
	*zstd.Encoder
}

// NewEncoderPool returns a sync.Pool that provides EncoderWrapper objects
// which embed a *zstd.Encoder.
func NewEncoderPool(opts ...zstd.EOption) *sync.Pool {
	p := &sync.Pool{}

	p.New = func() interface{} {
		e, _ := zstd.NewWriter(nil, opts...)
		ew := &EncoderWrapper{Encoder: e}

		runtime.SetFinalizer(ew, func(ew *EncoderWrapper) {
			// Ensure that resources are freed by the *zstd.Encoder.
			ew.Close()
		})

		return ew
	}

	return p
}
