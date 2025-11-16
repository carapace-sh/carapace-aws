package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listCrawlersCmd = &cobra.Command{
	Use:   "list-crawlers",
	Short: "Retrieves the names of all crawler resources in this Amazon Web Services account, or the resources with the specified tag.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listCrawlersCmd).Standalone()

	glue_listCrawlersCmd.Flags().String("max-results", "", "The maximum size of a list to return.")
	glue_listCrawlersCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation request.")
	glue_listCrawlersCmd.Flags().String("tags", "", "Specifies to return only these tagged resources.")
	glueCmd.AddCommand(glue_listCrawlersCmd)
}
