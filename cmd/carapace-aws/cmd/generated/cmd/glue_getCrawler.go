package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getCrawlerCmd = &cobra.Command{
	Use:   "get-crawler",
	Short: "Retrieves metadata for a specified crawler.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getCrawlerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getCrawlerCmd).Standalone()

		glue_getCrawlerCmd.Flags().String("name", "", "The name of the crawler to retrieve metadata for.")
		glue_getCrawlerCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_getCrawlerCmd)
}
