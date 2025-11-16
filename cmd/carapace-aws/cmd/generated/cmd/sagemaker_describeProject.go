package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeProjectCmd = &cobra.Command{
	Use:   "describe-project",
	Short: "Describes the details of a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeProjectCmd).Standalone()

		sagemaker_describeProjectCmd.Flags().String("project-name", "", "The name of the project to describe.")
		sagemaker_describeProjectCmd.MarkFlagRequired("project-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeProjectCmd)
}
