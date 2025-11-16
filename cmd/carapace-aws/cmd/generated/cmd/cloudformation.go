package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformationCmd = &cobra.Command{
	Use:   "cloudformation",
	Short: "CloudFormation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformationCmd).Standalone()

	})
	rootCmd.AddCommand(cloudformationCmd)
}
