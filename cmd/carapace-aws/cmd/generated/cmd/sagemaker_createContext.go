package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createContextCmd = &cobra.Command{
	Use:   "create-context",
	Short: "Creates a *context*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createContextCmd).Standalone()

	sagemaker_createContextCmd.Flags().String("context-name", "", "The name of the context.")
	sagemaker_createContextCmd.Flags().String("context-type", "", "The context type.")
	sagemaker_createContextCmd.Flags().String("description", "", "The description of the context.")
	sagemaker_createContextCmd.Flags().String("properties", "", "A list of properties to add to the context.")
	sagemaker_createContextCmd.Flags().String("source", "", "The source type, ID, and URI.")
	sagemaker_createContextCmd.Flags().String("tags", "", "A list of tags to apply to the context.")
	sagemaker_createContextCmd.MarkFlagRequired("context-name")
	sagemaker_createContextCmd.MarkFlagRequired("context-type")
	sagemaker_createContextCmd.MarkFlagRequired("source")
	sagemakerCmd.AddCommand(sagemaker_createContextCmd)
}
