package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_putEmailIdentityFeedbackAttributesCmd = &cobra.Command{
	Use:   "put-email-identity-feedback-attributes",
	Short: "Used to enable or disable feedback forwarding for an identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_putEmailIdentityFeedbackAttributesCmd).Standalone()

	pinpointEmail_putEmailIdentityFeedbackAttributesCmd.Flags().String("email-forwarding-enabled", "", "Sets the feedback forwarding configuration for the identity.")
	pinpointEmail_putEmailIdentityFeedbackAttributesCmd.Flags().String("email-identity", "", "The email identity that you want to configure bounce and complaint feedback forwarding for.")
	pinpointEmail_putEmailIdentityFeedbackAttributesCmd.MarkFlagRequired("email-identity")
	pinpointEmailCmd.AddCommand(pinpointEmail_putEmailIdentityFeedbackAttributesCmd)
}
