package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_updateContactListCmd = &cobra.Command{
	Use:   "update-contact-list",
	Short: "Updates contact list metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_updateContactListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_updateContactListCmd).Standalone()

		sesv2_updateContactListCmd.Flags().String("contact-list-name", "", "The name of the contact list.")
		sesv2_updateContactListCmd.Flags().String("description", "", "A description of what the contact list is about.")
		sesv2_updateContactListCmd.Flags().String("topics", "", "An interest group, theme, or label within a list.")
		sesv2_updateContactListCmd.MarkFlagRequired("contact-list-name")
	})
	sesv2Cmd.AddCommand(sesv2_updateContactListCmd)
}
