package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_getOutpostCmd = &cobra.Command{
	Use:   "get-outpost",
	Short: "Gets information about the specified Outpost.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_getOutpostCmd).Standalone()

	outposts_getOutpostCmd.Flags().String("outpost-id", "", "The ID or ARN of the Outpost.")
	outposts_getOutpostCmd.MarkFlagRequired("outpost-id")
	outpostsCmd.AddCommand(outposts_getOutpostCmd)
}
