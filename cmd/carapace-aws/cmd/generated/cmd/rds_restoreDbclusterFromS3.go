package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_restoreDbclusterFromS3Cmd = &cobra.Command{
	Use:   "restore-dbcluster-from-s3",
	Short: "Creates an Amazon Aurora DB cluster from MySQL data stored in an Amazon S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_restoreDbclusterFromS3Cmd).Standalone()

	rds_restoreDbclusterFromS3Cmd.Flags().String("availability-zones", "", "A list of Availability Zones (AZs) where instances in the restored DB cluster can be created.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("backtrack-window", "", "The target backtrack window, in seconds.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("backup-retention-period", "", "The number of days for which automated backups of the restored DB cluster are retained.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("character-set-name", "", "A value that indicates that the restored DB cluster should be associated with the specified CharacterSet.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("copy-tags-to-snapshot", "", "Specifies whether to copy all tags from the restored DB cluster to snapshots of the restored DB cluster.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("database-name", "", "The database name for the restored DB cluster.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("dbcluster-identifier", "", "The name of the DB cluster to create from the source data in the Amazon S3 bucket.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the DB cluster parameter group to associate with the restored DB cluster.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("dbsubnet-group-name", "", "A DB subnet group to associate with the restored DB cluster.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("deletion-protection", "", "Specifies whether to enable deletion protection for the DB cluster.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("domain", "", "Specify the Active Directory directory ID to restore the DB cluster in.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("domain-iamrole-name", "", "Specify the name of the IAM role to be used when making API calls to the Directory Service.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("enable-cloudwatch-logs-exports", "", "The list of logs that the restored DB cluster is to export to CloudWatch Logs.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("enable-iamdatabase-authentication", "", "Specifies whether to enable mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("engine", "", "The name of the database engine to be used for this DB cluster.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("engine-lifecycle-support", "", "The life cycle type for this DB cluster.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("engine-version", "", "The version number of the database engine to use.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("kms-key-id", "", "The Amazon Web Services KMS key identifier for an encrypted DB cluster.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("manage-master-user-password", "", "Specifies whether to manage the master user password with Amazon Web Services Secrets Manager.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("master-user-password", "", "The password for the master database user.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("master-user-secret-kms-key-id", "", "The Amazon Web Services KMS key identifier to encrypt a secret that is automatically generated and managed in Amazon Web Services Secrets Manager.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("master-username", "", "The name of the master user for the restored DB cluster.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("network-type", "", "The network type of the DB cluster.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("option-group-name", "", "A value that indicates that the restored DB cluster should be associated with the specified option group.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("port", "", "The port number on which the instances in the restored DB cluster accept connections.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("preferred-backup-window", "", "The daily time range during which automated backups are created if automated backups are enabled using the `BackupRetentionPeriod` parameter.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("preferred-maintenance-window", "", "The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).")
	rds_restoreDbclusterFromS3Cmd.Flags().String("s3-bucket-name", "", "The name of the Amazon S3 bucket that contains the data used to create the Amazon Aurora DB cluster.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("s3-ingestion-role-arn", "", "The Amazon Resource Name (ARN) of the Amazon Web Services Identity and Access Management (IAM) role that authorizes Amazon RDS to access the Amazon S3 bucket on your behalf.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("s3-prefix", "", "The prefix for all of the file names that contain the data used to create the Amazon Aurora DB cluster.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("serverless-v2-scaling-configuration", "", "")
	rds_restoreDbclusterFromS3Cmd.Flags().String("source-engine", "", "The identifier for the database engine that was backed up to create the files stored in the Amazon S3 bucket.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("source-engine-version", "", "The version of the database that the backup files were created from.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("storage-encrypted", "", "Specifies whether the restored DB cluster is encrypted.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("storage-type", "", "Specifies the storage type to be associated with the DB cluster.")
	rds_restoreDbclusterFromS3Cmd.Flags().String("tags", "", "")
	rds_restoreDbclusterFromS3Cmd.Flags().String("vpc-security-group-ids", "", "A list of EC2 VPC security groups to associate with the restored DB cluster.")
	rds_restoreDbclusterFromS3Cmd.MarkFlagRequired("dbcluster-identifier")
	rds_restoreDbclusterFromS3Cmd.MarkFlagRequired("engine")
	rds_restoreDbclusterFromS3Cmd.MarkFlagRequired("master-username")
	rds_restoreDbclusterFromS3Cmd.MarkFlagRequired("s3-bucket-name")
	rds_restoreDbclusterFromS3Cmd.MarkFlagRequired("s3-ingestion-role-arn")
	rds_restoreDbclusterFromS3Cmd.MarkFlagRequired("source-engine")
	rds_restoreDbclusterFromS3Cmd.MarkFlagRequired("source-engine-version")
	rdsCmd.AddCommand(rds_restoreDbclusterFromS3Cmd)
}
