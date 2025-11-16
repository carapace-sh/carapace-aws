package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_listExperimentResolvedTargetsCmd = &cobra.Command{
	Use:   "list-experiment-resolved-targets",
	Short: "Lists the resolved targets information of the specified experiment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_listExperimentResolvedTargetsCmd).Standalone()

	fis_listExperimentResolvedTargetsCmd.Flags().String("experiment-id", "", "The ID of the experiment.")
	fis_listExperimentResolvedTargetsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	fis_listExperimentResolvedTargetsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	fis_listExperimentResolvedTargetsCmd.Flags().String("target-name", "", "The name of the target.")
	fis_listExperimentResolvedTargetsCmd.MarkFlagRequired("experiment-id")
	fisCmd.AddCommand(fis_listExperimentResolvedTargetsCmd)
}
