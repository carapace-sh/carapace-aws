package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53profilesCmd = &cobra.Command{
	Use:   "route53profiles",
	Short: "With Amazon Route 53 Profiles you can share Route 53 configurations with VPCs and AWS accounts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53profilesCmd).Standalone()

	rootCmd.AddCommand(route53profilesCmd)
}
