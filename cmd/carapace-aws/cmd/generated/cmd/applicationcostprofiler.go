package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationcostprofilerCmd = &cobra.Command{
	Use:   "applicationcostprofiler",
	Short: "This reference provides descriptions of the AWS Application Cost Profiler API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationcostprofilerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationcostprofilerCmd).Standalone()

	})
	rootCmd.AddCommand(applicationcostprofilerCmd)
}
