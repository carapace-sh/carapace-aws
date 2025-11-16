package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_createLayoutCmd = &cobra.Command{
	Use:   "create-layout",
	Short: "Creates a layout in the Cases domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_createLayoutCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_createLayoutCmd).Standalone()

		connectcases_createLayoutCmd.Flags().String("content", "", "Information about which fields will be present in the layout, and information about the order of the fields.")
		connectcases_createLayoutCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
		connectcases_createLayoutCmd.Flags().String("name", "", "The name of the layout.")
		connectcases_createLayoutCmd.MarkFlagRequired("content")
		connectcases_createLayoutCmd.MarkFlagRequired("domain-id")
		connectcases_createLayoutCmd.MarkFlagRequired("name")
	})
	connectcasesCmd.AddCommand(connectcases_createLayoutCmd)
}
