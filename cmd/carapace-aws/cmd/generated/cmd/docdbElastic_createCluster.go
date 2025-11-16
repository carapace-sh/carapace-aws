package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElastic_createClusterCmd = &cobra.Command{
	Use:   "create-cluster",
	Short: "Creates a new Amazon DocumentDB elastic cluster and returns its cluster structure.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElastic_createClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdbElastic_createClusterCmd).Standalone()

		docdbElastic_createClusterCmd.Flags().String("admin-user-name", "", "The name of the Amazon DocumentDB elastic clusters administrator.")
		docdbElastic_createClusterCmd.Flags().String("admin-user-password", "", "The password for the Amazon DocumentDB elastic clusters administrator.")
		docdbElastic_createClusterCmd.Flags().String("auth-type", "", "The authentication type used to determine where to fetch the password used for accessing the elastic cluster.")
		docdbElastic_createClusterCmd.Flags().String("backup-retention-period", "", "The number of days for which automatic snapshots are retained.")
		docdbElastic_createClusterCmd.Flags().String("client-token", "", "The client token for the elastic cluster.")
		docdbElastic_createClusterCmd.Flags().String("cluster-name", "", "The name of the new elastic cluster.")
		docdbElastic_createClusterCmd.Flags().String("kms-key-id", "", "The KMS key identifier to use to encrypt the new elastic cluster.")
		docdbElastic_createClusterCmd.Flags().String("preferred-backup-window", "", "The daily time range during which automated backups are created if automated backups are enabled, as determined by the `backupRetentionPeriod`.")
		docdbElastic_createClusterCmd.Flags().String("preferred-maintenance-window", "", "The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).")
		docdbElastic_createClusterCmd.Flags().String("shard-capacity", "", "The number of vCPUs assigned to each elastic cluster shard.")
		docdbElastic_createClusterCmd.Flags().String("shard-count", "", "The number of shards assigned to the elastic cluster.")
		docdbElastic_createClusterCmd.Flags().String("shard-instance-count", "", "The number of replica instances applying to all shards in the elastic cluster.")
		docdbElastic_createClusterCmd.Flags().String("subnet-ids", "", "The Amazon EC2 subnet IDs for the new elastic cluster.")
		docdbElastic_createClusterCmd.Flags().String("tags", "", "The tags to be assigned to the new elastic cluster.")
		docdbElastic_createClusterCmd.Flags().String("vpc-security-group-ids", "", "A list of EC2 VPC security groups to associate with the new elastic cluster.")
		docdbElastic_createClusterCmd.MarkFlagRequired("admin-user-name")
		docdbElastic_createClusterCmd.MarkFlagRequired("admin-user-password")
		docdbElastic_createClusterCmd.MarkFlagRequired("auth-type")
		docdbElastic_createClusterCmd.MarkFlagRequired("cluster-name")
		docdbElastic_createClusterCmd.MarkFlagRequired("shard-capacity")
		docdbElastic_createClusterCmd.MarkFlagRequired("shard-count")
	})
	docdbElasticCmd.AddCommand(docdbElastic_createClusterCmd)
}
