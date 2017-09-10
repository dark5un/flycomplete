package main

import "github.com/posener/complete"

func main() {
	fly := complete.Command{
		Sub: complete.Commands{
			"abort-build":       complete.Command{Flags: complete.Flags{"--job=": complete.PredictAnything, "--build=": complete.PredictAnything}},
			"builds":            complete.Command{Flags: complete.Flags{"--count=": complete.PredictAnything, "--job=": complete.PredictAnything}},
			"check-resource":    complete.Command{Flags: complete.Flags{"--resource=": complete.PredictAnything, "--from=": complete.PredictAnything}},
			"checklist":         complete.Command{Flags: complete.Flags{"--pipeline": complete.PredictAnything}},
			"containers":        complete.Command{},
			"destroy-pipeline":  complete.Command{Flags: complete.Flags{"--pipeline=": complete.PredictAnything, "--non-interactive": complete.PredictAnything}},
			"destroy-team":      complete.Command{Flags: complete.Flags{"--team-name=": complete.PredictAnything}},
			"execute":           complete.Command{Flags: complete.Flags{"--config=": complete.PredictAnything, "--privileged": complete.PredictAnything, "--exclude-ignored": complete.PredictAnything, "--input=": complete.PredictAnything, "--inputs-from=": complete.PredictAnything, "--output=": complete.PredictAnything, "--tag=": complete.PredictAnything}},
			"expose-pipeline":   complete.Command{Flags: complete.Flags{"--pipeline=": complete.PredictAnything}},
			"get-pipeline":      complete.Command{Flags: complete.Flags{"--pipeline=": complete.PredictAnything, "--json": complete.PredictAnything}},
			"hide-pipeline":     complete.Command{Flags: complete.Flags{"--pipeline=": complete.PredictAnything}},
			"hijack":            complete.Command{Flags: complete.Flags{"--job=": complete.PredictAnything, "--check=": complete.PredictAnything, "--build=": complete.PredictAnything, "--step=": complete.PredictAnything, "--attempt=": complete.PredictAnything}},
			"login":             complete.Command{Flags: complete.Flags{"--concourse-url=": complete.PredictAnything, "--insecure": complete.PredictAnything, "--username=": complete.PredictAnything, "--password=": complete.PredictAnything, "--team-name=": complete.PredictAnything, "--ca-cert=": complete.PredictAnything}},
			"logout":            complete.Command{Flags: complete.Flags{"--all": complete.PredictAnything}},
			"pause-job":         complete.Command{Flags: complete.Flags{"--job=": complete.PredictAnything}},
			"pause-pipeline":    complete.Command{Flags: complete.Flags{"--pipeline=": complete.PredictAnything}},
			"pause-resource":    complete.Command{Flags: complete.Flags{"--resource=": complete.PredictAnything}},
			"pipelines":         complete.Command{Flags: complete.Flags{"--all": complete.PredictAnything}},
			"prune-worker":      complete.Command{Flags: complete.Flags{"--worker=": complete.PredictAnything}},
			"rename-pipeline":   complete.Command{Flags: complete.Flags{"--old-name=": complete.PredictAnything, "--new-name=": complete.PredictAnything}},
			"set-pipeline":      complete.Command{Flags: complete.Flags{"--non-interactive": complete.PredictAnything, "--pipeline=": complete.PredictAnything, "--config=": complete.PredictAnything, "--var=": complete.PredictAnything, "--yaml-var=": complete.PredictAnything, "--load-vars-from=": complete.PredictAnything}},
			"set-team":          complete.Command{Flags: complete.Flags{"--team-name=": complete.PredictAnything, "--no-really-i-dont-want-any-auth": complete.PredictAnything, "--basic-auth-username=": complete.PredictAnything, "--basic-auth-password=": complete.PredictAnything, "--generic-oauth-display-name=": complete.PredictAnything, "--generic-oauth-client-id=": complete.PredictAnything, "--generic-oauth-client-secret=": complete.PredictAnything, "--generic-oauth-auth-url=": complete.PredictAnything, "--generic-oauth-auth-url-param=": complete.PredictAnything, "--generic-oauth-scope=": complete.PredictAnything, "--generic-oauth-token-url=": complete.PredictAnything, "--github-auth-client-id=": complete.PredictAnything, "--github-auth-client-secret=": complete.PredictAnything, "--github-auth-organization=": complete.PredictAnything, "--github-auth-team=": complete.PredictAnything, "--github-auth-user=": complete.PredictAnything, "--github-auth-auth-url=": complete.PredictAnything, "--github-auth-token-url=": complete.PredictAnything, "--github-auth-api-url=": complete.PredictAnything, "--uaa-auth-client-id=": complete.PredictAnything, "--uaa-auth-client-secret=": complete.PredictAnything, "--uaa-auth-auth-url=": complete.PredictAnything, "--uaa-auth-token-url=": complete.PredictAnything, "--uaa-auth-cf-space=": complete.PredictAnything, "--uaa-auth-cf-url=": complete.PredictAnything, "--uaa-auth-cf-ca-cert=": complete.PredictAnything}},
			"sync":              complete.Command{},
			"targets":           complete.Command{},
			"teams":             complete.Command{},
			"trigger-job":       complete.Command{Flags: complete.Flags{"--job=": complete.PredictAnything, "--watch": complete.PredictAnything}},
			"unpause-job":       complete.Command{Flags: complete.Flags{"--job=": complete.PredictAnything}},
			"unpause-pipeline":  complete.Command{Flags: complete.Flags{"--pipeline=": complete.PredictAnything}},
			"unpause-resource":  complete.Command{Flags: complete.Flags{"--resource=": complete.PredictAnything}},
			"validate-pipeline": complete.Command{Flags: complete.Flags{"--config=": complete.PredictAnything, "--strict": complete.PredictAnything}},
			"volumes":           complete.Command{Flags: complete.Flags{"--details": complete.PredictAnything}},
			"watch":             complete.Command{Flags: complete.Flags{"--job=": complete.PredictAnything, "--build=": complete.PredictAnything}},
			"workers":           complete.Command{Flags: complete.Flags{"--details": complete.PredictAnything}},
		},
		GlobalFlags: complete.Flags{
			"--help":                complete.PredictNothing,
			"--target=":             complete.PredictAnything,
			"--version":             complete.PredictNothing,
			"--print-table-headers": complete.PredictAnything,
		}}
	complete.New("fly", fly).Run()
}
