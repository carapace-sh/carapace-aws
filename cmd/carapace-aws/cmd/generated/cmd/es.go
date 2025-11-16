package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var esCmd = &cobra.Command{
	Use:   "es",
	Short: "Amazon Elasticsearch Configuration Service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(esCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(esCmd).Standalone()

	})
	rootCmd.AddCommand(esCmd)
}
