package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createUserProfileCmd = &cobra.Command{
	Use:   "create-user-profile",
	Short: "Creates a user profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createUserProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createUserProfileCmd).Standalone()

		sagemaker_createUserProfileCmd.Flags().String("domain-id", "", "The ID of the associated Domain.")
		sagemaker_createUserProfileCmd.Flags().String("single-sign-on-user-identifier", "", "A specifier for the type of value specified in SingleSignOnUserValue.")
		sagemaker_createUserProfileCmd.Flags().String("single-sign-on-user-value", "", "The username of the associated Amazon Web Services Single Sign-On User for this UserProfile.")
		sagemaker_createUserProfileCmd.Flags().String("tags", "", "Each tag consists of a key and an optional value.")
		sagemaker_createUserProfileCmd.Flags().String("user-profile-name", "", "A name for the UserProfile.")
		sagemaker_createUserProfileCmd.Flags().String("user-settings", "", "A collection of settings.")
		sagemaker_createUserProfileCmd.MarkFlagRequired("domain-id")
		sagemaker_createUserProfileCmd.MarkFlagRequired("user-profile-name")
	})
	sagemakerCmd.AddCommand(sagemaker_createUserProfileCmd)
}
