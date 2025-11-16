package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_deleteReplicationSetCmd = &cobra.Command{
	Use:   "delete-replication-set",
	Short: "Deletes all Regions in your replication set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_deleteReplicationSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmIncidents_deleteReplicationSetCmd).Standalone()

		ssmIncidents_deleteReplicationSetCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the replication set you're deleting.")
		ssmIncidents_deleteReplicationSetCmd.MarkFlagRequired("arn")
	})
	ssmIncidentsCmd.AddCommand(ssmIncidents_deleteReplicationSetCmd)
}
