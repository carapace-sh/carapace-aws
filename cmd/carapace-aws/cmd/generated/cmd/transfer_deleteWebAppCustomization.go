package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_deleteWebAppCustomizationCmd = &cobra.Command{
	Use:   "delete-web-app-customization",
	Short: "Deletes the `WebAppCustomization` object that corresponds to the web app ID specified.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_deleteWebAppCustomizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_deleteWebAppCustomizationCmd).Standalone()

		transfer_deleteWebAppCustomizationCmd.Flags().String("web-app-id", "", "Provide the unique identifier for the web app that contains the customizations that you are deleting.")
		transfer_deleteWebAppCustomizationCmd.MarkFlagRequired("web-app-id")
	})
	transferCmd.AddCommand(transfer_deleteWebAppCustomizationCmd)
}
