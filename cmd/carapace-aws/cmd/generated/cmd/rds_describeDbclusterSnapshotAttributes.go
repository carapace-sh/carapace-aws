package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbclusterSnapshotAttributesCmd = &cobra.Command{
	Use:   "describe-dbcluster-snapshot-attributes",
	Short: "Returns a list of DB cluster snapshot attribute names and values for a manual DB cluster snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbclusterSnapshotAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeDbclusterSnapshotAttributesCmd).Standalone()

		rds_describeDbclusterSnapshotAttributesCmd.Flags().String("dbcluster-snapshot-identifier", "", "The identifier for the DB cluster snapshot to describe the attributes for.")
		rds_describeDbclusterSnapshotAttributesCmd.MarkFlagRequired("dbcluster-snapshot-identifier")
	})
	rdsCmd.AddCommand(rds_describeDbclusterSnapshotAttributesCmd)
}
