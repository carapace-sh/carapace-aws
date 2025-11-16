package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbCmd = &cobra.Command{
	Use:   "docdb",
	Short: "Amazon DocumentDB is a fast, reliable, and fully managed database service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdbCmd).Standalone()

	})
	rootCmd.AddCommand(docdbCmd)
}
