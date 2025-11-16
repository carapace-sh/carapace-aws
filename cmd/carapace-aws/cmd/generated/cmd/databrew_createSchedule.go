package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_createScheduleCmd = &cobra.Command{
	Use:   "create-schedule",
	Short: "Creates a new schedule for one or more DataBrew jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_createScheduleCmd).Standalone()

	databrew_createScheduleCmd.Flags().String("cron-expression", "", "The date or dates and time or times when the jobs are to be run.")
	databrew_createScheduleCmd.Flags().String("job-names", "", "The name or names of one or more jobs to be run.")
	databrew_createScheduleCmd.Flags().String("name", "", "A unique name for the schedule.")
	databrew_createScheduleCmd.Flags().String("tags", "", "Metadata tags to apply to this schedule.")
	databrew_createScheduleCmd.MarkFlagRequired("cron-expression")
	databrew_createScheduleCmd.MarkFlagRequired("name")
	databrewCmd.AddCommand(databrew_createScheduleCmd)
}
