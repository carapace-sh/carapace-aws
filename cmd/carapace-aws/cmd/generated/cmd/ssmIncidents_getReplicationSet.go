package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_getReplicationSetCmd = &cobra.Command{
	Use:   "get-replication-set",
	Short: "Retrieve your Incident Manager replication set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_getReplicationSetCmd).Standalone()

	ssmIncidents_getReplicationSetCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the replication set you want to retrieve.")
	ssmIncidents_getReplicationSetCmd.MarkFlagRequired("arn")
	ssmIncidentsCmd.AddCommand(ssmIncidents_getReplicationSetCmd)
}
