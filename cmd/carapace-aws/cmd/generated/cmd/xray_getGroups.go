package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_getGroupsCmd = &cobra.Command{
	Use:   "get-groups",
	Short: "Retrieves all active group details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_getGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(xray_getGroupsCmd).Standalone()

		xray_getGroupsCmd.Flags().String("next-token", "", "Pagination token.")
	})
	xrayCmd.AddCommand(xray_getGroupsCmd)
}
