package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbsnapshotAttributesCmd = &cobra.Command{
	Use:   "describe-dbsnapshot-attributes",
	Short: "Returns a list of DB snapshot attribute names and values for a manual DB snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbsnapshotAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeDbsnapshotAttributesCmd).Standalone()

		rds_describeDbsnapshotAttributesCmd.Flags().String("dbsnapshot-identifier", "", "The identifier for the DB snapshot to describe the attributes for.")
		rds_describeDbsnapshotAttributesCmd.MarkFlagRequired("dbsnapshot-identifier")
	})
	rdsCmd.AddCommand(rds_describeDbsnapshotAttributesCmd)
}
