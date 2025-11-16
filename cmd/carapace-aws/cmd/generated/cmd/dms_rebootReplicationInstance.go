package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_rebootReplicationInstanceCmd = &cobra.Command{
	Use:   "reboot-replication-instance",
	Short: "Reboots a replication instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_rebootReplicationInstanceCmd).Standalone()

	dms_rebootReplicationInstanceCmd.Flags().String("force-failover", "", "If this parameter is `true`, the reboot is conducted through a Multi-AZ failover.")
	dms_rebootReplicationInstanceCmd.Flags().String("force-planned-failover", "", "If this parameter is `true`, the reboot is conducted through a planned Multi-AZ failover where resources are released and cleaned up prior to conducting the failover.")
	dms_rebootReplicationInstanceCmd.Flags().String("replication-instance-arn", "", "The Amazon Resource Name (ARN) of the replication instance.")
	dms_rebootReplicationInstanceCmd.MarkFlagRequired("replication-instance-arn")
	dmsCmd.AddCommand(dms_rebootReplicationInstanceCmd)
}
