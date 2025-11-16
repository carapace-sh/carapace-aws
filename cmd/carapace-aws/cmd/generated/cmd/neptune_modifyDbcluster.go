package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_modifyDbclusterCmd = &cobra.Command{
	Use:   "modify-dbcluster",
	Short: "Modify a setting for a DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_modifyDbclusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_modifyDbclusterCmd).Standalone()

		neptune_modifyDbclusterCmd.Flags().Bool("allow-major-version-upgrade", false, "A value that indicates whether upgrades between different major versions are allowed.")
		neptune_modifyDbclusterCmd.Flags().Bool("apply-immediately", false, "A value that specifies whether the modifications in this request and any pending modifications are asynchronously applied as soon as possible, regardless of the `PreferredMaintenanceWindow` setting for the DB cluster.")
		neptune_modifyDbclusterCmd.Flags().String("backup-retention-period", "", "The number of days for which automated backups are retained.")
		neptune_modifyDbclusterCmd.Flags().String("cloudwatch-logs-export-configuration", "", "The configuration setting for the log types to be enabled for export to CloudWatch Logs for a specific DB cluster.")
		neptune_modifyDbclusterCmd.Flags().String("copy-tags-to-snapshot", "", "*If set to `true`, tags are copied to any snapshot of the DB cluster that is created.*")
		neptune_modifyDbclusterCmd.Flags().String("dbcluster-identifier", "", "The DB cluster identifier for the cluster being modified.")
		neptune_modifyDbclusterCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the DB cluster parameter group to use for the DB cluster.")
		neptune_modifyDbclusterCmd.Flags().String("dbinstance-parameter-group-name", "", "The name of the DB parameter group to apply to all instances of the DB cluster.")
		neptune_modifyDbclusterCmd.Flags().String("deletion-protection", "", "A value that indicates whether the DB cluster has deletion protection enabled.")
		neptune_modifyDbclusterCmd.Flags().String("enable-iamdatabase-authentication", "", "True to enable mapping of Amazon Identity and Access Management (IAM) accounts to database accounts, and otherwise false.")
		neptune_modifyDbclusterCmd.Flags().String("engine-version", "", "The version number of the database engine to which you want to upgrade.")
		neptune_modifyDbclusterCmd.Flags().String("master-user-password", "", "Not supported by Neptune.")
		neptune_modifyDbclusterCmd.Flags().String("new-dbcluster-identifier", "", "The new DB cluster identifier for the DB cluster when renaming a DB cluster.")
		neptune_modifyDbclusterCmd.Flags().Bool("no-allow-major-version-upgrade", false, "A value that indicates whether upgrades between different major versions are allowed.")
		neptune_modifyDbclusterCmd.Flags().Bool("no-apply-immediately", false, "A value that specifies whether the modifications in this request and any pending modifications are asynchronously applied as soon as possible, regardless of the `PreferredMaintenanceWindow` setting for the DB cluster.")
		neptune_modifyDbclusterCmd.Flags().String("option-group-name", "", "*Not supported by Neptune.*")
		neptune_modifyDbclusterCmd.Flags().String("port", "", "The port number on which the DB cluster accepts connections.")
		neptune_modifyDbclusterCmd.Flags().String("preferred-backup-window", "", "The daily time range during which automated backups are created if automated backups are enabled, using the `BackupRetentionPeriod` parameter.")
		neptune_modifyDbclusterCmd.Flags().String("preferred-maintenance-window", "", "The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).")
		neptune_modifyDbclusterCmd.Flags().String("serverless-v2-scaling-configuration", "", "Contains the scaling configuration of a Neptune Serverless DB cluster.")
		neptune_modifyDbclusterCmd.Flags().String("storage-type", "", "The storage type to associate with the DB cluster.")
		neptune_modifyDbclusterCmd.Flags().String("vpc-security-group-ids", "", "A list of VPC security groups that the DB cluster will belong to.")
		neptune_modifyDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
		neptune_modifyDbclusterCmd.Flag("no-allow-major-version-upgrade").Hidden = true
		neptune_modifyDbclusterCmd.Flag("no-apply-immediately").Hidden = true
	})
	neptuneCmd.AddCommand(neptune_modifyDbclusterCmd)
}
