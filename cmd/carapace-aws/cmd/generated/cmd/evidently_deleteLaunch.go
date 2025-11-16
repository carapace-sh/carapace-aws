package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_deleteLaunchCmd = &cobra.Command{
	Use:   "delete-launch",
	Short: "Deletes an Evidently launch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_deleteLaunchCmd).Standalone()

	evidently_deleteLaunchCmd.Flags().String("launch", "", "The name of the launch to delete.")
	evidently_deleteLaunchCmd.Flags().String("project", "", "The name or ARN of the project that contains the launch to delete.")
	evidently_deleteLaunchCmd.MarkFlagRequired("launch")
	evidently_deleteLaunchCmd.MarkFlagRequired("project")
	evidentlyCmd.AddCommand(evidently_deleteLaunchCmd)
}
