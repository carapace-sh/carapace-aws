package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_disassociatePhoneNumberContactFlowCmd = &cobra.Command{
	Use:   "disassociate-phone-number-contact-flow",
	Short: "Removes the flow association from a phone number claimed to your Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_disassociatePhoneNumberContactFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_disassociatePhoneNumberContactFlowCmd).Standalone()

		connect_disassociatePhoneNumberContactFlowCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_disassociatePhoneNumberContactFlowCmd.Flags().String("phone-number-id", "", "A unique identifier for the phone number.")
		connect_disassociatePhoneNumberContactFlowCmd.MarkFlagRequired("instance-id")
		connect_disassociatePhoneNumberContactFlowCmd.MarkFlagRequired("phone-number-id")
	})
	connectCmd.AddCommand(connect_disassociatePhoneNumberContactFlowCmd)
}
