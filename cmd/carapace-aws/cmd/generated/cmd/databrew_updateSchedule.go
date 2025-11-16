package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_updateScheduleCmd = &cobra.Command{
	Use:   "update-schedule",
	Short: "Modifies the definition of an existing DataBrew schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_updateScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(databrew_updateScheduleCmd).Standalone()

		databrew_updateScheduleCmd.Flags().String("cron-expression", "", "The date or dates and time or times when the jobs are to be run.")
		databrew_updateScheduleCmd.Flags().String("job-names", "", "The name or names of one or more jobs to be run for this schedule.")
		databrew_updateScheduleCmd.Flags().String("name", "", "The name of the schedule to update.")
		databrew_updateScheduleCmd.MarkFlagRequired("cron-expression")
		databrew_updateScheduleCmd.MarkFlagRequired("name")
	})
	databrewCmd.AddCommand(databrew_updateScheduleCmd)
}
