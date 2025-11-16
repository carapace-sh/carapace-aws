package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_restoreDbinstanceToPointInTimeCmd = &cobra.Command{
	Use:   "restore-dbinstance-to-point-in-time",
	Short: "Restores a DB instance to an arbitrary point in time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_restoreDbinstanceToPointInTimeCmd).Standalone()

	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("allocated-storage", "", "The amount of storage (in gibibytes) to allocate initially for the DB instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("auto-minor-version-upgrade", "", "Specifies whether minor version upgrades are applied automatically to the DB instance during the maintenance window.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("availability-zone", "", "The Availability Zone (AZ) where the DB instance will be created.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("backup-target", "", "The location for storing automated backups and manual snapshots for the restored DB instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("cacertificate-identifier", "", "The CA certificate identifier to use for the DB instance's server certificate.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("copy-tags-to-snapshot", "", "Specifies whether to copy all tags from the restored DB instance to snapshots of the DB instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("custom-iam-instance-profile", "", "The instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("dbinstance-class", "", "The compute and memory capacity of the Amazon RDS DB instance, for example db.m4.large. Not all DB instance classes are available in all Amazon Web Services Regions, or for all database engines.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("dbname", "", "The database name for the restored DB instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("dbparameter-group-name", "", "The name of the DB parameter group to associate with this DB instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("dbsubnet-group-name", "", "The DB subnet group name to use for the new instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("dedicated-log-volume", "", "Specifies whether to enable a dedicated log volume (DLV) for the DB instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("deletion-protection", "", "Specifies whether the DB instance has deletion protection enabled.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("domain", "", "The Active Directory directory ID to restore the DB instance in.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("domain-auth-secret-arn", "", "The ARN for the Secrets Manager secret with the credentials for the user joining the domain.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("domain-dns-ips", "", "The IPv4 DNS IP addresses of your primary and secondary Active Directory domain controllers.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("domain-fqdn", "", "The fully qualified domain name (FQDN) of an Active Directory domain.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("domain-iamrole-name", "", "The name of the IAM role to use when making API calls to the Directory Service.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("domain-ou", "", "The Active Directory organizational unit for your DB instance to join.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("enable-cloudwatch-logs-exports", "", "The list of logs that the restored DB instance is to export to CloudWatch Logs.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("enable-customer-owned-ip", "", "Specifies whether to enable a customer-owned IP address (CoIP) for an RDS on Outposts DB instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("enable-iamdatabase-authentication", "", "Specifies whether to enable mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("engine", "", "The database engine to use for the new instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("engine-lifecycle-support", "", "The life cycle type for this DB instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("iops", "", "The amount of Provisioned IOPS (input/output operations per second) to initially allocate for the DB instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("license-model", "", "The license model information for the restored DB instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("manage-master-user-password", "", "Specifies whether to manage the master user password with Amazon Web Services Secrets Manager in the restored DB instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("master-user-secret-kms-key-id", "", "The Amazon Web Services KMS key identifier to encrypt a secret that is automatically generated and managed in Amazon Web Services Secrets Manager.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("max-allocated-storage", "", "The upper limit in gibibytes (GiB) to which Amazon RDS can automatically scale the storage of the DB instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("multi-az", "", "Secifies whether the DB instance is a Multi-AZ deployment.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("network-type", "", "The network type of the DB instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().Bool("no-use-latest-restorable-time", false, "Specifies whether the DB instance is restored from the latest backup time.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("option-group-name", "", "The name of the option group to use for the restored DB instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("port", "", "The port number on which the database accepts connections.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("processor-features", "", "The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("publicly-accessible", "", "Specifies whether the DB instance is publicly accessible.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("restore-time", "", "The date and time to restore from.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("source-dbi-resource-id", "", "The resource ID of the source DB instance from which to restore.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("source-dbinstance-automated-backups-arn", "", "The Amazon Resource Name (ARN) of the replicated automated backups from which to restore, for example, `arn:aws:rds:us-east-1:123456789012:auto-backup:ab-L2IJCEXJP7XQ7HOJ4SIEXAMPLE`.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("source-dbinstance-identifier", "", "The identifier of the source DB instance from which to restore.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("storage-throughput", "", "The storage throughput value for the DB instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("storage-type", "", "The storage type to associate with the DB instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("tags", "", "")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("target-dbinstance-identifier", "", "The name of the new DB instance to create.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("tde-credential-arn", "", "The ARN from the key store with which to associate the instance for TDE encryption.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("tde-credential-password", "", "The password for the given ARN from the key store in order to access the device.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("use-default-processor-features", "", "Specifies whether the DB instance class of the DB instance uses its default processor features.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().Bool("use-latest-restorable-time", false, "Specifies whether the DB instance is restored from the latest backup time.")
	rds_restoreDbinstanceToPointInTimeCmd.Flags().String("vpc-security-group-ids", "", "A list of EC2 VPC security groups to associate with this DB instance.")
	rds_restoreDbinstanceToPointInTimeCmd.Flag("no-use-latest-restorable-time").Hidden = true
	rds_restoreDbinstanceToPointInTimeCmd.MarkFlagRequired("target-dbinstance-identifier")
	rdsCmd.AddCommand(rds_restoreDbinstanceToPointInTimeCmd)
}
