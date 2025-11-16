package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_getLoaderJobStatusCmd = &cobra.Command{
	Use:   "get-loader-job-status",
	Short: "Gets status information about a specified load job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_getLoaderJobStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_getLoaderJobStatusCmd).Standalone()

		neptunedata_getLoaderJobStatusCmd.Flags().Bool("details", false, "Flag indicating whether or not to include details beyond the overall status (`TRUE` or `FALSE`; the default is `FALSE`).")
		neptunedata_getLoaderJobStatusCmd.Flags().Bool("errors", false, "Flag indicating whether or not to include a list of errors encountered (`TRUE` or `FALSE`; the default is `FALSE`).")
		neptunedata_getLoaderJobStatusCmd.Flags().String("errors-per-page", "", "The number of errors returned in each page (a positive integer; the default is `10`).")
		neptunedata_getLoaderJobStatusCmd.Flags().String("load-id", "", "The load ID of the load job to get the status of.")
		neptunedata_getLoaderJobStatusCmd.Flags().Bool("no-details", false, "Flag indicating whether or not to include details beyond the overall status (`TRUE` or `FALSE`; the default is `FALSE`).")
		neptunedata_getLoaderJobStatusCmd.Flags().Bool("no-errors", false, "Flag indicating whether or not to include a list of errors encountered (`TRUE` or `FALSE`; the default is `FALSE`).")
		neptunedata_getLoaderJobStatusCmd.Flags().String("page", "", "The error page number (a positive integer; the default is `1`).")
		neptunedata_getLoaderJobStatusCmd.MarkFlagRequired("load-id")
		neptunedata_getLoaderJobStatusCmd.Flag("no-details").Hidden = true
		neptunedata_getLoaderJobStatusCmd.Flag("no-errors").Hidden = true
	})
	neptunedataCmd.AddCommand(neptunedata_getLoaderJobStatusCmd)
}
