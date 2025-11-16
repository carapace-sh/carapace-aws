package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteAssociationCmd = &cobra.Command{
	Use:   "delete-association",
	Short: "Deletes an association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteAssociationCmd).Standalone()

		sagemaker_deleteAssociationCmd.Flags().String("destination-arn", "", "The Amazon Resource Name (ARN) of the destination.")
		sagemaker_deleteAssociationCmd.Flags().String("source-arn", "", "The ARN of the source.")
		sagemaker_deleteAssociationCmd.MarkFlagRequired("destination-arn")
		sagemaker_deleteAssociationCmd.MarkFlagRequired("source-arn")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteAssociationCmd)
}
