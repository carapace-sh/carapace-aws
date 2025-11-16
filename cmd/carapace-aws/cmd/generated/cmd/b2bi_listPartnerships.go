package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_listPartnershipsCmd = &cobra.Command{
	Use:   "list-partnerships",
	Short: "Lists the partnerships associated with your Amazon Web Services account for your current or specified region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_listPartnershipsCmd).Standalone()

	b2bi_listPartnershipsCmd.Flags().String("max-results", "", "Specifies the maximum number of capabilities to return.")
	b2bi_listPartnershipsCmd.Flags().String("next-token", "", "When additional results are obtained from the command, a `NextToken` parameter is returned in the output.")
	b2bi_listPartnershipsCmd.Flags().String("profile-id", "", "Specifies the unique, system-generated identifier for the profile connected to this partnership.")
	b2biCmd.AddCommand(b2bi_listPartnershipsCmd)
}
