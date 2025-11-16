package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_createContactListCmd = &cobra.Command{
	Use:   "create-contact-list",
	Short: "Creates a contact list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_createContactListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_createContactListCmd).Standalone()

		sesv2_createContactListCmd.Flags().String("contact-list-name", "", "The name of the contact list.")
		sesv2_createContactListCmd.Flags().String("description", "", "A description of what the contact list is about.")
		sesv2_createContactListCmd.Flags().String("tags", "", "The tags associated with a contact list.")
		sesv2_createContactListCmd.Flags().String("topics", "", "An interest group, theme, or label within a list.")
		sesv2_createContactListCmd.MarkFlagRequired("contact-list-name")
	})
	sesv2Cmd.AddCommand(sesv2_createContactListCmd)
}
