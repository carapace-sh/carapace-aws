package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putConfigurationSetTrackingOptionsCmd = &cobra.Command{
	Use:   "put-configuration-set-tracking-options",
	Short: "Specify a custom domain to use for open and click tracking elements in email that you send.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putConfigurationSetTrackingOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_putConfigurationSetTrackingOptionsCmd).Standalone()

		sesv2_putConfigurationSetTrackingOptionsCmd.Flags().String("configuration-set-name", "", "The name of the configuration set.")
		sesv2_putConfigurationSetTrackingOptionsCmd.Flags().String("custom-redirect-domain", "", "The domain to use to track open and click events.")
		sesv2_putConfigurationSetTrackingOptionsCmd.Flags().String("https-policy", "", "")
		sesv2_putConfigurationSetTrackingOptionsCmd.MarkFlagRequired("configuration-set-name")
	})
	sesv2Cmd.AddCommand(sesv2_putConfigurationSetTrackingOptionsCmd)
}
