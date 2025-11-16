package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElastic_listClusterSnapshotsCmd = &cobra.Command{
	Use:   "list-cluster-snapshots",
	Short: "Returns information about snapshots for a specified elastic cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElastic_listClusterSnapshotsCmd).Standalone()

	docdbElastic_listClusterSnapshotsCmd.Flags().String("cluster-arn", "", "The ARN identifier of the elastic cluster.")
	docdbElastic_listClusterSnapshotsCmd.Flags().String("max-results", "", "The maximum number of elastic cluster snapshot results to receive in the response.")
	docdbElastic_listClusterSnapshotsCmd.Flags().String("next-token", "", "A pagination token provided by a previous request.")
	docdbElastic_listClusterSnapshotsCmd.Flags().String("snapshot-type", "", "The type of cluster snapshots to be returned.")
	docdbElasticCmd.AddCommand(docdbElastic_listClusterSnapshotsCmd)
}
