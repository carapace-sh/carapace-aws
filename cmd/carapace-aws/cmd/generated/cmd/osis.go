package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osisCmd = &cobra.Command{
	Use:   "osis",
	Short: "Use the Amazon OpenSearch Ingestion API to create and manage ingestion pipelines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osisCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(osisCmd).Standalone()

	})
	rootCmd.AddCommand(osisCmd)
}
