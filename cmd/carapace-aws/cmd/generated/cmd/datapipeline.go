package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datapipelineCmd = &cobra.Command{
	Use:   "datapipeline",
	Short: "AWS Data Pipeline configures and manages a data-driven workflow called a pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datapipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datapipelineCmd).Standalone()

	})
	rootCmd.AddCommand(datapipelineCmd)
}
