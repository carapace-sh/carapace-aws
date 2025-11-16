package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_getExperimentResultsCmd = &cobra.Command{
	Use:   "get-experiment-results",
	Short: "Retrieves the results of a running or completed experiment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_getExperimentResultsCmd).Standalone()

	evidently_getExperimentResultsCmd.Flags().String("base-stat", "", "The statistic used to calculate experiment results.")
	evidently_getExperimentResultsCmd.Flags().String("end-time", "", "The date and time that the experiment ended, if it is completed.")
	evidently_getExperimentResultsCmd.Flags().String("experiment", "", "The name of the experiment to retrieve the results of.")
	evidently_getExperimentResultsCmd.Flags().String("metric-names", "", "The names of the experiment metrics that you want to see the results of.")
	evidently_getExperimentResultsCmd.Flags().String("period", "", "In seconds, the amount of time to aggregate results together.")
	evidently_getExperimentResultsCmd.Flags().String("project", "", "The name or ARN of the project that contains the experiment that you want to see the results of.")
	evidently_getExperimentResultsCmd.Flags().String("report-names", "", "The names of the report types that you want to see.")
	evidently_getExperimentResultsCmd.Flags().String("result-stats", "", "The statistics that you want to see in the returned results.")
	evidently_getExperimentResultsCmd.Flags().String("start-time", "", "The date and time that the experiment started.")
	evidently_getExperimentResultsCmd.Flags().String("treatment-names", "", "The names of the experiment treatments that you want to see the results for.")
	evidently_getExperimentResultsCmd.MarkFlagRequired("experiment")
	evidently_getExperimentResultsCmd.MarkFlagRequired("metric-names")
	evidently_getExperimentResultsCmd.MarkFlagRequired("project")
	evidently_getExperimentResultsCmd.MarkFlagRequired("treatment-names")
	evidentlyCmd.AddCommand(evidently_getExperimentResultsCmd)
}
