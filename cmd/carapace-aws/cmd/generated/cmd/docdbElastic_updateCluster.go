package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElastic_updateClusterCmd = &cobra.Command{
	Use:   "update-cluster",
	Short: "Modifies an elastic cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElastic_updateClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdbElastic_updateClusterCmd).Standalone()

		docdbElastic_updateClusterCmd.Flags().String("admin-user-password", "", "The password associated with the elastic cluster administrator.")
		docdbElastic_updateClusterCmd.Flags().String("auth-type", "", "The authentication type used to determine where to fetch the password used for accessing the elastic cluster.")
		docdbElastic_updateClusterCmd.Flags().String("backup-retention-period", "", "The number of days for which automatic snapshots are retained.")
		docdbElastic_updateClusterCmd.Flags().String("client-token", "", "The client token for the elastic cluster.")
		docdbElastic_updateClusterCmd.Flags().String("cluster-arn", "", "The ARN identifier of the elastic cluster.")
		docdbElastic_updateClusterCmd.Flags().String("preferred-backup-window", "", "The daily time range during which automated backups are created if automated backups are enabled, as determined by the `backupRetentionPeriod`.")
		docdbElastic_updateClusterCmd.Flags().String("preferred-maintenance-window", "", "The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).")
		docdbElastic_updateClusterCmd.Flags().String("shard-capacity", "", "The number of vCPUs assigned to each elastic cluster shard.")
		docdbElastic_updateClusterCmd.Flags().String("shard-count", "", "The number of shards assigned to the elastic cluster.")
		docdbElastic_updateClusterCmd.Flags().String("shard-instance-count", "", "The number of replica instances applying to all shards in the elastic cluster.")
		docdbElastic_updateClusterCmd.Flags().String("subnet-ids", "", "The Amazon EC2 subnet IDs for the elastic cluster.")
		docdbElastic_updateClusterCmd.Flags().String("vpc-security-group-ids", "", "A list of EC2 VPC security groups to associate with the elastic cluster.")
		docdbElastic_updateClusterCmd.MarkFlagRequired("cluster-arn")
	})
	docdbElasticCmd.AddCommand(docdbElastic_updateClusterCmd)
}
