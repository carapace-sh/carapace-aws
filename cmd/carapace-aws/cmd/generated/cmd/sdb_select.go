package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sdb_selectCmd = &cobra.Command{
	Use:   "select",
	Short: "The `Select` operation returns a set of attributes for `ItemNames` that match the select expression.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sdb_selectCmd).Standalone()

	sdb_selectCmd.Flags().Bool("consistent-read", false, "Determines whether or not strong consistency should be enforced when data is read from SimpleDB.")
	sdb_selectCmd.Flags().String("next-token", "", "A string informing Amazon SimpleDB where to start the next list of `ItemNames`.")
	sdb_selectCmd.Flags().Bool("no-consistent-read", false, "Determines whether or not strong consistency should be enforced when data is read from SimpleDB.")
	sdb_selectCmd.Flags().String("select-expression", "", "The expression used to query the domain.")
	sdb_selectCmd.Flag("no-consistent-read").Hidden = true
	sdb_selectCmd.MarkFlagRequired("select-expression")
	sdbCmd.AddCommand(sdb_selectCmd)
}
