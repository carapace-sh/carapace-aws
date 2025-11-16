package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_restoreDbclusterFromSnapshotCmd = &cobra.Command{
	Use:   "restore-dbcluster-from-snapshot",
	Short: "Creates a new cluster from a snapshot or cluster snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_restoreDbclusterFromSnapshotCmd).Standalone()

	docdb_restoreDbclusterFromSnapshotCmd.Flags().String("availability-zones", "", "Provides the list of Amazon EC2 Availability Zones that instances in the restored DB cluster can be created in.")
	docdb_restoreDbclusterFromSnapshotCmd.Flags().String("dbcluster-identifier", "", "The name of the cluster to create from the snapshot or cluster snapshot.")
	docdb_restoreDbclusterFromSnapshotCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the DB cluster parameter group to associate with this DB cluster.")
	docdb_restoreDbclusterFromSnapshotCmd.Flags().String("dbsubnet-group-name", "", "The name of the subnet group to use for the new cluster.")
	docdb_restoreDbclusterFromSnapshotCmd.Flags().String("deletion-protection", "", "Specifies whether this cluster can be deleted.")
	docdb_restoreDbclusterFromSnapshotCmd.Flags().String("enable-cloudwatch-logs-exports", "", "A list of log types that must be enabled for exporting to Amazon CloudWatch Logs.")
	docdb_restoreDbclusterFromSnapshotCmd.Flags().String("engine", "", "The database engine to use for the new cluster.")
	docdb_restoreDbclusterFromSnapshotCmd.Flags().String("engine-version", "", "The version of the database engine to use for the new cluster.")
	docdb_restoreDbclusterFromSnapshotCmd.Flags().String("kms-key-id", "", "The KMS key identifier to use when restoring an encrypted cluster from a DB snapshot or cluster snapshot.")
	docdb_restoreDbclusterFromSnapshotCmd.Flags().String("network-type", "", "The network type of the cluster.")
	docdb_restoreDbclusterFromSnapshotCmd.Flags().String("port", "", "The port number on which the new cluster accepts connections.")
	docdb_restoreDbclusterFromSnapshotCmd.Flags().String("serverless-v2-scaling-configuration", "", "Contains the scaling configuration of an Amazon DocumentDB Serverless cluster.")
	docdb_restoreDbclusterFromSnapshotCmd.Flags().String("snapshot-identifier", "", "The identifier for the snapshot or cluster snapshot to restore from.")
	docdb_restoreDbclusterFromSnapshotCmd.Flags().String("storage-type", "", "The storage type to associate with the DB cluster.")
	docdb_restoreDbclusterFromSnapshotCmd.Flags().String("tags", "", "The tags to be assigned to the restored cluster.")
	docdb_restoreDbclusterFromSnapshotCmd.Flags().String("vpc-security-group-ids", "", "A list of virtual private cloud (VPC) security groups that the new cluster will belong to.")
	docdb_restoreDbclusterFromSnapshotCmd.MarkFlagRequired("dbcluster-identifier")
	docdb_restoreDbclusterFromSnapshotCmd.MarkFlagRequired("engine")
	docdb_restoreDbclusterFromSnapshotCmd.MarkFlagRequired("snapshot-identifier")
	docdbCmd.AddCommand(docdb_restoreDbclusterFromSnapshotCmd)
}
