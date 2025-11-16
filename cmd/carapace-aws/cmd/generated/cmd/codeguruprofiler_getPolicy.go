package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_getPolicyCmd = &cobra.Command{
	Use:   "get-policy",
	Short: "Returns the JSON-formatted resource-based policy on a profiling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_getPolicyCmd).Standalone()

	codeguruprofiler_getPolicyCmd.Flags().String("profiling-group-name", "", "The name of the profiling group.")
	codeguruprofiler_getPolicyCmd.MarkFlagRequired("profiling-group-name")
	codeguruprofilerCmd.AddCommand(codeguruprofiler_getPolicyCmd)
}
