package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getGroupCmd = &cobra.Command{
	Use:   "get-group",
	Short: "Retrieves information about a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_getGroupCmd).Standalone()

		greengrass_getGroupCmd.Flags().String("group-id", "", "The ID of the Greengrass group.")
		greengrass_getGroupCmd.MarkFlagRequired("group-id")
	})
	greengrassCmd.AddCommand(greengrass_getGroupCmd)
}
