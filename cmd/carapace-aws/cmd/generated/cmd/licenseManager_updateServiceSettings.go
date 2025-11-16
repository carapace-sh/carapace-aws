package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManager_updateServiceSettingsCmd = &cobra.Command{
	Use:   "update-service-settings",
	Short: "Updates License Manager settings for the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManager_updateServiceSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManager_updateServiceSettingsCmd).Standalone()

		licenseManager_updateServiceSettingsCmd.Flags().String("enable-cross-accounts-discovery", "", "Activates cross-account discovery.")
		licenseManager_updateServiceSettingsCmd.Flags().String("organization-configuration", "", "Enables integration with Organizations for cross-account discovery.")
		licenseManager_updateServiceSettingsCmd.Flags().String("s3-bucket-arn", "", "Amazon Resource Name (ARN) of the Amazon S3 bucket where the License Manager information is stored.")
		licenseManager_updateServiceSettingsCmd.Flags().String("sns-topic-arn", "", "Amazon Resource Name (ARN) of the Amazon SNS topic used for License Manager alerts.")
	})
	licenseManagerCmd.AddCommand(licenseManager_updateServiceSettingsCmd)
}
