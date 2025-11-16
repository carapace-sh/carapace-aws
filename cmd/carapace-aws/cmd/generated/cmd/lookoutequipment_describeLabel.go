package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_describeLabelCmd = &cobra.Command{
	Use:   "describe-label",
	Short: "Returns the name of the label.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_describeLabelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_describeLabelCmd).Standalone()

		lookoutequipment_describeLabelCmd.Flags().String("label-group-name", "", "Returns the name of the group containing the label.")
		lookoutequipment_describeLabelCmd.Flags().String("label-id", "", "Returns the ID of the label.")
		lookoutequipment_describeLabelCmd.MarkFlagRequired("label-group-name")
		lookoutequipment_describeLabelCmd.MarkFlagRequired("label-id")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_describeLabelCmd)
}
