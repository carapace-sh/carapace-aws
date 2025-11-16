package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_describeDbsubnetGroupsCmd = &cobra.Command{
	Use:   "describe-dbsubnet-groups",
	Short: "Returns a list of `DBSubnetGroup` descriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_describeDbsubnetGroupsCmd).Standalone()

	docdb_describeDbsubnetGroupsCmd.Flags().String("dbsubnet-group-name", "", "The name of the subnet group to return details for.")
	docdb_describeDbsubnetGroupsCmd.Flags().String("filters", "", "This parameter is not currently supported.")
	docdb_describeDbsubnetGroupsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	docdb_describeDbsubnetGroupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	docdbCmd.AddCommand(docdb_describeDbsubnetGroupsCmd)
}
