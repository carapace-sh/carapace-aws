package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_updateOutpostCmd = &cobra.Command{
	Use:   "update-outpost",
	Short: "Updates an Outpost.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_updateOutpostCmd).Standalone()

	outposts_updateOutpostCmd.Flags().String("description", "", "")
	outposts_updateOutpostCmd.Flags().String("name", "", "")
	outposts_updateOutpostCmd.Flags().String("outpost-id", "", "The ID or ARN of the Outpost.")
	outposts_updateOutpostCmd.Flags().String("supported-hardware-type", "", "The type of hardware for this Outpost.")
	outposts_updateOutpostCmd.MarkFlagRequired("outpost-id")
	outpostsCmd.AddCommand(outposts_updateOutpostCmd)
}
