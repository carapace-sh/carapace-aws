package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags that are assigned to an Amazon Cognito user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_listTagsForResourceCmd).Standalone()

	cognitoIdp_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the user pool that the tags are assigned to.")
	cognitoIdp_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	cognitoIdpCmd.AddCommand(cognitoIdp_listTagsForResourceCmd)
}
