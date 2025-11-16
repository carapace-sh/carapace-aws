package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listAppInputSourcesCmd = &cobra.Command{
	Use:   "list-app-input-sources",
	Short: "Lists all the input sources of the Resilience Hub application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listAppInputSourcesCmd).Standalone()

	resiliencehub_listAppInputSourcesCmd.Flags().String("app-arn", "", "Amazon Resource Name (ARN) of the Resilience Hub application.")
	resiliencehub_listAppInputSourcesCmd.Flags().String("app-version", "", "Resilience Hub application version.")
	resiliencehub_listAppInputSourcesCmd.Flags().String("max-results", "", "Maximum number of input sources to be displayed per Resilience Hub application.")
	resiliencehub_listAppInputSourcesCmd.Flags().String("next-token", "", "Null, or the token from a previous call to get the next set of results.")
	resiliencehub_listAppInputSourcesCmd.MarkFlagRequired("app-arn")
	resiliencehub_listAppInputSourcesCmd.MarkFlagRequired("app-version")
	resiliencehubCmd.AddCommand(resiliencehub_listAppInputSourcesCmd)
}
