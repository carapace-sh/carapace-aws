package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_updateLayoutCmd = &cobra.Command{
	Use:   "update-layout",
	Short: "Updates the attributes of an existing layout.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_updateLayoutCmd).Standalone()

	connectcases_updateLayoutCmd.Flags().String("content", "", "Information about which fields will be present in the layout, the order of the fields.")
	connectcases_updateLayoutCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
	connectcases_updateLayoutCmd.Flags().String("layout-id", "", "The unique identifier of the layout.")
	connectcases_updateLayoutCmd.Flags().String("name", "", "The name of the layout.")
	connectcases_updateLayoutCmd.MarkFlagRequired("domain-id")
	connectcases_updateLayoutCmd.MarkFlagRequired("layout-id")
	connectcasesCmd.AddCommand(connectcases_updateLayoutCmd)
}
