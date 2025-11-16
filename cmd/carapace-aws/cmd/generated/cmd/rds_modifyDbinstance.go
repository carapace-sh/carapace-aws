package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyDbinstanceCmd = &cobra.Command{
	Use:   "modify-dbinstance",
	Short: "Modifies settings for a DB instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyDbinstanceCmd).Standalone()

	rds_modifyDbinstanceCmd.Flags().String("allocated-storage", "", "The new amount of storage in gibibytes (GiB) to allocate for the DB instance.")
	rds_modifyDbinstanceCmd.Flags().Bool("allow-major-version-upgrade", false, "Specifies whether major version upgrades are allowed.")
	rds_modifyDbinstanceCmd.Flags().Bool("apply-immediately", false, "Specifies whether the modifications in this request and any pending modifications are asynchronously applied as soon as possible, regardless of the `PreferredMaintenanceWindow` setting for the DB instance.")
	rds_modifyDbinstanceCmd.Flags().String("auto-minor-version-upgrade", "", "Specifies whether minor version upgrades are applied automatically to the DB instance during the maintenance window.")
	rds_modifyDbinstanceCmd.Flags().String("automation-mode", "", "The automation mode of the RDS Custom DB instance.")
	rds_modifyDbinstanceCmd.Flags().String("aws-backup-recovery-point-arn", "", "The Amazon Resource Name (ARN) of the recovery point in Amazon Web Services Backup.")
	rds_modifyDbinstanceCmd.Flags().String("backup-retention-period", "", "The number of days to retain automated backups.")
	rds_modifyDbinstanceCmd.Flags().String("cacertificate-identifier", "", "The CA certificate identifier to use for the DB instance's server certificate.")
	rds_modifyDbinstanceCmd.Flags().String("certificate-rotation-restart", "", "Specifies whether the DB instance is restarted when you rotate your SSL/TLS certificate.")
	rds_modifyDbinstanceCmd.Flags().String("cloudwatch-logs-export-configuration", "", "The log types to be enabled for export to CloudWatch Logs for a specific DB instance.")
	rds_modifyDbinstanceCmd.Flags().String("copy-tags-to-snapshot", "", "Specifies whether to copy all tags from the DB instance to snapshots of the DB instance.")
	rds_modifyDbinstanceCmd.Flags().String("database-insights-mode", "", "Specifies the mode of Database Insights to enable for the DB instance.")
	rds_modifyDbinstanceCmd.Flags().String("dbinstance-class", "", "The new compute and memory capacity of the DB instance, for example `db.m4.large`. Not all DB instance classes are available in all Amazon Web Services Regions, or for all database engines.")
	rds_modifyDbinstanceCmd.Flags().String("dbinstance-identifier", "", "The identifier of DB instance to modify.")
	rds_modifyDbinstanceCmd.Flags().String("dbparameter-group-name", "", "The name of the DB parameter group to apply to the DB instance.")
	rds_modifyDbinstanceCmd.Flags().String("dbport-number", "", "The port number on which the database accepts connections.")
	rds_modifyDbinstanceCmd.Flags().String("dbsecurity-groups", "", "A list of DB security groups to authorize on this DB instance.")
	rds_modifyDbinstanceCmd.Flags().String("dbsubnet-group-name", "", "The new DB subnet group for the DB instance.")
	rds_modifyDbinstanceCmd.Flags().String("dedicated-log-volume", "", "Indicates whether the DB instance has a dedicated log volume (DLV) enabled.")
	rds_modifyDbinstanceCmd.Flags().String("deletion-protection", "", "Specifies whether the DB instance has deletion protection enabled.")
	rds_modifyDbinstanceCmd.Flags().String("disable-domain", "", "Specifies whether to remove the DB instance from the Active Directory domain.")
	rds_modifyDbinstanceCmd.Flags().String("domain", "", "The Active Directory directory ID to move the DB instance to.")
	rds_modifyDbinstanceCmd.Flags().String("domain-auth-secret-arn", "", "The ARN for the Secrets Manager secret with the credentials for the user joining the domain.")
	rds_modifyDbinstanceCmd.Flags().String("domain-dns-ips", "", "The IPv4 DNS IP addresses of your primary and secondary Active Directory domain controllers.")
	rds_modifyDbinstanceCmd.Flags().String("domain-fqdn", "", "The fully qualified domain name (FQDN) of an Active Directory domain.")
	rds_modifyDbinstanceCmd.Flags().String("domain-iamrole-name", "", "The name of the IAM role to use when making API calls to the Directory Service.")
	rds_modifyDbinstanceCmd.Flags().String("domain-ou", "", "The Active Directory organizational unit for your DB instance to join.")
	rds_modifyDbinstanceCmd.Flags().String("enable-customer-owned-ip", "", "Specifies whether to enable a customer-owned IP address (CoIP) for an RDS on Outposts DB instance.")
	rds_modifyDbinstanceCmd.Flags().String("enable-iamdatabase-authentication", "", "Specifies whether to enable mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts.")
	rds_modifyDbinstanceCmd.Flags().String("enable-performance-insights", "", "Specifies whether to enable Performance Insights for the DB instance.")
	rds_modifyDbinstanceCmd.Flags().String("engine", "", "The target Oracle DB engine when you convert a non-CDB to a CDB.")
	rds_modifyDbinstanceCmd.Flags().String("engine-version", "", "The version number of the database engine to upgrade to.")
	rds_modifyDbinstanceCmd.Flags().String("iops", "", "The new Provisioned IOPS (I/O operations per second) value for the RDS instance.")
	rds_modifyDbinstanceCmd.Flags().String("license-model", "", "The license model for the DB instance.")
	rds_modifyDbinstanceCmd.Flags().String("manage-master-user-password", "", "Specifies whether to manage the master user password with Amazon Web Services Secrets Manager.")
	rds_modifyDbinstanceCmd.Flags().String("master-user-authentication-type", "", "Specifies the authentication type for the master user.")
	rds_modifyDbinstanceCmd.Flags().String("master-user-password", "", "The new password for the master user.")
	rds_modifyDbinstanceCmd.Flags().String("master-user-secret-kms-key-id", "", "The Amazon Web Services KMS key identifier to encrypt a secret that is automatically generated and managed in Amazon Web Services Secrets Manager.")
	rds_modifyDbinstanceCmd.Flags().String("max-allocated-storage", "", "The upper limit in gibibytes (GiB) to which Amazon RDS can automatically scale the storage of the DB instance.")
	rds_modifyDbinstanceCmd.Flags().String("monitoring-interval", "", "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.")
	rds_modifyDbinstanceCmd.Flags().String("monitoring-role-arn", "", "The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs.")
	rds_modifyDbinstanceCmd.Flags().String("multi-az", "", "Specifies whether the DB instance is a Multi-AZ deployment.")
	rds_modifyDbinstanceCmd.Flags().String("multi-tenant", "", "Specifies whether the to convert your DB instance from the single-tenant conﬁguration to the multi-tenant conﬁguration.")
	rds_modifyDbinstanceCmd.Flags().String("network-type", "", "The network type of the DB instance.")
	rds_modifyDbinstanceCmd.Flags().String("new-dbinstance-identifier", "", "The new identifier for the DB instance when renaming a DB instance.")
	rds_modifyDbinstanceCmd.Flags().Bool("no-allow-major-version-upgrade", false, "Specifies whether major version upgrades are allowed.")
	rds_modifyDbinstanceCmd.Flags().Bool("no-apply-immediately", false, "Specifies whether the modifications in this request and any pending modifications are asynchronously applied as soon as possible, regardless of the `PreferredMaintenanceWindow` setting for the DB instance.")
	rds_modifyDbinstanceCmd.Flags().String("option-group-name", "", "The option group to associate the DB instance with.")
	rds_modifyDbinstanceCmd.Flags().String("performance-insights-kmskey-id", "", "The Amazon Web Services KMS key identifier for encryption of Performance Insights data.")
	rds_modifyDbinstanceCmd.Flags().String("performance-insights-retention-period", "", "The number of days to retain Performance Insights data.")
	rds_modifyDbinstanceCmd.Flags().String("preferred-backup-window", "", "The daily time range during which automated backups are created if automated backups are enabled, as determined by the `BackupRetentionPeriod` parameter.")
	rds_modifyDbinstanceCmd.Flags().String("preferred-maintenance-window", "", "The weekly time range during which system maintenance can occur, which might result in an outage.")
	rds_modifyDbinstanceCmd.Flags().String("processor-features", "", "The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.")
	rds_modifyDbinstanceCmd.Flags().String("promotion-tier", "", "The order of priority in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance.")
	rds_modifyDbinstanceCmd.Flags().String("publicly-accessible", "", "Specifies whether the DB instance is publicly accessible.")
	rds_modifyDbinstanceCmd.Flags().String("replica-mode", "", "The open mode of a replica database.")
	rds_modifyDbinstanceCmd.Flags().String("resume-full-automation-mode-minutes", "", "The number of minutes to pause the automation.")
	rds_modifyDbinstanceCmd.Flags().String("rotate-master-user-password", "", "Specifies whether to rotate the secret managed by Amazon Web Services Secrets Manager for the master user password.")
	rds_modifyDbinstanceCmd.Flags().String("storage-throughput", "", "The storage throughput value for the DB instance.")
	rds_modifyDbinstanceCmd.Flags().String("storage-type", "", "The storage type to associate with the DB instance.")
	rds_modifyDbinstanceCmd.Flags().String("tde-credential-arn", "", "The ARN from the key store with which to associate the instance for TDE encryption.")
	rds_modifyDbinstanceCmd.Flags().String("tde-credential-password", "", "The password for the given ARN from the key store in order to access the device.")
	rds_modifyDbinstanceCmd.Flags().String("use-default-processor-features", "", "Specifies whether the DB instance class of the DB instance uses its default processor features.")
	rds_modifyDbinstanceCmd.Flags().String("vpc-security-group-ids", "", "A list of Amazon EC2 VPC security groups to associate with this DB instance.")
	rds_modifyDbinstanceCmd.MarkFlagRequired("dbinstance-identifier")
	rds_modifyDbinstanceCmd.Flag("no-allow-major-version-upgrade").Hidden = true
	rds_modifyDbinstanceCmd.Flag("no-apply-immediately").Hidden = true
	rdsCmd.AddCommand(rds_modifyDbinstanceCmd)
}
