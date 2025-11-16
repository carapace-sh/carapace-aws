package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_restoreDbclusterFromSnapshotCmd = &cobra.Command{
	Use:   "restore-dbcluster-from-snapshot",
	Short: "Creates a new DB cluster from a DB snapshot or DB cluster snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_restoreDbclusterFromSnapshotCmd).Standalone()

	rds_restoreDbclusterFromSnapshotCmd.Flags().String("availability-zones", "", "Provides the list of Availability Zones (AZs) where instances in the restored DB cluster can be created.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("backtrack-window", "", "The target backtrack window, in seconds.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("copy-tags-to-snapshot", "", "Specifies whether to copy all tags from the restored DB cluster to snapshots of the restored DB cluster.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("database-name", "", "The database name for the restored DB cluster.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("dbcluster-identifier", "", "The name of the DB cluster to create from the DB snapshot or DB cluster snapshot.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("dbcluster-instance-class", "", "The compute and memory capacity of the each DB instance in the Multi-AZ DB cluster, for example db.m6gd.xlarge. Not all DB instance classes are available in all Amazon Web Services Regions, or for all database engines.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the DB cluster parameter group to associate with this DB cluster.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("dbsubnet-group-name", "", "The name of the DB subnet group to use for the new DB cluster.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("deletion-protection", "", "Specifies whether to enable deletion protection for the DB cluster.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("domain", "", "The Active Directory directory ID to restore the DB cluster in.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("domain-iamrole-name", "", "The name of the IAM role to be used when making API calls to the Directory Service.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("enable-cloudwatch-logs-exports", "", "The list of logs that the restored DB cluster is to export to Amazon CloudWatch Logs.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("enable-iamdatabase-authentication", "", "Specifies whether to enable mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("enable-performance-insights", "", "Specifies whether to turn on Performance Insights for the DB cluster.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("engine", "", "The database engine to use for the new DB cluster.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("engine-lifecycle-support", "", "The life cycle type for this DB cluster.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("engine-mode", "", "The DB engine mode of the DB cluster, either `provisioned` or `serverless`.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("engine-version", "", "The version of the database engine to use for the new DB cluster.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("iops", "", "The amount of Provisioned IOPS (input/output operations per second) to be initially allocated for each DB instance in the Multi-AZ DB cluster.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("kms-key-id", "", "The Amazon Web Services KMS key identifier to use when restoring an encrypted DB cluster from a DB snapshot or DB cluster snapshot.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("monitoring-interval", "", "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB cluster.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("monitoring-role-arn", "", "The Amazon Resource Name (ARN) for the IAM role that permits RDS to send Enhanced Monitoring metrics to Amazon CloudWatch Logs.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("network-type", "", "The network type of the DB cluster.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("option-group-name", "", "The name of the option group to use for the restored DB cluster.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("performance-insights-kmskey-id", "", "The Amazon Web Services KMS key identifier for encryption of Performance Insights data.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("performance-insights-retention-period", "", "The number of days to retain Performance Insights data.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("port", "", "The port number on which the new DB cluster accepts connections.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("publicly-accessible", "", "Specifies whether the DB cluster is publicly accessible.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("rds-custom-cluster-configuration", "", "Reserved for future use.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("scaling-configuration", "", "For DB clusters in `serverless` DB engine mode, the scaling properties of the DB cluster.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("serverless-v2-scaling-configuration", "", "")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("snapshot-identifier", "", "The identifier for the DB snapshot or DB cluster snapshot to restore from.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("storage-type", "", "Specifies the storage type to be associated with the DB cluster.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("tags", "", "The tags to be assigned to the restored DB cluster.")
	rds_restoreDbclusterFromSnapshotCmd.Flags().String("vpc-security-group-ids", "", "A list of VPC security groups that the new DB cluster will belong to.")
	rds_restoreDbclusterFromSnapshotCmd.MarkFlagRequired("dbcluster-identifier")
	rds_restoreDbclusterFromSnapshotCmd.MarkFlagRequired("engine")
	rds_restoreDbclusterFromSnapshotCmd.MarkFlagRequired("snapshot-identifier")
	rdsCmd.AddCommand(rds_restoreDbclusterFromSnapshotCmd)
}
