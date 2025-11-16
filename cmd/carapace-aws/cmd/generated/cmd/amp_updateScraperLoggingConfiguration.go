package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_updateScraperLoggingConfigurationCmd = &cobra.Command{
	Use:   "update-scraper-logging-configuration",
	Short: "Updates the logging configuration for a Amazon Managed Service for Prometheus scraper.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_updateScraperLoggingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_updateScraperLoggingConfigurationCmd).Standalone()

		amp_updateScraperLoggingConfigurationCmd.Flags().String("logging-destination", "", "The destination where scraper logs will be sent.")
		amp_updateScraperLoggingConfigurationCmd.Flags().String("scraper-components", "", "The list of scraper components to configure for logging.")
		amp_updateScraperLoggingConfigurationCmd.Flags().String("scraper-id", "", "The ID of the scraper whose logging configuration will be updated.")
		amp_updateScraperLoggingConfigurationCmd.MarkFlagRequired("logging-destination")
		amp_updateScraperLoggingConfigurationCmd.MarkFlagRequired("scraper-id")
	})
	ampCmd.AddCommand(amp_updateScraperLoggingConfigurationCmd)
}
