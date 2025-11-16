package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_createReplicationSetCmd = &cobra.Command{
	Use:   "create-replication-set",
	Short: "A replication set replicates and encrypts your data to the provided Regions with the provided KMS key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_createReplicationSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmIncidents_createReplicationSetCmd).Standalone()

		ssmIncidents_createReplicationSetCmd.Flags().String("client-token", "", "A token that ensures that the operation is called only once with the specified details.")
		ssmIncidents_createReplicationSetCmd.Flags().String("regions", "", "The Regions that Incident Manager replicates your data to.")
		ssmIncidents_createReplicationSetCmd.Flags().String("tags", "", "A list of tags to add to the replication set.")
		ssmIncidents_createReplicationSetCmd.MarkFlagRequired("regions")
	})
	ssmIncidentsCmd.AddCommand(ssmIncidents_createReplicationSetCmd)
}
