package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_getDevEnvironmentCmd = &cobra.Command{
	Use:   "get-dev-environment",
	Short: "Returns information about a Dev Environment for a source repository in a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_getDevEnvironmentCmd).Standalone()

	codecatalyst_getDevEnvironmentCmd.Flags().String("id", "", "The system-generated unique ID of the Dev Environment for which you want to view information.")
	codecatalyst_getDevEnvironmentCmd.Flags().String("project-name", "", "The name of the project in the space.")
	codecatalyst_getDevEnvironmentCmd.Flags().String("space-name", "", "The name of the space.")
	codecatalyst_getDevEnvironmentCmd.MarkFlagRequired("id")
	codecatalyst_getDevEnvironmentCmd.MarkFlagRequired("project-name")
	codecatalyst_getDevEnvironmentCmd.MarkFlagRequired("space-name")
	codecatalystCmd.AddCommand(codecatalyst_getDevEnvironmentCmd)
}
