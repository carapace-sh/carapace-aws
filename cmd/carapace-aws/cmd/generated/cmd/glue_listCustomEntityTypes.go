package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listCustomEntityTypesCmd = &cobra.Command{
	Use:   "list-custom-entity-types",
	Short: "Lists all the custom patterns that have been created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listCustomEntityTypesCmd).Standalone()

	glue_listCustomEntityTypesCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	glue_listCustomEntityTypesCmd.Flags().String("next-token", "", "A paginated token to offset the results.")
	glue_listCustomEntityTypesCmd.Flags().String("tags", "", "A list of key-value pair tags.")
	glueCmd.AddCommand(glue_listCustomEntityTypesCmd)
}
