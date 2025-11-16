package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElastic_getClusterSnapshotCmd = &cobra.Command{
	Use:   "get-cluster-snapshot",
	Short: "Returns information about a specific elastic cluster snapshot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElastic_getClusterSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdbElastic_getClusterSnapshotCmd).Standalone()

		docdbElastic_getClusterSnapshotCmd.Flags().String("snapshot-arn", "", "The ARN identifier of the elastic cluster snapshot.")
		docdbElastic_getClusterSnapshotCmd.MarkFlagRequired("snapshot-arn")
	})
	docdbElasticCmd.AddCommand(docdbElastic_getClusterSnapshotCmd)
}
