package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmContacts_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags a contact or escalation plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmContacts_tagResourceCmd).Standalone()

	ssmContacts_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the contact or escalation plan.")
	ssmContacts_tagResourceCmd.Flags().String("tags", "", "A list of tags that you are adding to the contact or escalation plan.")
	ssmContacts_tagResourceCmd.MarkFlagRequired("resource-arn")
	ssmContacts_tagResourceCmd.MarkFlagRequired("tags")
	ssmContactsCmd.AddCommand(ssmContacts_tagResourceCmd)
}
