package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_createDbclusterCmd = &cobra.Command{
	Use:   "create-dbcluster",
	Short: "Creates a new Amazon Aurora DB cluster or Multi-AZ DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_createDbclusterCmd).Standalone()

	rds_createDbclusterCmd.Flags().String("allocated-storage", "", "The amount of storage in gibibytes (GiB) to allocate to each DB instance in the Multi-AZ DB cluster.")
	rds_createDbclusterCmd.Flags().String("auto-minor-version-upgrade", "", "Specifies whether minor engine upgrades are applied automatically to the DB cluster during the maintenance window.")
	rds_createDbclusterCmd.Flags().String("availability-zones", "", "A list of Availability Zones (AZs) where you specifically want to create DB instances in the DB cluster.")
	rds_createDbclusterCmd.Flags().String("backtrack-window", "", "The target backtrack window, in seconds.")
	rds_createDbclusterCmd.Flags().String("backup-retention-period", "", "The number of days for which automated backups are retained.")
	rds_createDbclusterCmd.Flags().String("cacertificate-identifier", "", "The CA certificate identifier to use for the DB cluster's server certificate.")
	rds_createDbclusterCmd.Flags().String("character-set-name", "", "The name of the character set (`CharacterSet`) to associate the DB cluster with.")
	rds_createDbclusterCmd.Flags().String("cluster-scalability-type", "", "Specifies the scalability mode of the Aurora DB cluster.")
	rds_createDbclusterCmd.Flags().String("copy-tags-to-snapshot", "", "Specifies whether to copy all tags from the DB cluster to snapshots of the DB cluster.")
	rds_createDbclusterCmd.Flags().String("database-insights-mode", "", "The mode of Database Insights to enable for the DB cluster.")
	rds_createDbclusterCmd.Flags().String("database-name", "", "The name for your database of up to 64 alphanumeric characters.")
	rds_createDbclusterCmd.Flags().String("dbcluster-identifier", "", "The identifier for this DB cluster.")
	rds_createDbclusterCmd.Flags().String("dbcluster-instance-class", "", "The compute and memory capacity of each DB instance in the Multi-AZ DB cluster, for example `db.m6gd.xlarge`. Not all DB instance classes are available in all Amazon Web Services Regions, or for all database engines.")
	rds_createDbclusterCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the DB cluster parameter group to associate with this DB cluster.")
	rds_createDbclusterCmd.Flags().String("dbsubnet-group-name", "", "A DB subnet group to associate with this DB cluster.")
	rds_createDbclusterCmd.Flags().String("dbsystem-id", "", "Reserved for future use.")
	rds_createDbclusterCmd.Flags().String("deletion-protection", "", "Specifies whether the DB cluster has deletion protection enabled.")
	rds_createDbclusterCmd.Flags().String("domain", "", "The Active Directory directory ID to create the DB cluster in.")
	rds_createDbclusterCmd.Flags().String("domain-iamrole-name", "", "The name of the IAM role to use when making API calls to the Directory Service.")
	rds_createDbclusterCmd.Flags().String("enable-cloudwatch-logs-exports", "", "The list of log types that need to be enabled for exporting to CloudWatch Logs.")
	rds_createDbclusterCmd.Flags().String("enable-global-write-forwarding", "", "Specifies whether to enable this DB cluster to forward write operations to the primary cluster of a global cluster (Aurora global database).")
	rds_createDbclusterCmd.Flags().String("enable-http-endpoint", "", "Specifies whether to enable the HTTP endpoint for the DB cluster.")
	rds_createDbclusterCmd.Flags().String("enable-iamdatabase-authentication", "", "Specifies whether to enable mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts.")
	rds_createDbclusterCmd.Flags().String("enable-limitless-database", "", "Specifies whether to enable Aurora Limitless Database.")
	rds_createDbclusterCmd.Flags().String("enable-local-write-forwarding", "", "Specifies whether read replicas can forward write operations to the writer DB instance in the DB cluster.")
	rds_createDbclusterCmd.Flags().String("enable-performance-insights", "", "Specifies whether to turn on Performance Insights for the DB cluster.")
	rds_createDbclusterCmd.Flags().String("engine", "", "The database engine to use for this DB cluster.")
	rds_createDbclusterCmd.Flags().String("engine-lifecycle-support", "", "The life cycle type for this DB cluster.")
	rds_createDbclusterCmd.Flags().String("engine-mode", "", "The DB engine mode of the DB cluster, either `provisioned` or `serverless`.")
	rds_createDbclusterCmd.Flags().String("engine-version", "", "The version number of the database engine to use.")
	rds_createDbclusterCmd.Flags().String("global-cluster-identifier", "", "The global cluster ID of an Aurora cluster that becomes the primary cluster in the new global database cluster.")
	rds_createDbclusterCmd.Flags().String("iops", "", "The amount of Provisioned IOPS (input/output operations per second) to be initially allocated for each DB instance in the Multi-AZ DB cluster.")
	rds_createDbclusterCmd.Flags().String("kms-key-id", "", "The Amazon Web Services KMS key identifier for an encrypted DB cluster.")
	rds_createDbclusterCmd.Flags().String("manage-master-user-password", "", "Specifies whether to manage the master user password with Amazon Web Services Secrets Manager.")
	rds_createDbclusterCmd.Flags().String("master-user-authentication-type", "", "Specifies the authentication type for the master user.")
	rds_createDbclusterCmd.Flags().String("master-user-password", "", "The password for the master database user.")
	rds_createDbclusterCmd.Flags().String("master-user-secret-kms-key-id", "", "The Amazon Web Services KMS key identifier to encrypt a secret that is automatically generated and managed in Amazon Web Services Secrets Manager.")
	rds_createDbclusterCmd.Flags().String("master-username", "", "The name of the master user for the DB cluster.")
	rds_createDbclusterCmd.Flags().String("monitoring-interval", "", "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB cluster.")
	rds_createDbclusterCmd.Flags().String("monitoring-role-arn", "", "The Amazon Resource Name (ARN) for the IAM role that permits RDS to send Enhanced Monitoring metrics to Amazon CloudWatch Logs.")
	rds_createDbclusterCmd.Flags().String("network-type", "", "The network type of the DB cluster.")
	rds_createDbclusterCmd.Flags().String("option-group-name", "", "The option group to associate the DB cluster with.")
	rds_createDbclusterCmd.Flags().String("performance-insights-kmskey-id", "", "The Amazon Web Services KMS key identifier for encryption of Performance Insights data.")
	rds_createDbclusterCmd.Flags().String("performance-insights-retention-period", "", "The number of days to retain Performance Insights data.")
	rds_createDbclusterCmd.Flags().String("port", "", "The port number on which the instances in the DB cluster accept connections.")
	rds_createDbclusterCmd.Flags().String("pre-signed-url", "", "When you are replicating a DB cluster from one Amazon Web Services GovCloud (US) Region to another, an URL that contains a Signature Version 4 signed request for the `CreateDBCluster` operation to be called in the source Amazon Web Services Region where the DB cluster is replicated from.")
	rds_createDbclusterCmd.Flags().String("preferred-backup-window", "", "The daily time range during which automated backups are created if automated backups are enabled using the `BackupRetentionPeriod` parameter.")
	rds_createDbclusterCmd.Flags().String("preferred-maintenance-window", "", "The weekly time range during which system maintenance can occur.")
	rds_createDbclusterCmd.Flags().String("publicly-accessible", "", "Specifies whether the DB cluster is publicly accessible.")
	rds_createDbclusterCmd.Flags().String("rds-custom-cluster-configuration", "", "Reserved for future use.")
	rds_createDbclusterCmd.Flags().String("replication-source-identifier", "", "The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this DB cluster is created as a read replica.")
	rds_createDbclusterCmd.Flags().String("scaling-configuration", "", "For DB clusters in `serverless` DB engine mode, the scaling properties of the DB cluster.")
	rds_createDbclusterCmd.Flags().String("serverless-v2-scaling-configuration", "", "")
	rds_createDbclusterCmd.Flags().String("storage-encrypted", "", "Specifies whether the DB cluster is encrypted.")
	rds_createDbclusterCmd.Flags().String("storage-type", "", "The storage type to associate with the DB cluster.")
	rds_createDbclusterCmd.Flags().String("tags", "", "Tags to assign to the DB cluster.")
	rds_createDbclusterCmd.Flags().String("vpc-security-group-ids", "", "A list of EC2 VPC security groups to associate with this DB cluster.")
	rds_createDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
	rds_createDbclusterCmd.MarkFlagRequired("engine")
	rdsCmd.AddCommand(rds_createDbclusterCmd)
}
