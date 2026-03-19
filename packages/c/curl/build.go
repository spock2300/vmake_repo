package main

import "gitee.com/spock2300/vmake/pkg/api"

func Main(b *api.Builder) {
	b.OnPackage(func(ctx *api.PackageContext) {
		ctx.SetHomepage("https://curl.se/").
			SetDescription("Command line tool and library for transferring data with URLs").
			SetLicense("curl").
			SetGit(
				"https://gitee.com/mirrors/curl.git",
				"https://github.com/curl/curl.git",
			).
			AddVersion("8.6.0", "curl-8_6_0").
			AddVersion("8.5.0", "curl-8_5_0").
			DeclarePackages("official/zlib", "official/mbedtls", "official/openssl", "official/wolfssl").
			SetLibs("curl")
	})

	b.OnConfig(func(ctx *api.ConfigContext) {
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

	b.OnRequire(func(ctx *api.RequireContext) {
		ctx.AddRequires(
			"official/zlib",
			"official/mbedtls",
			"official/openssl",
			"official/wolfssl",
		)
	})

	b.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("curl").
			SetKind(api.TargetVoid).
			SetBuildFunc(func(pctx *api.PackageContext) error {
				ssl := pctx.String("ssl")
				zlib := pctx.Dep("official/zlib")

				args := []string{
					"-DBUILD_SHARED_LIBS=" + pctx.BoolStr("shared"),
					"-DCURL_DISABLE_LDAP=ON",
					"-DCURL_DISABLE_LDAPS=ON",
					"-DBUILD_CURL_EXE=OFF",
					"-DENABLE_TESTING=OFF",
					"-DZLIB_INCLUDE_DIR=" + zlib.IncludeDir,
					"-DZLIB_LIBRARY=" + zlib.LibDir + "/libz.a",
				}

				switch ssl {
				case "mbedtls":
					mbedtls := pctx.Dep("official/mbedtls")
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
					wolfssl := pctx.Dep("official/wolfssl")
					args = append(args,
						"-DCURL_USE_WOLFSSL=ON",
						"-DCURL_USE_OPENSSL=OFF",
						"-DCURL_USE_MBEDTLS=OFF",
						"-DWOLFSSL_INCLUDE_DIRS="+wolfssl.IncludeDir,
						"-DWOLFSSL_LIBRARY="+wolfssl.LibDir+"/libwolfssl.a",
					)
				default:
					openssl := pctx.Dep("official/openssl")
					args = append(args,
						"-DCURL_USE_OPENSSL=ON",
						"-DCURL_USE_MBEDTLS=OFF",
						"-DCURL_USE_WOLFSSL=OFF",
						"-DOPENSSL_INCLUDE_DIR="+openssl.IncludeDir,
						"-DOPENSSL_CRYPTO_LIBRARY="+openssl.LibDir+"/libcrypto.a",
						"-DOPENSSL_SSL_LIBRARY="+openssl.LibDir+"/libssl.a",
					)
				}

				pctx.CMakeConfigure(args...)
				pctx.CMakeBuild("-j4")
				pctx.CMakeInstall()
				return nil
			})
	})
}
