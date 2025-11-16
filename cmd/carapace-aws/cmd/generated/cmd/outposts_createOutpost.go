package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_createOutpostCmd = &cobra.Command{
	Use:   "create-outpost",
	Short: "Creates an Outpost.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_createOutpostCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outposts_createOutpostCmd).Standalone()

		outposts_createOutpostCmd.Flags().String("availability-zone", "", "")
		outposts_createOutpostCmd.Flags().String("availability-zone-id", "", "")
		outposts_createOutpostCmd.Flags().String("description", "", "")
		outposts_createOutpostCmd.Flags().String("name", "", "")
		outposts_createOutpostCmd.Flags().String("site-id", "", "The ID or the Amazon Resource Name (ARN) of the site.")
		outposts_createOutpostCmd.Flags().String("supported-hardware-type", "", "The type of hardware for this Outpost.")
		outposts_createOutpostCmd.Flags().String("tags", "", "The tags to apply to the Outpost.")
		outposts_createOutpostCmd.MarkFlagRequired("name")
		outposts_createOutpostCmd.MarkFlagRequired("site-id")
	})
	outpostsCmd.AddCommand(outposts_createOutpostCmd)
}
