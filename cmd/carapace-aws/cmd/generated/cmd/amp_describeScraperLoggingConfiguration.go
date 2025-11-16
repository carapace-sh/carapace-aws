package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_describeScraperLoggingConfigurationCmd = &cobra.Command{
	Use:   "describe-scraper-logging-configuration",
	Short: "Describes the logging configuration for a Amazon Managed Service for Prometheus scraper.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_describeScraperLoggingConfigurationCmd).Standalone()

	amp_describeScraperLoggingConfigurationCmd.Flags().String("scraper-id", "", "The ID of the scraper whose logging configuration will be described.")
	amp_describeScraperLoggingConfigurationCmd.MarkFlagRequired("scraper-id")
	ampCmd.AddCommand(amp_describeScraperLoggingConfigurationCmd)
}
