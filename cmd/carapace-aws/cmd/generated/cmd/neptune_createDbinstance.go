package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_createDbinstanceCmd = &cobra.Command{
	Use:   "create-dbinstance",
	Short: "Creates a new DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_createDbinstanceCmd).Standalone()

	neptune_createDbinstanceCmd.Flags().String("allocated-storage", "", "Not supported by Neptune.")
	neptune_createDbinstanceCmd.Flags().String("auto-minor-version-upgrade", "", "Indicates that minor engine upgrades are applied automatically to the DB instance during the maintenance window.")
	neptune_createDbinstanceCmd.Flags().String("availability-zone", "", "The EC2 Availability Zone that the DB instance is created in")
	neptune_createDbinstanceCmd.Flags().String("backup-retention-period", "", "The number of days for which automated backups are retained.")
	neptune_createDbinstanceCmd.Flags().String("character-set-name", "", "*(Not supported by Neptune)*")
	neptune_createDbinstanceCmd.Flags().String("copy-tags-to-snapshot", "", "True to copy all tags from the DB instance to snapshots of the DB instance, and otherwise false.")
	neptune_createDbinstanceCmd.Flags().String("dbcluster-identifier", "", "The identifier of the DB cluster that the instance will belong to.")
	neptune_createDbinstanceCmd.Flags().String("dbinstance-class", "", "The compute and memory capacity of the DB instance, for example, `db.m4.large`. Not all DB instance classes are available in all Amazon Regions.")
	neptune_createDbinstanceCmd.Flags().String("dbinstance-identifier", "", "The DB instance identifier.")
	neptune_createDbinstanceCmd.Flags().String("dbname", "", "Not supported.")
	neptune_createDbinstanceCmd.Flags().String("dbparameter-group-name", "", "The name of the DB parameter group to associate with this DB instance.")
	neptune_createDbinstanceCmd.Flags().String("dbsecurity-groups", "", "A list of DB security groups to associate with this DB instance.")
	neptune_createDbinstanceCmd.Flags().String("dbsubnet-group-name", "", "A DB subnet group to associate with this DB instance.")
	neptune_createDbinstanceCmd.Flags().String("deletion-protection", "", "A value that indicates whether the DB instance has deletion protection enabled.")
	neptune_createDbinstanceCmd.Flags().String("domain", "", "Specify the Active Directory Domain to create the instance in.")
	neptune_createDbinstanceCmd.Flags().String("domain-iamrole-name", "", "Specify the name of the IAM role to be used when making API calls to the Directory Service.")
	neptune_createDbinstanceCmd.Flags().String("enable-cloudwatch-logs-exports", "", "The list of log types that need to be enabled for exporting to CloudWatch Logs.")
	neptune_createDbinstanceCmd.Flags().String("enable-iamdatabase-authentication", "", "Not supported by Neptune (ignored).")
	neptune_createDbinstanceCmd.Flags().String("enable-performance-insights", "", "*(Not supported by Neptune)*")
	neptune_createDbinstanceCmd.Flags().String("engine", "", "The name of the database engine to be used for this instance.")
	neptune_createDbinstanceCmd.Flags().String("engine-version", "", "The version number of the database engine to use.")
	neptune_createDbinstanceCmd.Flags().String("iops", "", "The amount of Provisioned IOPS (input/output operations per second) to be initially allocated for the DB instance.")
	neptune_createDbinstanceCmd.Flags().String("kms-key-id", "", "The Amazon KMS key identifier for an encrypted DB instance.")
	neptune_createDbinstanceCmd.Flags().String("license-model", "", "License model information for this DB instance.")
	neptune_createDbinstanceCmd.Flags().String("master-user-password", "", "Not supported by Neptune.")
	neptune_createDbinstanceCmd.Flags().String("master-username", "", "Not supported by Neptune.")
	neptune_createDbinstanceCmd.Flags().String("monitoring-interval", "", "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.")
	neptune_createDbinstanceCmd.Flags().String("monitoring-role-arn", "", "The ARN for the IAM role that permits Neptune to send enhanced monitoring metrics to Amazon CloudWatch Logs.")
	neptune_createDbinstanceCmd.Flags().String("multi-az", "", "Specifies if the DB instance is a Multi-AZ deployment.")
	neptune_createDbinstanceCmd.Flags().String("option-group-name", "", "*(Not supported by Neptune)*")
	neptune_createDbinstanceCmd.Flags().String("performance-insights-kmskey-id", "", "*(Not supported by Neptune)*")
	neptune_createDbinstanceCmd.Flags().String("port", "", "The port number on which the database accepts connections.")
	neptune_createDbinstanceCmd.Flags().String("preferred-backup-window", "", "The daily time range during which automated backups are created.")
	neptune_createDbinstanceCmd.Flags().String("preferred-maintenance-window", "", "The time range each week during which system maintenance can occur, in Universal Coordinated Time (UTC).")
	neptune_createDbinstanceCmd.Flags().String("promotion-tier", "", "A value that specifies the order in which an Read Replica is promoted to the primary instance after a failure of the existing primary instance.")
	neptune_createDbinstanceCmd.Flags().String("publicly-accessible", "", "Indicates whether the DB instance is publicly accessible.")
	neptune_createDbinstanceCmd.Flags().String("storage-encrypted", "", "Specifies whether the DB instance is encrypted.")
	neptune_createDbinstanceCmd.Flags().String("storage-type", "", "Not applicable.")
	neptune_createDbinstanceCmd.Flags().String("tags", "", "The tags to assign to the new instance.")
	neptune_createDbinstanceCmd.Flags().String("tde-credential-arn", "", "The ARN from the key store with which to associate the instance for TDE encryption.")
	neptune_createDbinstanceCmd.Flags().String("tde-credential-password", "", "The password for the given ARN from the key store in order to access the device.")
	neptune_createDbinstanceCmd.Flags().String("timezone", "", "The time zone of the DB instance.")
	neptune_createDbinstanceCmd.Flags().String("vpc-security-group-ids", "", "A list of EC2 VPC security groups to associate with this DB instance.")
	neptune_createDbinstanceCmd.MarkFlagRequired("dbcluster-identifier")
	neptune_createDbinstanceCmd.MarkFlagRequired("dbinstance-class")
	neptune_createDbinstanceCmd.MarkFlagRequired("dbinstance-identifier")
	neptune_createDbinstanceCmd.MarkFlagRequired("engine")
	neptuneCmd.AddCommand(neptune_createDbinstanceCmd)
}
