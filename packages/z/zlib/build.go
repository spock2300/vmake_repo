package main

import "gitee.com/spock2300/vmake/pkg/api"

func Main(b *api.Builder) {
	b.OnPackage(func(ctx *api.PackageContext) {
		ctx.SetHomepage("http://www.zlib.net").
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

	b.OnConfig(func(ctx *api.ConfigContext) {
		ctx.Option("shared").
			SetType(api.OptionBool).
			SetDefault(false).
			SetDescription("Build shared library")
	})

	b.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("zlib").
			SetKind(api.TargetVoid).
			SetBuildFunc(func(pctx *api.PackageContext) error {
				pctx.CMakeConfigure(
					"-DBUILD_SHARED_LIBS=" + pctx.BoolStr("shared"),
				)
				pctx.CMakeBuild("-j4")
				pctx.CMakeInstall()
				return nil
			})
	})
}
