package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrCmd = &cobra.Command{
	Use:   "ecr",
	Short: "Amazon Elastic Container Registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecrCmd).Standalone()

	})
	rootCmd.AddCommand(ecrCmd)
}
