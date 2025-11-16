package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_deleteContactListCmd = &cobra.Command{
	Use:   "delete-contact-list",
	Short: "Deletes a contact list and all of the contacts on that list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_deleteContactListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_deleteContactListCmd).Standalone()

		sesv2_deleteContactListCmd.Flags().String("contact-list-name", "", "The name of the contact list.")
		sesv2_deleteContactListCmd.MarkFlagRequired("contact-list-name")
	})
	sesv2Cmd.AddCommand(sesv2_deleteContactListCmd)
}
