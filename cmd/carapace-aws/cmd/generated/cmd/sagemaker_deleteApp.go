package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteAppCmd = &cobra.Command{
	Use:   "delete-app",
	Short: "Used to stop and delete an app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteAppCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteAppCmd).Standalone()

		sagemaker_deleteAppCmd.Flags().String("app-name", "", "The name of the app.")
		sagemaker_deleteAppCmd.Flags().String("app-type", "", "The type of app.")
		sagemaker_deleteAppCmd.Flags().String("domain-id", "", "The domain ID.")
		sagemaker_deleteAppCmd.Flags().String("space-name", "", "The name of the space.")
		sagemaker_deleteAppCmd.Flags().String("user-profile-name", "", "The user profile name.")
		sagemaker_deleteAppCmd.MarkFlagRequired("app-name")
		sagemaker_deleteAppCmd.MarkFlagRequired("app-type")
		sagemaker_deleteAppCmd.MarkFlagRequired("domain-id")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteAppCmd)
}
