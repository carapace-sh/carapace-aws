package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_deleteLabelCmd = &cobra.Command{
	Use:   "delete-label",
	Short: "Deletes a label.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_deleteLabelCmd).Standalone()

	lookoutequipment_deleteLabelCmd.Flags().String("label-group-name", "", "The name of the label group that contains the label that you want to delete.")
	lookoutequipment_deleteLabelCmd.Flags().String("label-id", "", "The ID of the label that you want to delete.")
	lookoutequipment_deleteLabelCmd.MarkFlagRequired("label-group-name")
	lookoutequipment_deleteLabelCmd.MarkFlagRequired("label-id")
	lookoutequipmentCmd.AddCommand(lookoutequipment_deleteLabelCmd)
}
