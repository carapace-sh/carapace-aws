package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_getOutpostInstanceTypesCmd = &cobra.Command{
	Use:   "get-outpost-instance-types",
	Short: "Gets the instance types for the specified Outpost.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_getOutpostInstanceTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outposts_getOutpostInstanceTypesCmd).Standalone()

		outposts_getOutpostInstanceTypesCmd.Flags().String("max-results", "", "")
		outposts_getOutpostInstanceTypesCmd.Flags().String("next-token", "", "")
		outposts_getOutpostInstanceTypesCmd.Flags().String("outpost-id", "", "The ID or ARN of the Outpost.")
		outposts_getOutpostInstanceTypesCmd.MarkFlagRequired("outpost-id")
	})
	outpostsCmd.AddCommand(outposts_getOutpostInstanceTypesCmd)
}
