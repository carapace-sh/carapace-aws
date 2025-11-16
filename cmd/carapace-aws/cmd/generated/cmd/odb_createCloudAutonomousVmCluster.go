package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_createCloudAutonomousVmClusterCmd = &cobra.Command{
	Use:   "create-cloud-autonomous-vm-cluster",
	Short: "Creates a new Autonomous VM cluster in the specified Exadata infrastructure.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_createCloudAutonomousVmClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_createCloudAutonomousVmClusterCmd).Standalone()

		odb_createCloudAutonomousVmClusterCmd.Flags().String("autonomous-data-storage-size-in-tbs", "", "The data disk group size to be allocated for Autonomous Databases, in terabytes (TB).")
		odb_createCloudAutonomousVmClusterCmd.Flags().String("client-token", "", "A client-provided token to ensure idempotency of the request.")
		odb_createCloudAutonomousVmClusterCmd.Flags().String("cloud-exadata-infrastructure-id", "", "The unique identifier of the Exadata infrastructure where the VM cluster will be created.")
		odb_createCloudAutonomousVmClusterCmd.Flags().String("cpu-core-count-per-node", "", "The number of CPU cores to be enabled per VM cluster node.")
		odb_createCloudAutonomousVmClusterCmd.Flags().String("db-servers", "", "The list of database servers to be used for the Autonomous VM cluster.")
		odb_createCloudAutonomousVmClusterCmd.Flags().String("description", "", "A user-provided description of the Autonomous VM cluster.")
		odb_createCloudAutonomousVmClusterCmd.Flags().String("display-name", "", "The display name for the Autonomous VM cluster.")
		odb_createCloudAutonomousVmClusterCmd.Flags().Bool("is-mtls-enabled-vm-cluster", false, "Specifies whether to enable mutual TLS (mTLS) authentication for the Autonomous VM cluster.")
		odb_createCloudAutonomousVmClusterCmd.Flags().String("license-model", "", "The Oracle license model to apply to the Autonomous VM cluster.")
		odb_createCloudAutonomousVmClusterCmd.Flags().String("maintenance-window", "", "The scheduling details for the maintenance window.")
		odb_createCloudAutonomousVmClusterCmd.Flags().String("memory-per-oracle-compute-unit-in-gbs", "", "The amount of memory to be allocated per OCPU, in GB.")
		odb_createCloudAutonomousVmClusterCmd.Flags().Bool("no-is-mtls-enabled-vm-cluster", false, "Specifies whether to enable mutual TLS (mTLS) authentication for the Autonomous VM cluster.")
		odb_createCloudAutonomousVmClusterCmd.Flags().String("odb-network-id", "", "The unique identifier of the ODB network to be used for the VM cluster.")
		odb_createCloudAutonomousVmClusterCmd.Flags().String("scan-listener-port-non-tls", "", "The SCAN listener port for non-TLS (TCP) protocol.")
		odb_createCloudAutonomousVmClusterCmd.Flags().String("scan-listener-port-tls", "", "The SCAN listener port for TLS (TCP) protocol.")
		odb_createCloudAutonomousVmClusterCmd.Flags().String("tags", "", "Free-form tags for this resource.")
		odb_createCloudAutonomousVmClusterCmd.Flags().String("time-zone", "", "The time zone to use for the Autonomous VM cluster.")
		odb_createCloudAutonomousVmClusterCmd.Flags().String("total-container-databases", "", "The total number of Autonomous CDBs that you can create in the Autonomous VM cluster.")
		odb_createCloudAutonomousVmClusterCmd.MarkFlagRequired("autonomous-data-storage-size-in-tbs")
		odb_createCloudAutonomousVmClusterCmd.MarkFlagRequired("cloud-exadata-infrastructure-id")
		odb_createCloudAutonomousVmClusterCmd.MarkFlagRequired("cpu-core-count-per-node")
		odb_createCloudAutonomousVmClusterCmd.MarkFlagRequired("display-name")
		odb_createCloudAutonomousVmClusterCmd.MarkFlagRequired("memory-per-oracle-compute-unit-in-gbs")
		odb_createCloudAutonomousVmClusterCmd.Flag("no-is-mtls-enabled-vm-cluster").Hidden = true
		odb_createCloudAutonomousVmClusterCmd.MarkFlagRequired("odb-network-id")
		odb_createCloudAutonomousVmClusterCmd.MarkFlagRequired("total-container-databases")
	})
	odbCmd.AddCommand(odb_createCloudAutonomousVmClusterCmd)
}
