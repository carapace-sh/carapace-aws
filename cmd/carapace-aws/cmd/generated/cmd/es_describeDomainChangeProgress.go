package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_describeDomainChangeProgressCmd = &cobra.Command{
	Use:   "describe-domain-change-progress",
	Short: "Returns information about the current blue/green deployment happening on a domain, including a change ID, status, and progress stages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_describeDomainChangeProgressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_describeDomainChangeProgressCmd).Standalone()

		es_describeDomainChangeProgressCmd.Flags().String("change-id", "", "The specific change ID for which you want to get progress information.")
		es_describeDomainChangeProgressCmd.Flags().String("domain-name", "", "The domain you want to get the progress information about.")
		es_describeDomainChangeProgressCmd.MarkFlagRequired("domain-name")
	})
	esCmd.AddCommand(es_describeDomainChangeProgressCmd)
}
