package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearningCmd = &cobra.Command{
	Use:   "machinelearning",
	Short: "Definition of the public APIs exposed by Amazon Machine Learning",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearningCmd).Standalone()

	rootCmd.AddCommand(machinelearningCmd)
}
