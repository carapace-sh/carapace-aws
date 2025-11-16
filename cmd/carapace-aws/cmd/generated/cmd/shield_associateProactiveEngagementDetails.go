package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_associateProactiveEngagementDetailsCmd = &cobra.Command{
	Use:   "associate-proactive-engagement-details",
	Short: "Initializes proactive engagement and sets the list of contacts for the Shield Response Team (SRT) to use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_associateProactiveEngagementDetailsCmd).Standalone()

	shield_associateProactiveEngagementDetailsCmd.Flags().String("emergency-contact-list", "", "A list of email addresses and phone numbers that the Shield Response Team (SRT) can use to contact you for escalations to the SRT and to initiate proactive customer support.")
	shield_associateProactiveEngagementDetailsCmd.MarkFlagRequired("emergency-contact-list")
	shieldCmd.AddCommand(shield_associateProactiveEngagementDetailsCmd)
}
