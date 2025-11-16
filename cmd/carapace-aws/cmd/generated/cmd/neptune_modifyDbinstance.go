package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_modifyDbinstanceCmd = &cobra.Command{
	Use:   "modify-dbinstance",
	Short: "Modifies settings for a DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_modifyDbinstanceCmd).Standalone()

	neptune_modifyDbinstanceCmd.Flags().String("allocated-storage", "", "Not supported by Neptune.")
	neptune_modifyDbinstanceCmd.Flags().Bool("allow-major-version-upgrade", false, "Indicates that major version upgrades are allowed.")
	neptune_modifyDbinstanceCmd.Flags().Bool("apply-immediately", false, "Specifies whether the modifications in this request and any pending modifications are asynchronously applied as soon as possible, regardless of the `PreferredMaintenanceWindow` setting for the DB instance.")
	neptune_modifyDbinstanceCmd.Flags().String("auto-minor-version-upgrade", "", "Indicates that minor version upgrades are applied automatically to the DB instance during the maintenance window.")
	neptune_modifyDbinstanceCmd.Flags().String("backup-retention-period", "", "Not applicable.")
	neptune_modifyDbinstanceCmd.Flags().String("cacertificate-identifier", "", "Indicates the certificate that needs to be associated with the instance.")
	neptune_modifyDbinstanceCmd.Flags().String("cloudwatch-logs-export-configuration", "", "The configuration setting for the log types to be enabled for export to CloudWatch Logs for a specific DB instance or DB cluster.")
	neptune_modifyDbinstanceCmd.Flags().String("copy-tags-to-snapshot", "", "True to copy all tags from the DB instance to snapshots of the DB instance, and otherwise false.")
	neptune_modifyDbinstanceCmd.Flags().String("dbinstance-class", "", "The new compute and memory capacity of the DB instance, for example, `db.m4.large`. Not all DB instance classes are available in all Amazon Regions.")
	neptune_modifyDbinstanceCmd.Flags().String("dbinstance-identifier", "", "The DB instance identifier.")
	neptune_modifyDbinstanceCmd.Flags().String("dbparameter-group-name", "", "The name of the DB parameter group to apply to the DB instance.")
	neptune_modifyDbinstanceCmd.Flags().String("dbport-number", "", "The port number on which the database accepts connections.")
	neptune_modifyDbinstanceCmd.Flags().String("dbsecurity-groups", "", "A list of DB security groups to authorize on this DB instance.")
	neptune_modifyDbinstanceCmd.Flags().String("dbsubnet-group-name", "", "The new DB subnet group for the DB instance.")
	neptune_modifyDbinstanceCmd.Flags().String("deletion-protection", "", "A value that indicates whether the DB instance has deletion protection enabled.")
	neptune_modifyDbinstanceCmd.Flags().String("domain", "", "Not supported.")
	neptune_modifyDbinstanceCmd.Flags().String("domain-iamrole-name", "", "Not supported")
	neptune_modifyDbinstanceCmd.Flags().String("enable-iamdatabase-authentication", "", "True to enable mapping of Amazon Identity and Access Management (IAM) accounts to database accounts, and otherwise false.")
	neptune_modifyDbinstanceCmd.Flags().String("enable-performance-insights", "", "*(Not supported by Neptune)*")
	neptune_modifyDbinstanceCmd.Flags().String("engine-version", "", "The version number of the database engine to upgrade to.")
	neptune_modifyDbinstanceCmd.Flags().String("iops", "", "The new Provisioned IOPS (I/O operations per second) value for the instance.")
	neptune_modifyDbinstanceCmd.Flags().String("license-model", "", "Not supported by Neptune.")
	neptune_modifyDbinstanceCmd.Flags().String("master-user-password", "", "Not supported by Neptune.")
	neptune_modifyDbinstanceCmd.Flags().String("monitoring-interval", "", "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.")
	neptune_modifyDbinstanceCmd.Flags().String("monitoring-role-arn", "", "The ARN for the IAM role that permits Neptune to send enhanced monitoring metrics to Amazon CloudWatch Logs.")
	neptune_modifyDbinstanceCmd.Flags().String("multi-az", "", "Specifies if the DB instance is a Multi-AZ deployment.")
	neptune_modifyDbinstanceCmd.Flags().String("new-dbinstance-identifier", "", "The new DB instance identifier for the DB instance when renaming a DB instance.")
	neptune_modifyDbinstanceCmd.Flags().Bool("no-allow-major-version-upgrade", false, "Indicates that major version upgrades are allowed.")
	neptune_modifyDbinstanceCmd.Flags().Bool("no-apply-immediately", false, "Specifies whether the modifications in this request and any pending modifications are asynchronously applied as soon as possible, regardless of the `PreferredMaintenanceWindow` setting for the DB instance.")
	neptune_modifyDbinstanceCmd.Flags().String("option-group-name", "", "*(Not supported by Neptune)*")
	neptune_modifyDbinstanceCmd.Flags().String("performance-insights-kmskey-id", "", "*(Not supported by Neptune)*")
	neptune_modifyDbinstanceCmd.Flags().String("preferred-backup-window", "", "The daily time range during which automated backups are created if automated backups are enabled.")
	neptune_modifyDbinstanceCmd.Flags().String("preferred-maintenance-window", "", "The weekly time range (in UTC) during which system maintenance can occur, which might result in an outage.")
	neptune_modifyDbinstanceCmd.Flags().String("promotion-tier", "", "A value that specifies the order in which a Read Replica is promoted to the primary instance after a failure of the existing primary instance.")
	neptune_modifyDbinstanceCmd.Flags().String("publicly-accessible", "", "Indicates whether the DB instance is publicly accessible.")
	neptune_modifyDbinstanceCmd.Flags().String("storage-type", "", "Not applicable.")
	neptune_modifyDbinstanceCmd.Flags().String("tde-credential-arn", "", "The ARN from the key store with which to associate the instance for TDE encryption.")
	neptune_modifyDbinstanceCmd.Flags().String("tde-credential-password", "", "The password for the given ARN from the key store in order to access the device.")
	neptune_modifyDbinstanceCmd.Flags().String("vpc-security-group-ids", "", "A list of EC2 VPC security groups to authorize on this DB instance.")
	neptune_modifyDbinstanceCmd.MarkFlagRequired("dbinstance-identifier")
	neptune_modifyDbinstanceCmd.Flag("no-allow-major-version-upgrade").Hidden = true
	neptune_modifyDbinstanceCmd.Flag("no-apply-immediately").Hidden = true
	neptuneCmd.AddCommand(neptune_modifyDbinstanceCmd)
}
