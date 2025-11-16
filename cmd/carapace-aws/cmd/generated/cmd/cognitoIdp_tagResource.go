package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns a set of tags to an Amazon Cognito user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_tagResourceCmd).Standalone()

		cognitoIdp_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the user pool to assign the tags to.")
		cognitoIdp_tagResourceCmd.Flags().String("tags", "", "An array of tag keys and values that you want to assign to the user pool.")
		cognitoIdp_tagResourceCmd.MarkFlagRequired("resource-arn")
		cognitoIdp_tagResourceCmd.MarkFlagRequired("tags")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_tagResourceCmd)
}
