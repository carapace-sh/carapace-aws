package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags that are assigned to an Amazon Cognito identity pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdentity_listTagsForResourceCmd).Standalone()

		cognitoIdentity_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the identity pool that the tags are assigned to.")
		cognitoIdentity_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	cognitoIdentityCmd.AddCommand(cognitoIdentity_listTagsForResourceCmd)
}
