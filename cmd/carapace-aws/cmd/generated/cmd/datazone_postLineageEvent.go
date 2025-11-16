package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_postLineageEventCmd = &cobra.Command{
	Use:   "post-lineage-event",
	Short: "Posts a data lineage event.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_postLineageEventCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_postLineageEventCmd).Standalone()

		datazone_postLineageEventCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_postLineageEventCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to post a data lineage event.")
		datazone_postLineageEventCmd.Flags().String("event", "", "The data lineage event that you want to post.")
		datazone_postLineageEventCmd.MarkFlagRequired("domain-identifier")
		datazone_postLineageEventCmd.MarkFlagRequired("event")
	})
	datazoneCmd.AddCommand(datazone_postLineageEventCmd)
}
