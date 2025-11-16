package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbsubnetGroupsCmd = &cobra.Command{
	Use:   "describe-dbsubnet-groups",
	Short: "Returns a list of DBSubnetGroup descriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbsubnetGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeDbsubnetGroupsCmd).Standalone()

		rds_describeDbsubnetGroupsCmd.Flags().String("dbsubnet-group-name", "", "The name of the DB subnet group to return details for.")
		rds_describeDbsubnetGroupsCmd.Flags().String("filters", "", "This parameter isn't currently supported.")
		rds_describeDbsubnetGroupsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous DescribeDBSubnetGroups request.")
		rds_describeDbsubnetGroupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	rdsCmd.AddCommand(rds_describeDbsubnetGroupsCmd)
}
