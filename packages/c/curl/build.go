package main

import "github.com/spock2300/vmake/pkg/api"

func Main(p *api.Package) {
	p.OnPackage(func(p *api.Package) {
		p.SetHomepage("https://curl.se/").
			SetDescription("Command line tool and library for transferring data with URLs").
			SetLicense("curl").
			SetGit(
				"https://gitee.com/mirrors/curl.git",
				"https://github.com/curl/curl.git",
			).
			AddVersion("8.6.0", "curl-8_6_0").
			AddVersion("8.5.0", "curl-8_5_0")
	})

	p.OnConfig(func(ctx *api.ConfigContext) {
		ctx.Option("shared").
			SetType(api.OptionBool).
			SetDefault(false).
			SetDescription("Build shared library")

		ctx.Option("ssl").
			SetType(api.OptionChoice).
			SetDefault("mbedtls").
			SetValues("mbedtls", "openssl", "wolfssl").
			SetDescription("SSL backend")
	})

	p.OnRequire(func(ctx *api.RequireContext) {
		ctx.AddRequires("official/zlib")
		ctx.AddRequires(ctx.Equal("ssl", "mbedtls", "official/mbedtls"))
		ctx.AddRequires(ctx.Equal("ssl", "openssl", "official/openssl"))
		ctx.AddRequires(ctx.Equal("ssl", "wolfssl", "official/wolfssl"))
	})

	p.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("curl").
			SetKind(api.TargetVoid).
			AddProvidedLibs("curl").
			AddDeps("official/zlib").
			AddDeps(ctx.Equal("ssl", "mbedtls", "official/mbedtls")).
			AddDeps(ctx.Equal("ssl", "openssl", "official/openssl")).
			AddDeps(ctx.Equal("ssl", "wolfssl", "official/wolfssl")).
			SetBuildFunc(func(p *api.Package) error {
				ssl := p.String("ssl")
				deps := p.Deps()
				zlib := deps["official/zlib"]

				args := []string{
					"-DBUILD_SHARED_LIBS=" + p.BoolStr("shared"),
					"-DCURL_DISABLE_LDAP=ON",
					"-DCURL_DISABLE_LDAPS=ON",
					"-DBUILD_CURL_EXE=OFF",
					"-DENABLE_TESTING=OFF",
					"-DCMAKE_DEBUG_POSTFIX=",
					"-DZLIB_INCLUDE_DIR=" + zlib.IncludeDir,
					"-DZLIB_LIBRARY=" + zlib.LibDir + "/libz.a",
				}

				switch ssl {
				case "mbedtls":
					mbedtls := deps["official/mbedtls"]
					args = append(args,
						"-DCURL_USE_MBEDTLS=ON",
						"-DCURL_USE_OPENSSL=OFF",
						"-DCURL_USE_WOLFSSL=OFF",
						"-DMBEDTLS_INCLUDE_DIRS="+mbedtls.IncludeDir,
						"-DMBEDTLS_LIBRARY="+mbedtls.LibDir+"/libmbedtls.a",
						"-DMBEDX509_LIBRARY="+mbedtls.LibDir+"/libmbedx509.a",
						"-DMBEDCRYPTO_LIBRARY="+mbedtls.LibDir+"/libmbedcrypto.a",
					)
				case "wolfssl":
					wolfssl := deps["official/wolfssl"]
					args = append(args,
						"-DCURL_USE_WOLFSSL=ON",
						"-DCURL_USE_OPENSSL=OFF",
						"-DCURL_USE_MBEDTLS=OFF",
						"-DWolfSSL_INCLUDE_DIR="+wolfssl.IncludeDir,
						"-DWolfSSL_LIBRARY="+wolfssl.LibDir+"/libwolfssl.a",
					)
				default:
					openssl := deps["official/openssl"]
					args = append(args,
						"-DCURL_USE_OPENSSL=ON",
						"-DCURL_USE_MBEDTLS=OFF",
						"-DCURL_USE_WOLFSSL=OFF",
						"-DOPENSSL_INCLUDE_DIR="+openssl.IncludeDir,
						"-DOPENSSL_CRYPTO_LIBRARY="+openssl.LibDir+"/libcrypto.a",
						"-DOPENSSL_SSL_LIBRARY="+openssl.LibDir+"/libssl.a",
					)
				}

				p.CMakeConfigure(args...)
				p.CMakeBuild("-j4")
				p.CMakeInstall()
				return nil
			})
	})
}
