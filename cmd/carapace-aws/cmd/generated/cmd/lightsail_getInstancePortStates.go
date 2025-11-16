package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getInstancePortStatesCmd = &cobra.Command{
	Use:   "get-instance-port-states",
	Short: "Returns the firewall port states for a specific Amazon Lightsail instance, the IP addresses allowed to connect to the instance through the ports, and the protocol.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getInstancePortStatesCmd).Standalone()

	lightsail_getInstancePortStatesCmd.Flags().String("instance-name", "", "The name of the instance for which to return firewall port states.")
	lightsail_getInstancePortStatesCmd.MarkFlagRequired("instance-name")
	lightsailCmd.AddCommand(lightsail_getInstancePortStatesCmd)
}
