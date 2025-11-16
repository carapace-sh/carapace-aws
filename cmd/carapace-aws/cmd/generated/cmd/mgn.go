package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgnCmd = &cobra.Command{
	Use:   "mgn",
	Short: "The Application Migration Service service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgnCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgnCmd).Standalone()

	})
	rootCmd.AddCommand(mgnCmd)
}
