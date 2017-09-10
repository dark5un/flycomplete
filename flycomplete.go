package main

import "github.com/posener/complete"

func main() {
	abort_build := complete.Command{
		Flags: complete.Flags{
			"--job=":   complete.PredictAnything,
			"-j":       complete.PredictAnything,
			"--build=": complete.PredictAnything,
			"-b":       complete.PredictAnything,
		},
	}
	builds := complete.Command{
		Flags: complete.Flags{
			"--count=": complete.PredictAnything,
			"-c":       complete.PredictAnything,
			"--job=":   complete.PredictAnything,
			"-j":       complete.PredictAnything,
		},
	}
	check_resource := complete.Command{
		Flags: complete.Flags{
			"--resource=": complete.PredictAnything,
			"-r":          complete.PredictAnything,
			"--from=":     complete.PredictAnything,
			"-f":          complete.PredictAnything,
		},
	}
	checklist := complete.Command{
		Flags: complete.Flags{
			"--pipeline": complete.PredictAnything,
			"-p":         complete.PredictAnything,
		},
	}
	containers := complete.Command{}
	destroy_pipeline := complete.Command{
		Flags: complete.Flags{
			"--pipeline=":       complete.PredictAnything,
			"-p":                complete.PredictAnything,
			"--non-interactive": complete.PredictAnything,
			"-n":                complete.PredictAnything,
		},
	}
	destroy_team := complete.Command{
		Flags: complete.Flags{
			"--team-name=": complete.PredictAnything,
			"-n":           complete.PredictAnything,
		},
	}
	execute := complete.Command{
		Flags: complete.Flags{
			"--config=":         complete.PredictAnything,
			"-c":                complete.PredictAnything,
			"--privileged":      complete.PredictAnything,
			"-p":                complete.PredictAnything,
			"--exclude-ignored": complete.PredictAnything,
			"-x":                complete.PredictAnything,
			"--input=":          complete.PredictAnything,
			"-i":                complete.PredictAnything,
			"--inputs-from=":    complete.PredictAnything,
			"-j":                complete.PredictAnything,
			"--output=":         complete.PredictAnything,
			"-o":                complete.PredictAnything,
			"--tag=":            complete.PredictAnything,
		},
	}
	expose_pipeline := complete.Command{
		Flags: complete.Flags{
			"--pipeline=": complete.PredictAnything,
			"-p":          complete.PredictAnything,
		},
	}
	get_pipeline := complete.Command{
		Flags: complete.Flags{
			"--pipeline=": complete.PredictAnything,
			"-p":          complete.PredictAnything,
			"--json":      complete.PredictAnything,
			"-j":          complete.PredictAnything,
		},
	}
	hide_pipeline := complete.Command{
		Flags: complete.Flags{
			"--pipeline=": complete.PredictAnything,
			"-p":          complete.PredictAnything,
		},
	}
	hijack := complete.Command{
		Flags: complete.Flags{
			"--job=":     complete.PredictAnything,
			"-j":         complete.PredictAnything,
			"--check=":   complete.PredictAnything,
			"-c":         complete.PredictAnything,
			"--build=":   complete.PredictAnything,
			"-b":         complete.PredictAnything,
			"--step=":    complete.PredictAnything,
			"-s":         complete.PredictAnything,
			"--attempt=": complete.PredictAnything,
			"-a":         complete.PredictAnything,
		},
	}
	login := complete.Command{
		Flags: complete.Flags{
			"--concourse-url=": complete.PredictAnything,
			"-c":               complete.PredictAnything,
			"--insecure":       complete.PredictAnything,
			"-k":               complete.PredictAnything,
			"--username=":      complete.PredictAnything,
			"-u":               complete.PredictAnything,
			"--password=":      complete.PredictAnything,
			"-p":               complete.PredictAnything,
			"--team-name=":     complete.PredictAnything,
			"-n":               complete.PredictAnything,
			"--ca-cert=":       complete.PredictAnything,
		},
	}
	logout := complete.Command{
		Flags: complete.Flags{
			"--all": complete.PredictAnything,
			"-a":    complete.PredictAnything,
		},
	}
	pause_job := complete.Command{
		Flags: complete.Flags{
			"--job=": complete.PredictAnything,
			"-j":     complete.PredictAnything,
		},
	}
	pause_pipeline := complete.Command{
		Flags: complete.Flags{
			"--pipeline=": complete.PredictAnything,
			"-p":          complete.PredictAnything,
		},
	}
	pause_resource := complete.Command{
		Flags: complete.Flags{
			"--resource=": complete.PredictAnything,
			"-r":          complete.PredictAnything,
		},
	}
	pipelines := complete.Command{
		Flags: complete.Flags{
			"--all": complete.PredictAnything,
			"-a":    complete.PredictAnything,
		},
	}
	prune_worker := complete.Command{
		Flags: complete.Flags{
			"--worker=": complete.PredictAnything,
			"-w":        complete.PredictAnything,
		},
	}
	rename_pipeline := complete.Command{
		Flags: complete.Flags{
			"--old-name=": complete.PredictAnything,
			"-o":          complete.PredictAnything,
			"--new-name=": complete.PredictAnything,
			"-n":          complete.PredictAnything,
		},
	}
	set_pipeline := complete.Command{
		Flags: complete.Flags{
			"--non-interactive": complete.PredictAnything,
			"-n":                complete.PredictAnything,
			"--pipeline=":       complete.PredictAnything,
			"-p":                complete.PredictAnything,
			"--config=":         complete.PredictAnything,
			"-c":                complete.PredictAnything,
			"--var=":            complete.PredictAnything,
			"-v":                complete.PredictAnything,
			"--yaml-var=":       complete.PredictAnything,
			"-y":                complete.PredictAnything,
			"--load-vars-from=": complete.PredictAnything,
			"-l":                complete.PredictAnything,
		},
	}
	set_team := complete.Command{
		Flags: complete.Flags{
			"--team-name=": complete.PredictAnything,
			"-n":           complete.PredictAnything,
			"--no-really-i-dont-want-any-auth": complete.PredictAnything,
			"--basic-auth-username=":           complete.PredictAnything,
			"--basic-auth-password=":           complete.PredictAnything,
			"--generic-oauth-display-name=":    complete.PredictAnything,
			"--generic-oauth-client-id=":       complete.PredictAnything,
			"--generic-oauth-client-secret=":   complete.PredictAnything,
			"--generic-oauth-auth-url=":        complete.PredictAnything,
			"--generic-oauth-auth-url-param=":  complete.PredictAnything,
			"--generic-oauth-scope=":           complete.PredictAnything,
			"--generic-oauth-token-url=":       complete.PredictAnything,
			"--github-auth-client-id=":         complete.PredictAnything,
			"--github-auth-client-secret=":     complete.PredictAnything,
			"--github-auth-organization=":      complete.PredictAnything,
			"--github-auth-team=":              complete.PredictAnything,
			"--github-auth-user=":              complete.PredictAnything,
			"--github-auth-auth-url=":          complete.PredictAnything,
			"--github-auth-token-url=":         complete.PredictAnything,
			"--github-auth-api-url=":           complete.PredictAnything,
			"--uaa-auth-client-id=":            complete.PredictAnything,
			"--uaa-auth-client-secret=":        complete.PredictAnything,
			"--uaa-auth-auth-url=":             complete.PredictAnything,
			"--uaa-auth-token-url=":            complete.PredictAnything,
			"--uaa-auth-cf-space=":             complete.PredictAnything,
			"--uaa-auth-cf-url=":               complete.PredictAnything,
			"--uaa-auth-cf-ca-cert=":           complete.PredictAnything,
		},
	}
	sync := complete.Command{}
	targets := complete.Command{}
	teams := complete.Command{}
	trigger_job := complete.Command{
		Flags: complete.Flags{
			"--job=":  complete.PredictAnything,
			"-j":      complete.PredictAnything,
			"--watch": complete.PredictAnything,
			"-w":      complete.PredictAnything,
		},
	}
	unpause_job := complete.Command{
		Flags: complete.Flags{
			"--job=": complete.PredictAnything,
			"-j":     complete.PredictAnything,
		},
	}
	unpause_pipeline := complete.Command{
		Flags: complete.Flags{
			"--pipeline=": complete.PredictAnything,
			"-p":          complete.PredictAnything,
		},
	}
	unpause_resource := complete.Command{
		Flags: complete.Flags{
			"--resource=": complete.PredictAnything,
			"-r":          complete.PredictAnything,
		},
	}
	validate_pipeline := complete.Command{
		Flags: complete.Flags{
			"--config=": complete.PredictAnything,
			"-c":        complete.PredictAnything,
			"--strict":  complete.PredictAnything,
			"-s":        complete.PredictAnything,
		},
	}
	volumes := complete.Command{
		Flags: complete.Flags{
			"--details": complete.PredictAnything,
			"-d":        complete.PredictAnything,
		},
	}
	watch := complete.Command{
		Flags: complete.Flags{
			"--job=":   complete.PredictAnything,
			"-j":       complete.PredictAnything,
			"--build=": complete.PredictAnything,
			"-b":       complete.PredictAnything,
		},
	}
	workers := complete.Command{
		Flags: complete.Flags{
			"--details": complete.PredictAnything,
			"-d":        complete.PredictAnything,
		},
	}
	fly := complete.Command{
		Sub: complete.Commands{
			"abort-build":       abort_build,
			"builds":            builds,
			"check-resource":    check_resource,
			"checklist":         checklist,
			"containers":        containers,
			"destroy-pipeline":  destroy_pipeline,
			"destroy-team":      destroy_team,
			"execute":           execute,
			"expose-pipeline":   expose_pipeline,
			"get-pipeline":      get_pipeline,
			"hide-pipeline":     hide_pipeline,
			"hijack":            hijack,
			"login":             login,
			"logout":            logout,
			"pause-job":         pause_job,
			"pause-pipeline":    pause_pipeline,
			"pause-resource":    pause_resource,
			"pipelines":         pipelines,
			"prune-worker":      prune_worker,
			"rename-pipeline":   rename_pipeline,
			"set-pipeline":      set_pipeline,
			"set-team":          set_team,
			"sync":              sync,
			"targets":           targets,
			"teams":             teams,
			"trigger-job":       trigger_job,
			"unpause-job":       unpause_job,
			"unpause-pipeline":  unpause_pipeline,
			"unpause-resource":  unpause_resource,
			"validate-pipeline": validate_pipeline,
			"volumes":           volumes,
			"watch":             watch,
			"workers":           workers,
		},
		GlobalFlags: complete.Flags{
			"--help":                complete.PredictNothing,
			"--target=":             complete.PredictAnything,
			"--version":             complete.PredictNothing,
			"--print-table-headers": complete.PredictAnything,
		}}
	complete.New("fly", fly).Run()
}
