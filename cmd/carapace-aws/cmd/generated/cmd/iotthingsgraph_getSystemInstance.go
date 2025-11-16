package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_getSystemInstanceCmd = &cobra.Command{
	Use:   "get-system-instance",
	Short: "Gets a system instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_getSystemInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotthingsgraph_getSystemInstanceCmd).Standalone()

		iotthingsgraph_getSystemInstanceCmd.Flags().String("id", "", "The ID of the system deployment instance.")
		iotthingsgraph_getSystemInstanceCmd.MarkFlagRequired("id")
	})
	iotthingsgraphCmd.AddCommand(iotthingsgraph_getSystemInstanceCmd)
}
