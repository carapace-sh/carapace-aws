package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getParametersByPathCmd = &cobra.Command{
	Use:   "get-parameters-by-path",
	Short: "Retrieve information about one or more parameters under a specified level in a hierarchy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getParametersByPathCmd).Standalone()

	ssm_getParametersByPathCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_getParametersByPathCmd.Flags().String("next-token", "", "A token to start the list.")
	ssm_getParametersByPathCmd.Flags().Bool("no-recursive", false, "Retrieve all parameters within a hierarchy.")
	ssm_getParametersByPathCmd.Flags().Bool("no-with-decryption", false, "Retrieve all parameters in a hierarchy with their value decrypted.")
	ssm_getParametersByPathCmd.Flags().String("parameter-filters", "", "Filters to limit the request results.")
	ssm_getParametersByPathCmd.Flags().String("path", "", "The hierarchy for the parameter.")
	ssm_getParametersByPathCmd.Flags().Bool("recursive", false, "Retrieve all parameters within a hierarchy.")
	ssm_getParametersByPathCmd.Flags().Bool("with-decryption", false, "Retrieve all parameters in a hierarchy with their value decrypted.")
	ssm_getParametersByPathCmd.Flag("no-recursive").Hidden = true
	ssm_getParametersByPathCmd.Flag("no-with-decryption").Hidden = true
	ssm_getParametersByPathCmd.MarkFlagRequired("path")
	ssmCmd.AddCommand(ssm_getParametersByPathCmd)
}
