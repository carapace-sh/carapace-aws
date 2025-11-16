package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeOptionGroupsCmd = &cobra.Command{
	Use:   "describe-option-groups",
	Short: "Describes the available option groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeOptionGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeOptionGroupsCmd).Standalone()

		rds_describeOptionGroupsCmd.Flags().String("engine-name", "", "A filter to only include option groups associated with this database engine.")
		rds_describeOptionGroupsCmd.Flags().String("filters", "", "This parameter isn't currently supported.")
		rds_describeOptionGroupsCmd.Flags().String("major-engine-version", "", "Filters the list of option groups to only include groups associated with a specific database engine version.")
		rds_describeOptionGroupsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous DescribeOptionGroups request.")
		rds_describeOptionGroupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		rds_describeOptionGroupsCmd.Flags().String("option-group-name", "", "The name of the option group to describe.")
	})
	rdsCmd.AddCommand(rds_describeOptionGroupsCmd)
}
