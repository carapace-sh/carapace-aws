package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_restoreDbclusterToPointInTimeCmd = &cobra.Command{
	Use:   "restore-dbcluster-to-point-in-time",
	Short: "Restores a DB cluster to an arbitrary point in time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_restoreDbclusterToPointInTimeCmd).Standalone()

	rds_restoreDbclusterToPointInTimeCmd.Flags().String("backtrack-window", "", "The target backtrack window, in seconds.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("copy-tags-to-snapshot", "", "Specifies whether to copy all tags from the restored DB cluster to snapshots of the restored DB cluster.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("dbcluster-identifier", "", "The name of the new DB cluster to be created.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("dbcluster-instance-class", "", "The compute and memory capacity of the each DB instance in the Multi-AZ DB cluster, for example db.m6gd.xlarge. Not all DB instance classes are available in all Amazon Web Services Regions, or for all database engines.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the custom DB cluster parameter group to associate with this DB cluster.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("dbsubnet-group-name", "", "The DB subnet group name to use for the new DB cluster.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("deletion-protection", "", "Specifies whether to enable deletion protection for the DB cluster.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("domain", "", "The Active Directory directory ID to restore the DB cluster in.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("domain-iamrole-name", "", "The name of the IAM role to be used when making API calls to the Directory Service.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("enable-cloudwatch-logs-exports", "", "The list of logs that the restored DB cluster is to export to CloudWatch Logs.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("enable-iamdatabase-authentication", "", "Specifies whether to enable mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("enable-performance-insights", "", "Specifies whether to turn on Performance Insights for the DB cluster.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("engine-lifecycle-support", "", "The life cycle type for this DB cluster.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("engine-mode", "", "The engine mode of the new cluster.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("iops", "", "The amount of Provisioned IOPS (input/output operations per second) to be initially allocated for each DB instance in the Multi-AZ DB cluster.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("kms-key-id", "", "The Amazon Web Services KMS key identifier to use when restoring an encrypted DB cluster from an encrypted DB cluster.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("monitoring-interval", "", "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB cluster.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("monitoring-role-arn", "", "The Amazon Resource Name (ARN) for the IAM role that permits RDS to send Enhanced Monitoring metrics to Amazon CloudWatch Logs.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("network-type", "", "The network type of the DB cluster.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().Bool("no-use-latest-restorable-time", false, "Specifies whether to restore the DB cluster to the latest restorable backup time.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("option-group-name", "", "The name of the option group for the new DB cluster.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("performance-insights-kmskey-id", "", "The Amazon Web Services KMS key identifier for encryption of Performance Insights data.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("performance-insights-retention-period", "", "The number of days to retain Performance Insights data.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("port", "", "The port number on which the new DB cluster accepts connections.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("publicly-accessible", "", "Specifies whether the DB cluster is publicly accessible.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("rds-custom-cluster-configuration", "", "Reserved for future use.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("restore-to-time", "", "The date and time to restore the DB cluster to.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("restore-type", "", "The type of restore to be performed.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("scaling-configuration", "", "For DB clusters in `serverless` DB engine mode, the scaling properties of the DB cluster.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("serverless-v2-scaling-configuration", "", "")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("source-db-cluster-resource-id", "", "The resource ID of the source DB cluster from which to restore.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("source-dbcluster-identifier", "", "The identifier of the source DB cluster from which to restore.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("storage-type", "", "Specifies the storage type to be associated with the DB cluster.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("tags", "", "")
	rds_restoreDbclusterToPointInTimeCmd.Flags().Bool("use-latest-restorable-time", false, "Specifies whether to restore the DB cluster to the latest restorable backup time.")
	rds_restoreDbclusterToPointInTimeCmd.Flags().String("vpc-security-group-ids", "", "A list of VPC security groups that the new DB cluster belongs to.")
	rds_restoreDbclusterToPointInTimeCmd.MarkFlagRequired("dbcluster-identifier")
	rds_restoreDbclusterToPointInTimeCmd.Flag("no-use-latest-restorable-time").Hidden = true
	rdsCmd.AddCommand(rds_restoreDbclusterToPointInTimeCmd)
}
