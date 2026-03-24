package main

import "gitee.com/spock2300/vmake/pkg/api"

func Main(p *api.Package) {
	p.OnPackage(func(p *api.Package) {
		p.SetHomepage("https://codeplea.com/tinyexpr").
			SetDescription("Tiny expression evaluator for C").
			SetLicense("zlib").
			SetGit("https://github.com/codeplea/tinyexpr.git").
			SetLibs("tinyexpr", "m").
			AddVersion("1.0.0", "9907207e5def0fabdb60c443517b0d9e9d521393")
	})

	p.OnBuild(func(ctx *api.BuildContext) {
		ctx.Target("tinyexpr").
			SetKind(api.TargetStatic).
			AddFiles("tinyexpr.c").
			AddPublicIncludes(".")
	})
}
