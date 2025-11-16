package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_retryDataReplicationCmd = &cobra.Command{
	Use:   "retry-data-replication",
	Short: "Causes the data replication initiation sequence to begin immediately upon next Handshake for specified SourceServer IDs, regardless of when the previous initiation started.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_retryDataReplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_retryDataReplicationCmd).Standalone()

		mgn_retryDataReplicationCmd.Flags().String("account-id", "", "Retry data replication for Account ID.")
		mgn_retryDataReplicationCmd.Flags().String("source-server-id", "", "Retry data replication for Source Server ID.")
		mgn_retryDataReplicationCmd.MarkFlagRequired("source-server-id")
	})
	mgnCmd.AddCommand(mgn_retryDataReplicationCmd)
}
