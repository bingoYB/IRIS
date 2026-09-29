package js

import "github.com/versenilvis/iris/spec"

func pnpmScriptGenerator(_ []string, _ string, _ string) []spec.Suggestion {
	scripts, found := packageScriptSuggestions()
	if !found {
		return nil
	}

	registered := spec.Registry["pnpm"]
	if registered == nil {
		return scripts
	}

	builtins := make(map[string]bool, len(registered.Subcommands))
	for _, sub := range registered.Subcommands {
		builtins[sub.Name] = true
	}

	var results []spec.Suggestion
	for _, script := range scripts {
		if !builtins[script.Cmd] {
			results = append(results, script)
		}
	}
	return results
}

func init() {
	spec.Register(&spec.Spec{
		Name:        "pnpm",
		Description: "fast node packages",
		Generator:   pnpmScriptGenerator,
		Subcommands: []spec.Subcommand{
			{Name: "install", Description: "install packages", Options: []spec.Option{
				{Name: "--frozen-lockfile", Description: "no lockfile update"},
				{Name: "--prod", Description: "production only"},
			}},
			{Name: "add", Description: "add package", Options: []spec.Option{
				{Name: "-D", Description: "save as devDependency"},
				{Name: "-g", Description: "install globally"},
				{Name: "-E", Description: "exact version"},
			}},
			{Name: "remove", Description: "remove package"},
			{Name: "run", Description: "run script", Generator: NpmScriptGenerator},
			{Name: "build", Description: "build project"},
			{Name: "dev", Description: "start dev server"},
			{Name: "test", Description: "run tests"},
			{Name: "update", Description: "update packages", Options: []spec.Option{
				{Name: "--latest", Description: "update to latest"},
				{Name: "-i", Description: "interactive"},
			}},
			{Name: "store", Description: "manage store", Subcommands: []spec.Subcommand{
				{Name: "prune", Description: "remove unreferenced"},
				{Name: "status", Description: "check store"},
				{Name: "path", Description: "print store path"},
			}},
			{Name: "exec", Description: "exec shell command"},
			{Name: "dlx", Description: "fetch and run package"},
			{Name: "ls", Description: "list packages"},
			{Name: "audit", Description: "security audit"},
			{Name: "patch", Description: "patch package"},
			{Name: "init", Description: "create package.json"},
			{Name: "publish", Description: "publish package"},
		},
		Options: []spec.Option{
			{Name: "--filter", Description: "filter by package"},
			{Name: "-r", Description: "recursive"},
			{Name: "--workspace-root", Description: "run in workspace root"},
		},
	})
}
