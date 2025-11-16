package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_startWebRtccontactCmd = &cobra.Command{
	Use:   "start-web-rtccontact",
	Short: "Places an inbound in-app, web, or video call to a contact, and then initiates the flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_startWebRtccontactCmd).Standalone()

	connect_startWebRtccontactCmd.Flags().String("allowed-capabilities", "", "Information about the video sharing capabilities of the participants (customer, agent).")
	connect_startWebRtccontactCmd.Flags().String("attributes", "", "A custom key-value pair using an attribute map.")
	connect_startWebRtccontactCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_startWebRtccontactCmd.Flags().String("contact-flow-id", "", "The identifier of the flow for the call.")
	connect_startWebRtccontactCmd.Flags().String("description", "", "A description of the task that is shown to an agent in the Contact Control Panel (CCP).")
	connect_startWebRtccontactCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_startWebRtccontactCmd.Flags().String("participant-details", "", "")
	connect_startWebRtccontactCmd.Flags().String("references", "", "A formatted URL that is shown to an agent in the Contact Control Panel (CCP).")
	connect_startWebRtccontactCmd.Flags().String("related-contact-id", "", "The unique identifier for an Amazon Connect contact.")
	connect_startWebRtccontactCmd.MarkFlagRequired("contact-flow-id")
	connect_startWebRtccontactCmd.MarkFlagRequired("instance-id")
	connect_startWebRtccontactCmd.MarkFlagRequired("participant-details")
	connectCmd.AddCommand(connect_startWebRtccontactCmd)
}
