package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_getStateTemplateCmd = &cobra.Command{
	Use:   "get-state-template",
	Short: "Retrieves information about a state template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_getStateTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_getStateTemplateCmd).Standalone()

		iotfleetwise_getStateTemplateCmd.Flags().String("identifier", "", "The unique ID of the state template.")
		iotfleetwise_getStateTemplateCmd.MarkFlagRequired("identifier")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_getStateTemplateCmd)
}
