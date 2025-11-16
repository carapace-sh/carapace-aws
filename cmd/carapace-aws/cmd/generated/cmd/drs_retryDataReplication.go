package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_retryDataReplicationCmd = &cobra.Command{
	Use:   "retry-data-replication",
	Short: "WARNING: RetryDataReplication is deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_retryDataReplicationCmd).Standalone()

	drs_retryDataReplicationCmd.Flags().String("source-server-id", "", "The ID of the Source Server whose data replication should be retried.")
	drs_retryDataReplicationCmd.MarkFlagRequired("source-server-id")
	drsCmd.AddCommand(drs_retryDataReplicationCmd)
}
