package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_restoreDbclusterToPointInTimeCmd = &cobra.Command{
	Use:   "restore-dbcluster-to-point-in-time",
	Short: "Restores a DB cluster to an arbitrary point in time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_restoreDbclusterToPointInTimeCmd).Standalone()

	neptune_restoreDbclusterToPointInTimeCmd.Flags().String("dbcluster-identifier", "", "The name of the new DB cluster to be created.")
	neptune_restoreDbclusterToPointInTimeCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the DB cluster parameter group to associate with the new DB cluster.")
	neptune_restoreDbclusterToPointInTimeCmd.Flags().String("dbsubnet-group-name", "", "The DB subnet group name to use for the new DB cluster.")
	neptune_restoreDbclusterToPointInTimeCmd.Flags().String("deletion-protection", "", "A value that indicates whether the DB cluster has deletion protection enabled.")
	neptune_restoreDbclusterToPointInTimeCmd.Flags().String("enable-cloudwatch-logs-exports", "", "The list of logs that the restored DB cluster is to export to CloudWatch Logs.")
	neptune_restoreDbclusterToPointInTimeCmd.Flags().String("enable-iamdatabase-authentication", "", "True to enable mapping of Amazon Identity and Access Management (IAM) accounts to database accounts, and otherwise false.")
	neptune_restoreDbclusterToPointInTimeCmd.Flags().String("kms-key-id", "", "The Amazon KMS key identifier to use when restoring an encrypted DB cluster from an encrypted DB cluster.")
	neptune_restoreDbclusterToPointInTimeCmd.Flags().Bool("no-use-latest-restorable-time", false, "A value that is set to `true` to restore the DB cluster to the latest restorable backup time, and `false` otherwise.")
	neptune_restoreDbclusterToPointInTimeCmd.Flags().String("option-group-name", "", "*(Not supported by Neptune)*")
	neptune_restoreDbclusterToPointInTimeCmd.Flags().String("port", "", "The port number on which the new DB cluster accepts connections.")
	neptune_restoreDbclusterToPointInTimeCmd.Flags().String("restore-to-time", "", "The date and time to restore the DB cluster to.")
	neptune_restoreDbclusterToPointInTimeCmd.Flags().String("restore-type", "", "The type of restore to be performed.")
	neptune_restoreDbclusterToPointInTimeCmd.Flags().String("serverless-v2-scaling-configuration", "", "Contains the scaling configuration of a Neptune Serverless DB cluster.")
	neptune_restoreDbclusterToPointInTimeCmd.Flags().String("source-dbcluster-identifier", "", "The identifier of the source DB cluster from which to restore.")
	neptune_restoreDbclusterToPointInTimeCmd.Flags().String("storage-type", "", "Specifies the storage type to be associated with the DB cluster.")
	neptune_restoreDbclusterToPointInTimeCmd.Flags().String("tags", "", "The tags to be applied to the restored DB cluster.")
	neptune_restoreDbclusterToPointInTimeCmd.Flags().Bool("use-latest-restorable-time", false, "A value that is set to `true` to restore the DB cluster to the latest restorable backup time, and `false` otherwise.")
	neptune_restoreDbclusterToPointInTimeCmd.Flags().String("vpc-security-group-ids", "", "A list of VPC security groups that the new DB cluster belongs to.")
	neptune_restoreDbclusterToPointInTimeCmd.MarkFlagRequired("dbcluster-identifier")
	neptune_restoreDbclusterToPointInTimeCmd.Flag("no-use-latest-restorable-time").Hidden = true
	neptune_restoreDbclusterToPointInTimeCmd.MarkFlagRequired("source-dbcluster-identifier")
	neptuneCmd.AddCommand(neptune_restoreDbclusterToPointInTimeCmd)
}
