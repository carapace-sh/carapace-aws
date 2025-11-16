package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createHubCmd = &cobra.Command{
	Use:   "create-hub",
	Short: "Create a hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createHubCmd).Standalone()

	sagemaker_createHubCmd.Flags().String("hub-description", "", "A description of the hub.")
	sagemaker_createHubCmd.Flags().String("hub-display-name", "", "The display name of the hub.")
	sagemaker_createHubCmd.Flags().String("hub-name", "", "The name of the hub to create.")
	sagemaker_createHubCmd.Flags().String("hub-search-keywords", "", "The searchable keywords for the hub.")
	sagemaker_createHubCmd.Flags().String("s3-storage-config", "", "The Amazon S3 storage configuration for the hub.")
	sagemaker_createHubCmd.Flags().String("tags", "", "Any tags to associate with the hub.")
	sagemaker_createHubCmd.MarkFlagRequired("hub-description")
	sagemaker_createHubCmd.MarkFlagRequired("hub-name")
	sagemakerCmd.AddCommand(sagemaker_createHubCmd)
}
