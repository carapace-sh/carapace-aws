package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudcontrolCmd = &cobra.Command{
	Use:   "cloudcontrol",
	Short: "For more information about Amazon Web Services Cloud Control API, see the [Amazon Web Services Cloud Control API User Guide](https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/what-is-cloudcontrolapi.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudcontrolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudcontrolCmd).Standalone()

	})
	rootCmd.AddCommand(cloudcontrolCmd)
}
