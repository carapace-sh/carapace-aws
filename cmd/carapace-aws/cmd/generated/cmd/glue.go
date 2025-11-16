package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glueCmd = &cobra.Command{
	Use:   "glue",
	Short: "Glue\n\nDefines the public endpoint for the Glue service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glueCmd).Standalone()

	})
	rootCmd.AddCommand(glueCmd)
}
