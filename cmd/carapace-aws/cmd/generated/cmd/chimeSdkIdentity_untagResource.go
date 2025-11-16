package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from the specified Amazon Chime SDK identity resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkIdentity_untagResourceCmd).Standalone()

		chimeSdkIdentity_untagResourceCmd.Flags().String("resource-arn", "", "The resource ARN.")
		chimeSdkIdentity_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys.")
		chimeSdkIdentity_untagResourceCmd.MarkFlagRequired("resource-arn")
		chimeSdkIdentity_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_untagResourceCmd)
}
