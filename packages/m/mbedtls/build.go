package main

import "gitee.com/spock2300/vmake/pkg/api"

func Main(b *api.Builder) {
	b.OnPackage(func(ctx *api.PackageContext) {
		ctx.SetHomepage("https://www.trustedfirmware.org/projects/mbed-tls/").
			SetDescription("Lightweight, open source SSL/TLS library").
			SetLicense("Apache-2.0").
			SetGit("https://github.com/Mbed-TLS/mbedtls.git").
			SetLibs("mbedtls", "mbedx509", "mbedcrypto").
			AddVersion("3.6.0", "v3.6.0").
			AddVersion("2.28.8", "v2.28.8")
	})

	b.OnConfig(func(ctx *api.ConfigContext) {
		ctx.Option("shared").
			SetType(api.OptionBool).
			SetDefault(false).
			SetDescription("Build shared library")
	})

	b.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("mbedtls").
			SetKind(api.TargetVoid).
			SetBuildFunc(func(pctx *api.PackageContext) error {
				args := []string{
					"-DUSE_SHARED_MBEDTLS_LIBRARY=" + pctx.BoolStr("shared"),
					"-DENABLE_PROGRAMS=OFF",
					"-DENABLE_TESTING=OFF",
				}
				pctx.CMakeConfigure(args...)
				pctx.CMakeBuild("-j4")
				pctx.CMakeInstall()
				return nil
			})
	})
}
