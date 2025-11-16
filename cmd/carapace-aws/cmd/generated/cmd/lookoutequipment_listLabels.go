package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_listLabelsCmd = &cobra.Command{
	Use:   "list-labels",
	Short: "Provides a list of labels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_listLabelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_listLabelsCmd).Standalone()

		lookoutequipment_listLabelsCmd.Flags().String("equipment", "", "Lists the labels that pertain to a particular piece of equipment.")
		lookoutequipment_listLabelsCmd.Flags().String("fault-code", "", "Returns labels with a particular fault code.")
		lookoutequipment_listLabelsCmd.Flags().String("interval-end-time", "", "Returns all labels with a start time earlier than the end time given.")
		lookoutequipment_listLabelsCmd.Flags().String("interval-start-time", "", "Returns all the labels with a end time equal to or later than the start time given.")
		lookoutequipment_listLabelsCmd.Flags().String("label-group-name", "", "Returns the name of the label group.")
		lookoutequipment_listLabelsCmd.Flags().String("max-results", "", "Specifies the maximum number of labels to list.")
		lookoutequipment_listLabelsCmd.Flags().String("next-token", "", "An opaque pagination token indicating where to continue the listing of label groups.")
		lookoutequipment_listLabelsCmd.MarkFlagRequired("label-group-name")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_listLabelsCmd)
}
