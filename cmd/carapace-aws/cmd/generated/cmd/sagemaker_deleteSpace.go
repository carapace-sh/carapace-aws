package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteSpaceCmd = &cobra.Command{
	Use:   "delete-space",
	Short: "Used to delete a space.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteSpaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteSpaceCmd).Standalone()

		sagemaker_deleteSpaceCmd.Flags().String("domain-id", "", "The ID of the associated domain.")
		sagemaker_deleteSpaceCmd.Flags().String("space-name", "", "The name of the space.")
		sagemaker_deleteSpaceCmd.MarkFlagRequired("domain-id")
		sagemaker_deleteSpaceCmd.MarkFlagRequired("space-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteSpaceCmd)
}
