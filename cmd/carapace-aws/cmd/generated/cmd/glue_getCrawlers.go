package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getCrawlersCmd = &cobra.Command{
	Use:   "get-crawlers",
	Short: "Retrieves metadata for all crawlers defined in the customer account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getCrawlersCmd).Standalone()

	glue_getCrawlersCmd.Flags().String("max-results", "", "The number of crawlers to return on each call.")
	glue_getCrawlersCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation request.")
	glueCmd.AddCommand(glue_getCrawlersCmd)
}
