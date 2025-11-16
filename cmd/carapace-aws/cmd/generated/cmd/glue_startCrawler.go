package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_startCrawlerCmd = &cobra.Command{
	Use:   "start-crawler",
	Short: "Starts a crawl using the specified crawler, regardless of what is scheduled.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_startCrawlerCmd).Standalone()

	glue_startCrawlerCmd.Flags().String("name", "", "Name of the crawler to start.")
	glue_startCrawlerCmd.MarkFlagRequired("name")
	glueCmd.AddCommand(glue_startCrawlerCmd)
}
