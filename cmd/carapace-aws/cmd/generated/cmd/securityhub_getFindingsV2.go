package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getFindingsV2Cmd = &cobra.Command{
	Use:   "get-findings-v2",
	Short: "Return a list of findings that match the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getFindingsV2Cmd).Standalone()

	securityhub_getFindingsV2Cmd.Flags().String("filters", "", "The finding attributes used to define a condition to filter the returned OCSF findings.")
	securityhub_getFindingsV2Cmd.Flags().String("max-results", "", "The maximum number of results to return.")
	securityhub_getFindingsV2Cmd.Flags().String("next-token", "", "The token required for pagination.")
	securityhub_getFindingsV2Cmd.Flags().String("sort-criteria", "", "The finding attributes used to sort the list of returned findings.")
	securityhubCmd.AddCommand(securityhub_getFindingsV2Cmd)
}
