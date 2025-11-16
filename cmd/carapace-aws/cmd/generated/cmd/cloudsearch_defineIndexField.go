package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_defineIndexFieldCmd = &cobra.Command{
	Use:   "define-index-field",
	Short: "Configures an `IndexField` for the search domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_defineIndexFieldCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearch_defineIndexFieldCmd).Standalone()

		cloudsearch_defineIndexFieldCmd.Flags().String("domain-name", "", "")
		cloudsearch_defineIndexFieldCmd.Flags().String("index-field", "", "The index field and field options you want to configure.")
		cloudsearch_defineIndexFieldCmd.MarkFlagRequired("domain-name")
		cloudsearch_defineIndexFieldCmd.MarkFlagRequired("index-field")
	})
	cloudsearchCmd.AddCommand(cloudsearch_defineIndexFieldCmd)
}
