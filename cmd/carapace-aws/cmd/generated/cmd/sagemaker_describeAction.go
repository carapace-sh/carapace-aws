package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeActionCmd = &cobra.Command{
	Use:   "describe-action",
	Short: "Describes an action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeActionCmd).Standalone()

	sagemaker_describeActionCmd.Flags().String("action-name", "", "The name of the action to describe.")
	sagemaker_describeActionCmd.MarkFlagRequired("action-name")
	sagemakerCmd.AddCommand(sagemaker_describeActionCmd)
}
