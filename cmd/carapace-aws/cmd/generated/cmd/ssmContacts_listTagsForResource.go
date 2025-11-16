package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags of a contact, escalation plan, rotation, or on-call schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_listTagsForResourceCmd).Standalone()

	ssmContacts_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the contact, escalation plan, rotation, or on-call schedule.")
	ssmContacts_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	ssmContactsCmd.AddCommand(ssmContacts_listTagsForResourceCmd)
}
