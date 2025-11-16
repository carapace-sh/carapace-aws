package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_createGroupCmd = &cobra.Command{
	Use:   "create-group",
	Short: "Creates a group resource with a name and a filter expression.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_createGroupCmd).Standalone()

	xray_createGroupCmd.Flags().String("filter-expression", "", "The filter expression defining criteria by which to group traces.")
	xray_createGroupCmd.Flags().String("group-name", "", "The case-sensitive name of the new group.")
	xray_createGroupCmd.Flags().String("insights-configuration", "", "The structure containing configurations related to insights.")
	xray_createGroupCmd.Flags().String("tags", "", "A map that contains one or more tag keys and tag values to attach to an X-Ray group.")
	xray_createGroupCmd.MarkFlagRequired("group-name")
	xrayCmd.AddCommand(xray_createGroupCmd)
}
