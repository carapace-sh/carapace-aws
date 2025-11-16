package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_describeDbclustersCmd = &cobra.Command{
	Use:   "describe-dbclusters",
	Short: "Returns information about provisioned Amazon DocumentDB clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_describeDbclustersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_describeDbclustersCmd).Standalone()

		docdb_describeDbclustersCmd.Flags().String("dbcluster-identifier", "", "The user-provided cluster identifier.")
		docdb_describeDbclustersCmd.Flags().String("filters", "", "A filter that specifies one or more clusters to describe.")
		docdb_describeDbclustersCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		docdb_describeDbclustersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	docdbCmd.AddCommand(docdb_describeDbclustersCmd)
}
