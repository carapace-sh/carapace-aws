package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_listEnvironmentsCmd = &cobra.Command{
	Use:   "list-environments",
	Short: "Lists the runtime environments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_listEnvironmentsCmd).Standalone()

	m2_listEnvironmentsCmd.Flags().String("engine-type", "", "The engine type for the runtime environment.")
	m2_listEnvironmentsCmd.Flags().String("max-results", "", "The maximum number of runtime environments to return.")
	m2_listEnvironmentsCmd.Flags().String("names", "", "The names of the runtime environments.")
	m2_listEnvironmentsCmd.Flags().String("next-token", "", "A pagination token to control the number of runtime environments displayed in the list.")
	m2Cmd.AddCommand(m2_listEnvironmentsCmd)
}
