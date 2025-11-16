package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElastic_restoreClusterFromSnapshotCmd = &cobra.Command{
	Use:   "restore-cluster-from-snapshot",
	Short: "Restores an elastic cluster from a snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElastic_restoreClusterFromSnapshotCmd).Standalone()

	docdbElastic_restoreClusterFromSnapshotCmd.Flags().String("cluster-name", "", "The name of the elastic cluster.")
	docdbElastic_restoreClusterFromSnapshotCmd.Flags().String("kms-key-id", "", "The KMS key identifier to use to encrypt the new Amazon DocumentDB elastic clusters cluster.")
	docdbElastic_restoreClusterFromSnapshotCmd.Flags().String("shard-capacity", "", "The capacity of each shard in the new restored elastic cluster.")
	docdbElastic_restoreClusterFromSnapshotCmd.Flags().String("shard-instance-count", "", "The number of replica instances applying to all shards in the elastic cluster.")
	docdbElastic_restoreClusterFromSnapshotCmd.Flags().String("snapshot-arn", "", "The ARN identifier of the elastic cluster snapshot.")
	docdbElastic_restoreClusterFromSnapshotCmd.Flags().String("subnet-ids", "", "The Amazon EC2 subnet IDs for the elastic cluster.")
	docdbElastic_restoreClusterFromSnapshotCmd.Flags().String("tags", "", "A list of the tag names to be assigned to the restored elastic cluster, in the form of an array of key-value pairs in which the key is the tag name and the value is the key value.")
	docdbElastic_restoreClusterFromSnapshotCmd.Flags().String("vpc-security-group-ids", "", "A list of EC2 VPC security groups to associate with the elastic cluster.")
	docdbElastic_restoreClusterFromSnapshotCmd.MarkFlagRequired("cluster-name")
	docdbElastic_restoreClusterFromSnapshotCmd.MarkFlagRequired("snapshot-arn")
	docdbElasticCmd.AddCommand(docdbElastic_restoreClusterFromSnapshotCmd)
}
