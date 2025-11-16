package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltowerCmd = &cobra.Command{
	Use:   "controltower",
	Short: "Amazon Web Services Control Tower offers application programming interface (API) operations that support programmatic interaction with these types of resources:",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltowerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controltowerCmd).Standalone()

	})
	rootCmd.AddCommand(controltowerCmd)
}
