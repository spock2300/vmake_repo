package main

import "gitee.com/spock2300/vmake/pkg/api"

func Package(p *api.Package) {
	p.SetHomepage("https://curl.se/").
		SetDescription("Command line tool and library for transferring data with URLs").
		SetLicense("curl").
		SetGit("https://github.com/curl/curl.git").
		SetLibs("curl").
		DeclarePackages("official/zlib", "official/mbedtls", "official/openssl", "official/wolfssl")

	p.AddVersion("8.6.0", "curl-8_6_0")
	p.AddVersion("8.5.0", "curl-8_5_0")

	p.Option("shared").
		SetType(api.OptionBool).
		SetDefault(false).
		SetDescription("Build shared library")

	p.Option("ssl").
		SetType(api.OptionString).
		SetDefault("mbedtls").
		SetDescription("SSL backend: mbedtls, openssl, or wolfssl")

	p.Build(func(ctx *api.PackageContext) {
		ssl := ctx.String("ssl")
		zlib := ctx.Dep("official/zlib")

		args := []string{
			"-DBUILD_SHARED_LIBS=" + ctx.BoolStr("shared"),
			"-DCURL_DISABLE_LDAP=ON",
			"-DCURL_DISABLE_LDAPS=ON",
			"-DBUILD_CURL_EXE=OFF",
			"-DENABLE_TESTING=OFF",
			"-DZLIB_INCLUDE_DIR=" + zlib.IncludeDir,
			"-DZLIB_LIBRARY=" + zlib.LibDir + "/libz.a",
		}

		switch ssl {
		case "mbedtls":
			mbedtls := ctx.Dep("official/mbedtls")
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
			wolfssl := ctx.Dep("official/wolfssl")
			args = append(args,
				"-DCURL_USE_WOLFSSL=ON",
				"-DCURL_USE_OPENSSL=OFF",
				"-DCURL_USE_MBEDTLS=OFF",
				"-DWOLFSSL_INCLUDE_DIRS="+wolfssl.IncludeDir,
				"-DWOLFSSL_LIBRARY="+wolfssl.LibDir+"/libwolfssl.a",
			)
		default:
			openssl := ctx.Dep("official/openssl")
			args = append(args,
				"-DCURL_USE_OPENSSL=ON",
				"-DCURL_USE_MBEDTLS=OFF",
				"-DCURL_USE_WOLFSSL=OFF",
				"-DOPENSSL_INCLUDE_DIR="+openssl.IncludeDir,
				"-DOPENSSL_CRYPTO_LIBRARY="+openssl.LibDir+"/libcrypto.a",
				"-DOPENSSL_SSL_LIBRARY="+openssl.LibDir+"/libssl.a",
			)
		}

		ctx.CMakeConfigure(args...)
		ctx.CMakeBuild("-j4")
		ctx.CMakeInstall()
	})
}
