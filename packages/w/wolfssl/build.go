package main

import "gitee.com/spock2300/vmake/pkg/api"

func Main(p *api.Package) {
	p.OnPackage(func(p *api.Package) {
		p.SetHomepage("https://www.wolfssl.com/").
			SetDescription("Lightweight SSL/TLS library optimized for embedded systems").
			SetLicense("GPL-2.0-or-later").
			SetGit(
				"https://gitee.com/mirrors/wolfssl.git",
				"https://github.com/wolfSSL/wolfssl.git",
			).
			AddVersion("5.7.2", "v5.7.2-stable").
			AddVersion("5.6.6", "v5.6.6-stable").
			SetLibs("wolfssl")
	})

	p.OnConfig(func(ctx *api.ConfigContext) {
		ctx.Option("shared").
			SetType(api.OptionBool).
			SetDefault(false).
			SetDescription("Build shared library")
	})

	p.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("wolfssl").
			SetKind(api.TargetVoid).
			SetBuildFunc(func(p *api.Package) error {
				p.CMakeConfigure(
					"-DBUILD_SHARED_LIBS="+p.BoolStr("shared"),
					"-DWOLFSSL_EXAMPLES=OFF",
					"-DWOLFSSL_CRYPT_TESTS=OFF",
				)
				p.CMakeBuild("-j4")
				p.CMakeInstall()
				return nil
			})
	})
}
