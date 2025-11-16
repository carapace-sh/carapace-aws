package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from the specified Amazon Cognito identity pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdentity_untagResourceCmd).Standalone()

		cognitoIdentity_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the identity pool.")
		cognitoIdentity_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to remove from the user pool.")
		cognitoIdentity_untagResourceCmd.MarkFlagRequired("resource-arn")
		cognitoIdentity_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	cognitoIdentityCmd.AddCommand(cognitoIdentity_untagResourceCmd)
}
