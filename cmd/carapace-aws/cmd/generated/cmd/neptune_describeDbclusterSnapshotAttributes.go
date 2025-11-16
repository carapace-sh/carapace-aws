package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_describeDbclusterSnapshotAttributesCmd = &cobra.Command{
	Use:   "describe-dbcluster-snapshot-attributes",
	Short: "Returns a list of DB cluster snapshot attribute names and values for a manual DB cluster snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_describeDbclusterSnapshotAttributesCmd).Standalone()

	neptune_describeDbclusterSnapshotAttributesCmd.Flags().String("dbcluster-snapshot-identifier", "", "The identifier for the DB cluster snapshot to describe the attributes for.")
	neptune_describeDbclusterSnapshotAttributesCmd.MarkFlagRequired("dbcluster-snapshot-identifier")
	neptuneCmd.AddCommand(neptune_describeDbclusterSnapshotAttributesCmd)
}
