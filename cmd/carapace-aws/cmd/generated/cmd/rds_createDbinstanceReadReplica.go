package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_createDbinstanceReadReplicaCmd = &cobra.Command{
	Use:   "create-dbinstance-read-replica",
	Short: "Creates a new DB instance that acts as a read replica for an existing source DB instance or Multi-AZ DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_createDbinstanceReadReplicaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_createDbinstanceReadReplicaCmd).Standalone()

		rds_createDbinstanceReadReplicaCmd.Flags().String("allocated-storage", "", "The amount of storage (in gibibytes) to allocate initially for the read replica.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("auto-minor-version-upgrade", "", "Specifies whether to automatically apply minor engine upgrades to the read replica during the maintenance window.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("availability-zone", "", "The Availability Zone (AZ) where the read replica will be created.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("backup-target", "", "The location where RDS stores automated backups and manual snapshots.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("cacertificate-identifier", "", "The CA certificate identifier to use for the read replica's server certificate.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("copy-tags-to-snapshot", "", "Specifies whether to copy all tags from the read replica to snapshots of the read replica.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("custom-iam-instance-profile", "", "The instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("database-insights-mode", "", "The mode of Database Insights to enable for the read replica.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("dbinstance-class", "", "The compute and memory capacity of the read replica, for example db.m4.large. Not all DB instance classes are available in all Amazon Web Services Regions, or for all database engines.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("dbinstance-identifier", "", "The DB instance identifier of the read replica.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("dbparameter-group-name", "", "The name of the DB parameter group to associate with this read replica DB instance.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("dbsubnet-group-name", "", "A DB subnet group for the DB instance.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("dedicated-log-volume", "", "Indicates whether the DB instance has a dedicated log volume (DLV) enabled.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("deletion-protection", "", "Specifies whether to enable deletion protection for the DB instance.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("domain", "", "The Active Directory directory ID to create the DB instance in.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("domain-auth-secret-arn", "", "The ARN for the Secrets Manager secret with the credentials for the user joining the domain.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("domain-dns-ips", "", "The IPv4 DNS IP addresses of your primary and secondary Active Directory domain controllers.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("domain-fqdn", "", "The fully qualified domain name (FQDN) of an Active Directory domain.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("domain-iamrole-name", "", "The name of the IAM role to use when making API calls to the Directory Service.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("domain-ou", "", "The Active Directory organizational unit for your DB instance to join.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("enable-cloudwatch-logs-exports", "", "The list of logs that the new DB instance is to export to CloudWatch Logs.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("enable-customer-owned-ip", "", "Specifies whether to enable a customer-owned IP address (CoIP) for an RDS on Outposts read replica.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("enable-iamdatabase-authentication", "", "Specifies whether to enable mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("enable-performance-insights", "", "Specifies whether to enable Performance Insights for the read replica.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("iops", "", "The amount of Provisioned IOPS (input/output operations per second) to initially allocate for the DB instance.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("kms-key-id", "", "The Amazon Web Services KMS key identifier for an encrypted read replica.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("max-allocated-storage", "", "The upper limit in gibibytes (GiB) to which Amazon RDS can automatically scale the storage of the DB instance.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("monitoring-interval", "", "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the read replica.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("monitoring-role-arn", "", "The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("multi-az", "", "Specifies whether the read replica is in a Multi-AZ deployment.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("network-type", "", "The network type of the DB instance.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("option-group-name", "", "The option group to associate the DB instance with.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("performance-insights-kmskey-id", "", "The Amazon Web Services KMS key identifier for encryption of Performance Insights data.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("performance-insights-retention-period", "", "The number of days to retain Performance Insights data.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("port", "", "The port number that the DB instance uses for connections.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("pre-signed-url", "", "When you are creating a read replica from one Amazon Web Services GovCloud (US) Region to another or from one China Amazon Web Services Region to another, the URL that contains a Signature Version 4 signed request for the `CreateDBInstanceReadReplica` API operation in the source Amazon Web Services Region that contains the source DB instance.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("processor-features", "", "The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("publicly-accessible", "", "Specifies whether the DB instance is publicly accessible.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("replica-mode", "", "The open mode of the replica database.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("source-dbcluster-identifier", "", "The identifier of the Multi-AZ DB cluster that will act as the source for the read replica.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("source-dbinstance-identifier", "", "The identifier of the DB instance that will act as the source for the read replica.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("storage-throughput", "", "Specifies the storage throughput value for the read replica.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("storage-type", "", "The storage type to associate with the read replica.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("tags", "", "")
		rds_createDbinstanceReadReplicaCmd.Flags().String("upgrade-storage-config", "", "Whether to upgrade the storage file system configuration on the read replica.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("use-default-processor-features", "", "Specifies whether the DB instance class of the DB instance uses its default processor features.")
		rds_createDbinstanceReadReplicaCmd.Flags().String("vpc-security-group-ids", "", "A list of Amazon EC2 VPC security groups to associate with the read replica.")
		rds_createDbinstanceReadReplicaCmd.MarkFlagRequired("dbinstance-identifier")
	})
	rdsCmd.AddCommand(rds_createDbinstanceReadReplicaCmd)
}
