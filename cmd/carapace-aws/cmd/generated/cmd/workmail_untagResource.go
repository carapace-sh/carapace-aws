package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Untags the specified tags from the specified WorkMail organization resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_untagResourceCmd).Standalone()

		workmail_untagResourceCmd.Flags().String("resource-arn", "", "The resource ARN.")
		workmail_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys.")
		workmail_untagResourceCmd.MarkFlagRequired("resource-arn")
		workmail_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	workmailCmd.AddCommand(workmail_untagResourceCmd)
}
