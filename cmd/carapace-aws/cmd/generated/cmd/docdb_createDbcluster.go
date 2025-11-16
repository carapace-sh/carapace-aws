package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_createDbclusterCmd = &cobra.Command{
	Use:   "create-dbcluster",
	Short: "Creates a new Amazon DocumentDB cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_createDbclusterCmd).Standalone()

	docdb_createDbclusterCmd.Flags().String("availability-zones", "", "A list of Amazon EC2 Availability Zones that instances in the cluster can be created in.")
	docdb_createDbclusterCmd.Flags().String("backup-retention-period", "", "The number of days for which automated backups are retained.")
	docdb_createDbclusterCmd.Flags().String("dbcluster-identifier", "", "The cluster identifier.")
	docdb_createDbclusterCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the cluster parameter group to associate with this cluster.")
	docdb_createDbclusterCmd.Flags().String("dbsubnet-group-name", "", "A subnet group to associate with this cluster.")
	docdb_createDbclusterCmd.Flags().String("deletion-protection", "", "Specifies whether this cluster can be deleted.")
	docdb_createDbclusterCmd.Flags().String("enable-cloudwatch-logs-exports", "", "A list of log types that need to be enabled for exporting to Amazon CloudWatch Logs.")
	docdb_createDbclusterCmd.Flags().String("engine", "", "The name of the database engine to be used for this cluster.")
	docdb_createDbclusterCmd.Flags().String("engine-version", "", "The version number of the database engine to use.")
	docdb_createDbclusterCmd.Flags().String("global-cluster-identifier", "", "The cluster identifier of the new global cluster.")
	docdb_createDbclusterCmd.Flags().String("kms-key-id", "", "The KMS key identifier for an encrypted cluster.")
	docdb_createDbclusterCmd.Flags().String("manage-master-user-password", "", "Specifies whether to manage the master user password with Amazon Web Services Secrets Manager.")
	docdb_createDbclusterCmd.Flags().String("master-user-password", "", "The password for the master database user.")
	docdb_createDbclusterCmd.Flags().String("master-user-secret-kms-key-id", "", "The Amazon Web Services KMS key identifier to encrypt a secret that is automatically generated and managed in Amazon Web Services Secrets Manager.")
	docdb_createDbclusterCmd.Flags().String("master-username", "", "The name of the master user for the cluster.")
	docdb_createDbclusterCmd.Flags().String("network-type", "", "The network type of the cluster.")
	docdb_createDbclusterCmd.Flags().String("port", "", "The port number on which the instances in the cluster accept connections.")
	docdb_createDbclusterCmd.Flags().String("pre-signed-url", "", "Not currently supported.")
	docdb_createDbclusterCmd.Flags().String("preferred-backup-window", "", "The daily time range during which automated backups are created if automated backups are enabled using the `BackupRetentionPeriod` parameter.")
	docdb_createDbclusterCmd.Flags().String("preferred-maintenance-window", "", "The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).")
	docdb_createDbclusterCmd.Flags().String("serverless-v2-scaling-configuration", "", "Contains the scaling configuration of an Amazon DocumentDB Serverless cluster.")
	docdb_createDbclusterCmd.Flags().String("storage-encrypted", "", "Specifies whether the cluster is encrypted.")
	docdb_createDbclusterCmd.Flags().String("storage-type", "", "The storage type to associate with the DB cluster.")
	docdb_createDbclusterCmd.Flags().String("tags", "", "The tags to be assigned to the cluster.")
	docdb_createDbclusterCmd.Flags().String("vpc-security-group-ids", "", "A list of EC2 VPC security groups to associate with this cluster.")
	docdb_createDbclusterCmd.MarkFlagRequired("dbcluster-identifier")
	docdb_createDbclusterCmd.MarkFlagRequired("engine")
	docdbCmd.AddCommand(docdb_createDbclusterCmd)
}
