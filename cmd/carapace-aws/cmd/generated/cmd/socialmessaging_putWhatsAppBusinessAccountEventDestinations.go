package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessaging_putWhatsAppBusinessAccountEventDestinationsCmd = &cobra.Command{
	Use:   "put-whats-app-business-account-event-destinations",
	Short: "Add an event destination to log event data from WhatsApp for a WhatsApp Business Account (WABA).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessaging_putWhatsAppBusinessAccountEventDestinationsCmd).Standalone()

	socialmessaging_putWhatsAppBusinessAccountEventDestinationsCmd.Flags().String("event-destinations", "", "An array of `WhatsAppBusinessAccountEventDestination` event destinations.")
	socialmessaging_putWhatsAppBusinessAccountEventDestinationsCmd.Flags().String("id", "", "The unique identifier of your WhatsApp Business Account.")
	socialmessaging_putWhatsAppBusinessAccountEventDestinationsCmd.MarkFlagRequired("event-destinations")
	socialmessaging_putWhatsAppBusinessAccountEventDestinationsCmd.MarkFlagRequired("id")
	socialmessagingCmd.AddCommand(socialmessaging_putWhatsAppBusinessAccountEventDestinationsCmd)
}
