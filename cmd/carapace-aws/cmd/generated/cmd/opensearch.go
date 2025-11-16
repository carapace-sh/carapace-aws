package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchCmd = &cobra.Command{
	Use:   "opensearch",
	Short: "Use the Amazon OpenSearch Service configuration API to create, configure, and manage OpenSearch Service domains.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchCmd).Standalone()

	rootCmd.AddCommand(opensearchCmd)
}
