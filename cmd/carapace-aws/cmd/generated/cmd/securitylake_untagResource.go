package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags (keys and values) from an Amazon Security Lake resource: a subscriber, or the data lake configuration for your Amazon Web Services account in a particular Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_untagResourceCmd).Standalone()

	securitylake_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Amazon Security Lake resource to remove one or more tags from.")
	securitylake_untagResourceCmd.Flags().String("tag-keys", "", "A list of one or more tag keys.")
	securitylake_untagResourceCmd.MarkFlagRequired("resource-arn")
	securitylake_untagResourceCmd.MarkFlagRequired("tag-keys")
	securitylakeCmd.AddCommand(securitylake_untagResourceCmd)
}
