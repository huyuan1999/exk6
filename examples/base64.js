// @Time    : 2023/11/30 7:50 PM
// @Author  : HuYuan
// @File    : base64.js
// @Email   : hy2803660215@163.com

import base64 from 'k6/x/base64'

export default function () {
    const localData = base64.localFileEncode("/etc/fstab")
    console.log("localFileEncode: ", localData)

    const httpData = base64.httpFileEncode("https://mirrors.tuna.tsinghua.edu.cn/static/img/logo-small.png")
    console.log("httpData: ", httpData)
}
