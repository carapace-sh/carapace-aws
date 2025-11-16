package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_listSolFunctionInstancesCmd = &cobra.Command{
	Use:   "list-sol-function-instances",
	Short: "Lists network function instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_listSolFunctionInstancesCmd).Standalone()

	tnb_listSolFunctionInstancesCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
	tnb_listSolFunctionInstancesCmd.Flags().String("next-token", "", "The token for the next page of results.")
	tnbCmd.AddCommand(tnb_listSolFunctionInstancesCmd)
}
