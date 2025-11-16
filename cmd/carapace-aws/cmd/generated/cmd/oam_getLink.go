package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oam_getLinkCmd = &cobra.Command{
	Use:   "get-link",
	Short: "Returns complete information about one link.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oam_getLinkCmd).Standalone()

	oam_getLinkCmd.Flags().String("identifier", "", "The ARN of the link to retrieve information for.")
	oam_getLinkCmd.Flags().String("include-tags", "", "Specifies whether to include the tags associated with the link in the response.")
	oam_getLinkCmd.MarkFlagRequired("identifier")
	oamCmd.AddCommand(oam_getLinkCmd)
}
