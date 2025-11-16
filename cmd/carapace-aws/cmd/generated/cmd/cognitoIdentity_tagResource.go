package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns a set of tags to the specified Amazon Cognito identity pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_tagResourceCmd).Standalone()

	cognitoIdentity_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the identity pool.")
	cognitoIdentity_tagResourceCmd.Flags().String("tags", "", "The tags to assign to the identity pool.")
	cognitoIdentity_tagResourceCmd.MarkFlagRequired("resource-arn")
	cognitoIdentity_tagResourceCmd.MarkFlagRequired("tags")
	cognitoIdentityCmd.AddCommand(cognitoIdentity_tagResourceCmd)
}
