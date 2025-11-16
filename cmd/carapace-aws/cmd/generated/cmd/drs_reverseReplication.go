package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_reverseReplicationCmd = &cobra.Command{
	Use:   "reverse-replication",
	Short: "Start replication to origin / target region - applies only to protected instances that originated in EC2.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_reverseReplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_reverseReplicationCmd).Standalone()

		drs_reverseReplicationCmd.Flags().String("recovery-instance-id", "", "The ID of the Recovery Instance that we want to reverse the replication for.")
		drs_reverseReplicationCmd.MarkFlagRequired("recovery-instance-id")
	})
	drsCmd.AddCommand(drs_reverseReplicationCmd)
}
