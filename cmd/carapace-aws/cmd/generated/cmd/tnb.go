package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnbCmd = &cobra.Command{
	Use:   "tnb",
	Short: "Amazon Web Services Telco Network Builder (TNB) is a network automation service that helps you deploy and manage telecom networks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnbCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnbCmd).Standalone()

	})
	rootCmd.AddCommand(tnbCmd)
}
