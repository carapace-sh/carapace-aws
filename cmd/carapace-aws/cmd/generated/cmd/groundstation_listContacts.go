package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_listContactsCmd = &cobra.Command{
	Use:   "list-contacts",
	Short: "Returns a list of contacts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_listContactsCmd).Standalone()

	groundstation_listContactsCmd.Flags().String("end-time", "", "End time of a contact in UTC.")
	groundstation_listContactsCmd.Flags().String("ephemeris", "", "Filter for selecting contacts that use a specific ephemeris\".")
	groundstation_listContactsCmd.Flags().String("ground-station", "", "Name of a ground station.")
	groundstation_listContactsCmd.Flags().String("max-results", "", "Maximum number of contacts returned.")
	groundstation_listContactsCmd.Flags().String("mission-profile-arn", "", "ARN of a mission profile.")
	groundstation_listContactsCmd.Flags().String("next-token", "", "Next token returned in the request of a previous `ListContacts` call.")
	groundstation_listContactsCmd.Flags().String("satellite-arn", "", "ARN of a satellite.")
	groundstation_listContactsCmd.Flags().String("start-time", "", "Start time of a contact in UTC.")
	groundstation_listContactsCmd.Flags().String("status-list", "", "Status of a contact reservation.")
	groundstation_listContactsCmd.MarkFlagRequired("end-time")
	groundstation_listContactsCmd.MarkFlagRequired("start-time")
	groundstation_listContactsCmd.MarkFlagRequired("status-list")
	groundstationCmd.AddCommand(groundstation_listContactsCmd)
}
