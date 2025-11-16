package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmContacts_untagResourceCmd).Standalone()

		ssmContacts_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the contact or escalation plan.")
		ssmContacts_untagResourceCmd.Flags().String("tag-keys", "", "The key of the tag that you want to remove.")
		ssmContacts_untagResourceCmd.MarkFlagRequired("resource-arn")
		ssmContacts_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	ssmContactsCmd.AddCommand(ssmContacts_untagResourceCmd)
}
