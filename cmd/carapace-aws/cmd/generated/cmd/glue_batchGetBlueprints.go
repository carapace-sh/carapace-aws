package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_batchGetBlueprintsCmd = &cobra.Command{
	Use:   "batch-get-blueprints",
	Short: "Retrieves information about a list of blueprints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_batchGetBlueprintsCmd).Standalone()

	glue_batchGetBlueprintsCmd.Flags().String("include-blueprint", "", "Specifies whether or not to include the blueprint in the response.")
	glue_batchGetBlueprintsCmd.Flags().String("include-parameter-spec", "", "Specifies whether or not to include the parameters, as a JSON string, for the blueprint in the response.")
	glue_batchGetBlueprintsCmd.Flags().String("names", "", "A list of blueprint names.")
	glue_batchGetBlueprintsCmd.MarkFlagRequired("names")
	glueCmd.AddCommand(glue_batchGetBlueprintsCmd)
}
