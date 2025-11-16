package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_associatePhoneNumberContactFlowCmd = &cobra.Command{
	Use:   "associate-phone-number-contact-flow",
	Short: "Associates a flow with a phone number claimed to your Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_associatePhoneNumberContactFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_associatePhoneNumberContactFlowCmd).Standalone()

		connect_associatePhoneNumberContactFlowCmd.Flags().String("contact-flow-id", "", "The identifier of the flow.")
		connect_associatePhoneNumberContactFlowCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_associatePhoneNumberContactFlowCmd.Flags().String("phone-number-id", "", "A unique identifier for the phone number.")
		connect_associatePhoneNumberContactFlowCmd.MarkFlagRequired("contact-flow-id")
		connect_associatePhoneNumberContactFlowCmd.MarkFlagRequired("instance-id")
		connect_associatePhoneNumberContactFlowCmd.MarkFlagRequired("phone-number-id")
	})
	connectCmd.AddCommand(connect_associatePhoneNumberContactFlowCmd)
}
