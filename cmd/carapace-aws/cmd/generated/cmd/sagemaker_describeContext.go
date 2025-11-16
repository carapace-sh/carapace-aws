package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeContextCmd = &cobra.Command{
	Use:   "describe-context",
	Short: "Describes a context.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeContextCmd).Standalone()

	sagemaker_describeContextCmd.Flags().String("context-name", "", "The name of the context to describe.")
	sagemaker_describeContextCmd.MarkFlagRequired("context-name")
	sagemakerCmd.AddCommand(sagemaker_describeContextCmd)
}
