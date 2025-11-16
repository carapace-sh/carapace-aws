package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_getSellingSystemSettingsCmd = &cobra.Command{
	Use:   "get-selling-system-settings",
	Short: "Retrieves the currently set system settings, which include the IAM Role used for resource snapshot jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_getSellingSystemSettingsCmd).Standalone()

	partnercentralSelling_getSellingSystemSettingsCmd.Flags().String("catalog", "", "Specifies the catalog in which the settings are defined.")
	partnercentralSelling_getSellingSystemSettingsCmd.MarkFlagRequired("catalog")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_getSellingSystemSettingsCmd)
}
