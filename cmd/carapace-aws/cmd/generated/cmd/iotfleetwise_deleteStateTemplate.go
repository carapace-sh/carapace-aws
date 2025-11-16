package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_deleteStateTemplateCmd = &cobra.Command{
	Use:   "delete-state-template",
	Short: "Deletes a state template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_deleteStateTemplateCmd).Standalone()

	iotfleetwise_deleteStateTemplateCmd.Flags().String("identifier", "", "The unique ID of the state template.")
	iotfleetwise_deleteStateTemplateCmd.MarkFlagRequired("identifier")
	iotfleetwiseCmd.AddCommand(iotfleetwise_deleteStateTemplateCmd)
}
