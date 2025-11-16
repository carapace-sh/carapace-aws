package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspaceCmd = &cobra.Command{
	Use:   "repostspace",
	Short: "AWS re:Post Private is a private version of AWS re:Post for enterprises with Enterprise Support or Enterprise On-Ramp Support plans.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(repostspaceCmd).Standalone()

	})
	rootCmd.AddCommand(repostspaceCmd)
}
