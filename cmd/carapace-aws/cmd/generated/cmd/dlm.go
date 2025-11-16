package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dlmCmd = &cobra.Command{
	Use:   "dlm",
	Short: "Amazon Data Lifecycle Manager\n\nWith Amazon Data Lifecycle Manager, you can manage the lifecycle of your Amazon Web Services resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dlmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dlmCmd).Standalone()

	})
	rootCmd.AddCommand(dlmCmd)
}
