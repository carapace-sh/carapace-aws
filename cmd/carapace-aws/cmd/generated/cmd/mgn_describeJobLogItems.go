package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_describeJobLogItemsCmd = &cobra.Command{
	Use:   "describe-job-log-items",
	Short: "Retrieves detailed job log items with paging.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_describeJobLogItemsCmd).Standalone()

	mgn_describeJobLogItemsCmd.Flags().String("account-id", "", "Request to describe Job log Account ID.")
	mgn_describeJobLogItemsCmd.Flags().String("job-id", "", "Request to describe Job log job ID.")
	mgn_describeJobLogItemsCmd.Flags().String("max-results", "", "Request to describe Job log item maximum results.")
	mgn_describeJobLogItemsCmd.Flags().String("next-token", "", "Request to describe Job log next token.")
	mgn_describeJobLogItemsCmd.MarkFlagRequired("job-id")
	mgnCmd.AddCommand(mgn_describeJobLogItemsCmd)
}
