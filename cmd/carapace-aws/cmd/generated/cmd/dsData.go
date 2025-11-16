package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsDataCmd = &cobra.Command{
	Use:   "ds-data",
	Short: "Amazon Web Services Directory Service Data is an extension of Directory Service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsDataCmd).Standalone()

	rootCmd.AddCommand(dsDataCmd)
}
