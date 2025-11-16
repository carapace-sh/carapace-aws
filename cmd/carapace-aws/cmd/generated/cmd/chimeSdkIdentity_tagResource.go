package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Applies the specified tags to the specified Amazon Chime SDK identity resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_tagResourceCmd).Standalone()

	chimeSdkIdentity_tagResourceCmd.Flags().String("resource-arn", "", "The resource ARN.")
	chimeSdkIdentity_tagResourceCmd.Flags().String("tags", "", "The tag key-value pairs.")
	chimeSdkIdentity_tagResourceCmd.MarkFlagRequired("resource-arn")
	chimeSdkIdentity_tagResourceCmd.MarkFlagRequired("tags")
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_tagResourceCmd)
}
