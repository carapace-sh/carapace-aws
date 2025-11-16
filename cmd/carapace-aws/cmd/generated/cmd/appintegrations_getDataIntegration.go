package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_getDataIntegrationCmd = &cobra.Command{
	Use:   "get-data-integration",
	Short: "Returns information about the DataIntegration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_getDataIntegrationCmd).Standalone()

	appintegrations_getDataIntegrationCmd.Flags().String("identifier", "", "A unique identifier.")
	appintegrations_getDataIntegrationCmd.MarkFlagRequired("identifier")
	appintegrationsCmd.AddCommand(appintegrations_getDataIntegrationCmd)
}
