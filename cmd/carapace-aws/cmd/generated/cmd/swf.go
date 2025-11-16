package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swfCmd = &cobra.Command{
	Use:   "swf",
	Short: "Amazon Simple Workflow Service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swfCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(swfCmd).Standalone()

	})
	rootCmd.AddCommand(swfCmd)
}
