package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_stopReplicationCmd = &cobra.Command{
	Use:   "stop-replication",
	Short: "Stops replication for a Source Server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_stopReplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_stopReplicationCmd).Standalone()

		drs_stopReplicationCmd.Flags().String("source-server-id", "", "The ID of the Source Server to stop replication for.")
		drs_stopReplicationCmd.MarkFlagRequired("source-server-id")
	})
	drsCmd.AddCommand(drs_stopReplicationCmd)
}
