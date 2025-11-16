package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyDbclusterCmd = &cobra.Command{
	Use:   "modify-dbcluster",
	Short: "Modifies the settings of an Amazon Aurora DB cluster or a Multi-AZ DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyDbclusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_modifyDbclusterCmd).Standalone()

		rds_modifyDbclusterCmd.Flags().String("allocated-storage", "", "The amount of storage in gibibytes (GiB) to allocate to each DB instance in the Multi-AZ DB cluster.")
		rds_modifyDbclusterCmd.Flags().Bool("allow-engine-mode-change", false, "Specifies whether engine mode changes from `serverless` to `provisioned` are allowed.")
		rds_modifyDbclusterCmd.Flags().Bool("allow-major-version-upgrade", false, "Specifies whether major version upgrades are allowed.")
		rds_modifyDbclusterCmd.Flags().Bool("apply-immediately", false, "Specifies whether the modifications in this request are asynchronously applied as soon as possible, regardless of the `PreferredMaintenanceWindow` setting for the DB cluster.")
		rds_modifyDbclusterCmd.Flags().String("auto-minor-version-upgrade", "", "Specifies whether minor engine upgrades are applied automatically to the DB cluster during the maintenance window.")
		rds_modifyDbclusterCmd.Flags().String("aws-backup-recovery-point-arn", "", "The Amazon Resource Name (ARN) of the recovery point in Amazon Web Services Backup.")
		rds_modifyDbclusterCmd.Flags().String("backtrack-window", "", "The target backtrack window, in seconds.")
		rds_modifyDbclusterCmd.Flags().String("backup-retention-period", "", "The number of days for which automated backups are retained.")
		rds_modifyDbclusterCmd.Flags().String("cacertificate-identifier", "", "The CA certificate identifier to use for the DB cluster's server certificate.")
		rds_modifyDbclusterCmd.Flags().String("cloudwatch-logs-export-configuration", "", "The configuration setting for the log types to be enabled for export to CloudWatch Logs for a specific DB cluster.")
		rds_modifyDbclusterCmd.Flags().String("copy-tags-to-snapshot", "", "Specifies whether to copy all tags from the DB cluster to snapshots of the DB cluster.")
		rds_modifyDbclusterCmd.Flags().String("database-insights-mode", "", "Specifies the mode of Database Insights to enable for the DB cluster.")
		rds_modifyDbclusterCmd.Flags().String("dbcluster-identifier", "", "The DB cluster identifier for the cluster being modified.")
		rds_modifyDbclusterCmd.Flags().String("dbcluster-instance-class", "", "The compute and memory capacity of each DB instance in the Multi-AZ DB cluster, for example `db.m6gd.xlarge`. Not all DB instance classes are available in all Amazon Web Services Regions, or for all database engines.")
		rds_modifyDbclusterCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the DB cluster parameter group to use for the DB cluster.")
		rds_modifyDbclusterCmd.Flags().String("dbinstance-parameter-group-name", "", "The name of the DB parameter group to apply to all instances of the DB cluster.")
		rds_modifyDbclusterCmd.Flags().String("deletion-protection", "", "Specifies whether the DB cluster has deletion protection enabled.")
		rds_modifyDbclusterCmd.Flags().String("domain", "", "The Active Directory directory ID to move the DB cluster to.")
		rds_modifyDbclusterCmd.Flags().String("domain-iamrole-name", "", "The name of the IAM role to use when making API calls to the Directory Service.")
		rds_modifyDbclusterCmd.Flags().String("enable-global-write-forwarding", "", "Specifies whether to enable this DB cluster to forward write operations to the primary cluster of a global cluster (Aurora global database).")
		rds_modifyDbclusterCmd.Flags().String("enable-http-endpoint", "", "Specifies whether to enable the HTTP endpoint for an Aurora Serverless v1 DB cluster.")
		rds_modifyDbclusterCmd.Flags().String("enable-iamdatabase-authentication", "", "Specifies whether to enable mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts.")
		rds_modifyDbclusterCmd.Flags().String("enable-limitless-database", "", "Specifies whether to enable Aurora Limitless Database.")
		rds_modifyDbclusterCmd.Flags().String("enable-local-write-forwarding", "", "Specifies whether read replicas can forward write operations to the writer DB instance in the DB cluster.")
		rds_modifyDbclusterCmd.Flags().String("enable-performance-insights", "", "Specifies whether to turn on Performance Insights for the DB cluster.")
		rds_modifyDbclusterCmd.Flags().String("engine-mode", "", "The DB engine mode of the DB cluster, either `provisioned` or `serverless`.")
		rds_modifyDbclusterCmd.Flags().String("engine-version", "", "The version number of the database engine to which you want to upgrade.")
		rds_modifyDbclusterCmd.Flags().String("iops", "", "The amount of Provisioned IOPS (input/output operations per second) to be initially allocated for each DB instance in the Multi-AZ DB cluster.")
		rds_modifyDbclusterCmd.Flags().String("manage-master-user-password", "", "Specifies whether to manage the master user password with Amazon Web Services Secrets Manager.")
		rds_modifyDbclusterCmd.Flags().String("master-user-authentication-type", "", "Specifies the authentication type for the master user.")
		rds_modifyDbclusterCmd.Flags().String("master-user-password", "", "The new password for the master database user.")
		rds_modifyDbclusterCmd.Flags().String("master-user-secret-kms-key-id", "", "The Amazon Web Services KMS key identifier to encrypt a secret that is automatically generated and managed in Amazon Web Services Secrets Manager.")
		rds_modifyDbclusterCmd.Flags().String("monitoring-interval", "", "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB cluster.")
		rds_modifyDbclusterCmd.Flags().String("monitoring-role-arn", "", "The Amazon Resource Name (ARN) for the IAM role that permits RDS to send Enhanced Monitoring metrics to Amazon CloudWatch Logs.")
		rds_modifyDbclusterCmd.Flags().String("network-type", "", "The network type of the DB cluster.")
		rds_modifyDbclusterCmd.Flags().String("new-dbcluster-identifier", "", "The new DB cluster identifier for the DB cluster when renaming a DB cluster.")
		rds_modifyDbclusterCmd.Flags().Bool("no-allow-engine-mode-change", false, "Specifies whether engine mode changes from `serverless` to `provisioned` are allowed.")
		rds_modifyDbclusterCmd.Flags().Bool("no-allow-major-version-upgrade", false, "Specifies whether major version upgrades are allowed.")
		rds_modifyDbclusterCmd.Flags().Bool("no-apply-immediately", false, "Specifies whether the modifications in this request are asynchronously applied as soon as possible, regardless of the `PreferredMaintenanceWindow` setting for the DB cluster.")
		rds_modifyDbclusterCmd.Flags().String("option-group-name", "", "The option group to associate the DB cluster with.")
		rds_modifyDbclusterCmd.Flags().String("performance-insights-kmskey-id", "", "The Amazon Web Services KMS key identifier for encryption of Performance Insights data.")
		rds_modifyDbclusterCmd.Flags().String("performance-insights-retention-period", "", "The number of days to retain Performance Insights data.")
		rds_modifyDbclusterCmd.Flags().String("port", "", "The port number on which the DB cluster accepts connections.")
		rds_modifyDbclusterCmd.Flags().String("preferred-backup-window", "", "The daily time range during which automated backups are created if automated backups are enabled, using the `BackupRetentionPeriod` parameter.")
		rds_modifyDbclusterCmd.Flags().String("preferred-maintenance-window", "", "The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).")
		rds_modifyDbclusterCmd.Flags().String("rotate-master-user-password", "", "Specifies whether to rotate the secret managed by Amazon Web Services Secrets Manager for the master user password.")
		rds_modifyDbclusterCmd.Flags().String("scaling-configuration", "", "The scaling properties of the DB cluster.")
		rds_modifyDbclusterCmd.Flags().String("serverless-v2-scaling-configuration", "", "")
		rds_modifyDbclusterCmd.Flags().String("storage-type", "", "The storage type to associate with the DB cluster.")
		rds_modifyDbclusterCmd.Flags().String("vpc-security-group-ids", "", "A list of EC2 VPC security groups to associate with this DB cluster.")
		rds_modifyDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
		rds_modifyDbclusterCmd.Flag("no-allow-engine-mode-change").Hidden = true
		rds_modifyDbclusterCmd.Flag("no-allow-major-version-upgrade").Hidden = true
		rds_modifyDbclusterCmd.Flag("no-apply-immediately").Hidden = true
	})
	rdsCmd.AddCommand(rds_modifyDbclusterCmd)
}
