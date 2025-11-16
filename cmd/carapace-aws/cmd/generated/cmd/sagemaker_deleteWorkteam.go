package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteWorkteamCmd = &cobra.Command{
	Use:   "delete-workteam",
	Short: "Deletes an existing work team.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteWorkteamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteWorkteamCmd).Standalone()

		sagemaker_deleteWorkteamCmd.Flags().String("workteam-name", "", "The name of the work team to delete.")
		sagemaker_deleteWorkteamCmd.MarkFlagRequired("workteam-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteWorkteamCmd)
}
