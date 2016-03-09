// Copyright 2014 fabux authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package fabcore

import (
	"strings"

	"github.com/Unknwon/goconfig"
)

// GetConfig parse the configuration file and return
func GetConfig(configfile string) *goconfig.ConfigFile {
	c, err := goconfig.LoadConfigFile(configfile)
	if err == nil {
		return c
	}
	return nil
}

func GetGlobal(c *goconfig.ConfigFile) map[string]string {
	global := make(map[string]string)
	cmds := c.GetKeyList("global")
	for _, v := range cmds {
		row, _ := c.GetValue("global", v)
		global[strings.Trim(v, "#")] = row
	}
	return global
}

func GetHost(c *goconfig.ConfigFile) map[string][]string {
	server := make(map[string][]string)
	host := c.GetKeyList("host")
	for _, v := range host {
		row, _ := c.GetValue("host", v)
		server[strings.Trim(v, "#")] = strings.Split(row, "`")
	}
	return server
}

func GetLocal(c *goconfig.ConfigFile) map[string]string {
	local := make(map[string]string)
	dir := c.GetKeyList("local")
	for _, v := range dir {
		row, _ := c.GetValue("local", v)
		local[strings.Trim(v, "#")] = row
	}
	return local
}
