package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Given tag IDs that you previously assigned to a user pool, removes them.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_untagResourceCmd).Standalone()

	cognitoIdp_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the user pool that the tags are assigned to.")
	cognitoIdp_untagResourceCmd.Flags().String("tag-keys", "", "An array of tag keys that you want to remove from the user pool.")
	cognitoIdp_untagResourceCmd.MarkFlagRequired("resource-arn")
	cognitoIdp_untagResourceCmd.MarkFlagRequired("tag-keys")
	cognitoIdpCmd.AddCommand(cognitoIdp_untagResourceCmd)
}
