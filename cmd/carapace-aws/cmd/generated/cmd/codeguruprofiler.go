package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofilerCmd = &cobra.Command{
	Use:   "codeguruprofiler",
	Short: "This section provides documentation for the Amazon CodeGuru Profiler API operations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofilerCmd).Standalone()

	rootCmd.AddCommand(codeguruprofilerCmd)
}
