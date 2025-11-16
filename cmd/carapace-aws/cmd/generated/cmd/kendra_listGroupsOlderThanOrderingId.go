package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_listGroupsOlderThanOrderingIdCmd = &cobra.Command{
	Use:   "list-groups-older-than-ordering-id",
	Short: "Provides a list of groups that are mapped to users before a given ordering or timestamp identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_listGroupsOlderThanOrderingIdCmd).Standalone()

	kendra_listGroupsOlderThanOrderingIdCmd.Flags().String("data-source-id", "", "The identifier of the data source for getting a list of groups mapped to users before a given ordering timestamp identifier.")
	kendra_listGroupsOlderThanOrderingIdCmd.Flags().String("index-id", "", "The identifier of the index for getting a list of groups mapped to users before a given ordering or timestamp identifier.")
	kendra_listGroupsOlderThanOrderingIdCmd.Flags().String("max-results", "", "The maximum number of returned groups that are mapped to users before a given ordering or timestamp identifier.")
	kendra_listGroupsOlderThanOrderingIdCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more data to retrieve), Amazon Kendra returns a pagination token in the response.")
	kendra_listGroupsOlderThanOrderingIdCmd.Flags().String("ordering-id", "", "The timestamp identifier used for the latest `PUT` or `DELETE` action for mapping users to their groups.")
	kendra_listGroupsOlderThanOrderingIdCmd.MarkFlagRequired("index-id")
	kendra_listGroupsOlderThanOrderingIdCmd.MarkFlagRequired("ordering-id")
	kendraCmd.AddCommand(kendra_listGroupsOlderThanOrderingIdCmd)
}
