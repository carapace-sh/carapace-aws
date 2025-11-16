package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listBlueprintsCmd = &cobra.Command{
	Use:   "list-blueprints",
	Short: "Lists all the blueprint names in an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listBlueprintsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_listBlueprintsCmd).Standalone()

		glue_listBlueprintsCmd.Flags().String("max-results", "", "The maximum size of a list to return.")
		glue_listBlueprintsCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation request.")
		glue_listBlueprintsCmd.Flags().String("tags", "", "Filters the list by an Amazon Web Services resource tag.")
	})
	glueCmd.AddCommand(glue_listBlueprintsCmd)
}
