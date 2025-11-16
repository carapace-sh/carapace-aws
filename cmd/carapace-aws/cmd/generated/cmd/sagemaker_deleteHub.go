package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteHubCmd = &cobra.Command{
	Use:   "delete-hub",
	Short: "Delete a hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteHubCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteHubCmd).Standalone()

		sagemaker_deleteHubCmd.Flags().String("hub-name", "", "The name of the hub to delete.")
		sagemaker_deleteHubCmd.MarkFlagRequired("hub-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteHubCmd)
}
