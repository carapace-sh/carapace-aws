package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getLineageNodeCmd = &cobra.Command{
	Use:   "get-lineage-node",
	Short: "Gets the data lineage node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getLineageNodeCmd).Standalone()

	datazone_getLineageNodeCmd.Flags().String("domain-identifier", "", "The ID of the domain in which you want to get the data lineage node.")
	datazone_getLineageNodeCmd.Flags().String("event-timestamp", "", "The event time stamp for which you want to get the data lineage node.")
	datazone_getLineageNodeCmd.Flags().String("identifier", "", "The ID of the data lineage node that you want to get.")
	datazone_getLineageNodeCmd.MarkFlagRequired("domain-identifier")
	datazone_getLineageNodeCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_getLineageNodeCmd)
}
