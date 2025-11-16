package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeHubContentCmd = &cobra.Command{
	Use:   "describe-hub-content",
	Short: "Describe the content of a hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeHubContentCmd).Standalone()

	sagemaker_describeHubContentCmd.Flags().String("hub-content-name", "", "The name of the content to describe.")
	sagemaker_describeHubContentCmd.Flags().String("hub-content-type", "", "The type of content in the hub.")
	sagemaker_describeHubContentCmd.Flags().String("hub-content-version", "", "The version of the content to describe.")
	sagemaker_describeHubContentCmd.Flags().String("hub-name", "", "The name of the hub that contains the content to describe.")
	sagemaker_describeHubContentCmd.MarkFlagRequired("hub-content-name")
	sagemaker_describeHubContentCmd.MarkFlagRequired("hub-content-type")
	sagemaker_describeHubContentCmd.MarkFlagRequired("hub-name")
	sagemakerCmd.AddCommand(sagemaker_describeHubContentCmd)
}
