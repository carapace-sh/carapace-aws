package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getContactListCmd = &cobra.Command{
	Use:   "get-contact-list",
	Short: "Returns contact list metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getContactListCmd).Standalone()

	sesv2_getContactListCmd.Flags().String("contact-list-name", "", "The name of the contact list.")
	sesv2_getContactListCmd.MarkFlagRequired("contact-list-name")
	sesv2Cmd.AddCommand(sesv2_getContactListCmd)
}
