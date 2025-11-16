package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_updateEmergencyContactSettingsCmd = &cobra.Command{
	Use:   "update-emergency-contact-settings",
	Short: "Updates the details of the list of email addresses and phone numbers that the Shield Response Team (SRT) can use to contact you if you have proactive engagement enabled, for escalations to the SRT and to initiate proactive customer support.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_updateEmergencyContactSettingsCmd).Standalone()

	shield_updateEmergencyContactSettingsCmd.Flags().String("emergency-contact-list", "", "A list of email addresses and phone numbers that the Shield Response Team (SRT) can use to contact you if you have proactive engagement enabled, for escalations to the SRT and to initiate proactive customer support.")
	shieldCmd.AddCommand(shield_updateEmergencyContactSettingsCmd)
}
