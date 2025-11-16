package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_deleteLayoutCmd = &cobra.Command{
	Use:   "delete-layout",
	Short: "Deletes a layout from a cases template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_deleteLayoutCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_deleteLayoutCmd).Standalone()

		connectcases_deleteLayoutCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
		connectcases_deleteLayoutCmd.Flags().String("layout-id", "", "The unique identifier of the layout.")
		connectcases_deleteLayoutCmd.MarkFlagRequired("domain-id")
		connectcases_deleteLayoutCmd.MarkFlagRequired("layout-id")
	})
	connectcasesCmd.AddCommand(connectcases_deleteLayoutCmd)
}
