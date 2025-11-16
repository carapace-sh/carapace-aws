package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listVersionsCmd = &cobra.Command{
	Use:   "list-versions",
	Short: "Retrieves an array of all the encoder engine versions that are available in this AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_listVersionsCmd).Standalone()

	})
	medialiveCmd.AddCommand(medialive_listVersionsCmd)
}
