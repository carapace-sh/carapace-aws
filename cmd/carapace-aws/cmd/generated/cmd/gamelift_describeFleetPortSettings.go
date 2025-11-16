package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeFleetPortSettingsCmd = &cobra.Command{
	Use:   "describe-fleet-port-settings",
	Short: "**This API works with the following fleet types:** EC2, Container\n\nRetrieves a fleet's inbound connection permissions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeFleetPortSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_describeFleetPortSettingsCmd).Standalone()

		gamelift_describeFleetPortSettingsCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to retrieve port settings for.")
		gamelift_describeFleetPortSettingsCmd.Flags().String("location", "", "A remote location to check for status of port setting updates.")
		gamelift_describeFleetPortSettingsCmd.MarkFlagRequired("fleet-id")
	})
	gameliftCmd.AddCommand(gamelift_describeFleetPortSettingsCmd)
}
