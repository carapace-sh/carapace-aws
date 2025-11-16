package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_cancelContactCmd = &cobra.Command{
	Use:   "cancel-contact",
	Short: "Cancels a contact with a specified contact ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_cancelContactCmd).Standalone()

	groundstation_cancelContactCmd.Flags().String("contact-id", "", "UUID of a contact.")
	groundstation_cancelContactCmd.MarkFlagRequired("contact-id")
	groundstationCmd.AddCommand(groundstation_cancelContactCmd)
}
