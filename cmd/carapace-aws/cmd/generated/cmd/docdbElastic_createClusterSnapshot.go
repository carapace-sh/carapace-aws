package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElastic_createClusterSnapshotCmd = &cobra.Command{
	Use:   "create-cluster-snapshot",
	Short: "Creates a snapshot of an elastic cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElastic_createClusterSnapshotCmd).Standalone()

	docdbElastic_createClusterSnapshotCmd.Flags().String("cluster-arn", "", "The ARN identifier of the elastic cluster of which you want to create a snapshot.")
	docdbElastic_createClusterSnapshotCmd.Flags().String("snapshot-name", "", "The name of the new elastic cluster snapshot.")
	docdbElastic_createClusterSnapshotCmd.Flags().String("tags", "", "The tags to be assigned to the new elastic cluster snapshot.")
	docdbElastic_createClusterSnapshotCmd.MarkFlagRequired("cluster-arn")
	docdbElastic_createClusterSnapshotCmd.MarkFlagRequired("snapshot-name")
	docdbElasticCmd.AddCommand(docdbElastic_createClusterSnapshotCmd)
}
