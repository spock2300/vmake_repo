package main

import "gitee.com/spock2300/vmake/pkg/api"

func Package(p *api.Package) {
	p.SetHomepage("https://www.openssl.org/").
		SetDescription("TLS/SSL and crypto library").
		SetLicense("Apache-2.0").
		SetGit("https://github.com/openssl/openssl.git").
		SetLibs("ssl", "crypto")

	p.AddVersion("3.2.1", "openssl-3.2.1")
	p.AddVersion("3.0.13", "openssl-3.0.13")

	p.Option("shared").
		SetType(api.OptionBool).
		SetDefault(false).
		SetDescription("Build shared library")

	p.Build(func(ctx *api.PackageContext) {
		sharedOpt := "no-shared"
		if ctx.Bool("shared") {
			sharedOpt = "shared"
		}

		ctx.RunIn(ctx.SourceDir(), "./Configure",
			"--prefix="+ctx.InstallDir(),
			"--openssldir="+ctx.InstallDir()+"/ssl",
			sharedOpt,
			"no-tests",
			"no-apps",
		)
		ctx.RunIn(ctx.SourceDir(), "make", "-j4")
		ctx.RunIn(ctx.SourceDir(), "make", "install_sw", "install_ssldirs")
	})
}
