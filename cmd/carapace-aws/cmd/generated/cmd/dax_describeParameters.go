package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_describeParametersCmd = &cobra.Command{
	Use:   "describe-parameters",
	Short: "Returns the detailed parameter list for a particular parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_describeParametersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dax_describeParametersCmd).Standalone()

		dax_describeParametersCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		dax_describeParametersCmd.Flags().String("next-token", "", "An optional token returned from a prior request.")
		dax_describeParametersCmd.Flags().String("parameter-group-name", "", "The name of the parameter group.")
		dax_describeParametersCmd.Flags().String("source", "", "How the parameter is defined.")
		dax_describeParametersCmd.MarkFlagRequired("parameter-group-name")
	})
	daxCmd.AddCommand(dax_describeParametersCmd)
}
