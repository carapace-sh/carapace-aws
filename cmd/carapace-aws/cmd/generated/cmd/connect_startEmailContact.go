package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_startEmailContactCmd = &cobra.Command{
	Use:   "start-email-contact",
	Short: "Creates an inbound email contact and initiates a flow to start the email contact for the customer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_startEmailContactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_startEmailContactCmd).Standalone()

		connect_startEmailContactCmd.Flags().String("additional-recipients", "", "The additional recipients address of the email.")
		connect_startEmailContactCmd.Flags().String("attachments", "", "List of S3 presigned URLs of email attachments and their file name.")
		connect_startEmailContactCmd.Flags().String("attributes", "", "A custom key-value pair using an attribute map.")
		connect_startEmailContactCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_startEmailContactCmd.Flags().String("contact-flow-id", "", "The identifier of the flow for initiating the emails.")
		connect_startEmailContactCmd.Flags().String("description", "", "A description of the email contact.")
		connect_startEmailContactCmd.Flags().String("destination-email-address", "", "The email address associated with the Amazon Connect instance.")
		connect_startEmailContactCmd.Flags().String("email-message", "", "The email message body to be sent to the newly created email.")
		connect_startEmailContactCmd.Flags().String("from-email-address", "", "The email address of the customer.")
		connect_startEmailContactCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_startEmailContactCmd.Flags().String("name", "", "The name of a email that is shown to an agent in the Contact Control Panel (CCP).")
		connect_startEmailContactCmd.Flags().String("references", "", "A formatted URL that is shown to an agent in the Contact Control Panel (CCP).")
		connect_startEmailContactCmd.Flags().String("related-contact-id", "", "The contactId that is related to this contact.")
		connect_startEmailContactCmd.Flags().String("segment-attributes", "", "A set of system defined key-value pairs stored on individual contact segments using an attribute map.")
		connect_startEmailContactCmd.MarkFlagRequired("destination-email-address")
		connect_startEmailContactCmd.MarkFlagRequired("email-message")
		connect_startEmailContactCmd.MarkFlagRequired("from-email-address")
		connect_startEmailContactCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_startEmailContactCmd)
}
