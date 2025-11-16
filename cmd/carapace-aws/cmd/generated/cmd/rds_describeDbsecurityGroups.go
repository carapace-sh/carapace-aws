package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeDbsecurityGroupsCmd = &cobra.Command{
	Use:   "describe-dbsecurity-groups",
	Short: "Returns a list of `DBSecurityGroup` descriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeDbsecurityGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeDbsecurityGroupsCmd).Standalone()

		rds_describeDbsecurityGroupsCmd.Flags().String("dbsecurity-group-name", "", "The name of the DB security group to return details for.")
		rds_describeDbsecurityGroupsCmd.Flags().String("filters", "", "This parameter isn't currently supported.")
		rds_describeDbsecurityGroupsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeDBSecurityGroups` request.")
		rds_describeDbsecurityGroupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	rdsCmd.AddCommand(rds_describeDbsecurityGroupsCmd)
}
