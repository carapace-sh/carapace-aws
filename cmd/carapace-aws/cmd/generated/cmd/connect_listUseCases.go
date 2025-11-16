package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listUseCasesCmd = &cobra.Command{
	Use:   "list-use-cases",
	Short: "Lists the use cases for the integration association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listUseCasesCmd).Standalone()

	connect_listUseCasesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_listUseCasesCmd.Flags().String("integration-association-id", "", "The identifier for the integration association.")
	connect_listUseCasesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_listUseCasesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_listUseCasesCmd.MarkFlagRequired("instance-id")
	connect_listUseCasesCmd.MarkFlagRequired("integration-association-id")
	connectCmd.AddCommand(connect_listUseCasesCmd)
}
