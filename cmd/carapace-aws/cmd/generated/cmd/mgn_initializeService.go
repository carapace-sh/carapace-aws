package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_initializeServiceCmd = &cobra.Command{
	Use:   "initialize-service",
	Short: "Initialize Application Migration Service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_initializeServiceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_initializeServiceCmd).Standalone()

	})
	mgnCmd.AddCommand(mgn_initializeServiceCmd)
}
