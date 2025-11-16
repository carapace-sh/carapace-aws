package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_describeScheduleCmd = &cobra.Command{
	Use:   "describe-schedule",
	Short: "Returns the definition of a specific DataBrew schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_describeScheduleCmd).Standalone()

	databrew_describeScheduleCmd.Flags().String("name", "", "The name of the schedule to be described.")
	databrew_describeScheduleCmd.MarkFlagRequired("name")
	databrewCmd.AddCommand(databrew_describeScheduleCmd)
}
