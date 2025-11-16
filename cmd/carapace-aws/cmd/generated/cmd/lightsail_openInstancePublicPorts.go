package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_openInstancePublicPortsCmd = &cobra.Command{
	Use:   "open-instance-public-ports",
	Short: "Opens ports for a specific Amazon Lightsail instance, and specifies the IP addresses allowed to connect to the instance through the ports, and the protocol.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_openInstancePublicPortsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_openInstancePublicPortsCmd).Standalone()

		lightsail_openInstancePublicPortsCmd.Flags().String("instance-name", "", "The name of the instance for which to open ports.")
		lightsail_openInstancePublicPortsCmd.Flags().String("port-info", "", "An object to describe the ports to open for the specified instance.")
		lightsail_openInstancePublicPortsCmd.MarkFlagRequired("instance-name")
		lightsail_openInstancePublicPortsCmd.MarkFlagRequired("port-info")
	})
	lightsailCmd.AddCommand(lightsail_openInstancePublicPortsCmd)
}
