package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_getDomainCmd = &cobra.Command{
	Use:   "get-domain",
	Short: "Returns information about a specific domain if it exists.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_getDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_getDomainCmd).Standalone()

		connectcases_getDomainCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
		connectcases_getDomainCmd.MarkFlagRequired("domain-id")
	})
	connectcasesCmd.AddCommand(connectcases_getDomainCmd)
}
