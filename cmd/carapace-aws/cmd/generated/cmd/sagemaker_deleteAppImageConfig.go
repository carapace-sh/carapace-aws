package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteAppImageConfigCmd = &cobra.Command{
	Use:   "delete-app-image-config",
	Short: "Deletes an AppImageConfig.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteAppImageConfigCmd).Standalone()

	sagemaker_deleteAppImageConfigCmd.Flags().String("app-image-config-name", "", "The name of the AppImageConfig to delete.")
	sagemaker_deleteAppImageConfigCmd.MarkFlagRequired("app-image-config-name")
	sagemakerCmd.AddCommand(sagemaker_deleteAppImageConfigCmd)
}
