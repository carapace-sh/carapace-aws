package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listEventStreamsCmd = &cobra.Command{
	Use:   "list-event-streams",
	Short: "Returns a list of all the event streams in a specific domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listEventStreamsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(customerProfiles_listEventStreamsCmd).Standalone()

		customerProfiles_listEventStreamsCmd.Flags().String("domain-name", "", "The unique name of the domain.")
		customerProfiles_listEventStreamsCmd.Flags().String("max-results", "", "The maximum number of objects returned per page.")
		customerProfiles_listEventStreamsCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
		customerProfiles_listEventStreamsCmd.MarkFlagRequired("domain-name")
	})
	customerProfilesCmd.AddCommand(customerProfiles_listEventStreamsCmd)
}
