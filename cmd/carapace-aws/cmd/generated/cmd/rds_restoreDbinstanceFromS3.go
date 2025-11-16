package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_restoreDbinstanceFromS3Cmd = &cobra.Command{
	Use:   "restore-dbinstance-from-s3",
	Short: "Amazon Relational Database Service (Amazon RDS) supports importing MySQL databases by using backup files.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_restoreDbinstanceFromS3Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_restoreDbinstanceFromS3Cmd).Standalone()

		rds_restoreDbinstanceFromS3Cmd.Flags().String("allocated-storage", "", "The amount of storage (in gibibytes) to allocate initially for the DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("auto-minor-version-upgrade", "", "Specifies whether to automatically apply minor engine upgrades to the DB instance during the maintenance window.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("availability-zone", "", "The Availability Zone that the DB instance is created in.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("backup-retention-period", "", "The number of days for which automated backups are retained.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("cacertificate-identifier", "", "The CA certificate identifier to use for the DB instance's server certificate.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("copy-tags-to-snapshot", "", "Specifies whether to copy all tags from the DB instance to snapshots of the DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("database-insights-mode", "", "Specifies the mode of Database Insights to enable for the DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("dbinstance-class", "", "The compute and memory capacity of the DB instance, for example db.m4.large. Not all DB instance classes are available in all Amazon Web Services Regions, or for all database engines.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("dbinstance-identifier", "", "The DB instance identifier.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("dbname", "", "The name of the database to create when the DB instance is created.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("dbparameter-group-name", "", "The name of the DB parameter group to associate with this DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("dbsecurity-groups", "", "A list of DB security groups to associate with this DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("dbsubnet-group-name", "", "A DB subnet group to associate with this DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("dedicated-log-volume", "", "Specifies whether to enable a dedicated log volume (DLV) for the DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("deletion-protection", "", "Specifies whether to enable deletion protection for the DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("enable-cloudwatch-logs-exports", "", "The list of logs that the restored DB instance is to export to CloudWatch Logs.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("enable-iamdatabase-authentication", "", "Specifies whether to enable mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("enable-performance-insights", "", "Specifies whether to enable Performance Insights for the DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("engine", "", "The name of the database engine to be used for this instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("engine-lifecycle-support", "", "The life cycle type for this DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("engine-version", "", "The version number of the database engine to use.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("iops", "", "The amount of Provisioned IOPS (input/output operations per second) to allocate initially for the DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("kms-key-id", "", "The Amazon Web Services KMS key identifier for an encrypted DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("license-model", "", "The license model for this DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("manage-master-user-password", "", "Specifies whether to manage the master user password with Amazon Web Services Secrets Manager.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("master-user-password", "", "The password for the master user.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("master-user-secret-kms-key-id", "", "The Amazon Web Services KMS key identifier to encrypt a secret that is automatically generated and managed in Amazon Web Services Secrets Manager.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("master-username", "", "The name for the master user.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("max-allocated-storage", "", "The upper limit in gibibytes (GiB) to which Amazon RDS can automatically scale the storage of the DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("monitoring-interval", "", "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("monitoring-role-arn", "", "The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("multi-az", "", "Specifies whether the DB instance is a Multi-AZ deployment.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("network-type", "", "The network type of the DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("option-group-name", "", "The name of the option group to associate with this DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("performance-insights-kmskey-id", "", "The Amazon Web Services KMS key identifier for encryption of Performance Insights data.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("performance-insights-retention-period", "", "The number of days to retain Performance Insights data.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("port", "", "The port number on which the database accepts connections.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("preferred-backup-window", "", "The time range each day during which automated backups are created if automated backups are enabled.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("preferred-maintenance-window", "", "The time range each week during which system maintenance can occur, in Universal Coordinated Time (UTC).")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("processor-features", "", "The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("publicly-accessible", "", "Specifies whether the DB instance is publicly accessible.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("s3-bucket-name", "", "The name of your Amazon S3 bucket that contains your database backup file.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("s3-ingestion-role-arn", "", "An Amazon Web Services Identity and Access Management (IAM) role with a trust policy and a permissions policy that allows Amazon RDS to access your Amazon S3 bucket.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("s3-prefix", "", "The prefix of your Amazon S3 bucket.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("source-engine", "", "The name of the engine of your source database.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("source-engine-version", "", "The version of the database that the backup files were created from.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("storage-encrypted", "", "Specifies whether the new DB instance is encrypted or not.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("storage-throughput", "", "Specifies the storage throughput value for the DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("storage-type", "", "Specifies the storage type to be associated with the DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("tags", "", "A list of tags to associate with this DB instance.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("use-default-processor-features", "", "Specifies whether the DB instance class of the DB instance uses its default processor features.")
		rds_restoreDbinstanceFromS3Cmd.Flags().String("vpc-security-group-ids", "", "A list of VPC security groups to associate with this DB instance.")
		rds_restoreDbinstanceFromS3Cmd.MarkFlagRequired("dbinstance-class")
		rds_restoreDbinstanceFromS3Cmd.MarkFlagRequired("dbinstance-identifier")
		rds_restoreDbinstanceFromS3Cmd.MarkFlagRequired("engine")
		rds_restoreDbinstanceFromS3Cmd.MarkFlagRequired("s3-bucket-name")
		rds_restoreDbinstanceFromS3Cmd.MarkFlagRequired("s3-ingestion-role-arn")
		rds_restoreDbinstanceFromS3Cmd.MarkFlagRequired("source-engine")
		rds_restoreDbinstanceFromS3Cmd.MarkFlagRequired("source-engine-version")
	})
	rdsCmd.AddCommand(rds_restoreDbinstanceFromS3Cmd)
}
