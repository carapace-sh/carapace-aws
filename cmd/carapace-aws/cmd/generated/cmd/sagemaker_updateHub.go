package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateHubCmd = &cobra.Command{
	Use:   "update-hub",
	Short: "Update a hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateHubCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateHubCmd).Standalone()

		sagemaker_updateHubCmd.Flags().String("hub-description", "", "A description of the updated hub.")
		sagemaker_updateHubCmd.Flags().String("hub-display-name", "", "The display name of the hub.")
		sagemaker_updateHubCmd.Flags().String("hub-name", "", "The name of the hub to update.")
		sagemaker_updateHubCmd.Flags().String("hub-search-keywords", "", "The searchable keywords for the hub.")
		sagemaker_updateHubCmd.MarkFlagRequired("hub-name")
	})
	sagemakerCmd.AddCommand(sagemaker_updateHubCmd)
}
