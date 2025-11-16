package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_deleteLabelGroupCmd = &cobra.Command{
	Use:   "delete-label-group",
	Short: "Deletes a group of labels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_deleteLabelGroupCmd).Standalone()

	lookoutequipment_deleteLabelGroupCmd.Flags().String("label-group-name", "", "The name of the label group that you want to delete.")
	lookoutequipment_deleteLabelGroupCmd.MarkFlagRequired("label-group-name")
	lookoutequipmentCmd.AddCommand(lookoutequipment_deleteLabelGroupCmd)
}
