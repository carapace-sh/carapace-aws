package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_deleteEventTriggerCmd = &cobra.Command{
	Use:   "delete-event-trigger",
	Short: "Disable and deletes the Event Trigger.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_deleteEventTriggerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_deleteEventTriggerCmd).Standalone()

		customerProfiles_deleteEventTriggerCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_deleteEventTriggerCmd.Flags().String("event-trigger-name", "", "The unique name of the event trigger.")
		customerProfiles_deleteEventTriggerCmd.MarkFlagRequired("domain-name")
		customerProfiles_deleteEventTriggerCmd.MarkFlagRequired("event-trigger-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_deleteEventTriggerCmd)
}
