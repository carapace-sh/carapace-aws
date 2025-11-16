package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_describeSdiSourceCmd = &cobra.Command{
	Use:   "describe-sdi-source",
	Short: "Gets details about a SdiSource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_describeSdiSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_describeSdiSourceCmd).Standalone()

		medialive_describeSdiSourceCmd.Flags().String("sdi-source-id", "", "Get details about an SdiSource.")
		medialive_describeSdiSourceCmd.MarkFlagRequired("sdi-source-id")
	})
	medialiveCmd.AddCommand(medialive_describeSdiSourceCmd)
}
