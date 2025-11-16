package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getCrawlerMetricsCmd = &cobra.Command{
	Use:   "get-crawler-metrics",
	Short: "Retrieves metrics about specified crawlers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getCrawlerMetricsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getCrawlerMetricsCmd).Standalone()

		glue_getCrawlerMetricsCmd.Flags().String("crawler-name-list", "", "A list of the names of crawlers about which to retrieve metrics.")
		glue_getCrawlerMetricsCmd.Flags().String("max-results", "", "The maximum size of a list to return.")
		glue_getCrawlerMetricsCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
	})
	glueCmd.AddCommand(glue_getCrawlerMetricsCmd)
}
