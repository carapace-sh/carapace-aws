package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_updateFleetPortSettingsCmd = &cobra.Command{
	Use:   "update-fleet-port-settings",
	Short: "**This API works with the following fleet types:** EC2, Container\n\nUpdates permissions that allow inbound traffic to connect to game sessions in the fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_updateFleetPortSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_updateFleetPortSettingsCmd).Standalone()

		gamelift_updateFleetPortSettingsCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to update port settings for.")
		gamelift_updateFleetPortSettingsCmd.Flags().String("inbound-permission-authorizations", "", "A collection of port settings to be added to the fleet resource.")
		gamelift_updateFleetPortSettingsCmd.Flags().String("inbound-permission-revocations", "", "A collection of port settings to be removed from the fleet resource.")
		gamelift_updateFleetPortSettingsCmd.MarkFlagRequired("fleet-id")
	})
	gameliftCmd.AddCommand(gamelift_updateFleetPortSettingsCmd)
}
