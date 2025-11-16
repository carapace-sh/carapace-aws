package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_describeDbsubnetGroupsCmd = &cobra.Command{
	Use:   "describe-dbsubnet-groups",
	Short: "Returns a list of DBSubnetGroup descriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_describeDbsubnetGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_describeDbsubnetGroupsCmd).Standalone()

		neptune_describeDbsubnetGroupsCmd.Flags().String("dbsubnet-group-name", "", "The name of the DB subnet group to return details for.")
		neptune_describeDbsubnetGroupsCmd.Flags().String("filters", "", "This parameter is not currently supported.")
		neptune_describeDbsubnetGroupsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous DescribeDBSubnetGroups request.")
		neptune_describeDbsubnetGroupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	neptuneCmd.AddCommand(neptune_describeDbsubnetGroupsCmd)
}
