package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putEmailIdentityFeedbackAttributesCmd = &cobra.Command{
	Use:   "put-email-identity-feedback-attributes",
	Short: "Used to enable or disable feedback forwarding for an identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putEmailIdentityFeedbackAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_putEmailIdentityFeedbackAttributesCmd).Standalone()

		sesv2_putEmailIdentityFeedbackAttributesCmd.Flags().String("email-forwarding-enabled", "", "Sets the feedback forwarding configuration for the identity.")
		sesv2_putEmailIdentityFeedbackAttributesCmd.Flags().String("email-identity", "", "The email identity.")
		sesv2_putEmailIdentityFeedbackAttributesCmd.MarkFlagRequired("email-identity")
	})
	sesv2Cmd.AddCommand(sesv2_putEmailIdentityFeedbackAttributesCmd)
}
