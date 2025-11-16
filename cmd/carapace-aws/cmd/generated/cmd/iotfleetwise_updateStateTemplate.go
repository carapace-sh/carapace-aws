package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_updateStateTemplateCmd = &cobra.Command{
	Use:   "update-state-template",
	Short: "Updates a state template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_updateStateTemplateCmd).Standalone()

	iotfleetwise_updateStateTemplateCmd.Flags().String("data-extra-dimensions", "", "A list of vehicle attributes to associate with the payload published on the state template's MQTT topic.")
	iotfleetwise_updateStateTemplateCmd.Flags().String("description", "", "A brief description of the state template.")
	iotfleetwise_updateStateTemplateCmd.Flags().String("identifier", "", "The unique ID of the state template.")
	iotfleetwise_updateStateTemplateCmd.Flags().String("metadata-extra-dimensions", "", "A list of vehicle attributes to associate with user properties of the messages published on the state template's MQTT topic.")
	iotfleetwise_updateStateTemplateCmd.Flags().String("state-template-properties-to-add", "", "Add signals from which data is collected as part of the state template.")
	iotfleetwise_updateStateTemplateCmd.Flags().String("state-template-properties-to-remove", "", "Remove signals from which data is collected as part of the state template.")
	iotfleetwise_updateStateTemplateCmd.MarkFlagRequired("identifier")
	iotfleetwiseCmd.AddCommand(iotfleetwise_updateStateTemplateCmd)
}
