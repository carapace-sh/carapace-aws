package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getParameterCmd = &cobra.Command{
	Use:   "get-parameter",
	Short: "Get information about a single parameter by specifying the parameter name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getParameterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_getParameterCmd).Standalone()

		ssm_getParameterCmd.Flags().String("name", "", "The name or Amazon Resource Name (ARN) of the parameter that you want to query.")
		ssm_getParameterCmd.Flags().Bool("no-with-decryption", false, "Return decrypted values for secure string parameters.")
		ssm_getParameterCmd.Flags().Bool("with-decryption", false, "Return decrypted values for secure string parameters.")
		ssm_getParameterCmd.MarkFlagRequired("name")
		ssm_getParameterCmd.Flag("no-with-decryption").Hidden = true
	})
	ssmCmd.AddCommand(ssm_getParameterCmd)
}
