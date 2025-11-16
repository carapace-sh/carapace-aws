package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getContainerServicePowersCmd = &cobra.Command{
	Use:   "get-container-service-powers",
	Short: "Returns the list of powers that can be specified for your Amazon Lightsail container services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getContainerServicePowersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getContainerServicePowersCmd).Standalone()

	})
	lightsailCmd.AddCommand(lightsail_getContainerServicePowersCmd)
}
