package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteUserProfileCmd = &cobra.Command{
	Use:   "delete-user-profile",
	Short: "Deletes a user profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteUserProfileCmd).Standalone()

	sagemaker_deleteUserProfileCmd.Flags().String("domain-id", "", "The domain ID.")
	sagemaker_deleteUserProfileCmd.Flags().String("user-profile-name", "", "The user profile name.")
	sagemaker_deleteUserProfileCmd.MarkFlagRequired("domain-id")
	sagemaker_deleteUserProfileCmd.MarkFlagRequired("user-profile-name")
	sagemakerCmd.AddCommand(sagemaker_deleteUserProfileCmd)
}
