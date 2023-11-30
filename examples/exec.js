// @Time    : 2023/11/30 7:53 PM
// @Author  : HuYuan
// @File    : exec.js
// @Email   : hy2803660215@163.com

import { check, sleep } from 'k6';
import exec from 'k6/x/exec';

export default function () {
    const workdir = ""
    exec.setWorkDir(workdir)
    const result = exec.command("/bin/bash", "-c", "ls -a")
    console.log(result)
}
