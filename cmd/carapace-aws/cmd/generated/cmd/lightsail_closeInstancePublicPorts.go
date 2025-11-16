package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_closeInstancePublicPortsCmd = &cobra.Command{
	Use:   "close-instance-public-ports",
	Short: "Closes ports for a specific Amazon Lightsail instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_closeInstancePublicPortsCmd).Standalone()

	lightsail_closeInstancePublicPortsCmd.Flags().String("instance-name", "", "The name of the instance for which to close ports.")
	lightsail_closeInstancePublicPortsCmd.Flags().String("port-info", "", "An object to describe the ports to close for the specified instance.")
	lightsail_closeInstancePublicPortsCmd.MarkFlagRequired("instance-name")
	lightsail_closeInstancePublicPortsCmd.MarkFlagRequired("port-info")
	lightsailCmd.AddCommand(lightsail_closeInstancePublicPortsCmd)
}
