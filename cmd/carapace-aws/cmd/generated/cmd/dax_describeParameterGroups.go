package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_describeParameterGroupsCmd = &cobra.Command{
	Use:   "describe-parameter-groups",
	Short: "Returns a list of parameter group descriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_describeParameterGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dax_describeParameterGroupsCmd).Standalone()

		dax_describeParameterGroupsCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		dax_describeParameterGroupsCmd.Flags().String("next-token", "", "An optional token returned from a prior request.")
		dax_describeParameterGroupsCmd.Flags().String("parameter-group-names", "", "The names of the parameter groups.")
	})
	daxCmd.AddCommand(dax_describeParameterGroupsCmd)
}
