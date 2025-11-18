package main

import (
	"github.com/foomo/gofana/pkg/api"
	"github.com/foomo/gofana/pkg/library/foomo/keel"
	"github.com/foomo/gofana/pkg/library/opentelmetry/http"
	"github.com/foomo/gofana/pkg/plugin"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func main() {
	tags := []string{"foomo", "gofana", "example"}
	link := dashboard.NewDashboardLinkBuilder("Application").
		IncludeVars(true).
		AsDropdown(true).
		KeepTime(true).
		Tags(tags)

	plugin.Execute(
		// Define the resources folder UID
		"example_subfolder",
		// Define folders to generate
		api.Folder{
			UID:  "example",
			Name: "Example",
			Folders: []api.Folder{
				{
					UID:     "example_subfolder",
					Name:    "Subfolder",
					Folders: nil,
				},
			},
		},
		// Add grafana resources
		keel.NewServerDashboard("squadron-sesamy-site", "site-backend").Tags(tags).Link(link),
		http.NewServerDashboard("squadron-sesamy-site", "site-backend").Tags(tags).Link(link),
	)
}
