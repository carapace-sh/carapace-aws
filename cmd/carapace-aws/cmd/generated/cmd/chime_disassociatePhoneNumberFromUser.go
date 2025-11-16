package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_disassociatePhoneNumberFromUserCmd = &cobra.Command{
	Use:   "disassociate-phone-number-from-user",
	Short: "Disassociates the primary provisioned phone number from the specified Amazon Chime user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_disassociatePhoneNumberFromUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_disassociatePhoneNumberFromUserCmd).Standalone()

		chime_disassociatePhoneNumberFromUserCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_disassociatePhoneNumberFromUserCmd.Flags().String("user-id", "", "The user ID.")
		chime_disassociatePhoneNumberFromUserCmd.MarkFlagRequired("account-id")
		chime_disassociatePhoneNumberFromUserCmd.MarkFlagRequired("user-id")
	})
	chimeCmd.AddCommand(chime_disassociatePhoneNumberFromUserCmd)
}
