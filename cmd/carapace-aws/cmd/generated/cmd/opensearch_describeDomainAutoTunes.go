package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_describeDomainAutoTunesCmd = &cobra.Command{
	Use:   "describe-domain-auto-tunes",
	Short: "Returns the list of optimizations that Auto-Tune has made to an Amazon OpenSearch Service domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_describeDomainAutoTunesCmd).Standalone()

	opensearch_describeDomainAutoTunesCmd.Flags().String("domain-name", "", "Name of the domain that you want Auto-Tune details about.")
	opensearch_describeDomainAutoTunesCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
	opensearch_describeDomainAutoTunesCmd.Flags().String("next-token", "", "If your initial `DescribeDomainAutoTunes` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `DescribeDomainAutoTunes` operations, which returns results in the next page.")
	opensearch_describeDomainAutoTunesCmd.MarkFlagRequired("domain-name")
	opensearchCmd.AddCommand(opensearch_describeDomainAutoTunesCmd)
}
