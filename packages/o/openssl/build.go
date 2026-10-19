package main

import (
	"os"
	"path/filepath"

	"github.com/spock2300/vmake/pkg/api"
)

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

				if _, err := os.Stat(filepath.Join(p.BuildDir(), "Makefile")); err != nil {
					p.Run(filepath.Join(p.SrcDir(), "Configure"),
						"--prefix="+p.InstallDir(),
						"--openssldir="+p.InstallDir()+"/ssl",
						sharedOpt,
						"no-tests",
						"no-apps",
					)
				}
				if err := p.Make(); err != nil {
					return err
				}
				return p.Make("install_sw", "install_ssldirs")
			})
	})
}
