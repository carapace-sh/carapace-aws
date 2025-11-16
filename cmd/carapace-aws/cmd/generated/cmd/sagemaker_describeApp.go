package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeAppCmd = &cobra.Command{
	Use:   "describe-app",
	Short: "Describes the app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeAppCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeAppCmd).Standalone()

		sagemaker_describeAppCmd.Flags().String("app-name", "", "The name of the app.")
		sagemaker_describeAppCmd.Flags().String("app-type", "", "The type of app.")
		sagemaker_describeAppCmd.Flags().String("domain-id", "", "The domain ID.")
		sagemaker_describeAppCmd.Flags().String("space-name", "", "The name of the space.")
		sagemaker_describeAppCmd.Flags().String("user-profile-name", "", "The user profile name.")
		sagemaker_describeAppCmd.MarkFlagRequired("app-name")
		sagemaker_describeAppCmd.MarkFlagRequired("app-type")
		sagemaker_describeAppCmd.MarkFlagRequired("domain-id")
	})
	sagemakerCmd.AddCommand(sagemaker_describeAppCmd)
}
