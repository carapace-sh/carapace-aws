package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_listTargetAccountConfigurationsCmd = &cobra.Command{
	Use:   "list-target-account-configurations",
	Short: "Lists the target account configurations of the specified experiment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_listTargetAccountConfigurationsCmd).Standalone()

	fis_listTargetAccountConfigurationsCmd.Flags().String("experiment-template-id", "", "The ID of the experiment template.")
	fis_listTargetAccountConfigurationsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	fis_listTargetAccountConfigurationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	fis_listTargetAccountConfigurationsCmd.MarkFlagRequired("experiment-template-id")
	fisCmd.AddCommand(fis_listTargetAccountConfigurationsCmd)
}
