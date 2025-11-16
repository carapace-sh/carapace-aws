package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_describeWebAppCustomizationCmd = &cobra.Command{
	Use:   "describe-web-app-customization",
	Short: "Describes the web app customization object that's identified by `WebAppId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_describeWebAppCustomizationCmd).Standalone()

	transfer_describeWebAppCustomizationCmd.Flags().String("web-app-id", "", "Provide the unique identifier for the web app.")
	transfer_describeWebAppCustomizationCmd.MarkFlagRequired("web-app-id")
	transferCmd.AddCommand(transfer_describeWebAppCustomizationCmd)
}
