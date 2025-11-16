package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_reserveContactCmd = &cobra.Command{
	Use:   "reserve-contact",
	Short: "Reserves a contact using specified parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_reserveContactCmd).Standalone()

	groundstation_reserveContactCmd.Flags().String("end-time", "", "End time of a contact in UTC.")
	groundstation_reserveContactCmd.Flags().String("ground-station", "", "Name of a ground station.")
	groundstation_reserveContactCmd.Flags().String("mission-profile-arn", "", "ARN of a mission profile.")
	groundstation_reserveContactCmd.Flags().String("satellite-arn", "", "ARN of a satellite")
	groundstation_reserveContactCmd.Flags().String("start-time", "", "Start time of a contact in UTC.")
	groundstation_reserveContactCmd.Flags().String("tags", "", "Tags assigned to a contact.")
	groundstation_reserveContactCmd.Flags().String("tracking-overrides", "", "Tracking configuration overrides for the contact.")
	groundstation_reserveContactCmd.MarkFlagRequired("end-time")
	groundstation_reserveContactCmd.MarkFlagRequired("ground-station")
	groundstation_reserveContactCmd.MarkFlagRequired("mission-profile-arn")
	groundstation_reserveContactCmd.MarkFlagRequired("start-time")
	groundstationCmd.AddCommand(groundstation_reserveContactCmd)
}
