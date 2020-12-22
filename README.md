# zstdpool-syncpool

github.com/klauspost/compress/zstd is a native Go implementation of Zstandard,
with an API that leaks memory and goroutines if you don't call `Close()`
on every `Encoder` and `Decoder` that you create, and `Decoders` cannot be
reused after they are closed.

zstdpool-syncpool is a Go wrapper for github.com/klauspost/compress/zstd
and sync.Pool, which automatically frees resources if you forget to call
`Close()` and/or when items are dropped from the sync.Pool.

Background: https://github.com/klauspost/compress/issues/264

# Usage

```
import (
	"github.com/klauspost/compress/zstd"
	syncpool "github.com/mostynb/zstdpool-syncpool"
)

// Create a sync.Pool which returns wrapped *zstd.Decoder's.
decoderPool := syncpool.NewDecoderPool(zstd.WithDecoderConcurrency(1))

// Get a DecoderWrapper and use it.
decoder := decoderPool.Get().(*syncpool.DecoderWrapper) 
decoder.Reset(compressedDataReader)
_, err = io.Copy(uncompressedDataWriter, decoder)

// Return the decoder to the pool. If we forget this, then the zstd.Decoder
// won't leak resources.
decoderPool.Put(decoder)


// Create a sync.Pool which returns wrapped *zstd.Endoder's.
encoderPool := syncpool.NewEncoderPool(zstd.WithEncoderConcurrency(1))

// Get an EncoderWrapper and use it.
encoder := encoderPool.Get().(*syncpool.EncoderWrapper)
encoder.Reset(compressedDataWriter)
_, err = io.Copy(encoder, uncompressedDataReader)

// Return the encoder to the pool. If we forget this, then the zstd.Encoder
// won't leak resources.
encoderPool.Put(encoder)
```

# Contributing

Bug reports, feature requests, PRs welcome.

# License

Licensed under the Apache License, Version 2.0. See the LICENSE file.
