package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElasticCmd = &cobra.Command{
	Use:   "docdb-elastic",
	Short: "Amazon DocumentDB elastic clusters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElasticCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdbElasticCmd).Standalone()

	})
	rootCmd.AddCommand(docdbElasticCmd)
}
