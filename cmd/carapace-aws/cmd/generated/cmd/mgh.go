package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mghCmd = &cobra.Command{
	Use:   "mgh",
	Short: "The AWS Migration Hub API methods help to obtain server and application migration status and integrate your resource-specific migration tool by providing a programmatic interface to Migration Hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mghCmd).Standalone()

	rootCmd.AddCommand(mghCmd)
}
