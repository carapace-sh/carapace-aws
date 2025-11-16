package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_setIdentityFeedbackForwardingEnabledCmd = &cobra.Command{
	Use:   "set-identity-feedback-forwarding-enabled",
	Short: "Given an identity (an email address or a domain), enables or disables whether Amazon SES forwards bounce and complaint notifications as email.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_setIdentityFeedbackForwardingEnabledCmd).Standalone()

	ses_setIdentityFeedbackForwardingEnabledCmd.Flags().String("forwarding-enabled", "", "Sets whether Amazon SES forwards bounce and complaint notifications as email.")
	ses_setIdentityFeedbackForwardingEnabledCmd.Flags().String("identity", "", "The identity for which to set bounce and complaint notification forwarding.")
	ses_setIdentityFeedbackForwardingEnabledCmd.MarkFlagRequired("forwarding-enabled")
	ses_setIdentityFeedbackForwardingEnabledCmd.MarkFlagRequired("identity")
	sesCmd.AddCommand(ses_setIdentityFeedbackForwardingEnabledCmd)
}
