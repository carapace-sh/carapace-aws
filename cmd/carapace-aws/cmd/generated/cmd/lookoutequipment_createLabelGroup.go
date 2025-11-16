package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_createLabelGroupCmd = &cobra.Command{
	Use:   "create-label-group",
	Short: "Creates a group of labels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_createLabelGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_createLabelGroupCmd).Standalone()

		lookoutequipment_createLabelGroupCmd.Flags().String("client-token", "", "A unique identifier for the request to create a label group.")
		lookoutequipment_createLabelGroupCmd.Flags().String("fault-codes", "", "The acceptable fault codes (indicating the type of anomaly associated with the label) that can be used with this label group.")
		lookoutequipment_createLabelGroupCmd.Flags().String("label-group-name", "", "Names a group of labels.")
		lookoutequipment_createLabelGroupCmd.Flags().String("tags", "", "Tags that provide metadata about the label group you are creating.")
		lookoutequipment_createLabelGroupCmd.MarkFlagRequired("client-token")
		lookoutequipment_createLabelGroupCmd.MarkFlagRequired("label-group-name")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_createLabelGroupCmd)
}
