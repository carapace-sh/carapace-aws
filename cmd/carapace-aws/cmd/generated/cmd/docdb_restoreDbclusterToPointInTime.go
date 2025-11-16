package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_restoreDbclusterToPointInTimeCmd = &cobra.Command{
	Use:   "restore-dbcluster-to-point-in-time",
	Short: "Restores a cluster to an arbitrary point in time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_restoreDbclusterToPointInTimeCmd).Standalone()

	docdb_restoreDbclusterToPointInTimeCmd.Flags().String("dbcluster-identifier", "", "The name of the new cluster to be created.")
	docdb_restoreDbclusterToPointInTimeCmd.Flags().String("dbsubnet-group-name", "", "The subnet group name to use for the new cluster.")
	docdb_restoreDbclusterToPointInTimeCmd.Flags().String("deletion-protection", "", "Specifies whether this cluster can be deleted.")
	docdb_restoreDbclusterToPointInTimeCmd.Flags().String("enable-cloudwatch-logs-exports", "", "A list of log types that must be enabled for exporting to Amazon CloudWatch Logs.")
	docdb_restoreDbclusterToPointInTimeCmd.Flags().String("kms-key-id", "", "The KMS key identifier to use when restoring an encrypted cluster from an encrypted cluster.")
	docdb_restoreDbclusterToPointInTimeCmd.Flags().String("network-type", "", "The network type of the cluster.")
	docdb_restoreDbclusterToPointInTimeCmd.Flags().Bool("no-use-latest-restorable-time", false, "A value that is set to `true` to restore the cluster to the latest restorable backup time, and `false` otherwise.")
	docdb_restoreDbclusterToPointInTimeCmd.Flags().String("port", "", "The port number on which the new cluster accepts connections.")
	docdb_restoreDbclusterToPointInTimeCmd.Flags().String("restore-to-time", "", "The date and time to restore the cluster to.")
	docdb_restoreDbclusterToPointInTimeCmd.Flags().String("restore-type", "", "The type of restore to be performed.")
	docdb_restoreDbclusterToPointInTimeCmd.Flags().String("serverless-v2-scaling-configuration", "", "Contains the scaling configuration of an Amazon DocumentDB Serverless cluster.")
	docdb_restoreDbclusterToPointInTimeCmd.Flags().String("source-dbcluster-identifier", "", "The identifier of the source cluster from which to restore.")
	docdb_restoreDbclusterToPointInTimeCmd.Flags().String("storage-type", "", "The storage type to associate with the DB cluster.")
	docdb_restoreDbclusterToPointInTimeCmd.Flags().String("tags", "", "The tags to be assigned to the restored cluster.")
	docdb_restoreDbclusterToPointInTimeCmd.Flags().Bool("use-latest-restorable-time", false, "A value that is set to `true` to restore the cluster to the latest restorable backup time, and `false` otherwise.")
	docdb_restoreDbclusterToPointInTimeCmd.Flags().String("vpc-security-group-ids", "", "A list of VPC security groups that the new cluster belongs to.")
	docdb_restoreDbclusterToPointInTimeCmd.MarkFlagRequired("dbcluster-identifier")
	docdb_restoreDbclusterToPointInTimeCmd.Flag("no-use-latest-restorable-time").Hidden = true
	docdb_restoreDbclusterToPointInTimeCmd.MarkFlagRequired("source-dbcluster-identifier")
	docdbCmd.AddCommand(docdb_restoreDbclusterToPointInTimeCmd)
}
