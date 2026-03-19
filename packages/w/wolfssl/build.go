package main

import "gitee.com/spock2300/vmake/pkg/api"

func Main(b *api.Builder) {
	b.OnPackage(func(ctx *api.PackageContext) {
		ctx.SetHomepage("https://www.wolfssl.com/").
			SetDescription("Lightweight SSL/TLS library optimized for embedded systems").
			SetLicense("GPL-2.0-or-later").
			SetGit("https://github.com/wolfSSL/wolfssl.git").
			AddVersion("5.7.2", "v5.7.2-stable").
			AddVersion("5.6.6", "v5.6.6-stable").
			SetLibs("wolfssl")
	})

	b.OnConfig(func(ctx *api.ConfigContext) {
		ctx.Option("shared").
			SetType(api.OptionBool).
			SetDefault(false).
			SetDescription("Build shared library")
	})

	b.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("wolfssl").
			SetKind(api.TargetVoid).
			SetBuildFunc(func(pctx *api.PackageContext) error {
				pctx.CMakeConfigure(
					"-DBUILD_SHARED_LIBS="+pctx.BoolStr("shared"),
					"-DWOLFSSL_EXAMPLES=OFF",
					"-DWOLFSSL_CRYPT_TESTS=OFF",
				)
				pctx.CMakeBuild("-j4")
				pctx.CMakeInstall()
				return nil
			})
	})
}
