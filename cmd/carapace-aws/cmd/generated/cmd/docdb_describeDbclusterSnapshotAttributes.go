package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_describeDbclusterSnapshotAttributesCmd = &cobra.Command{
	Use:   "describe-dbcluster-snapshot-attributes",
	Short: "Returns a list of cluster snapshot attribute names and values for a manual DB cluster snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_describeDbclusterSnapshotAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_describeDbclusterSnapshotAttributesCmd).Standalone()

		docdb_describeDbclusterSnapshotAttributesCmd.Flags().String("dbcluster-snapshot-identifier", "", "The identifier for the cluster snapshot to describe the attributes for.")
		docdb_describeDbclusterSnapshotAttributesCmd.MarkFlagRequired("dbcluster-snapshot-identifier")
	})
	docdbCmd.AddCommand(docdb_describeDbclusterSnapshotAttributesCmd)
}
