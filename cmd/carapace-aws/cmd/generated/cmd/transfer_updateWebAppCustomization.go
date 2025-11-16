package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_updateWebAppCustomizationCmd = &cobra.Command{
	Use:   "update-web-app-customization",
	Short: "Assigns new customization properties to a web app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_updateWebAppCustomizationCmd).Standalone()

	transfer_updateWebAppCustomizationCmd.Flags().String("favicon-file", "", "Specify an icon file data string (in base64 encoding).")
	transfer_updateWebAppCustomizationCmd.Flags().String("logo-file", "", "Specify logo file data string (in base64 encoding).")
	transfer_updateWebAppCustomizationCmd.Flags().String("title", "", "Provide an updated title.")
	transfer_updateWebAppCustomizationCmd.Flags().String("web-app-id", "", "Provide the identifier of the web app that you are updating.")
	transfer_updateWebAppCustomizationCmd.MarkFlagRequired("web-app-id")
	transferCmd.AddCommand(transfer_updateWebAppCustomizationCmd)
}
