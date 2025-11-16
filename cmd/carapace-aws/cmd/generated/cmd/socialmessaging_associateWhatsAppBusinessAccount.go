package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_associateWhatsAppBusinessAccountCmd = &cobra.Command{
	Use:   "associate-whats-app-business-account",
	Short: "This is only used through the Amazon Web Services console during sign-up to associate your WhatsApp Business Account to your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_associateWhatsAppBusinessAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(socialmessaging_associateWhatsAppBusinessAccountCmd).Standalone()

		socialmessaging_associateWhatsAppBusinessAccountCmd.Flags().String("setup-finalization", "", "A JSON object that contains the phone numbers and WhatsApp Business Account to link to your account.")
		socialmessaging_associateWhatsAppBusinessAccountCmd.Flags().String("signup-callback", "", "Contains the callback access token.")
	})
	socialmessagingCmd.AddCommand(socialmessaging_associateWhatsAppBusinessAccountCmd)
}
