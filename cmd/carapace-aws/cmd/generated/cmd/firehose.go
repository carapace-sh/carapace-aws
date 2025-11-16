package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firehoseCmd = &cobra.Command{
	Use:   "firehose",
	Short: "Amazon Data Firehose",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firehoseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(firehoseCmd).Standalone()

	})
	rootCmd.AddCommand(firehoseCmd)
}
