package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_listLabelGroupsCmd = &cobra.Command{
	Use:   "list-label-groups",
	Short: "Returns a list of the label groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_listLabelGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_listLabelGroupsCmd).Standalone()

		lookoutequipment_listLabelGroupsCmd.Flags().String("label-group-name-begins-with", "", "The beginning of the name of the label groups to be listed.")
		lookoutequipment_listLabelGroupsCmd.Flags().String("max-results", "", "Specifies the maximum number of label groups to list.")
		lookoutequipment_listLabelGroupsCmd.Flags().String("next-token", "", "An opaque pagination token indicating where to continue the listing of label groups.")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_listLabelGroupsCmd)
}
