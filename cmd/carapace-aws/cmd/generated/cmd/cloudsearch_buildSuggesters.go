package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_buildSuggestersCmd = &cobra.Command{
	Use:   "build-suggesters",
	Short: "Indexes the search suggestions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_buildSuggestersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearch_buildSuggestersCmd).Standalone()

		cloudsearch_buildSuggestersCmd.Flags().String("domain-name", "", "")
		cloudsearch_buildSuggestersCmd.MarkFlagRequired("domain-name")
	})
	cloudsearchCmd.AddCommand(cloudsearch_buildSuggestersCmd)
}
