package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_resumeResourceCmd = &cobra.Command{
	Use:   "resume-resource",
	Short: "Resumes a stopped monitor resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_resumeResourceCmd).Standalone()

	forecast_resumeResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the monitor resource to resume.")
	forecast_resumeResourceCmd.MarkFlagRequired("resource-arn")
	forecastCmd.AddCommand(forecast_resumeResourceCmd)
}
