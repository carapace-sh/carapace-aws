package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_describePackagesCmd = &cobra.Command{
	Use:   "describe-packages",
	Short: "Describes all packages available to Amazon ES.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_describePackagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_describePackagesCmd).Standalone()

		es_describePackagesCmd.Flags().String("filters", "", "Only returns packages that match the `DescribePackagesFilterList` values.")
		es_describePackagesCmd.Flags().String("max-results", "", "Limits results to a maximum number of packages.")
		es_describePackagesCmd.Flags().String("next-token", "", "Used for pagination.")
	})
	esCmd.AddCommand(es_describePackagesCmd)
}
