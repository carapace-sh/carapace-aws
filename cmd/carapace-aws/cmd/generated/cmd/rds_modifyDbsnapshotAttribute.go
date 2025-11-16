package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_modifyDbsnapshotAttributeCmd = &cobra.Command{
	Use:   "modify-dbsnapshot-attribute",
	Short: "Adds an attribute and values to, or removes an attribute and values from, a manual DB snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_modifyDbsnapshotAttributeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_modifyDbsnapshotAttributeCmd).Standalone()

		rds_modifyDbsnapshotAttributeCmd.Flags().String("attribute-name", "", "The name of the DB snapshot attribute to modify.")
		rds_modifyDbsnapshotAttributeCmd.Flags().String("dbsnapshot-identifier", "", "The identifier for the DB snapshot to modify the attributes for.")
		rds_modifyDbsnapshotAttributeCmd.Flags().String("values-to-add", "", "A list of DB snapshot attributes to add to the attribute specified by `AttributeName`.")
		rds_modifyDbsnapshotAttributeCmd.Flags().String("values-to-remove", "", "A list of DB snapshot attributes to remove from the attribute specified by `AttributeName`.")
		rds_modifyDbsnapshotAttributeCmd.MarkFlagRequired("attribute-name")
		rds_modifyDbsnapshotAttributeCmd.MarkFlagRequired("dbsnapshot-identifier")
	})
	rdsCmd.AddCommand(rds_modifyDbsnapshotAttributeCmd)
}
