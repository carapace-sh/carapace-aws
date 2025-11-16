package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_describeSubnetGroupsCmd = &cobra.Command{
	Use:   "describe-subnet-groups",
	Short: "Returns a list of subnet group descriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_describeSubnetGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dax_describeSubnetGroupsCmd).Standalone()

		dax_describeSubnetGroupsCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		dax_describeSubnetGroupsCmd.Flags().String("next-token", "", "An optional token returned from a prior request.")
		dax_describeSubnetGroupsCmd.Flags().String("subnet-group-names", "", "The name of the subnet group.")
	})
	daxCmd.AddCommand(dax_describeSubnetGroupsCmd)
}
