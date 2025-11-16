package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeReplicationConfigsCmd = &cobra.Command{
	Use:   "describe-replication-configs",
	Short: "Returns one or more existing DMS Serverless replication configurations as a list of structures.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeReplicationConfigsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeReplicationConfigsCmd).Standalone()

		dms_describeReplicationConfigsCmd.Flags().String("filters", "", "Filters applied to the replication configs.")
		dms_describeReplicationConfigsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		dms_describeReplicationConfigsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	dmsCmd.AddCommand(dms_describeReplicationConfigsCmd)
}
