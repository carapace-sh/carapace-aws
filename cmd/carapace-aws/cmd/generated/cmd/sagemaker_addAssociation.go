package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_addAssociationCmd = &cobra.Command{
	Use:   "add-association",
	Short: "Creates an *association* between the source and the destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_addAssociationCmd).Standalone()

	sagemaker_addAssociationCmd.Flags().String("association-type", "", "The type of association.")
	sagemaker_addAssociationCmd.Flags().String("destination-arn", "", "The Amazon Resource Name (ARN) of the destination.")
	sagemaker_addAssociationCmd.Flags().String("source-arn", "", "The ARN of the source.")
	sagemaker_addAssociationCmd.MarkFlagRequired("destination-arn")
	sagemaker_addAssociationCmd.MarkFlagRequired("source-arn")
	sagemakerCmd.AddCommand(sagemaker_addAssociationCmd)
}
