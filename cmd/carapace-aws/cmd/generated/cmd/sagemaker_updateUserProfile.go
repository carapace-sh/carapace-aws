package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateUserProfileCmd = &cobra.Command{
	Use:   "update-user-profile",
	Short: "Updates a user profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateUserProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateUserProfileCmd).Standalone()

		sagemaker_updateUserProfileCmd.Flags().String("domain-id", "", "The domain ID.")
		sagemaker_updateUserProfileCmd.Flags().String("user-profile-name", "", "The user profile name.")
		sagemaker_updateUserProfileCmd.Flags().String("user-settings", "", "A collection of settings.")
		sagemaker_updateUserProfileCmd.MarkFlagRequired("domain-id")
		sagemaker_updateUserProfileCmd.MarkFlagRequired("user-profile-name")
	})
	sagemakerCmd.AddCommand(sagemaker_updateUserProfileCmd)
}
