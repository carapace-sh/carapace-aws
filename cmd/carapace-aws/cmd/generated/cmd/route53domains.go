package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domainsCmd = &cobra.Command{
	Use:   "route53domains",
	Short: "Amazon Route 53 API actions let you register domain names and perform related operations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domainsCmd).Standalone()

	rootCmd.AddCommand(route53domainsCmd)
}
