package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteContactMethodCmd = &cobra.Command{
	Use:   "delete-contact-method",
	Short: "Deletes a contact method.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteContactMethodCmd).Standalone()

	lightsail_deleteContactMethodCmd.Flags().String("protocol", "", "The protocol that will be deleted, such as `Email` or `SMS` (text messaging).")
	lightsail_deleteContactMethodCmd.MarkFlagRequired("protocol")
	lightsailCmd.AddCommand(lightsail_deleteContactMethodCmd)
}
