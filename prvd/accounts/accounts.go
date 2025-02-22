package accounts

import (
	"os"

	"github.com/provideplatform/provide-cli/prvd/common"
	"github.com/spf13/cobra"
)

var AccountID string

var AccountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "Manage signing identities & accounts",
	Long: `Various APIs are exposed to provide convenient access to elliptic-curve cryptography
(ECC) helper methods such as generating managed (custodial) keypairs.

For convenience, it is also possible to generate keypairs with this utility which you (or your application)
is then responsible for securing. You should securely store any keys generated using this API. If you are
looking for hierarchical deterministic support, check out the wallets API.`,
	Run: func(cmd *cobra.Command, args []string) {
		common.CmdExistsOrExit(cmd, args)

		generalPrompt(cmd, args, "")

		defer func() {
			if r := recover(); r != nil {
				os.Exit(1)
			}
		}()
	},
}

func init() {
	AccountsCmd.AddCommand(accountsListCmd)
	AccountsCmd.AddCommand(accountsInitCmd)
	AccountsCmd.Flags().BoolVarP(&optional, "optional", "", false, "List all the optional flags")
	AccountsCmd.Flags().BoolVarP(&paginate, "paginate", "", false, "List pagination flags")
}
