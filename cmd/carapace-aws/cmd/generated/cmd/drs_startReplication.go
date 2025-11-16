package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_startReplicationCmd = &cobra.Command{
	Use:   "start-replication",
	Short: "Starts replication for a stopped Source Server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_startReplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_startReplicationCmd).Standalone()

		drs_startReplicationCmd.Flags().String("source-server-id", "", "The ID of the Source Server to start replication for.")
		drs_startReplicationCmd.MarkFlagRequired("source-server-id")
	})
	drsCmd.AddCommand(drs_startReplicationCmd)
}
