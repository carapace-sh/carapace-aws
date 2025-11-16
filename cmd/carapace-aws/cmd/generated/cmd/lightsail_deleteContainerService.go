package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteContainerServiceCmd = &cobra.Command{
	Use:   "delete-container-service",
	Short: "Deletes your Amazon Lightsail container service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteContainerServiceCmd).Standalone()

	lightsail_deleteContainerServiceCmd.Flags().String("service-name", "", "The name of the container service to delete.")
	lightsail_deleteContainerServiceCmd.MarkFlagRequired("service-name")
	lightsailCmd.AddCommand(lightsail_deleteContainerServiceCmd)
}
