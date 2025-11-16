package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_listSubscribersCmd = &cobra.Command{
	Use:   "list-subscribers",
	Short: "Lists all subscribers for the specific Amazon Security Lake account ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_listSubscribersCmd).Standalone()

	securitylake_listSubscribersCmd.Flags().String("max-results", "", "The maximum number of accounts for which the configuration is displayed.")
	securitylake_listSubscribersCmd.Flags().String("next-token", "", "If nextToken is returned, there are more results available.")
	securitylakeCmd.AddCommand(securitylake_listSubscribersCmd)
}
