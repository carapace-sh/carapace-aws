package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_stopCrawlerCmd = &cobra.Command{
	Use:   "stop-crawler",
	Short: "If the specified crawler is running, stops the crawl.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_stopCrawlerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_stopCrawlerCmd).Standalone()

		glue_stopCrawlerCmd.Flags().String("name", "", "Name of the crawler to stop.")
		glue_stopCrawlerCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_stopCrawlerCmd)
}
