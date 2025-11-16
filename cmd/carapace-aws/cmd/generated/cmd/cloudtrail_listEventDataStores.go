package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_listEventDataStoresCmd = &cobra.Command{
	Use:   "list-event-data-stores",
	Short: "Returns information about all event data stores in the account, in the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_listEventDataStoresCmd).Standalone()

	cloudtrail_listEventDataStoresCmd.Flags().String("max-results", "", "The maximum number of event data stores to display on a single page.")
	cloudtrail_listEventDataStoresCmd.Flags().String("next-token", "", "A token you can use to get the next page of event data store results.")
	cloudtrailCmd.AddCommand(cloudtrail_listEventDataStoresCmd)
}
