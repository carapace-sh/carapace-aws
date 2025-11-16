package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_deleteOutpostCmd = &cobra.Command{
	Use:   "delete-outpost",
	Short: "Deletes the specified Outpost.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_deleteOutpostCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outposts_deleteOutpostCmd).Standalone()

		outposts_deleteOutpostCmd.Flags().String("outpost-id", "", "The ID or ARN of the Outpost.")
		outposts_deleteOutpostCmd.MarkFlagRequired("outpost-id")
	})
	outpostsCmd.AddCommand(outposts_deleteOutpostCmd)
}
