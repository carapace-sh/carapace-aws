package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteProjectCmd = &cobra.Command{
	Use:   "delete-project",
	Short: "Delete the specified project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteProjectCmd).Standalone()

		sagemaker_deleteProjectCmd.Flags().String("project-name", "", "The name of the project to delete.")
		sagemaker_deleteProjectCmd.MarkFlagRequired("project-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteProjectCmd)
}
