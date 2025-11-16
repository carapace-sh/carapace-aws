package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateSpaceCmd = &cobra.Command{
	Use:   "update-space",
	Short: "Updates the settings of a space.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateSpaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateSpaceCmd).Standalone()

		sagemaker_updateSpaceCmd.Flags().String("domain-id", "", "The ID of the associated domain.")
		sagemaker_updateSpaceCmd.Flags().String("space-display-name", "", "The name of the space that appears in the Amazon SageMaker Studio UI.")
		sagemaker_updateSpaceCmd.Flags().String("space-name", "", "The name of the space.")
		sagemaker_updateSpaceCmd.Flags().String("space-settings", "", "A collection of space settings.")
		sagemaker_updateSpaceCmd.MarkFlagRequired("domain-id")
		sagemaker_updateSpaceCmd.MarkFlagRequired("space-name")
	})
	sagemakerCmd.AddCommand(sagemaker_updateSpaceCmd)
}
