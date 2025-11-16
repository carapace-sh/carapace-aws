package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_createLabelCmd = &cobra.Command{
	Use:   "create-label",
	Short: "Creates a label for an event.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_createLabelCmd).Standalone()

	lookoutequipment_createLabelCmd.Flags().String("client-token", "", "A unique identifier for the request to create a label.")
	lookoutequipment_createLabelCmd.Flags().String("end-time", "", "The end time of the labeled event.")
	lookoutequipment_createLabelCmd.Flags().String("equipment", "", "Indicates that a label pertains to a particular piece of equipment.")
	lookoutequipment_createLabelCmd.Flags().String("fault-code", "", "Provides additional information about the label.")
	lookoutequipment_createLabelCmd.Flags().String("label-group-name", "", "The name of a group of labels.")
	lookoutequipment_createLabelCmd.Flags().String("notes", "", "Metadata providing additional information about the label.")
	lookoutequipment_createLabelCmd.Flags().String("rating", "", "Indicates whether a labeled event represents an anomaly.")
	lookoutequipment_createLabelCmd.Flags().String("start-time", "", "The start time of the labeled event.")
	lookoutequipment_createLabelCmd.MarkFlagRequired("client-token")
	lookoutequipment_createLabelCmd.MarkFlagRequired("end-time")
	lookoutequipment_createLabelCmd.MarkFlagRequired("label-group-name")
	lookoutequipment_createLabelCmd.MarkFlagRequired("rating")
	lookoutequipment_createLabelCmd.MarkFlagRequired("start-time")
	lookoutequipmentCmd.AddCommand(lookoutequipment_createLabelCmd)
}
