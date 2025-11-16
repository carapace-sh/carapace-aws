package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags applied to an Amazon Chime SDK identity resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_listTagsForResourceCmd).Standalone()

	chimeSdkIdentity_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
	chimeSdkIdentity_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_listTagsForResourceCmd)
}
