package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilderCmd = &cobra.Command{
	Use:   "amplifyuibuilder",
	Short: "The Amplify UI Builder API provides a programmatic interface for creating and configuring user interface (UI) component libraries and themes for use in your Amplify applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilderCmd).Standalone()

	rootCmd.AddCommand(amplifyuibuilderCmd)
}
