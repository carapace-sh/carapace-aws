package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_startOutpostDecommissionCmd = &cobra.Command{
	Use:   "start-outpost-decommission",
	Short: "Starts the decommission process to return the Outposts racks or servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_startOutpostDecommissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outposts_startOutpostDecommissionCmd).Standalone()

		outposts_startOutpostDecommissionCmd.Flags().String("outpost-identifier", "", "The ID or ARN of the Outpost that you want to decommission.")
		outposts_startOutpostDecommissionCmd.Flags().String("validate-only", "", "Validates the request without starting the decommission process.")
		outposts_startOutpostDecommissionCmd.MarkFlagRequired("outpost-identifier")
	})
	outpostsCmd.AddCommand(outposts_startOutpostDecommissionCmd)
}
