package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_getSignedBluinsightsUrlCmd = &cobra.Command{
	Use:   "get-signed-bluinsights-url",
	Short: "Gets a single sign-on URL that can be used to connect to AWS Blu Insights.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_getSignedBluinsightsUrlCmd).Standalone()

	m2Cmd.AddCommand(m2_getSignedBluinsightsUrlCmd)
}
