package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oam_getSinkCmd = &cobra.Command{
	Use:   "get-sink",
	Short: "Returns complete information about one monitoring account sink.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oam_getSinkCmd).Standalone()

	oam_getSinkCmd.Flags().String("identifier", "", "The ARN of the sink to retrieve information for.")
	oam_getSinkCmd.Flags().String("include-tags", "", "Specifies whether to include the tags associated with the sink in the response.")
	oam_getSinkCmd.MarkFlagRequired("identifier")
	oamCmd.AddCommand(oam_getSinkCmd)
}
