package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_listProfilesCmd = &cobra.Command{
	Use:   "list-profiles",
	Short: "Lists the profiles associated with your Amazon Web Services account for your current or specified region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_listProfilesCmd).Standalone()

	b2bi_listProfilesCmd.Flags().String("max-results", "", "Specifies the maximum number of profiles to return.")
	b2bi_listProfilesCmd.Flags().String("next-token", "", "When additional results are obtained from the command, a `NextToken` parameter is returned in the output.")
	b2biCmd.AddCommand(b2bi_listProfilesCmd)
}
