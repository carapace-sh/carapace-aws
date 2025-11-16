package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_listEngineVersionsCmd = &cobra.Command{
	Use:   "list-engine-versions",
	Short: "Lists the available engine versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_listEngineVersionsCmd).Standalone()

	m2_listEngineVersionsCmd.Flags().String("engine-type", "", "The type of target platform.")
	m2_listEngineVersionsCmd.Flags().String("max-results", "", "The maximum number of objects to return.")
	m2_listEngineVersionsCmd.Flags().String("next-token", "", "A pagination token returned from a previous call to this operation.")
	m2Cmd.AddCommand(m2_listEngineVersionsCmd)
}
