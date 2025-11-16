package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmCmd = &cobra.Command{
	Use:   "cloudhsm",
	Short: "AWS CloudHSM Service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudhsmCmd).Standalone()

	})
	rootCmd.AddCommand(cloudhsmCmd)
}
