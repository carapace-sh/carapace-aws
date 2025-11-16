package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_describeSourceServersCmd = &cobra.Command{
	Use:   "describe-source-servers",
	Short: "Lists all Source Servers or multiple Source Servers filtered by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_describeSourceServersCmd).Standalone()

	drs_describeSourceServersCmd.Flags().String("filters", "", "A set of filters by which to return Source Servers.")
	drs_describeSourceServersCmd.Flags().String("max-results", "", "Maximum number of Source Servers to retrieve.")
	drs_describeSourceServersCmd.Flags().String("next-token", "", "The token of the next Source Server to retrieve.")
	drsCmd.AddCommand(drs_describeSourceServersCmd)
}
