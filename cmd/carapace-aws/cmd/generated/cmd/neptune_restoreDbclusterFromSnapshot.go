package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_restoreDbclusterFromSnapshotCmd = &cobra.Command{
	Use:   "restore-dbcluster-from-snapshot",
	Short: "Creates a new DB cluster from a DB snapshot or DB cluster snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_restoreDbclusterFromSnapshotCmd).Standalone()

	neptune_restoreDbclusterFromSnapshotCmd.Flags().String("availability-zones", "", "Provides the list of EC2 Availability Zones that instances in the restored DB cluster can be created in.")
	neptune_restoreDbclusterFromSnapshotCmd.Flags().String("copy-tags-to-snapshot", "", "*If set to `true`, tags are copied to any snapshot of the restored DB cluster that is created.*")
	neptune_restoreDbclusterFromSnapshotCmd.Flags().String("database-name", "", "Not supported.")
	neptune_restoreDbclusterFromSnapshotCmd.Flags().String("dbcluster-identifier", "", "The name of the DB cluster to create from the DB snapshot or DB cluster snapshot.")
	neptune_restoreDbclusterFromSnapshotCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the DB cluster parameter group to associate with the new DB cluster.")
	neptune_restoreDbclusterFromSnapshotCmd.Flags().String("dbsubnet-group-name", "", "The name of the DB subnet group to use for the new DB cluster.")
	neptune_restoreDbclusterFromSnapshotCmd.Flags().String("deletion-protection", "", "A value that indicates whether the DB cluster has deletion protection enabled.")
	neptune_restoreDbclusterFromSnapshotCmd.Flags().String("enable-cloudwatch-logs-exports", "", "The list of logs that the restored DB cluster is to export to Amazon CloudWatch Logs.")
	neptune_restoreDbclusterFromSnapshotCmd.Flags().String("enable-iamdatabase-authentication", "", "True to enable mapping of Amazon Identity and Access Management (IAM) accounts to database accounts, and otherwise false.")
	neptune_restoreDbclusterFromSnapshotCmd.Flags().String("engine", "", "The database engine to use for the new DB cluster.")
	neptune_restoreDbclusterFromSnapshotCmd.Flags().String("engine-version", "", "The version of the database engine to use for the new DB cluster.")
	neptune_restoreDbclusterFromSnapshotCmd.Flags().String("kms-key-id", "", "The Amazon KMS key identifier to use when restoring an encrypted DB cluster from a DB snapshot or DB cluster snapshot.")
	neptune_restoreDbclusterFromSnapshotCmd.Flags().String("option-group-name", "", "*(Not supported by Neptune)*")
	neptune_restoreDbclusterFromSnapshotCmd.Flags().String("port", "", "The port number on which the new DB cluster accepts connections.")
	neptune_restoreDbclusterFromSnapshotCmd.Flags().String("serverless-v2-scaling-configuration", "", "Contains the scaling configuration of a Neptune Serverless DB cluster.")
	neptune_restoreDbclusterFromSnapshotCmd.Flags().String("snapshot-identifier", "", "The identifier for the DB snapshot or DB cluster snapshot to restore from.")
	neptune_restoreDbclusterFromSnapshotCmd.Flags().String("storage-type", "", "Specifies the storage type to be associated with the DB cluster.")
	neptune_restoreDbclusterFromSnapshotCmd.Flags().String("tags", "", "The tags to be assigned to the restored DB cluster.")
	neptune_restoreDbclusterFromSnapshotCmd.Flags().String("vpc-security-group-ids", "", "A list of VPC security groups that the new DB cluster will belong to.")
	neptune_restoreDbclusterFromSnapshotCmd.MarkFlagRequired("dbcluster-identifier")
	neptune_restoreDbclusterFromSnapshotCmd.MarkFlagRequired("engine")
	neptune_restoreDbclusterFromSnapshotCmd.MarkFlagRequired("snapshot-identifier")
	neptuneCmd.AddCommand(neptune_restoreDbclusterFromSnapshotCmd)
}
