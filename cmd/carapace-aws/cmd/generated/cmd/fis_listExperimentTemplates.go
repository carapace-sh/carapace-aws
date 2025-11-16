package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_listExperimentTemplatesCmd = &cobra.Command{
	Use:   "list-experiment-templates",
	Short: "Lists your experiment templates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_listExperimentTemplatesCmd).Standalone()

	fis_listExperimentTemplatesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	fis_listExperimentTemplatesCmd.Flags().String("next-token", "", "The token for the next page of results.")
	fisCmd.AddCommand(fis_listExperimentTemplatesCmd)
}
