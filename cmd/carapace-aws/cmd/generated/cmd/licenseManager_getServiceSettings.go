package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_getServiceSettingsCmd = &cobra.Command{
	Use:   "get-service-settings",
	Short: "Gets the License Manager settings for the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_getServiceSettingsCmd).Standalone()

	licenseManagerCmd.AddCommand(licenseManager_getServiceSettingsCmd)
}
