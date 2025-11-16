package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createSpaceCmd = &cobra.Command{
	Use:   "create-space",
	Short: "Creates a private space or a space used for real time collaboration in a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createSpaceCmd).Standalone()

	sagemaker_createSpaceCmd.Flags().String("domain-id", "", "The ID of the associated domain.")
	sagemaker_createSpaceCmd.Flags().String("ownership-settings", "", "A collection of ownership settings.")
	sagemaker_createSpaceCmd.Flags().String("space-display-name", "", "The name of the space that appears in the SageMaker Studio UI.")
	sagemaker_createSpaceCmd.Flags().String("space-name", "", "The name of the space.")
	sagemaker_createSpaceCmd.Flags().String("space-settings", "", "A collection of space settings.")
	sagemaker_createSpaceCmd.Flags().String("space-sharing-settings", "", "A collection of space sharing settings.")
	sagemaker_createSpaceCmd.Flags().String("tags", "", "Tags to associated with the space.")
	sagemaker_createSpaceCmd.MarkFlagRequired("domain-id")
	sagemaker_createSpaceCmd.MarkFlagRequired("space-name")
	sagemakerCmd.AddCommand(sagemaker_createSpaceCmd)
}
