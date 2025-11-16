package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_batchGetCrawlersCmd = &cobra.Command{
	Use:   "batch-get-crawlers",
	Short: "Returns a list of resource metadata for a given list of crawler names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_batchGetCrawlersCmd).Standalone()

	glue_batchGetCrawlersCmd.Flags().String("crawler-names", "", "A list of crawler names, which might be the names returned from the `ListCrawlers` operation.")
	glue_batchGetCrawlersCmd.MarkFlagRequired("crawler-names")
	glueCmd.AddCommand(glue_batchGetCrawlersCmd)
}
