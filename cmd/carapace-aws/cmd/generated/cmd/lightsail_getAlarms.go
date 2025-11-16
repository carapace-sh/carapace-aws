package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getAlarmsCmd = &cobra.Command{
	Use:   "get-alarms",
	Short: "Returns information about the configured alarms.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getAlarmsCmd).Standalone()

	lightsail_getAlarmsCmd.Flags().String("alarm-name", "", "The name of the alarm.")
	lightsail_getAlarmsCmd.Flags().String("monitored-resource-name", "", "The name of the Lightsail resource being monitored by the alarm.")
	lightsail_getAlarmsCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	lightsailCmd.AddCommand(lightsail_getAlarmsCmd)
}
