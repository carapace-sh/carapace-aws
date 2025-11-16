package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_updateProjectCmd = &cobra.Command{
	Use:   "update-project",
	Short: "Updates the description of an existing project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_updateProjectCmd).Standalone()

	evidently_updateProjectCmd.Flags().String("app-config-resource", "", "Use this parameter if the project will use client-side evaluation powered by AppConfig.")
	evidently_updateProjectCmd.Flags().String("description", "", "An optional description of the project.")
	evidently_updateProjectCmd.Flags().String("project", "", "The name or ARN of the project to update.")
	evidently_updateProjectCmd.MarkFlagRequired("project")
	evidentlyCmd.AddCommand(evidently_updateProjectCmd)
}
