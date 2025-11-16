package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createActionCmd = &cobra.Command{
	Use:   "create-action",
	Short: "Creates an *action*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createActionCmd).Standalone()

		sagemaker_createActionCmd.Flags().String("action-name", "", "The name of the action.")
		sagemaker_createActionCmd.Flags().String("action-type", "", "The action type.")
		sagemaker_createActionCmd.Flags().String("description", "", "The description of the action.")
		sagemaker_createActionCmd.Flags().String("metadata-properties", "", "")
		sagemaker_createActionCmd.Flags().String("properties", "", "A list of properties to add to the action.")
		sagemaker_createActionCmd.Flags().String("source", "", "The source type, ID, and URI.")
		sagemaker_createActionCmd.Flags().String("status", "", "The status of the action.")
		sagemaker_createActionCmd.Flags().String("tags", "", "A list of tags to apply to the action.")
		sagemaker_createActionCmd.MarkFlagRequired("action-name")
		sagemaker_createActionCmd.MarkFlagRequired("action-type")
		sagemaker_createActionCmd.MarkFlagRequired("source")
	})
	sagemakerCmd.AddCommand(sagemaker_createActionCmd)
}
