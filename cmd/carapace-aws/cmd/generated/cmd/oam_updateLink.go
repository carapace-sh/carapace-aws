package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oam_updateLinkCmd = &cobra.Command{
	Use:   "update-link",
	Short: "Use this operation to change what types of data are shared from a source account to its linked monitoring account sink.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oam_updateLinkCmd).Standalone()

	oam_updateLinkCmd.Flags().String("identifier", "", "The ARN of the link that you want to update.")
	oam_updateLinkCmd.Flags().String("include-tags", "", "Specifies whether to include the tags associated with the link in the response after the update operation.")
	oam_updateLinkCmd.Flags().String("link-configuration", "", "Use this structure to filter which metric namespaces and which log groups are to be shared from the source account to the monitoring account.")
	oam_updateLinkCmd.Flags().String("resource-types", "", "An array of strings that define which types of data that the source account will send to the monitoring account.")
	oam_updateLinkCmd.MarkFlagRequired("identifier")
	oam_updateLinkCmd.MarkFlagRequired("resource-types")
	oamCmd.AddCommand(oam_updateLinkCmd)
}
