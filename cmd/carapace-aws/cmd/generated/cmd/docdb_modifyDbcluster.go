package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_modifyDbclusterCmd = &cobra.Command{
	Use:   "modify-dbcluster",
	Short: "Modifies a setting for an Amazon DocumentDB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_modifyDbclusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_modifyDbclusterCmd).Standalone()

		docdb_modifyDbclusterCmd.Flags().Bool("allow-major-version-upgrade", false, "A value that indicates whether major version upgrades are allowed.")
		docdb_modifyDbclusterCmd.Flags().Bool("apply-immediately", false, "A value that specifies whether the changes in this request and any pending changes are asynchronously applied as soon as possible, regardless of the `PreferredMaintenanceWindow` setting for the cluster.")
		docdb_modifyDbclusterCmd.Flags().String("backup-retention-period", "", "The number of days for which automated backups are retained.")
		docdb_modifyDbclusterCmd.Flags().String("cloudwatch-logs-export-configuration", "", "The configuration setting for the log types to be enabled for export to Amazon CloudWatch Logs for a specific instance or cluster.")
		docdb_modifyDbclusterCmd.Flags().String("dbcluster-identifier", "", "The cluster identifier for the cluster that is being modified.")
		docdb_modifyDbclusterCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the cluster parameter group to use for the cluster.")
		docdb_modifyDbclusterCmd.Flags().String("deletion-protection", "", "Specifies whether this cluster can be deleted.")
		docdb_modifyDbclusterCmd.Flags().String("engine-version", "", "The version number of the database engine to which you want to upgrade.")
		docdb_modifyDbclusterCmd.Flags().String("manage-master-user-password", "", "Specifies whether to manage the master user password with Amazon Web Services Secrets Manager.")
		docdb_modifyDbclusterCmd.Flags().String("master-user-password", "", "The password for the master database user.")
		docdb_modifyDbclusterCmd.Flags().String("master-user-secret-kms-key-id", "", "The Amazon Web Services KMS key identifier to encrypt a secret that is automatically generated and managed in Amazon Web Services Secrets Manager.")
		docdb_modifyDbclusterCmd.Flags().String("network-type", "", "The network type of the cluster.")
		docdb_modifyDbclusterCmd.Flags().String("new-dbcluster-identifier", "", "The new cluster identifier for the cluster when renaming a cluster.")
		docdb_modifyDbclusterCmd.Flags().Bool("no-allow-major-version-upgrade", false, "A value that indicates whether major version upgrades are allowed.")
		docdb_modifyDbclusterCmd.Flags().Bool("no-apply-immediately", false, "A value that specifies whether the changes in this request and any pending changes are asynchronously applied as soon as possible, regardless of the `PreferredMaintenanceWindow` setting for the cluster.")
		docdb_modifyDbclusterCmd.Flags().String("port", "", "The port number on which the cluster accepts connections.")
		docdb_modifyDbclusterCmd.Flags().String("preferred-backup-window", "", "The daily time range during which automated backups are created if automated backups are enabled, using the `BackupRetentionPeriod` parameter.")
		docdb_modifyDbclusterCmd.Flags().String("preferred-maintenance-window", "", "The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).")
		docdb_modifyDbclusterCmd.Flags().String("rotate-master-user-password", "", "Specifies whether to rotate the secret managed by Amazon Web Services Secrets Manager for the master user password.")
		docdb_modifyDbclusterCmd.Flags().String("serverless-v2-scaling-configuration", "", "Contains the scaling configuration of an Amazon DocumentDB Serverless cluster.")
		docdb_modifyDbclusterCmd.Flags().String("storage-type", "", "The storage type to associate with the DB cluster.")
		docdb_modifyDbclusterCmd.Flags().String("vpc-security-group-ids", "", "A list of virtual private cloud (VPC) security groups that the cluster will belong to.")
		docdb_modifyDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
		docdb_modifyDbclusterCmd.Flag("no-allow-major-version-upgrade").Hidden = true
		docdb_modifyDbclusterCmd.Flag("no-apply-immediately").Hidden = true
	})
	docdbCmd.AddCommand(docdb_modifyDbclusterCmd)
}
