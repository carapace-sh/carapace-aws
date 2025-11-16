package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_createCloudVmClusterCmd = &cobra.Command{
	Use:   "create-cloud-vm-cluster",
	Short: "Creates a VM cluster on the specified Exadata infrastructure.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_createCloudVmClusterCmd).Standalone()

	odb_createCloudVmClusterCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	odb_createCloudVmClusterCmd.Flags().String("cloud-exadata-infrastructure-id", "", "The unique identifier of the Exadata infrastructure for this VM cluster.")
	odb_createCloudVmClusterCmd.Flags().String("cluster-name", "", "A name for the Grid Infrastructure cluster.")
	odb_createCloudVmClusterCmd.Flags().String("cpu-core-count", "", "The number of CPU cores to enable on the VM cluster.")
	odb_createCloudVmClusterCmd.Flags().String("data-collection-options", "", "The set of preferences for the various diagnostic collection options for the VM cluster.")
	odb_createCloudVmClusterCmd.Flags().String("data-storage-size-in-tbs", "", "The size of the data disk group, in terabytes (TBs), to allocate for the VM cluster.")
	odb_createCloudVmClusterCmd.Flags().String("db-node-storage-size-in-gbs", "", "The amount of local node storage, in gigabytes (GBs), to allocate for the VM cluster.")
	odb_createCloudVmClusterCmd.Flags().String("db-servers", "", "The list of database servers for the VM cluster.")
	odb_createCloudVmClusterCmd.Flags().String("display-name", "", "A user-friendly name for the VM cluster.")
	odb_createCloudVmClusterCmd.Flags().String("gi-version", "", "A valid software version of Oracle Grid Infrastructure (GI).")
	odb_createCloudVmClusterCmd.Flags().String("hostname", "", "The host name for the VM cluster.")
	odb_createCloudVmClusterCmd.Flags().Bool("is-local-backup-enabled", false, "Specifies whether to enable database backups to local Exadata storage for the VM cluster.")
	odb_createCloudVmClusterCmd.Flags().Bool("is-sparse-diskgroup-enabled", false, "Specifies whether to create a sparse disk group for the VM cluster.")
	odb_createCloudVmClusterCmd.Flags().String("license-model", "", "The Oracle license model to apply to the VM cluster.")
	odb_createCloudVmClusterCmd.Flags().String("memory-size-in-gbs", "", "The amount of memory, in gigabytes (GBs), to allocate for the VM cluster.")
	odb_createCloudVmClusterCmd.Flags().Bool("no-is-local-backup-enabled", false, "Specifies whether to enable database backups to local Exadata storage for the VM cluster.")
	odb_createCloudVmClusterCmd.Flags().Bool("no-is-sparse-diskgroup-enabled", false, "Specifies whether to create a sparse disk group for the VM cluster.")
	odb_createCloudVmClusterCmd.Flags().String("odb-network-id", "", "The unique identifier of the ODB network for the VM cluster.")
	odb_createCloudVmClusterCmd.Flags().String("scan-listener-port-tcp", "", "The port number for TCP connections to the single client access name (SCAN) listener.")
	odb_createCloudVmClusterCmd.Flags().String("ssh-public-keys", "", "The public key portion of one or more key pairs used for SSH access to the VM cluster.")
	odb_createCloudVmClusterCmd.Flags().String("system-version", "", "The version of the operating system of the image for the VM cluster.")
	odb_createCloudVmClusterCmd.Flags().String("tags", "", "The list of resource tags to apply to the VM cluster.")
	odb_createCloudVmClusterCmd.Flags().String("time-zone", "", "The time zone for the VM cluster.")
	odb_createCloudVmClusterCmd.MarkFlagRequired("cloud-exadata-infrastructure-id")
	odb_createCloudVmClusterCmd.MarkFlagRequired("cpu-core-count")
	odb_createCloudVmClusterCmd.MarkFlagRequired("display-name")
	odb_createCloudVmClusterCmd.MarkFlagRequired("gi-version")
	odb_createCloudVmClusterCmd.MarkFlagRequired("hostname")
	odb_createCloudVmClusterCmd.Flag("no-is-local-backup-enabled").Hidden = true
	odb_createCloudVmClusterCmd.Flag("no-is-sparse-diskgroup-enabled").Hidden = true
	odb_createCloudVmClusterCmd.MarkFlagRequired("odb-network-id")
	odb_createCloudVmClusterCmd.MarkFlagRequired("ssh-public-keys")
	odbCmd.AddCommand(odb_createCloudVmClusterCmd)
}
