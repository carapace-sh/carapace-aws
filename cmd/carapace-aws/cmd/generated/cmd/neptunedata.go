package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedataCmd = &cobra.Command{
	Use:   "neptunedata",
	Short: "Neptune Data API\n\nThe Amazon Neptune data API provides SDK support for more than 40 of Neptune's data operations, including data loading, query execution, data inquiry, and machine learning.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedataCmd).Standalone()

	rootCmd.AddCommand(neptunedataCmd)
}
