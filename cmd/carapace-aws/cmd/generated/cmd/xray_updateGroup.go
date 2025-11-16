package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_updateGroupCmd = &cobra.Command{
	Use:   "update-group",
	Short: "Updates a group resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_updateGroupCmd).Standalone()

	xray_updateGroupCmd.Flags().String("filter-expression", "", "The updated filter expression defining criteria by which to group traces.")
	xray_updateGroupCmd.Flags().String("group-arn", "", "The ARN that was generated upon creation.")
	xray_updateGroupCmd.Flags().String("group-name", "", "The case-sensitive name of the group.")
	xray_updateGroupCmd.Flags().String("insights-configuration", "", "The structure containing configurations related to insights.")
	xrayCmd.AddCommand(xray_updateGroupCmd)
}
