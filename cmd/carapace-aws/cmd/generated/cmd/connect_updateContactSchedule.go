package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateContactScheduleCmd = &cobra.Command{
	Use:   "update-contact-schedule",
	Short: "Updates the scheduled time of a task contact that is already scheduled.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateContactScheduleCmd).Standalone()

	connect_updateContactScheduleCmd.Flags().String("contact-id", "", "The identifier of the contact.")
	connect_updateContactScheduleCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateContactScheduleCmd.Flags().String("scheduled-time", "", "The timestamp, in Unix Epoch seconds format, at which to start running the inbound flow.")
	connect_updateContactScheduleCmd.MarkFlagRequired("contact-id")
	connect_updateContactScheduleCmd.MarkFlagRequired("instance-id")
	connect_updateContactScheduleCmd.MarkFlagRequired("scheduled-time")
	connectCmd.AddCommand(connect_updateContactScheduleCmd)
}
