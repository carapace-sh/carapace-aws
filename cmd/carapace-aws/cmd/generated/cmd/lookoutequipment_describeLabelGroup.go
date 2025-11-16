package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_describeLabelGroupCmd = &cobra.Command{
	Use:   "describe-label-group",
	Short: "Returns information about the label group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_describeLabelGroupCmd).Standalone()

	lookoutequipment_describeLabelGroupCmd.Flags().String("label-group-name", "", "Returns the name of the label group.")
	lookoutequipment_describeLabelGroupCmd.MarkFlagRequired("label-group-name")
	lookoutequipmentCmd.AddCommand(lookoutequipment_describeLabelGroupCmd)
}
