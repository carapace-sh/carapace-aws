package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getLineageEventCmd = &cobra.Command{
	Use:   "get-lineage-event",
	Short: "Describes the lineage event.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getLineageEventCmd).Standalone()

	datazone_getLineageEventCmd.Flags().String("domain-identifier", "", "The ID of the domain.")
	datazone_getLineageEventCmd.Flags().String("identifier", "", "The ID of the lineage event.")
	datazone_getLineageEventCmd.MarkFlagRequired("domain-identifier")
	datazone_getLineageEventCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_getLineageEventCmd)
}
