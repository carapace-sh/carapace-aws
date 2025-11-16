package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_restoreDbinstanceFromDbsnapshotCmd = &cobra.Command{
	Use:   "restore-dbinstance-from-dbsnapshot",
	Short: "Creates a new DB instance from a DB snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_restoreDbinstanceFromDbsnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_restoreDbinstanceFromDbsnapshotCmd).Standalone()

		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("allocated-storage", "", "The amount of storage (in gibibytes) to allocate initially for the DB instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("auto-minor-version-upgrade", "", "Specifies whether to automatically apply minor version upgrades to the DB instance during the maintenance window.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("availability-zone", "", "The Availability Zone (AZ) where the DB instance will be created.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("backup-target", "", "Specifies where automated backups and manual snapshots are stored for the restored DB instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("cacertificate-identifier", "", "The CA certificate identifier to use for the DB instance's server certificate.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("copy-tags-to-snapshot", "", "Specifies whether to copy all tags from the restored DB instance to snapshots of the DB instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("custom-iam-instance-profile", "", "The instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("dbcluster-snapshot-identifier", "", "The identifier for the Multi-AZ DB cluster snapshot to restore from.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("dbinstance-class", "", "The compute and memory capacity of the Amazon RDS DB instance, for example db.m4.large. Not all DB instance classes are available in all Amazon Web Services Regions, or for all database engines.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("dbinstance-identifier", "", "The name of the DB instance to create from the DB snapshot.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("dbname", "", "The name of the database for the restored DB instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("dbparameter-group-name", "", "The name of the DB parameter group to associate with this DB instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("dbsnapshot-identifier", "", "The identifier for the DB snapshot to restore from.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("dbsubnet-group-name", "", "The name of the DB subnet group to use for the new instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("dedicated-log-volume", "", "Specifies whether to enable a dedicated log volume (DLV) for the DB instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("deletion-protection", "", "Specifies whether to enable deletion protection for the DB instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("domain", "", "The Active Directory directory ID to restore the DB instance in.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("domain-auth-secret-arn", "", "The ARN for the Secrets Manager secret with the credentials for the user joining the domain.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("domain-dns-ips", "", "The IPv4 DNS IP addresses of your primary and secondary Active Directory domain controllers.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("domain-fqdn", "", "The fully qualified domain name (FQDN) of an Active Directory domain.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("domain-iamrole-name", "", "The name of the IAM role to use when making API calls to the Directory Service.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("domain-ou", "", "The Active Directory organizational unit for your DB instance to join.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("enable-cloudwatch-logs-exports", "", "The list of logs for the restored DB instance to export to CloudWatch Logs.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("enable-customer-owned-ip", "", "Specifies whether to enable a customer-owned IP address (CoIP) for an RDS on Outposts DB instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("enable-iamdatabase-authentication", "", "Specifies whether to enable mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("engine", "", "The database engine to use for the new instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("engine-lifecycle-support", "", "The life cycle type for this DB instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("iops", "", "Specifies the amount of provisioned IOPS for the DB instance, expressed in I/O operations per second.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("license-model", "", "License model information for the restored DB instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("manage-master-user-password", "", "Specifies whether to manage the master user password with Amazon Web Services Secrets Manager in the restored DB instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("master-user-secret-kms-key-id", "", "The Amazon Web Services KMS key identifier to encrypt a secret that is automatically generated and managed in Amazon Web Services Secrets Manager.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("multi-az", "", "Specifies whether the DB instance is a Multi-AZ deployment.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("network-type", "", "The network type of the DB instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("option-group-name", "", "The name of the option group to be used for the restored DB instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("port", "", "The port number on which the database accepts connections.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("processor-features", "", "The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("publicly-accessible", "", "Specifies whether the DB instance is publicly accessible.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("storage-throughput", "", "Specifies the storage throughput value for the DB instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("storage-type", "", "Specifies the storage type to be associated with the DB instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("tags", "", "")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("tde-credential-arn", "", "The ARN from the key store with which to associate the instance for TDE encryption.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("tde-credential-password", "", "The password for the given ARN from the key store in order to access the device.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("use-default-processor-features", "", "Specifies whether the DB instance class of the DB instance uses its default processor features.")
		rds_restoreDbinstanceFromDbsnapshotCmd.Flags().String("vpc-security-group-ids", "", "A list of EC2 VPC security groups to associate with this DB instance.")
		rds_restoreDbinstanceFromDbsnapshotCmd.MarkFlagRequired("dbinstance-identifier")
	})
	rdsCmd.AddCommand(rds_restoreDbinstanceFromDbsnapshotCmd)
}
