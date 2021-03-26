# Unauthenticated GCM in Go

The Go stdlib implementation of GCM is quite restrictive in that it makes it difficult to do things incorrectly.
Sometimes you don't have a choice though. This library allows you to do streaming GCM encryption without authenticating data.
You should not do this unless you absolutely must.

This was forked from here: github.com/spacemonkeygo/openssl
The original code was more general purpose, but I've gutted most things I don't need. 

## Building

If you want to build this, you must provide the build tag "gcm". If you don't provide it a no-op implementation will be compiled instead.
This is useful for if you want to just switch off the encryption entirely during development (since it's slightly a pain to setup OpenSSL on your dev machine).

### License

Copyright (C) 2017. See AUTHORS.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
