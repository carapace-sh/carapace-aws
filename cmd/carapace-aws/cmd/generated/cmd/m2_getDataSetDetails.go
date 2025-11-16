package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_getDataSetDetailsCmd = &cobra.Command{
	Use:   "get-data-set-details",
	Short: "Gets the details of a specific data set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_getDataSetDetailsCmd).Standalone()

	m2_getDataSetDetailsCmd.Flags().String("application-id", "", "The unique identifier of the application that this data set is associated with.")
	m2_getDataSetDetailsCmd.Flags().String("data-set-name", "", "The name of the data set.")
	m2_getDataSetDetailsCmd.MarkFlagRequired("application-id")
	m2_getDataSetDetailsCmd.MarkFlagRequired("data-set-name")
	m2Cmd.AddCommand(m2_getDataSetDetailsCmd)
}
