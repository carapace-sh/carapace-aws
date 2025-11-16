package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_describeEmergencyContactSettingsCmd = &cobra.Command{
	Use:   "describe-emergency-contact-settings",
	Short: "A list of email addresses and phone numbers that the Shield Response Team (SRT) can use to contact you if you have proactive engagement enabled, for escalations to the SRT and to initiate proactive customer support.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_describeEmergencyContactSettingsCmd).Standalone()

	shieldCmd.AddCommand(shield_describeEmergencyContactSettingsCmd)
}
