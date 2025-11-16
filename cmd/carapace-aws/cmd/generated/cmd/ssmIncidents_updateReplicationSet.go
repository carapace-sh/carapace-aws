package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_updateReplicationSetCmd = &cobra.Command{
	Use:   "update-replication-set",
	Short: "Add or delete Regions from your replication set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_updateReplicationSetCmd).Standalone()

	ssmIncidents_updateReplicationSetCmd.Flags().String("actions", "", "An action to add or delete a Region.")
	ssmIncidents_updateReplicationSetCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the replication set you're updating.")
	ssmIncidents_updateReplicationSetCmd.Flags().String("client-token", "", "A token that ensures that the operation is called only once with the specified details.")
	ssmIncidents_updateReplicationSetCmd.MarkFlagRequired("actions")
	ssmIncidents_updateReplicationSetCmd.MarkFlagRequired("arn")
	ssmIncidentsCmd.AddCommand(ssmIncidents_updateReplicationSetCmd)
}
