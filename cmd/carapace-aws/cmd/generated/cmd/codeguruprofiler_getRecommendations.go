package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_getRecommendationsCmd = &cobra.Command{
	Use:   "get-recommendations",
	Short: "Returns a list of [`Recommendation`](https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_Recommendation.html) objects that contain recommendations for a profiling group for a given time period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_getRecommendationsCmd).Standalone()

	codeguruprofiler_getRecommendationsCmd.Flags().String("end-time", "", "The start time of the profile to get analysis data about.")
	codeguruprofiler_getRecommendationsCmd.Flags().String("locale", "", "The language used to provide analysis.")
	codeguruprofiler_getRecommendationsCmd.Flags().String("profiling-group-name", "", "The name of the profiling group to get analysis data about.")
	codeguruprofiler_getRecommendationsCmd.Flags().String("start-time", "", "The end time of the profile to get analysis data about.")
	codeguruprofiler_getRecommendationsCmd.MarkFlagRequired("end-time")
	codeguruprofiler_getRecommendationsCmd.MarkFlagRequired("profiling-group-name")
	codeguruprofiler_getRecommendationsCmd.MarkFlagRequired("start-time")
	codeguruprofilerCmd.AddCommand(codeguruprofiler_getRecommendationsCmd)
}
