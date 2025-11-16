package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_listCrossAccountAttachmentsCmd = &cobra.Command{
	Use:   "list-cross-account-attachments",
	Short: "List the cross-account attachments that have been created in Global Accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_listCrossAccountAttachmentsCmd).Standalone()

	globalaccelerator_listCrossAccountAttachmentsCmd.Flags().String("max-results", "", "The number of cross-account attachment objects that you want to return with this call.")
	globalaccelerator_listCrossAccountAttachmentsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	globalacceleratorCmd.AddCommand(globalaccelerator_listCrossAccountAttachmentsCmd)
}
