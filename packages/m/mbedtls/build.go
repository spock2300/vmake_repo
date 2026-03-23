package main

import "gitee.com/spock2300/vmake/pkg/api"

func Main(p *api.Package) {
	p.OnPackage(func(p *api.Package) {
		p.SetHomepage("https://www.trustedfirmware.org/projects/mbed-tls/").
			SetDescription("Lightweight, open source SSL/TLS library").
			SetLicense("Apache-2.0").
			SetGit(
				"https://gitee.com/mirrors/mbedtls.git",
				"https://github.com/Mbed-TLS/mbedtls.git",
			).
			SetLibs("mbedtls", "mbedx509", "mbedcrypto").
			AddVersion("3.6.0", "v3.6.0").
			AddVersion("2.28.8", "v2.28.8")
	})

	p.OnConfig(func(ctx *api.ConfigContext) {
		ctx.Option("shared").
			SetType(api.OptionBool).
			SetDefault(false).
			SetDescription("Build shared library")
	})

	p.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("mbedtls").
			SetKind(api.TargetVoid).
			SetBuildFunc(func(p *api.Package) error {
				args := []string{
					"-DUSE_SHARED_MBEDTLS_LIBRARY=" + p.BoolStr("shared"),
					"-DENABLE_PROGRAMS=OFF",
					"-DENABLE_TESTING=OFF",
				}
				p.CMakeConfigure(args...)
				p.CMakeBuild("-j4")
				p.CMakeInstall()
				return nil
			})
	})
}
