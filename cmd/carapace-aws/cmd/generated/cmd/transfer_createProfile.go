package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_createProfileCmd = &cobra.Command{
	Use:   "create-profile",
	Short: "Creates the local or partner profile to use for AS2 transfers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_createProfileCmd).Standalone()

	transfer_createProfileCmd.Flags().String("as2-id", "", "The `As2Id` is the *AS2-name*, as defined in the [RFC 4130](https://datatracker.ietf.org/doc/html/rfc4130). For inbound transfers, this is the `AS2-From` header for the AS2 messages sent from the partner.")
	transfer_createProfileCmd.Flags().String("certificate-ids", "", "An array of identifiers for the imported certificates.")
	transfer_createProfileCmd.Flags().String("profile-type", "", "Determines the type of profile to create:")
	transfer_createProfileCmd.Flags().String("tags", "", "Key-value pairs that can be used to group and search for AS2 profiles.")
	transfer_createProfileCmd.MarkFlagRequired("as2-id")
	transfer_createProfileCmd.MarkFlagRequired("profile-type")
	transferCmd.AddCommand(transfer_createProfileCmd)
}
