package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getContainerServicesCmd = &cobra.Command{
	Use:   "get-container-services",
	Short: "Returns information about one or more of your Amazon Lightsail container services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getContainerServicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getContainerServicesCmd).Standalone()

		lightsail_getContainerServicesCmd.Flags().String("service-name", "", "The name of the container service for which to return information.")
	})
	lightsailCmd.AddCommand(lightsail_getContainerServicesCmd)
}
