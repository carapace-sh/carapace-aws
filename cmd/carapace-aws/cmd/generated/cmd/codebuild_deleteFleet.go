package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_deleteFleetCmd = &cobra.Command{
	Use:   "delete-fleet",
	Short: "Deletes a compute fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_deleteFleetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_deleteFleetCmd).Standalone()

		codebuild_deleteFleetCmd.Flags().String("arn", "", "The ARN of the compute fleet.")
		codebuild_deleteFleetCmd.MarkFlagRequired("arn")
	})
	codebuildCmd.AddCommand(codebuild_deleteFleetCmd)
}
