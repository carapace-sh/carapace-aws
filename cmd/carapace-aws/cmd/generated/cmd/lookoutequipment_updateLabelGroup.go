package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_updateLabelGroupCmd = &cobra.Command{
	Use:   "update-label-group",
	Short: "Updates the label group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_updateLabelGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_updateLabelGroupCmd).Standalone()

		lookoutequipment_updateLabelGroupCmd.Flags().String("fault-codes", "", "Updates the code indicating the type of anomaly associated with the label.")
		lookoutequipment_updateLabelGroupCmd.Flags().String("label-group-name", "", "The name of the label group to be updated.")
		lookoutequipment_updateLabelGroupCmd.MarkFlagRequired("label-group-name")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_updateLabelGroupCmd)
}
