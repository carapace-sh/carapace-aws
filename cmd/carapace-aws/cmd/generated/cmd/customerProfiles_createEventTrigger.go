package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_createEventTriggerCmd = &cobra.Command{
	Use:   "create-event-trigger",
	Short: "Creates an event trigger, which specifies the rules when to perform action based on customer's ingested data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_createEventTriggerCmd).Standalone()

	customerProfiles_createEventTriggerCmd.Flags().String("description", "", "The description of the event trigger.")
	customerProfiles_createEventTriggerCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_createEventTriggerCmd.Flags().String("event-trigger-conditions", "", "A list of conditions that determine when an event should trigger the destination.")
	customerProfiles_createEventTriggerCmd.Flags().String("event-trigger-limits", "", "Defines limits controlling whether an event triggers the destination, based on ingestion latency and the number of invocations per profile over specific time periods.")
	customerProfiles_createEventTriggerCmd.Flags().String("event-trigger-name", "", "The unique name of the event trigger.")
	customerProfiles_createEventTriggerCmd.Flags().String("object-type-name", "", "The unique name of the object type.")
	customerProfiles_createEventTriggerCmd.Flags().String("segment-filter", "", "The destination is triggered only for profiles that meet the criteria of a segment definition.")
	customerProfiles_createEventTriggerCmd.Flags().String("tags", "", "An array of key-value pairs to apply to this resource.")
	customerProfiles_createEventTriggerCmd.MarkFlagRequired("domain-name")
	customerProfiles_createEventTriggerCmd.MarkFlagRequired("event-trigger-conditions")
	customerProfiles_createEventTriggerCmd.MarkFlagRequired("event-trigger-name")
	customerProfiles_createEventTriggerCmd.MarkFlagRequired("object-type-name")
	customerProfilesCmd.AddCommand(customerProfiles_createEventTriggerCmd)
}
