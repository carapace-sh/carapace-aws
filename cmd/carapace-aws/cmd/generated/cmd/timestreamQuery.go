package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamQueryCmd = &cobra.Command{
	Use:   "timestream-query",
	Short: "Amazon Timestream Query",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamQueryCmd).Standalone()

	})
	rootCmd.AddCommand(timestreamQueryCmd)
}
