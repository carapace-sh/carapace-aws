package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createContainerServiceRegistryLoginCmd = &cobra.Command{
	Use:   "create-container-service-registry-login",
	Short: "Creates a temporary set of log in credentials that you can use to log in to the Docker process on your local machine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createContainerServiceRegistryLoginCmd).Standalone()

	lightsailCmd.AddCommand(lightsail_createContainerServiceRegistryLoginCmd)
}
