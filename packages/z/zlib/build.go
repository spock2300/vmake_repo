package main

import "gitee.com/spock2300/vmake/pkg/api"

func Main(p *api.Package) {
	p.OnPackage(func(p *api.Package) {
		p.SetHomepage("http://www.zlib.net").
			SetDescription("Compression library").
			SetLicense("zlib").
			SetGit(
				"https://gitee.com/mirrors/zlib.git",
				"https://github.com/madler/zlib.git",
			).
			SetLibs("z").
			AddVersion("1.3.1", "v1.3.1").
			AddVersion("1.2.13", "v1.2.13")
	})

	p.OnConfig(func(ctx *api.ConfigContext) {
		ctx.Option("shared").
			SetType(api.OptionBool).
			SetDefault(false).
			SetDescription("Build shared library")
	})

	p.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("zlib").
			SetKind(api.TargetVoid).
			SetBuildFunc(func(p *api.Package) error {
				p.CMakeConfigure(
					"-DBUILD_SHARED_LIBS=" + p.BoolStr("shared"),
				)
				p.CMakeBuild("-j4")
				p.CMakeInstall()
				return nil
			})
	})
}
