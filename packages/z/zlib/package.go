package main

import "gitee.com/spock2300/vmake/pkg/api"

func Package(p *api.Package) {
	p.SetHomepage("http://www.zlib.net").
		SetDescription("Compression library").
		SetLicense("zlib").
		SetGit("https://github.com/madler/zlib.git")

	p.AddVersion("1.3.1", "v1.3.1")
	p.AddVersion("1.2.13", "v1.2.13")

	p.Option("shared").
		SetType(api.OptionBool).
		SetDefault(false).
		SetDescription("Build shared library")

	p.Build(func(ctx *api.PackageContext) {
		ctx.CMakeConfigure(
			"-DBUILD_SHARED_LIBS=" + ctx.BoolStr("shared"),
		)
		ctx.CMakeBuild("-j4")
		ctx.CMakeInstall()
	})
}
