package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_describeAppLicenseUsageCmd = &cobra.Command{
	Use:   "describe-app-license-usage",
	Short: "Retrieves license included application usage information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_describeAppLicenseUsageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_describeAppLicenseUsageCmd).Standalone()

		appstream_describeAppLicenseUsageCmd.Flags().String("billing-period", "", "Billing period for the usage record.")
		appstream_describeAppLicenseUsageCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		appstream_describeAppLicenseUsageCmd.Flags().String("next-token", "", "Token for pagination of results.")
		appstream_describeAppLicenseUsageCmd.MarkFlagRequired("billing-period")
	})
	appstreamCmd.AddCommand(appstream_describeAppLicenseUsageCmd)
}
