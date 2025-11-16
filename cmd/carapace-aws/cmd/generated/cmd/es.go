package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var esCmd = &cobra.Command{
	Use:   "es",
	Short: "Amazon Elasticsearch Configuration Service\n\nUse the Amazon Elasticsearch Configuration API to create, configure, and manage Elasticsearch domains.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(esCmd).Standalone()

	rootCmd.AddCommand(esCmd)
}
