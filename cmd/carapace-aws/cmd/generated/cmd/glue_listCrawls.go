package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listCrawlsCmd = &cobra.Command{
	Use:   "list-crawls",
	Short: "Returns all the crawls of a specified crawler.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listCrawlsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_listCrawlsCmd).Standalone()

		glue_listCrawlsCmd.Flags().String("crawler-name", "", "The name of the crawler whose runs you want to retrieve.")
		glue_listCrawlsCmd.Flags().String("filters", "", "Filters the crawls by the criteria you specify in a list of `CrawlsFilter` objects.")
		glue_listCrawlsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		glue_listCrawlsCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
		glue_listCrawlsCmd.MarkFlagRequired("crawler-name")
	})
	glueCmd.AddCommand(glue_listCrawlsCmd)
}
