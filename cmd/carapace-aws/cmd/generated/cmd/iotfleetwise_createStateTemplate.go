package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_createStateTemplateCmd = &cobra.Command{
	Use:   "create-state-template",
	Short: "Creates a state template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_createStateTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_createStateTemplateCmd).Standalone()

		iotfleetwise_createStateTemplateCmd.Flags().String("data-extra-dimensions", "", "A list of vehicle attributes to associate with the payload published on the state template's MQTT topic.")
		iotfleetwise_createStateTemplateCmd.Flags().String("description", "", "A brief description of the state template.")
		iotfleetwise_createStateTemplateCmd.Flags().String("metadata-extra-dimensions", "", "A list of vehicle attributes to associate with user properties of the messages published on the state template's MQTT topic.")
		iotfleetwise_createStateTemplateCmd.Flags().String("name", "", "The name of the state template.")
		iotfleetwise_createStateTemplateCmd.Flags().String("signal-catalog-arn", "", "The ARN of the signal catalog associated with the state template.")
		iotfleetwise_createStateTemplateCmd.Flags().String("state-template-properties", "", "A list of signals from which data is collected.")
		iotfleetwise_createStateTemplateCmd.Flags().String("tags", "", "Metadata that can be used to manage the state template.")
		iotfleetwise_createStateTemplateCmd.MarkFlagRequired("name")
		iotfleetwise_createStateTemplateCmd.MarkFlagRequired("signal-catalog-arn")
		iotfleetwise_createStateTemplateCmd.MarkFlagRequired("state-template-properties")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_createStateTemplateCmd)
}
