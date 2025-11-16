package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getClassifiersCmd = &cobra.Command{
	Use:   "get-classifiers",
	Short: "Lists all classifier objects in the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getClassifiersCmd).Standalone()

	glue_getClassifiersCmd.Flags().String("max-results", "", "The size of the list to return (optional).")
	glue_getClassifiersCmd.Flags().String("next-token", "", "An optional continuation token.")
	glueCmd.AddCommand(glue_getClassifiersCmd)
}
