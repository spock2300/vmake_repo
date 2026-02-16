package main

import "gitee.com/spock2300/vmake/pkg/api"

func Package(p *api.Package) {
	p.SetHomepage("https://www.trustedfirmware.org/projects/mbed-tls/").
		SetDescription("Lightweight, open source SSL/TLS library").
		SetLicense("Apache-2.0").
		SetGit("https://github.com/Mbed-TLS/mbedtls.git").
		SetLibs("mbedtls", "mbedx509", "mbedcrypto")

	p.AddVersion("3.6.0", "v3.6.0")
	p.AddVersion("2.28.8", "v2.28.8")

	p.Option("shared").
		SetType(api.OptionBool).
		SetDefault(false).
		SetDescription("Build shared library")

	p.Build(func(ctx *api.PackageContext) {
		ctx.CMakeConfigure(
			"-DUSE_SHARED_MBEDTLS_LIBRARY="+ctx.BoolStr("shared"),
			"-DENABLE_PROGRAMS=OFF",
			"-DENABLE_TESTING=OFF",
		)
		ctx.CMakeBuild("-j4")
		ctx.CMakeInstall()
	})
}
