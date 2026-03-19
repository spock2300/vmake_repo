package main

import "gitee.com/spock2300/vmake/pkg/api"

func Main(b *api.Builder) {
	b.OnPackage(func(ctx *api.PackageContext) {
		ctx.SetHomepage("https://www.openssl.org/").
			SetDescription("TLS/SSL and crypto library").
			SetLicense("Apache-2.0").
			SetGit(
				"https://gitee.com/mirrors/openssl.git",
				"https://github.com/openssl/openssl.git",
			).
			AddVersion("3.2.1", "openssl-3.2.1").
			AddVersion("3.0.13", "openssl-3.0.13").
			SetLibs("ssl", "crypto")
	})

	b.OnConfig(func(ctx *api.ConfigContext) {
		ctx.Option("shared").
			SetType(api.OptionBool).
			SetDefault(false).
			SetDescription("Build shared library")
	})

	b.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("openssl").
			SetKind(api.TargetVoid).
			SetBuildFunc(func(pctx *api.PackageContext) error {
				sharedOpt := "no-shared"
				if pctx.Bool("shared") {
					sharedOpt = "shared"
				}

				pctx.RunIn(pctx.SourceDir(), "./Configure",
					"--prefix="+pctx.InstallDir(),
					"--openssldir="+pctx.InstallDir()+"/ssl",
					sharedOpt,
					"no-tests",
					"no-apps",
				)
				pctx.RunIn(pctx.SourceDir(), "make", "-j4")
				pctx.RunIn(pctx.SourceDir(), "make", "install_sw", "install_ssldirs")
				return nil
			})
	})
}
