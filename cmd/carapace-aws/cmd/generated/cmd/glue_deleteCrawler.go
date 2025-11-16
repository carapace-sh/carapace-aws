package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteCrawlerCmd = &cobra.Command{
	Use:   "delete-crawler",
	Short: "Removes a specified crawler from the Glue Data Catalog, unless the crawler state is `RUNNING`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteCrawlerCmd).Standalone()

	glue_deleteCrawlerCmd.Flags().String("name", "", "The name of the crawler to remove.")
	glue_deleteCrawlerCmd.MarkFlagRequired("name")
	glueCmd.AddCommand(glue_deleteCrawlerCmd)
}
