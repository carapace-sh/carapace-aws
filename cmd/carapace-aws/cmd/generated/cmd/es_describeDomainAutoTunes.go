package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_describeDomainAutoTunesCmd = &cobra.Command{
	Use:   "describe-domain-auto-tunes",
	Short: "Provides scheduled Auto-Tune action details for the Elasticsearch domain, such as Auto-Tune action type, description, severity, and scheduled date.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_describeDomainAutoTunesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_describeDomainAutoTunesCmd).Standalone()

		es_describeDomainAutoTunesCmd.Flags().String("domain-name", "", "Specifies the domain name for which you want Auto-Tune action details.")
		es_describeDomainAutoTunesCmd.Flags().String("max-results", "", "Set this value to limit the number of results returned.")
		es_describeDomainAutoTunesCmd.Flags().String("next-token", "", "NextToken is sent in case the earlier API call results contain the NextToken.")
		es_describeDomainAutoTunesCmd.MarkFlagRequired("domain-name")
	})
	esCmd.AddCommand(es_describeDomainAutoTunesCmd)
}
