// SiYuan - From thought to insight, with agents
// Copyright (c) 2020-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package util

// bypassKey 是用于跳过第三方同步服务提供商会员校验的特性键。
// 启动时调用 DisableFeature 将其登记到运行期禁用列表，InitConf 会把该列表拷贝到
// Conf.System.DisabledFeatures，供服务端 IsSubscriber/IsPaidUser 以及前端
// window.siyuan.config.system.disabledFeatures 统一读取。
const bypassKey = "sync-provider-membership-check-bypass"

// init 在进程启动时禁用第三方同步服务提供商的会员校验。
// 通过单独的 init 函数登记特性键，避免直接修改 working.go、kernel/mobile/kernel.go 等
// 上游频繁改动的文件，从而降低合并冲突概率。该函数在 model.InitConf 读取
// util.DisabledFeatures 之前执行（model 包导入 util，util 的初始化先完成）。
func init() {
	DisableFeature(bypassKey)
}
