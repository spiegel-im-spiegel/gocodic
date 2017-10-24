// Copyright © 2017 Spiegel
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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/gocodic"
	"github.com/spiegel-im-spiegel/gocodic/options"
	"github.com/spiegel-im-spiegel/gocodic/response"
)

// transCmd represents the trans command
var transCmd = &cobra.Command{
	Use:   "trans [flags] <word>",
	Short: "Ttansration API  for codic.jp",
	Long:  "Ttansration API  for codic.jp",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return os.ErrInvalid
		}
		opts, err := options.NewOptions(viper.GetString("token"))
		if err != nil {
			return err
		}
		pid := viper.GetInt("projectid")
		if pid > 0 {
			opts.Add(options.ProjectID(pid))
		}
		casing, err := options.NewCasingOption(viper.GetString("casing"))
		if err == nil {
			opts.Add(casing)
		}
		style, err := options.NewAcronymStyleOption(viper.GetString("style"))
		if err == nil {
			opts.Add(style)
		}
		for _, arg := range args {
			opts.Add(options.Text(arg))
		}

		res, err := gocodic.Translate(opts)
		if err != nil {
			return err
		}
		if !res.IsSuccess() {
			ed, err2 := response.DecodeError(res.Body())
			if err != nil {
				return err2
			}
			result = ""
			resultErr = ""
			for _, d := range ed.Errors {
				resultErr += d.Message + "\n"
			}
			return nil
		}
		sd, err := response.DecodeSuccessTrans(res.Body())
		if err != nil {
			return err
		}
		result = ""
		resultErr = ""
		for _, d := range sd {
			result += d.TranslatedText + "\n"
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(transCmd)

	transCmd.Flags().IntP("projectid", "p", 0, "project_id parameter")
	transCmd.Flags().StringP("casing", "c", "", "casing parameter")
	transCmd.Flags().StringP("style", "s", "", "acronym_style parameter")
	viper.BindPFlag("projectid", transCmd.Flags().Lookup("projectid"))
	viper.BindPFlag("casing", transCmd.Flags().Lookup("casing"))
	viper.BindPFlag("style", transCmd.Flags().Lookup("style"))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
