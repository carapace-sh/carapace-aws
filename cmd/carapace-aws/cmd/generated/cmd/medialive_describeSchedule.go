package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_describeScheduleCmd = &cobra.Command{
	Use:   "describe-schedule",
	Short: "Get a channel schedule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_describeScheduleCmd).Standalone()

	medialive_describeScheduleCmd.Flags().String("channel-id", "", "Id of the channel whose schedule is being updated.")
	medialive_describeScheduleCmd.Flags().String("max-results", "", "")
	medialive_describeScheduleCmd.Flags().String("next-token", "", "")
	medialive_describeScheduleCmd.MarkFlagRequired("channel-id")
	medialiveCmd.AddCommand(medialive_describeScheduleCmd)
}
