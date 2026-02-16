# VMake Official Repository

Official package repository for VMake build system.

## Usage

Add this repository to VMake:

```bash
vmake repo add official git@gitee.com:spock2300/vmake_repo.git
```

## Available Packages

| Package | Description | Build System | Options |
|---------|-------------|--------------|---------|
| sds | Simple Dynamic Strings for C | vmake Native | - |
| tinyexpr | Tiny expression evaluator for C | vmake Native | - |
| zlib | Compression library | CMake | shared |

## Usage Example

In your project's `build.go`:

```go
package main

import "gitee.com/spock2300/vmake/pkg/api"

func Main(b *api.Builder) {
    b.OnRequire(func(ctx *api.RequireContext) {
        ctx.AddRequires("official/zlib >=1.2.0")
    })
    
    b.OnBuild(func(ctx *api.BuildContext) {
        ctx.Target("app").
            SetKind(api.TargetBinary).
            AddFiles("src/*.c").
            AddPackages("official/zlib")
    })
}
```

## Package Details

### sds

Simple Dynamic Strings (SDS) is a string library for C that enhances standard C strings with dynamic memory management.

- Homepage: https://github.com/redis/sds
- License: BSD-2-Clause
- Versions: 2.0.0, 1.0.5

### tinyexpr

A very small recursive descent expression parser and evaluator for C.

- Homepage: https://codeplea.com/tinyexpr
- License: zlib
- Versions: 1.0.0

### zlib

A massively spiffy yet delicately unobtrusive compression library.

- Homepage: http://www.zlib.net
- License: zlib
- Versions: 1.3.1, 1.2.13
- Options:
  - `shared` (bool): Build shared library (default: false)
