package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_listExperimentsCmd = &cobra.Command{
	Use:   "list-experiments",
	Short: "Lists your experiments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_listExperimentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fis_listExperimentsCmd).Standalone()

		fis_listExperimentsCmd.Flags().String("experiment-template-id", "", "The ID of the experiment template.")
		fis_listExperimentsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		fis_listExperimentsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	})
	fisCmd.AddCommand(fis_listExperimentsCmd)
}
