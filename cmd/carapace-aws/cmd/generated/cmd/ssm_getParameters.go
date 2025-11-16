package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getParametersCmd = &cobra.Command{
	Use:   "get-parameters",
	Short: "Get information about one or more parameters by specifying multiple parameter names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getParametersCmd).Standalone()

	ssm_getParametersCmd.Flags().String("names", "", "The names or Amazon Resource Names (ARNs) of the parameters that you want to query.")
	ssm_getParametersCmd.Flags().Bool("no-with-decryption", false, "Return decrypted secure string value.")
	ssm_getParametersCmd.Flags().Bool("with-decryption", false, "Return decrypted secure string value.")
	ssm_getParametersCmd.MarkFlagRequired("names")
	ssm_getParametersCmd.Flag("no-with-decryption").Hidden = true
	ssmCmd.AddCommand(ssm_getParametersCmd)
}
