package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_modifyReplicationInstanceCmd = &cobra.Command{
	Use:   "modify-replication-instance",
	Short: "Modifies the replication instance to apply new settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_modifyReplicationInstanceCmd).Standalone()

	dms_modifyReplicationInstanceCmd.Flags().String("allocated-storage", "", "The amount of storage (in gigabytes) to be allocated for the replication instance.")
	dms_modifyReplicationInstanceCmd.Flags().Bool("allow-major-version-upgrade", false, "Indicates that major version upgrades are allowed.")
	dms_modifyReplicationInstanceCmd.Flags().Bool("apply-immediately", false, "Indicates whether the changes should be applied immediately or during the next maintenance window.")
	dms_modifyReplicationInstanceCmd.Flags().String("auto-minor-version-upgrade", "", "A value that indicates that minor version upgrades are applied automatically to the replication instance during the maintenance window.")
	dms_modifyReplicationInstanceCmd.Flags().String("engine-version", "", "The engine version number of the replication instance.")
	dms_modifyReplicationInstanceCmd.Flags().String("kerberos-authentication-settings", "", "Specifies the settings required for kerberos authentication when modifying a replication instance.")
	dms_modifyReplicationInstanceCmd.Flags().String("multi-az", "", "Specifies whether the replication instance is a Multi-AZ deployment.")
	dms_modifyReplicationInstanceCmd.Flags().String("network-type", "", "The type of IP address protocol used by a replication instance, such as IPv4 only or Dual-stack that supports both IPv4 and IPv6 addressing.")
	dms_modifyReplicationInstanceCmd.Flags().Bool("no-allow-major-version-upgrade", false, "Indicates that major version upgrades are allowed.")
	dms_modifyReplicationInstanceCmd.Flags().Bool("no-apply-immediately", false, "Indicates whether the changes should be applied immediately or during the next maintenance window.")
	dms_modifyReplicationInstanceCmd.Flags().String("preferred-maintenance-window", "", "The weekly time range (in UTC) during which system maintenance can occur, which might result in an outage.")
	dms_modifyReplicationInstanceCmd.Flags().String("replication-instance-arn", "", "The Amazon Resource Name (ARN) of the replication instance.")
	dms_modifyReplicationInstanceCmd.Flags().String("replication-instance-class", "", "The compute and memory capacity of the replication instance as defined for the specified replication instance class.")
	dms_modifyReplicationInstanceCmd.Flags().String("replication-instance-identifier", "", "The replication instance identifier.")
	dms_modifyReplicationInstanceCmd.Flags().String("vpc-security-group-ids", "", "Specifies the VPC security group to be used with the replication instance.")
	dms_modifyReplicationInstanceCmd.Flag("no-allow-major-version-upgrade").Hidden = true
	dms_modifyReplicationInstanceCmd.Flag("no-apply-immediately").Hidden = true
	dms_modifyReplicationInstanceCmd.MarkFlagRequired("replication-instance-arn")
	dmsCmd.AddCommand(dms_modifyReplicationInstanceCmd)
}
