package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElastic_deleteClusterSnapshotCmd = &cobra.Command{
	Use:   "delete-cluster-snapshot",
	Short: "Delete an elastic cluster snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElastic_deleteClusterSnapshotCmd).Standalone()

	docdbElastic_deleteClusterSnapshotCmd.Flags().String("snapshot-arn", "", "The ARN identifier of the elastic cluster snapshot that is to be deleted.")
	docdbElastic_deleteClusterSnapshotCmd.MarkFlagRequired("snapshot-arn")
	docdbElasticCmd.AddCommand(docdbElastic_deleteClusterSnapshotCmd)
}
