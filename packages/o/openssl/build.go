package main

import "gitee.com/spock2300/vmake/pkg/api"

func Main(p *api.Package) {
	p.OnPackage(func(p *api.Package) {
		p.SetHomepage("https://www.openssl.org/").
			SetDescription("TLS/SSL and crypto library").
			SetLicense("Apache-2.0").
			SetGit(
				"https://gitee.com/mirrors/openssl.git",
				"https://github.com/openssl/openssl.git",
			).
			AddVersion("3.2.1", "openssl-3.2.1").
			AddVersion("3.0.13", "openssl-3.0.13")
	})

	p.OnConfig(func(ctx *api.ConfigContext) {
		ctx.Option("shared").
			SetType(api.OptionBool).
			SetDefault(false).
			SetDescription("Build shared library")
	})

	p.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("openssl").
			SetKind(api.TargetVoid).
			AddProvidedLibs("ssl", "crypto").
			SetBuildFunc(func(p *api.Package) error {
				sharedOpt := "no-shared"
				if p.Bool("shared") {
					sharedOpt = "shared"
				}

				p.RunIn(p.SourceDir(), "./Configure",
					"--prefix="+p.InstallDir(),
					"--openssldir="+p.InstallDir()+"/ssl",
					sharedOpt,
					"no-tests",
					"no-apps",
				)
				p.RunIn(p.SourceDir(), "make", "-j4")
				p.RunIn(p.SourceDir(), "make", "install_sw", "install_ssldirs")
				return nil
			})
	})
}
