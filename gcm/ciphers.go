// Copyright (C) 2017. See AUTHORS.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build gcm

package openssl

// #include "shim.h"
import "C"

import (
	"crypto/cipher"
	"errors"
	"runtime"
)

type unauthenticatedGCM struct {
	ctx *cipherCtx
	enc bool
}

func NewUnauthenticatedGCM(isEncrypter bool, key []byte, iv []byte) (cipher.BlockMode, error) {
	ctx, err := newCipherCtx()
	if err != nil {
		return nil, err
	}

	keyptr := (*C.uchar)(&key[0])
	ivptr := (*C.uchar)(&iv[0])

	ccipher := C.EVP_aes_256_gcm()

	if isEncrypter {
		if 1 != C.EVP_EncryptInit_ex(ctx.ctx, ccipher, nil, keyptr, ivptr) {
			return nil, errors.New("Error encrypt init")
		}
	} else {
		if 1 != C.EVP_DecryptInit_ex(ctx.ctx, ccipher, nil, keyptr, ivptr) {
			return nil, errors.New("Error decrypt init")
		}
	}

	return unauthenticatedGCM{ctx: ctx, enc: isEncrypter}, nil
}

func (ugcm unauthenticatedGCM) BlockSize() int {
	return 16
}

func (ugcm unauthenticatedGCM) CryptBlocks(dst, src []byte) {
	srcPtr := (*C.uchar)(&src[0])
	srcLen := C.int(len(src))

	dstPtr := (*C.uchar)(&dst[0])
	dstLen := C.int(0)
	if ugcm.enc {
		if 1 != C.EVP_EncryptUpdate(ugcm.ctx.ctx, dstPtr, &dstLen, srcPtr, srcLen) {
			panic(errors.New("Error encrypt update"))
		}
	} else {
		if 1 != C.EVP_DecryptUpdate(ugcm.ctx.ctx, dstPtr, &dstLen, srcPtr, srcLen) {
			panic(errors.New("Error decrypt update"))
		}
	}
	if dstLen > srcLen {
		panic("Unexpected length difference, possible buffer overrun") // I don't think this is possible with GCM
	}
}

type cipherCtx struct {
	ctx *C.EVP_CIPHER_CTX
}

func newCipherCtx() (*cipherCtx, error) {
	cctx := C.EVP_CIPHER_CTX_new()
	if cctx == nil {
		return nil, errors.New("failed to allocate cipher context")
	}
	ctx := &cipherCtx{cctx}
	runtime.SetFinalizer(ctx, func(ctx *cipherCtx) {
		C.EVP_CIPHER_CTX_free(ctx.ctx)
	})
	return ctx, nil
}
