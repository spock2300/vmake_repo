package main

import "gitee.com/spock2300/vmake/pkg/api"

func Main(b *api.Builder) {
	b.OnPackage(func(ctx *api.PackageContext) {
		ctx.SetHomepage("https://codeplea.com/tinyexpr").
			SetDescription("Tiny expression evaluator for C").
			SetLicense("zlib").
			SetGit("https://github.com/codeplea/tinyexpr.git").
			AddVersion("1.0.0", "v1.0")
	})

	b.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("tinyexpr").
			SetKind(api.TargetStatic).
			AddFiles("tinyexpr.c").
			AddPublicIncludes(".")
	})
}
