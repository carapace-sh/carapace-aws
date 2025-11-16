package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var curCmd = &cobra.Command{
	Use:   "cur",
	Short: "You can use the Amazon Web Services Cost and Usage Report API to programmatically create, query, and delete Amazon Web Services Cost and Usage Report definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(curCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(curCmd).Standalone()

	})
	rootCmd.AddCommand(curCmd)
}
