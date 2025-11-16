package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2ModelsCmd = &cobra.Command{
	Use:   "lexv2-models",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2ModelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2ModelsCmd).Standalone()

	})
	rootCmd.AddCommand(lexv2ModelsCmd)
}
