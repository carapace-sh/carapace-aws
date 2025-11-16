package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_describeRuntimeVersionsCmd = &cobra.Command{
	Use:   "describe-runtime-versions",
	Short: "Returns a list of Synthetics canary runtime versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_describeRuntimeVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(synthetics_describeRuntimeVersionsCmd).Standalone()

		synthetics_describeRuntimeVersionsCmd.Flags().String("max-results", "", "Specify this parameter to limit how many runs are returned each time you use the `DescribeRuntimeVersions` operation.")
		synthetics_describeRuntimeVersionsCmd.Flags().String("next-token", "", "A token that indicates that there is more data available.")
	})
	syntheticsCmd.AddCommand(synthetics_describeRuntimeVersionsCmd)
}
