package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmv2Cmd = &cobra.Command{
	Use:   "cloudhsmv2",
	Short: "For more information about CloudHSM, see [CloudHSM](http://aws.amazon.com/cloudhsm/) and the [CloudHSM User Guide](https://docs.aws.amazon.com/cloudhsm/latest/userguide/).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmv2Cmd).Standalone()

	rootCmd.AddCommand(cloudhsmv2Cmd)
}
