package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_updateEventTriggerCmd = &cobra.Command{
	Use:   "update-event-trigger",
	Short: "Update the properties of an Event Trigger.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_updateEventTriggerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_updateEventTriggerCmd).Standalone()

		customerProfiles_updateEventTriggerCmd.Flags().String("description", "", "The description of the event trigger.")
		customerProfiles_updateEventTriggerCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_updateEventTriggerCmd.Flags().String("event-trigger-conditions", "", "A list of conditions that determine when an event should trigger the destination.")
		customerProfiles_updateEventTriggerCmd.Flags().String("event-trigger-limits", "", "Defines limits controlling whether an event triggers the destination, based on ingestion latency and the number of invocations per profile over specific time periods.")
		customerProfiles_updateEventTriggerCmd.Flags().String("event-trigger-name", "", "The unique name of the event trigger.")
		customerProfiles_updateEventTriggerCmd.Flags().String("object-type-name", "", "The unique name of the object type.")
		customerProfiles_updateEventTriggerCmd.Flags().String("segment-filter", "", "The destination is triggered only for profiles that meet the criteria of a segment definition.")
		customerProfiles_updateEventTriggerCmd.MarkFlagRequired("domain-name")
		customerProfiles_updateEventTriggerCmd.MarkFlagRequired("event-trigger-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_updateEventTriggerCmd)
}
