package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_createReplicationInstanceCmd = &cobra.Command{
	Use:   "create-replication-instance",
	Short: "Creates the replication instance using the specified parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_createReplicationInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_createReplicationInstanceCmd).Standalone()

		dms_createReplicationInstanceCmd.Flags().String("allocated-storage", "", "The amount of storage (in gigabytes) to be initially allocated for the replication instance.")
		dms_createReplicationInstanceCmd.Flags().String("auto-minor-version-upgrade", "", "A value that indicates whether minor engine upgrades are applied automatically to the replication instance during the maintenance window.")
		dms_createReplicationInstanceCmd.Flags().String("availability-zone", "", "The Availability Zone where the replication instance will be created.")
		dms_createReplicationInstanceCmd.Flags().String("dns-name-servers", "", "A list of custom DNS name servers supported for the replication instance to access your on-premise source or target database.")
		dms_createReplicationInstanceCmd.Flags().String("engine-version", "", "The engine version number of the replication instance.")
		dms_createReplicationInstanceCmd.Flags().String("kerberos-authentication-settings", "", "Specifies the settings required for kerberos authentication when creating the replication instance.")
		dms_createReplicationInstanceCmd.Flags().String("kms-key-id", "", "An KMS key identifier that is used to encrypt the data on the replication instance.")
		dms_createReplicationInstanceCmd.Flags().String("multi-az", "", "Specifies whether the replication instance is a Multi-AZ deployment.")
		dms_createReplicationInstanceCmd.Flags().String("network-type", "", "The type of IP address protocol used by a replication instance, such as IPv4 only or Dual-stack that supports both IPv4 and IPv6 addressing.")
		dms_createReplicationInstanceCmd.Flags().String("preferred-maintenance-window", "", "The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).")
		dms_createReplicationInstanceCmd.Flags().String("publicly-accessible", "", "Specifies the accessibility options for the replication instance.")
		dms_createReplicationInstanceCmd.Flags().String("replication-instance-class", "", "The compute and memory capacity of the replication instance as defined for the specified replication instance class.")
		dms_createReplicationInstanceCmd.Flags().String("replication-instance-identifier", "", "The replication instance identifier.")
		dms_createReplicationInstanceCmd.Flags().String("replication-subnet-group-identifier", "", "A subnet group to associate with the replication instance.")
		dms_createReplicationInstanceCmd.Flags().String("resource-identifier", "", "A friendly name for the resource identifier at the end of the `EndpointArn` response parameter that is returned in the created `Endpoint` object.")
		dms_createReplicationInstanceCmd.Flags().String("tags", "", "One or more tags to be assigned to the replication instance.")
		dms_createReplicationInstanceCmd.Flags().String("vpc-security-group-ids", "", "Specifies the VPC security group to be used with the replication instance.")
		dms_createReplicationInstanceCmd.MarkFlagRequired("replication-instance-class")
		dms_createReplicationInstanceCmd.MarkFlagRequired("replication-instance-identifier")
	})
	dmsCmd.AddCommand(dms_createReplicationInstanceCmd)
}
