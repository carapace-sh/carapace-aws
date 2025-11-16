package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oam_createLinkCmd = &cobra.Command{
	Use:   "create-link",
	Short: "Creates a link between a source account and a sink that you have created in a monitoring account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oam_createLinkCmd).Standalone()

	oam_createLinkCmd.Flags().String("label-template", "", "Specify a friendly human-readable name to use to identify this source account when you are viewing data from it in the monitoring account.")
	oam_createLinkCmd.Flags().String("link-configuration", "", "Use this structure to optionally create filters that specify that only some metric namespaces or log groups are to be shared from the source account to the monitoring account.")
	oam_createLinkCmd.Flags().String("resource-types", "", "An array of strings that define which types of data that the source account shares with the monitoring account.")
	oam_createLinkCmd.Flags().String("sink-identifier", "", "The ARN of the sink to use to create this link.")
	oam_createLinkCmd.Flags().String("tags", "", "Assigns one or more tags (key-value pairs) to the link.")
	oam_createLinkCmd.MarkFlagRequired("label-template")
	oam_createLinkCmd.MarkFlagRequired("resource-types")
	oam_createLinkCmd.MarkFlagRequired("sink-identifier")
	oamCmd.AddCommand(oam_createLinkCmd)
}
