package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_createEventStreamCmd = &cobra.Command{
	Use:   "create-event-stream",
	Short: "Creates an event stream, which is a subscription to real-time events, such as when profiles are created and updated through Amazon Connect Customer Profiles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_createEventStreamCmd).Standalone()

	customerProfiles_createEventStreamCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_createEventStreamCmd.Flags().String("event-stream-name", "", "The name of the event stream.")
	customerProfiles_createEventStreamCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	customerProfiles_createEventStreamCmd.Flags().String("uri", "", "The StreamARN of the destination to deliver profile events to.")
	customerProfiles_createEventStreamCmd.MarkFlagRequired("domain-name")
	customerProfiles_createEventStreamCmd.MarkFlagRequired("event-stream-name")
	customerProfiles_createEventStreamCmd.MarkFlagRequired("uri")
	customerProfilesCmd.AddCommand(customerProfiles_createEventStreamCmd)
}
