package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_deleteDevEnvironmentCmd = &cobra.Command{
	Use:   "delete-dev-environment",
	Short: "Deletes a Dev Environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_deleteDevEnvironmentCmd).Standalone()

	codecatalyst_deleteDevEnvironmentCmd.Flags().String("id", "", "The system-generated unique ID of the Dev Environment you want to delete.")
	codecatalyst_deleteDevEnvironmentCmd.Flags().String("project-name", "", "The name of the project in the space.")
	codecatalyst_deleteDevEnvironmentCmd.Flags().String("space-name", "", "The name of the space.")
	codecatalyst_deleteDevEnvironmentCmd.MarkFlagRequired("id")
	codecatalyst_deleteDevEnvironmentCmd.MarkFlagRequired("project-name")
	codecatalyst_deleteDevEnvironmentCmd.MarkFlagRequired("space-name")
	codecatalystCmd.AddCommand(codecatalyst_deleteDevEnvironmentCmd)
}
