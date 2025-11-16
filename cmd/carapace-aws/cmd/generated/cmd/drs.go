package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drsCmd = &cobra.Command{
	Use:   "drs",
	Short: "AWS Elastic Disaster Recovery Service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drsCmd).Standalone()

	rootCmd.AddCommand(drsCmd)
}
