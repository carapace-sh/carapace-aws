package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformationCmd = &cobra.Command{
	Use:   "lakeformation",
	Short: "Lake Formation\n\nDefines the public endpoint for the Lake Formation service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformationCmd).Standalone()

	rootCmd.AddCommand(lakeformationCmd)
}
