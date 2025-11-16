package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch\n\nUsing Batch, you can run batch computing workloads on the Amazon Web Services Cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(batchCmd).Standalone()

	})
	rootCmd.AddCommand(batchCmd)
}
