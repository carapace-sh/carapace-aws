package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_batchGetCustomEntityTypesCmd = &cobra.Command{
	Use:   "batch-get-custom-entity-types",
	Short: "Retrieves the details for the custom patterns specified by a list of names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_batchGetCustomEntityTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_batchGetCustomEntityTypesCmd).Standalone()

		glue_batchGetCustomEntityTypesCmd.Flags().String("names", "", "A list of names of the custom patterns that you want to retrieve.")
		glue_batchGetCustomEntityTypesCmd.MarkFlagRequired("names")
	})
	glueCmd.AddCommand(glue_batchGetCustomEntityTypesCmd)
}
