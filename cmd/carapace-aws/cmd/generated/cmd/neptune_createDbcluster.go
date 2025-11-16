package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_createDbclusterCmd = &cobra.Command{
	Use:   "create-dbcluster",
	Short: "Creates a new Amazon Neptune DB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_createDbclusterCmd).Standalone()

	neptune_createDbclusterCmd.Flags().String("availability-zones", "", "A list of EC2 Availability Zones that instances in the DB cluster can be created in.")
	neptune_createDbclusterCmd.Flags().String("backup-retention-period", "", "The number of days for which automated backups are retained.")
	neptune_createDbclusterCmd.Flags().String("character-set-name", "", "*(Not supported by Neptune)*")
	neptune_createDbclusterCmd.Flags().String("copy-tags-to-snapshot", "", "*If set to `true`, tags are copied to any snapshot of the DB cluster that is created.*")
	neptune_createDbclusterCmd.Flags().String("database-name", "", "The name for your database of up to 64 alpha-numeric characters.")
	neptune_createDbclusterCmd.Flags().String("dbcluster-identifier", "", "The DB cluster identifier.")
	neptune_createDbclusterCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the DB cluster parameter group to associate with this DB cluster.")
	neptune_createDbclusterCmd.Flags().String("dbsubnet-group-name", "", "A DB subnet group to associate with this DB cluster.")
	neptune_createDbclusterCmd.Flags().String("deletion-protection", "", "A value that indicates whether the DB cluster has deletion protection enabled.")
	neptune_createDbclusterCmd.Flags().String("enable-cloudwatch-logs-exports", "", "A list of the log types that this DB cluster should export to CloudWatch Logs.")
	neptune_createDbclusterCmd.Flags().String("enable-iamdatabase-authentication", "", "If set to `true`, enables Amazon Identity and Access Management (IAM) authentication for the entire DB cluster (this cannot be set at an instance level).")
	neptune_createDbclusterCmd.Flags().String("engine", "", "The name of the database engine to be used for this DB cluster.")
	neptune_createDbclusterCmd.Flags().String("engine-version", "", "The version number of the database engine to use for the new DB cluster.")
	neptune_createDbclusterCmd.Flags().String("global-cluster-identifier", "", "The ID of the Neptune global database to which this new DB cluster should be added.")
	neptune_createDbclusterCmd.Flags().String("kms-key-id", "", "The Amazon KMS key identifier for an encrypted DB cluster.")
	neptune_createDbclusterCmd.Flags().String("master-user-password", "", "Not supported by Neptune.")
	neptune_createDbclusterCmd.Flags().String("master-username", "", "Not supported by Neptune.")
	neptune_createDbclusterCmd.Flags().String("option-group-name", "", "*(Not supported by Neptune)*")
	neptune_createDbclusterCmd.Flags().String("port", "", "The port number on which the instances in the DB cluster accept connections.")
	neptune_createDbclusterCmd.Flags().String("pre-signed-url", "", "This parameter is not currently supported.")
	neptune_createDbclusterCmd.Flags().String("preferred-backup-window", "", "The daily time range during which automated backups are created if automated backups are enabled using the `BackupRetentionPeriod` parameter.")
	neptune_createDbclusterCmd.Flags().String("preferred-maintenance-window", "", "The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).")
	neptune_createDbclusterCmd.Flags().String("replication-source-identifier", "", "The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this DB cluster is created as a Read Replica.")
	neptune_createDbclusterCmd.Flags().String("serverless-v2-scaling-configuration", "", "Contains the scaling configuration of a Neptune Serverless DB cluster.")
	neptune_createDbclusterCmd.Flags().String("storage-encrypted", "", "Specifies whether the DB cluster is encrypted.")
	neptune_createDbclusterCmd.Flags().String("storage-type", "", "The storage type for the new DB cluster.")
	neptune_createDbclusterCmd.Flags().String("tags", "", "The tags to assign to the new DB cluster.")
	neptune_createDbclusterCmd.Flags().String("vpc-security-group-ids", "", "A list of EC2 VPC security groups to associate with this DB cluster.")
	neptune_createDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
	neptune_createDbclusterCmd.MarkFlagRequired("engine")
	neptuneCmd.AddCommand(neptune_createDbclusterCmd)
}
