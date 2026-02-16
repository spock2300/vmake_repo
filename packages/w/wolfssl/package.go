package main

import "gitee.com/spock2300/vmake/pkg/api"

func Package(p *api.Package) {
	p.SetHomepage("https://www.wolfssl.com/").
		SetDescription("Lightweight SSL/TLS library optimized for embedded systems").
		SetLicense("GPL-2.0-or-later").
		SetGit("https://github.com/wolfSSL/wolfssl.git").
		SetLibs("wolfssl")

	p.AddVersion("5.7.2", "v5.7.2-stable")
	p.AddVersion("5.6.6", "v5.6.6-stable")

	p.Option("shared").
		SetType(api.OptionBool).
		SetDefault(false).
		SetDescription("Build shared library")

	p.Build(func(ctx *api.PackageContext) {
		ctx.CMakeConfigure(
			"-DBUILD_SHARED_LIBS="+ctx.BoolStr("shared"),
			"-DWOLFSSL_EXAMPLES=OFF",
			"-DWOLFSSL_CRYPT_TESTS=OFF",
		)
		ctx.CMakeBuild("-j4")
		ctx.CMakeInstall()
	})
}
