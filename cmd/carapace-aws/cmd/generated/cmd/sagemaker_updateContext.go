package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateContextCmd = &cobra.Command{
	Use:   "update-context",
	Short: "Updates a context.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateContextCmd).Standalone()

	sagemaker_updateContextCmd.Flags().String("context-name", "", "The name of the context to update.")
	sagemaker_updateContextCmd.Flags().String("description", "", "The new description for the context.")
	sagemaker_updateContextCmd.Flags().String("properties", "", "The new list of properties.")
	sagemaker_updateContextCmd.Flags().String("properties-to-remove", "", "A list of properties to remove.")
	sagemaker_updateContextCmd.MarkFlagRequired("context-name")
	sagemakerCmd.AddCommand(sagemaker_updateContextCmd)
}
