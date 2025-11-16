package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_putInstancePublicPortsCmd = &cobra.Command{
	Use:   "put-instance-public-ports",
	Short: "Opens ports for a specific Amazon Lightsail instance, and specifies the IP addresses allowed to connect to the instance through the ports, and the protocol.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_putInstancePublicPortsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_putInstancePublicPortsCmd).Standalone()

		lightsail_putInstancePublicPortsCmd.Flags().String("instance-name", "", "The name of the instance for which to open ports.")
		lightsail_putInstancePublicPortsCmd.Flags().String("port-infos", "", "An array of objects to describe the ports to open for the specified instance.")
		lightsail_putInstancePublicPortsCmd.MarkFlagRequired("instance-name")
		lightsail_putInstancePublicPortsCmd.MarkFlagRequired("port-infos")
	})
	lightsailCmd.AddCommand(lightsail_putInstancePublicPortsCmd)
}
