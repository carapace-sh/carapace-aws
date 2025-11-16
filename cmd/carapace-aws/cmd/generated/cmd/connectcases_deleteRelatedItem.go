package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_deleteRelatedItemCmd = &cobra.Command{
	Use:   "delete-related-item",
	Short: "Deletes the related item resource under a case.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_deleteRelatedItemCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_deleteRelatedItemCmd).Standalone()

		connectcases_deleteRelatedItemCmd.Flags().String("case-id", "", "A unique identifier of the case.")
		connectcases_deleteRelatedItemCmd.Flags().String("domain-id", "", "A unique identifier of the Cases domain.")
		connectcases_deleteRelatedItemCmd.Flags().String("related-item-id", "", "A unique identifier of a related item.")
		connectcases_deleteRelatedItemCmd.MarkFlagRequired("case-id")
		connectcases_deleteRelatedItemCmd.MarkFlagRequired("domain-id")
		connectcases_deleteRelatedItemCmd.MarkFlagRequired("related-item-id")
	})
	connectcasesCmd.AddCommand(connectcases_deleteRelatedItemCmd)
}
