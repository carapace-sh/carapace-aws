package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_modifyDbclusterSnapshotAttributeCmd = &cobra.Command{
	Use:   "modify-dbcluster-snapshot-attribute",
	Short: "Adds an attribute and values to, or removes an attribute and values from, a manual DB cluster snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_modifyDbclusterSnapshotAttributeCmd).Standalone()

	neptune_modifyDbclusterSnapshotAttributeCmd.Flags().String("attribute-name", "", "The name of the DB cluster snapshot attribute to modify.")
	neptune_modifyDbclusterSnapshotAttributeCmd.Flags().String("dbcluster-snapshot-identifier", "", "The identifier for the DB cluster snapshot to modify the attributes for.")
	neptune_modifyDbclusterSnapshotAttributeCmd.Flags().String("values-to-add", "", "A list of DB cluster snapshot attributes to add to the attribute specified by `AttributeName`.")
	neptune_modifyDbclusterSnapshotAttributeCmd.Flags().String("values-to-remove", "", "A list of DB cluster snapshot attributes to remove from the attribute specified by `AttributeName`.")
	neptune_modifyDbclusterSnapshotAttributeCmd.MarkFlagRequired("attribute-name")
	neptune_modifyDbclusterSnapshotAttributeCmd.MarkFlagRequired("dbcluster-snapshot-identifier")
	neptuneCmd.AddCommand(neptune_modifyDbclusterSnapshotAttributeCmd)
}
