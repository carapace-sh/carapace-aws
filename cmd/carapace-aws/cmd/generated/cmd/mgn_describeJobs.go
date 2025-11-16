package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_describeJobsCmd = &cobra.Command{
	Use:   "describe-jobs",
	Short: "Returns a list of Jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_describeJobsCmd).Standalone()

	mgn_describeJobsCmd.Flags().String("account-id", "", "Request to describe job log items by Account ID.")
	mgn_describeJobsCmd.Flags().String("filters", "", "Request to describe Job log filters.")
	mgn_describeJobsCmd.Flags().String("max-results", "", "Request to describe job log items by max results.")
	mgn_describeJobsCmd.Flags().String("next-token", "", "Request to describe job log items by next token.")
	mgnCmd.AddCommand(mgn_describeJobsCmd)
}
