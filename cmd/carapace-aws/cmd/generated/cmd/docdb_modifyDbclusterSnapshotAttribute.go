package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_modifyDbclusterSnapshotAttributeCmd = &cobra.Command{
	Use:   "modify-dbcluster-snapshot-attribute",
	Short: "Adds an attribute and values to, or removes an attribute and values from, a manual cluster snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_modifyDbclusterSnapshotAttributeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_modifyDbclusterSnapshotAttributeCmd).Standalone()

		docdb_modifyDbclusterSnapshotAttributeCmd.Flags().String("attribute-name", "", "The name of the cluster snapshot attribute to modify.")
		docdb_modifyDbclusterSnapshotAttributeCmd.Flags().String("dbcluster-snapshot-identifier", "", "The identifier for the cluster snapshot to modify the attributes for.")
		docdb_modifyDbclusterSnapshotAttributeCmd.Flags().String("values-to-add", "", "A list of cluster snapshot attributes to add to the attribute specified by `AttributeName`.")
		docdb_modifyDbclusterSnapshotAttributeCmd.Flags().String("values-to-remove", "", "A list of cluster snapshot attributes to remove from the attribute specified by `AttributeName`.")
		docdb_modifyDbclusterSnapshotAttributeCmd.MarkFlagRequired("attribute-name")
		docdb_modifyDbclusterSnapshotAttributeCmd.MarkFlagRequired("dbcluster-snapshot-identifier")
	})
	docdbCmd.AddCommand(docdb_modifyDbclusterSnapshotAttributeCmd)
}
