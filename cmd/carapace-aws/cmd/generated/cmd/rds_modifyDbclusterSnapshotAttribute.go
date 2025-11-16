package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyDbclusterSnapshotAttributeCmd = &cobra.Command{
	Use:   "modify-dbcluster-snapshot-attribute",
	Short: "Adds an attribute and values to, or removes an attribute and values from, a manual DB cluster snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyDbclusterSnapshotAttributeCmd).Standalone()

	rds_modifyDbclusterSnapshotAttributeCmd.Flags().String("attribute-name", "", "The name of the DB cluster snapshot attribute to modify.")
	rds_modifyDbclusterSnapshotAttributeCmd.Flags().String("dbcluster-snapshot-identifier", "", "The identifier for the DB cluster snapshot to modify the attributes for.")
	rds_modifyDbclusterSnapshotAttributeCmd.Flags().String("values-to-add", "", "A list of DB cluster snapshot attributes to add to the attribute specified by `AttributeName`.")
	rds_modifyDbclusterSnapshotAttributeCmd.Flags().String("values-to-remove", "", "A list of DB cluster snapshot attributes to remove from the attribute specified by `AttributeName`.")
	rds_modifyDbclusterSnapshotAttributeCmd.MarkFlagRequired("attribute-name")
	rds_modifyDbclusterSnapshotAttributeCmd.MarkFlagRequired("dbcluster-snapshot-identifier")
	rdsCmd.AddCommand(rds_modifyDbclusterSnapshotAttributeCmd)
}
