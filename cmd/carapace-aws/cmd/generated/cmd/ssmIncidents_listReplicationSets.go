package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_listReplicationSetsCmd = &cobra.Command{
	Use:   "list-replication-sets",
	Short: "Lists details about the replication set configured in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_listReplicationSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmIncidents_listReplicationSetsCmd).Standalone()

		ssmIncidents_listReplicationSetsCmd.Flags().String("max-results", "", "The maximum number of results per page.")
		ssmIncidents_listReplicationSetsCmd.Flags().String("next-token", "", "The pagination token for the next set of items to return.")
	})
	ssmIncidentsCmd.AddCommand(ssmIncidents_listReplicationSetsCmd)
}
