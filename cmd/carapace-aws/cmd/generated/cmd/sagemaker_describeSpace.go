package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeSpaceCmd = &cobra.Command{
	Use:   "describe-space",
	Short: "Describes the space.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeSpaceCmd).Standalone()

	sagemaker_describeSpaceCmd.Flags().String("domain-id", "", "The ID of the associated domain.")
	sagemaker_describeSpaceCmd.Flags().String("space-name", "", "The name of the space.")
	sagemaker_describeSpaceCmd.MarkFlagRequired("domain-id")
	sagemaker_describeSpaceCmd.MarkFlagRequired("space-name")
	sagemakerCmd.AddCommand(sagemaker_describeSpaceCmd)
}
