package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createAppCmd = &cobra.Command{
	Use:   "create-app",
	Short: "Creates a running app for the specified UserProfile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createAppCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createAppCmd).Standalone()

		sagemaker_createAppCmd.Flags().String("app-name", "", "The name of the app.")
		sagemaker_createAppCmd.Flags().String("app-type", "", "The type of app.")
		sagemaker_createAppCmd.Flags().String("domain-id", "", "The domain ID.")
		sagemaker_createAppCmd.Flags().Bool("no-recovery-mode", false, "Indicates whether the application is launched in recovery mode.")
		sagemaker_createAppCmd.Flags().Bool("recovery-mode", false, "Indicates whether the application is launched in recovery mode.")
		sagemaker_createAppCmd.Flags().String("resource-spec", "", "The instance type and the Amazon Resource Name (ARN) of the SageMaker AI image created on the instance.")
		sagemaker_createAppCmd.Flags().String("space-name", "", "The name of the space.")
		sagemaker_createAppCmd.Flags().String("tags", "", "Each tag consists of a key and an optional value.")
		sagemaker_createAppCmd.Flags().String("user-profile-name", "", "The user profile name.")
		sagemaker_createAppCmd.MarkFlagRequired("app-name")
		sagemaker_createAppCmd.MarkFlagRequired("app-type")
		sagemaker_createAppCmd.MarkFlagRequired("domain-id")
		sagemaker_createAppCmd.Flag("no-recovery-mode").Hidden = true
	})
	sagemakerCmd.AddCommand(sagemaker_createAppCmd)
}
