package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_listCapabilitiesCmd = &cobra.Command{
	Use:   "list-capabilities",
	Short: "Lists the capabilities associated with your Amazon Web Services account for your current or specified region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_listCapabilitiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(b2bi_listCapabilitiesCmd).Standalone()

		b2bi_listCapabilitiesCmd.Flags().String("max-results", "", "Specifies the maximum number of capabilities to return.")
		b2bi_listCapabilitiesCmd.Flags().String("next-token", "", "When additional results are obtained from the command, a `NextToken` parameter is returned in the output.")
	})
	b2biCmd.AddCommand(b2bi_listCapabilitiesCmd)
}
