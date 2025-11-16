package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeUserProfileCmd = &cobra.Command{
	Use:   "describe-user-profile",
	Short: "Describes a user profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeUserProfileCmd).Standalone()

	sagemaker_describeUserProfileCmd.Flags().String("domain-id", "", "The domain ID.")
	sagemaker_describeUserProfileCmd.Flags().String("user-profile-name", "", "The user profile name.")
	sagemaker_describeUserProfileCmd.MarkFlagRequired("domain-id")
	sagemaker_describeUserProfileCmd.MarkFlagRequired("user-profile-name")
	sagemakerCmd.AddCommand(sagemaker_describeUserProfileCmd)
}
