package main

import "gitee.com/spock2300/vmake/pkg/api"

func Main(p *api.Package) {
	p.OnPackage(func(p *api.Package) {
		p.SetHomepage("https://codeplea.com/tinyexpr").
			SetDescription("Tiny expression evaluator for C").
			SetLicense("zlib").
			SetGit("https://github.com/codeplea/tinyexpr.git").
			AddVersion("1.0.0", "v1.0")
	})

	p.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("tinyexpr").
			SetKind(api.TargetStatic).
			AddFiles("tinyexpr.c").
			AddPublicIncludes(".")
	})
}
