package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_getGroupCmd = &cobra.Command{
	Use:   "get-group",
	Short: "Returns information about one group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_getGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(synthetics_getGroupCmd).Standalone()

		synthetics_getGroupCmd.Flags().String("group-identifier", "", "Specifies the group to return information for.")
		synthetics_getGroupCmd.MarkFlagRequired("group-identifier")
	})
	syntheticsCmd.AddCommand(synthetics_getGroupCmd)
}
