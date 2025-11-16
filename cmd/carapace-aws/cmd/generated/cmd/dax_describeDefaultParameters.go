package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_describeDefaultParametersCmd = &cobra.Command{
	Use:   "describe-default-parameters",
	Short: "Returns the default system parameter information for the DAX caching software.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_describeDefaultParametersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dax_describeDefaultParametersCmd).Standalone()

		dax_describeDefaultParametersCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		dax_describeDefaultParametersCmd.Flags().String("next-token", "", "An optional token returned from a prior request.")
	})
	daxCmd.AddCommand(dax_describeDefaultParametersCmd)
}
