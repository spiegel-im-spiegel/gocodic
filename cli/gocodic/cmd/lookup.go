// Copyright Â© 2017 Spiegel
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/gocodic"
	"github.com/spiegel-im-spiegel/gocodic/options"
	"github.com/spiegel-im-spiegel/gocodic/response"
)

// lookupCmd represents the lookup command
var lookupCmd = &cobra.Command{
	Use:   "lookup [flags] <query string>",
	Short: "Lookup CED API for codic.jp",
	Long:  "Lookup CED API for codic.jp",
	RunE: func(cmd *cobra.Command, args []string) error {
		eid, err := cmd.Flags().GetInt("entryid")
		if err != nil {
			return os.ErrInvalid
		}
		var opts *options.Options
		if eid > 0 {
			opts, err = options.NewOptions(options.CmdCED, viper.GetString("token"))
		} else {
			opts, err = options.NewOptions(options.CmdCEDQuery, viper.GetString("token"))
		}
		if err != nil {
			return err
		}
		if eid == 0 {
			if len(args) < 1 {
				return os.ErrInvalid
			}
			opts.Add(options.Query(args[0]))
			ct, err2 := cmd.Flags().GetInt("count")
			if err2 != nil {
				return os.ErrInvalid
			}
			if ct > 0 {
				opts.Add(options.Count(ct))
			}
		}
		jsonFlag := viper.GetBool("json")

		res, err := gocodic.LookupCED(opts, options.EntryID(eid))
		if err != nil {
			return err
		}
		if !res.IsSuccess() {
			result = ""
			resultErr = ""
			if jsonFlag {
				resultErr = res.String()
				return nil
			}
			ed, err2 := response.DecodeError(res.Body())
			if err != nil {
				return err2
			}
			for _, d := range ed.Errors {
				resultErr += d.Message + "\n"
			}
			return nil
		}
		result = ""
		resultErr = ""
		if jsonFlag {
			result = res.String()
			return nil
		}
		if opts.Cmd() == options.CmdCEDQuery {
			sd, err := response.DecodeSuccessPLookup(res.Body())
			if err != nil {
				return err
			}
			for _, d := range sd {
				result += fmt.Sprintf("%d:%s, %s\n", d.ID, d.Title, d.Digest)
			}
		} else {
			sd, err := response.DecodeSuccessEntry(res.Body())
			if err != nil {
				return err
			}
			result = fmt.Sprintf("%d:%s, %s\n", sd.ID, sd.Title, sd.Digest)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(lookupCmd)

	lookupCmd.Flags().IntP("count", "c", 0, "count parameter")
	lookupCmd.Flags().IntP("entryid", "e", 0, "CED entry ID")
}
