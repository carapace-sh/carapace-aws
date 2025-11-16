package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_listSchedulesCmd = &cobra.Command{
	Use:   "list-schedules",
	Short: "Lists the DataBrew schedules that are defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_listSchedulesCmd).Standalone()

	databrew_listSchedulesCmd.Flags().String("job-name", "", "The name of the job that these schedules apply to.")
	databrew_listSchedulesCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
	databrew_listSchedulesCmd.Flags().String("next-token", "", "The token returned by a previous call to retrieve the next set of results.")
	databrewCmd.AddCommand(databrew_listSchedulesCmd)
}
