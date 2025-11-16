package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_getLayoutCmd = &cobra.Command{
	Use:   "get-layout",
	Short: "Returns the details for the requested layout.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_getLayoutCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_getLayoutCmd).Standalone()

		connectcases_getLayoutCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
		connectcases_getLayoutCmd.Flags().String("layout-id", "", "The unique identifier of the layout.")
		connectcases_getLayoutCmd.MarkFlagRequired("domain-id")
		connectcases_getLayoutCmd.MarkFlagRequired("layout-id")
	})
	connectcasesCmd.AddCommand(connectcases_getLayoutCmd)
}
