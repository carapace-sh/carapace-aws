package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getBlueprintCmd = &cobra.Command{
	Use:   "get-blueprint",
	Short: "Retrieves the details of a blueprint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getBlueprintCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getBlueprintCmd).Standalone()

		glue_getBlueprintCmd.Flags().String("include-blueprint", "", "Specifies whether or not to include the blueprint in the response.")
		glue_getBlueprintCmd.Flags().String("include-parameter-spec", "", "Specifies whether or not to include the parameter specification.")
		glue_getBlueprintCmd.Flags().String("name", "", "The name of the blueprint.")
		glue_getBlueprintCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_getBlueprintCmd)
}
